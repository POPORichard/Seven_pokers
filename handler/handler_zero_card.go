package handler

import (
	"Seven_pokers/model"
	"Seven_pokers/tool"
)

// analysesWithZeroCard 分析有赖子的牌
func analysesWithZeroCard(handCards model.HandCards)model.HandCards {
	var tmpHandCardsForDecks model.HandCards
	var tmpHandCardsForSeries model.HandCards

	tmpHandCardsForDecks.Pokers = make([]model.Poker,7,7)
	tmpHandCardsForSeries.Pokers = make([]model.Poker,7,7)

	copy(tmpHandCardsForDecks.Pokers,handCards.Pokers)
	copy(tmpHandCardsForSeries.Pokers,handCards.Pokers)

	tmpHandCardsForDecks = analyseDecksWithZeroCard(tmpHandCardsForDecks)
	tmpHandCardsForDecks = getLevelByDeck(tmpHandCardsForDecks)

	tmpHandCardsForSeries = analyseContinueWithCardZero(tmpHandCardsForSeries)

	if tmpHandCardsForSeries.Level > tmpHandCardsForDecks.Level {
		handCards = tmpHandCardsForSeries
	}else {
		sortByDeck(&tmpHandCardsForDecks)
		handCards = tmpHandCardsForDecks
	}

	return handCards
}

// analyseDecksWithZeroCard 分析有赖子的牌的Decks
func analyseDecksWithZeroCard(handCards model.HandCards)model.HandCards {
	finish := false
	handCards.Deck = tool.GetDeck(handCards.Pokers)
	decks := handCards.Deck
	for i := range decks {
		if decks[i].Quantity == 4 {
			if decks[i].Face == 14 {
				if decks[i+1].Face == 13 {
					handCards.Pokers[6].Face = 12
				} else {
					handCards.Pokers[6].Face = 13
				}
			} else {
				handCards.Pokers[6].Face = decks[i].Face + 1
			}
			finish = true
			break
		}
	}

	if !finish {
		for i := range decks {
			if decks[i].Quantity == 3 {
				handCards.Pokers[6].Face = decks[i].Face
				finish = true
				break
			}
		}
	}

	if !finish {
		for i := range decks {
			if decks[i].Quantity == 2 {
				handCards.Pokers[6].Face = decks[i].Face
				finish = true
				break
			}
		}
	}

	if !finish {
		handCards.Pokers[6].Face = handCards.Pokers[0].Face
		finish = true
	}

	if !finish {
		panic("Error")
	}

	handCards.Pokers = tool.Sort(handCards.Pokers)
	handCards.Deck = tool.GetDeck(handCards.Pokers)

	return  handCards

}

// analyseContinueWithCardZero 分析带有赖子的牌的连续性并处理A
func analyseContinueWithCardZero(handCards model.HandCards)(tmpHandCards model.HandCards) {
	//因同花的等级高于顺子则先判断同花
	flush, color, length := tool.CheckFlush(handCards.Pokers)
	handCards.Pokers[6].Color = color
	// 可以达到同花的条件时
	if flush {
		flushSeries := make([]model.Poker, 0, 7)
		for i := range handCards.Pokers {
			if handCards.Pokers[i].Color == color {
				flushSeries = append(flushSeries, handCards.Pokers[i])
			}
		}

		// 这里可以在最开始处理A因没有重复的牌
		if flushSeries[0].Face == 14 {
			flushSeries = handleA(flushSeries)
			flushSeries = tool.Sort(flushSeries)
		}

		series := analyseSeries(flushSeries)
		var tmpSeries []model.Series

		// 根据series判断cardZero的插入位置
		for i := range series {
			if series[i].Length == 5 {
				if flushSeries[length-1].Face == 0 {
					if flushSeries[series[i].Pinter].Face == 14 {
						flushSeries[length-1].Face = flushSeries[series[i].Pinter+series[i].Length-1].Face - 1
					} else {
						flushSeries[length-1].Face = flushSeries[series[i].Pinter].Face + 1
					}

				}
				flushSeries = tool.Sort(flushSeries)
				tmpSeries = analyseSeries(flushSeries)
				target:=0
				for a := range tmpSeries{
					if tmpSeries[a].Length > 4{target = a}
				}
				for t := 0; t < 5; t++ {
					tmpHandCards.Pokers = append(tmpHandCards.Pokers, flushSeries[t+tmpSeries[target].Pinter])
				}
				tmpHandCards.Level = 9
				tmpHandCards.Finish = true
				return
			}

			if series[i].Length == 4 {
				if flushSeries[series[i].Pinter].Face == 14 {
					flushSeries[length-1].Face = flushSeries[series[i].Pinter+series[i].Length-1].Face - 1
				} else {
					flushSeries[length-1].Face = flushSeries[series[i].Pinter].Face + 1
				}
				flushSeries = tool.Sort(flushSeries)
				tmpSeries = analyseSeries(flushSeries)
				target:=0
				for a := range tmpSeries{
					if tmpSeries[a].Length > 4{target = a}
				}
				for t := 0; t < 5; t++ {
					tmpHandCards.Pokers = append(tmpHandCards.Pokers,flushSeries[t+tmpSeries[target].Pinter])
				}
				tmpHandCards.Level = 9
				tmpHandCards.Finish = true
				return
			}

			if i+2 < len(series) {
				if series[i].Length > 1 && series[i+1].Length > 1 && flushSeries[series[i].Pinter].Face-3 == flushSeries[series[i+1].Pinter].Face {
					flushSeries[length-1].Face = flushSeries[series[i].Pinter].Face - 2
					tmpSeries = analyseSeries(flushSeries)
					for t := 0; t < 5; t++ {
						tmpHandCards.Pokers = append(tmpHandCards.Pokers,flushSeries[t+tmpSeries[i].Pinter])
					}
					tmpHandCards.Level = 9
					tmpHandCards.Finish = true
					return
				}
				if series[i].Length > 0 && series[i+1].Length > 2 && flushSeries[series[i].Pinter].Face-2 == flushSeries[series[i+1].Pinter].Face {
					flushSeries[length-1].Face = flushSeries[series[i].Pinter].Face - 1
					flushSeries = tool.Sort(flushSeries)
					tmpSeries = analyseSeries(flushSeries)
					target:=0
					for a := range tmpSeries{
						if tmpSeries[a].Length > 4{target = a}
					}
					for t := 0; t < 5; t++ {
						tmpHandCards.Pokers = append(tmpHandCards.Pokers,flushSeries[t+tmpSeries[target].Pinter])
					}
					tmpHandCards.Level = 9
					tmpHandCards.Finish = true
					return
				}
				if series[i].Length > 2 && series[i+1].Length > 0 && flushSeries[series[i].Pinter].Face-4 == flushSeries[series[i+1].Pinter].Face {
					flushSeries[length-1].Face = flushSeries[series[i].Pinter].Face - 3
					flushSeries = tool.Sort(flushSeries)
					tmpSeries = analyseSeries(flushSeries)
					target:=0
					for a := range tmpSeries{
						if tmpSeries[a].Length > 4{target = a}
					}
					for t := 0; t < 5; t++ {
						tmpHandCards.Pokers = append(tmpHandCards.Pokers, flushSeries[t+tmpSeries[target].Pinter])
					}
					tmpHandCards.Level = 9
					tmpHandCards.Finish = true
					return
				}
			}
		}

		//若没有cardZero的位置使其连续则cardZero变为最大值
		flushSeries[len(flushSeries)-1].Face = 14
		tmpHandCards.Pokers = append(tmpHandCards.Pokers, flushSeries[len(flushSeries)-1])
		for t := 0; t < 4; t++ {
			tmpHandCards.Pokers = append(tmpHandCards.Pokers, flushSeries[t])
		}
		tmpHandCards.Level = 6
		return 

	}

	//在非同花的情况下需要注意重复的牌
	series := analyseSeries(handCards.Pokers)
	var tmpSeries []model.Series
	length = len(handCards.Pokers)
	// 根据series判断cardZero的插入位置
	for i := range series {
		if series[i].Length >4 {
			if handCards.Pokers[length-1].Face == 0 {
				if handCards.Pokers[series[i].Pinter].Face == 14 {
					handCards.Pokers[length-1].Face = handCards.Pokers[series[i].Pinter+series[i].Length-1].Face - 1
				} else {
					handCards.Pokers[length-1].Face = handCards.Pokers[series[i].Pinter].Face + 1
				}
			}
			handCards.Pokers = tool.Sort(handCards.Pokers)
			tmpSeries = analyseSeries(handCards.Pokers)
			target:=0
			for a := range tmpSeries{
				if tmpSeries[a].Length > 4{target = a}
			}
			add :=0
			tmp := 0
			for t := 0; t < 5; t++ {
				if handCards.Pokers[t+tmpSeries[target].Pinter+add].Face == tmp{add++}
				tmp = handCards.Pokers[t+tmpSeries[target].Pinter+add].Face
				tmpHandCards.Pokers = append(tmpHandCards.Pokers,handCards.Pokers[t+tmpSeries[target].Pinter+add])
			}
			tmpHandCards.Level = 5
			return
		}

		if series[i].Length == 4 {
			if handCards.Pokers[series[i].Pinter].Face == 14 {
				handCards.Pokers[length-1].Face = handCards.Pokers[series[i].Pinter+series[i].Length-1].Face - 1
			} else {
				handCards.Pokers[length-1].Face = handCards.Pokers[series[i].Pinter].Face + 1
			}
			handCards.Pokers = tool.Sort(handCards.Pokers)
			tmpSeries = analyseSeries(handCards.Pokers)
			target:=0
			for a := range tmpSeries{
				if tmpSeries[a].Length > 4{target = a}
			}
			add :=0
			tmp := 0
			for t := 0; t < 5; t++ {
				if handCards.Pokers[t+tmpSeries[target].Pinter+add].Face == tmp{add++}
				tmp = handCards.Pokers[t+tmpSeries[target].Pinter+add].Face
				tmpHandCards.Pokers = append(tmpHandCards.Pokers,handCards.Pokers[t+tmpSeries[target].Pinter+add])
			}
			tmpHandCards.Level = 5
			tmpHandCards.Finish = true
			return
		}

		if i+2 < len(series) {
			if series[i].Length > 1 && series[i+1].Length > 1 && handCards.Pokers[series[i].Pinter].Face-3 == handCards.Pokers[series[i+1].Pinter].Face {
				handCards.Pokers[length-1].Face = handCards.Pokers[series[i].Pinter].Face - 2
				handCards.Pokers = tool.Sort(handCards.Pokers)
				tmpSeries = analyseSeries(handCards.Pokers)
				target:=0
				for a := range tmpSeries{
					if tmpSeries[a].Length > 4{target = a}
				}
				add :=0
				tmp := 0
				for t := 0; t < 5; t++ {
					if handCards.Pokers[t+tmpSeries[target].Pinter+add].Face == tmp{add++}
					tmp = handCards.Pokers[t+tmpSeries[target].Pinter+add].Face
					tmpHandCards.Pokers = append(tmpHandCards.Pokers,handCards.Pokers[t+tmpSeries[target].Pinter+add])
				}
				tmpHandCards.Level = 5
				tmpHandCards.Finish = true
				return
			}
			if series[i].Length > 0 && series[i+1].Length > 2 && handCards.Pokers[series[i].Pinter].Face-2 == handCards.Pokers[series[i+1].Pinter].Face {
				handCards.Pokers[length-1].Face = handCards.Pokers[series[i].Pinter].Face - 1
				handCards.Pokers = tool.Sort(handCards.Pokers)
				tmpSeries = analyseSeries(handCards.Pokers)
				target:=0
				for a := range tmpSeries{
					if tmpSeries[a].Length  > 4{target = a}
				}
				add :=0
				tmp := 0
				for t := 0; t < 5; t++ {
					if handCards.Pokers[t+tmpSeries[target].Pinter+add].Face == tmp{add++}
					tmp = handCards.Pokers[t+tmpSeries[target].Pinter+add].Face
					tmpHandCards.Pokers = append(tmpHandCards.Pokers,handCards.Pokers[t+tmpSeries[target].Pinter+add])
				}
				tmpHandCards.Level = 5
				tmpHandCards.Finish = true
				return
			}
			if series[i].Length > 2 && series[i+1].Length > 0 && handCards.Pokers[series[i].Pinter].Face-4 == handCards.Pokers[series[i+1].Pinter].Face {
				handCards.Pokers[length-1].Face = handCards.Pokers[series[i].Pinter].Face - 3
				handCards.Pokers = tool.Sort(handCards.Pokers)
				tmpSeries = analyseSeries(handCards.Pokers)
				target:=0
				for a := range tmpSeries{
					if tmpSeries[a].Length > 4{target = a}
				}
				add :=0
				tmp := 0
				for t := 0; t < 5; t++ {
					if handCards.Pokers[t+tmpSeries[target].Pinter+add].Face == tmp{add++}
					tmp = handCards.Pokers[t+tmpSeries[target].Pinter+add].Face
					tmpHandCards.Pokers = append(tmpHandCards.Pokers,handCards.Pokers[t+tmpSeries[target].Pinter+add])
				}
				tmpHandCards.Level = 5
				tmpHandCards.Finish = true
				return
			}
		}
	}

	// 处理A
	if handCards.Pokers[0].Face == 14 {
		handCards.Pokers = handleA(handCards.Pokers)
		handCards.Pokers = tool.Sort(handCards.Pokers)
		//A被处理了
		if handCards.Pokers[0].Face != 14{
			tmpSeries = analyseSeries(handCards.Pokers)
			target:=0
			for a := range tmpSeries{
				if tmpSeries[a].Length > 4{target = a}
			}
			add :=0
			tmp := 0
			for t := 0; t < 5; t++ {
				if handCards.Pokers[t+tmpSeries[target].Pinter+add].Face == tmp{add++}
				tmp = handCards.Pokers[t+tmpSeries[target].Pinter+add].Face
				tmpHandCards.Pokers = append(tmpHandCards.Pokers,handCards.Pokers[t+tmpSeries[target].Pinter+add])
			}
			tmpHandCards.Level = 5
		}
		return
	}

	tmpHandCards = handCards
	return
}

// analyseSeries 分析扑克牌以获得连续性特征:series
func analyseSeries(pokers []model.Poker) (series []model.Series) {
	length := len(pokers)
	jumpOvertimes := 0

	i:=0
	t := 1
	for i = 0; i < length-1; i++ {
		if pokers[i].Face-1 == pokers[i+1].Face {
			t++
		} else if pokers[i].Face == pokers[i+1].Face {
			jumpOvertimes++
			continue
		} else {
			ser := model.Series{
				Length: t,
				Pinter: i - t - jumpOvertimes + 1,
			}
			series = append(series, ser)
			jumpOvertimes = 0
			t = 1
		}
	}

	series = append(series, model.Series{
		Length: t,
		Pinter: i - t - jumpOvertimes + 1,
	})

	return series
}

// handleA 在有赖子的情况下处理A
func handleA(pokers []model.Poker) []model.Poker {
	a := [4]bool{false}

	for i := range pokers {
		if pokers[i].Face == 5 {
			a[0] = true
		}
		if pokers[i].Face == 4 {
			a[1] = true
		}
		if pokers[i].Face == 3 {
			a[2] = true
		}
		if pokers[i].Face == 2 {
			a[3] = true
		}
	}
	t := 0
	need := 0
	for i := range a {
		if a[i] {
			t++
		} else {
			need = 5 - i
		}
	}

	if t > 2 {
		pokers[0].Face = 1
		pokers[len(pokers)-1].Face = need
	}

	return pokers
}
