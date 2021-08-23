package handler

import (
	"Seven_pokers/model"
	"Seven_pokers/tool"
)
// CreateTurn 创建一局游戏
func CreateTurn(data *model.Data)model.Turn{
	turn := tool.PutCardIntoHand(data)
	turn.Alice.Pokers = tool.Sort(turn.Alice.Pokers)
	turn.Bob.Pokers = tool.Sort(turn.Bob.Pokers)
	return turn
}

// Analyse 分析该轮游戏双方手牌
func Analyse(turn *model.Turn){
	// 处理Alice的牌
	if turn.Alice.Pokers[len(turn.Alice.Pokers)-1].Face == 0{
		turn.Alice = analysesWithZeroCard(turn.Alice)
	}else {
		analyseDecks(&turn.Alice)
		turn.Alice = getLevelByDeck(turn.Alice)
		if !turn.Alice.Finish{
			analyseFlush(&turn.Alice)
		}
		if !turn.Alice.Finish{
			analyseContinue(&turn.Alice)
		}
		if !turn.Alice.Finish{
			sortByDeck(&turn.Alice)
		}
	}

	// 处理Bob的牌
	if turn.Bob.Pokers[len(turn.Bob.Pokers)-1].Face == 0{
		turn.Bob = analysesWithZeroCard(turn.Bob)
	}else {
		analyseDecks(&turn.Bob)
		turn.Bob = getLevelByDeck(turn.Bob)
		if !turn.Bob.Finish{
			analyseFlush(&turn.Bob)
		}
		if !turn.Bob.Finish{
			analyseContinue(&turn.Bob)
		}
		if !turn.Bob.Finish{
			sortByDeck(&turn.Bob)
		}
	}
}

// JudgeWinner 判断赢家
func JudgeWinner(turn *model.Turn)*model.Turn{
	if turn.Alice.Level > turn.Bob.Level{
		turn.Winner = 1
	}else if turn.Bob.Level > turn.Alice.Level{
		turn.Winner = 2
	}else{
		turn.Winner = tool.CompareEachCard(turn.Alice.Pokers,turn.Bob.Pokers)
	}
	return turn
}

// analyseDecks 分析牌生成牌组Deck
func analyseDecks(handCards *model.HandCards){
	handCards.Deck = tool.GetDeck(handCards.Pokers)
}

// getLevelByDeck 根据Decks判断手牌等级
func getLevelByDeck(handCards model.HandCards) model.HandCards{
	fourOfAKind := 0

	two := 0
	three := 0
	for i := range handCards.Deck{
		if handCards.Deck[i].Quantity == 4{fourOfAKind++}
		if handCards.Deck[i].Quantity == 3{three++}
		if handCards.Deck[i].Quantity == 2{two++}
	}
	if fourOfAKind == 1{
		handCards.Level = 8
		return handCards
	}
	if three == 2{
		handCards.Level = 7
		return handCards
	}
	if three == 1 && two > 0{
		handCards.Level = 7
		return handCards
	}
	if three == 1{
		handCards.Level = 4
		return handCards
	}
	if two > 1{
		handCards.Level = 3
		return handCards
	}
	if two == 1{
		handCards.Level = 2
		return handCards
	}
	handCards.Level = 1
	return handCards

}

// sortByDeck 按Deck选择牌
func sortByDeck(handCardsCards *model.HandCards) {
	newPokers := make([]model.Poker,0,5)
	decks := tool.SortDeck(handCardsCards.Deck)
	num := 0
	if len(handCardsCards.Pokers)>5 && handCardsCards.Level == 3{
		newPokers = append(newPokers,handCardsCards.Pokers[decks[0].Pinter])
		newPokers = append(newPokers,handCardsCards.Pokers[decks[0].Pinter+1])
		newPokers = append(newPokers,handCardsCards.Pokers[decks[1].Pinter])
		newPokers = append(newPokers,handCardsCards.Pokers[decks[1].Pinter+1])
		if decks[2].Face > decks[3].Face{
			newPokers = append(newPokers,handCardsCards.Pokers[decks[2].Pinter])
		}else{
			newPokers = append(newPokers,handCardsCards.Pokers[decks[3].Pinter])
		}

	}else if len(handCardsCards.Pokers)>5 && handCardsCards.Level == 8{
		newPokers = append(newPokers,handCardsCards.Pokers[decks[0].Pinter])
		newPokers = append(newPokers,handCardsCards.Pokers[decks[0].Pinter+1])
		newPokers = append(newPokers,handCardsCards.Pokers[decks[0].Pinter+2])
		newPokers = append(newPokers,handCardsCards.Pokers[decks[0].Pinter+3])
		if decks[1].Face > decks[2].Face{
			newPokers = append(newPokers,handCardsCards.Pokers[decks[1].Pinter])
		}else{
			newPokers = append(newPokers,handCardsCards.Pokers[decks[2].Pinter])
		}

	} else{
	loop:
		for i:= range decks{
			for t:=0;t<decks[i].Quantity;t++{
				if num<5 {
					poker := handCardsCards.Pokers[decks[i].Pinter+t]
					newPokers = append(newPokers,poker)
					num++
				}else {
					break loop
				}
			}
		}
	}


	handCardsCards.Deck = decks
	handCardsCards.Pokers = newPokers
}

// analyseFlush 判断同花
func analyseFlush (handCards *model.HandCards){
	flush, color,length := tool.CheckFlush(handCards.Pokers)
	if flush {
		newPokers := make([]model.Poker, 0, 7)
		//将花色相同的牌放入新的切片中
		for i := range handCards.Pokers {
			if handCards.Pokers[i].Color != color {
				continue
			}
			newPokers = append(newPokers, handCards.Pokers[i])
		}

		_,pointer,result := tool.CheckContinueLength(newPokers)
		if result > 4{
			if length ==5{
				handCards.Pokers = newPokers
				handCards.Level = 9
				handCards.Finish = true
				return
			}
			handCards.Pokers = newPokers[pointer:pointer+5]
			handCards.Level = 9
			handCards.Finish = true
			return
		}
		// A牌需要变为1时
		if  result == 4 && newPokers[pointer].Face ==5 && newPokers[0].Face == 14{
			straightPokers := make([]model.Poker,0,5)
			for i:=0;i<4;i++{
				if i==0{
					straightPokers =append(straightPokers,newPokers[pointer+i])
					continue
				}
				if straightPokers[i-1].Face == newPokers[pointer+i].Face{
					pointer++
					i--
				}else{
					straightPokers =append(straightPokers,newPokers[pointer+i])
				}
			}
			newPokers[0].Face = 1
			straightPokers = append(straightPokers,newPokers[0])
			handCards.Pokers = straightPokers
			handCards.Level = 9
			handCards.Finish = true
		}
		if handCards.Level<6{
			handCards.Level = 6
			handCards.Pokers = newPokers[0:5]
			handCards.Finish = true
		}
		return

	}
}

// analyseContinue判断连续&处理A
func analyseContinue(handCards *model.HandCards){
	isStraight,pointer,result := tool.CheckContinueLength(handCards.Pokers)
	if isStraight && result > 4{
		newPokers := make([]model.Poker,0,5)

		for i:=0;i<5;i++{
			if i==0{
				newPokers =append(newPokers,handCards.Pokers[pointer+i])
				continue
			}
			if newPokers[i-1].Face == handCards.Pokers[pointer+i].Face{
				pointer++
				i--
			}else{
				newPokers =append(newPokers,handCards.Pokers[pointer+i])
			}
		}
		handCards.Pokers = newPokers
		handCards.Level = 5
		handCards.Finish = true
	}

	// A牌需要变为1时
	if isStraight && result == 4 && handCards.Pokers[pointer].Face ==5 && handCards.Pokers[0].Face == 14{
		newPokers := make([]model.Poker,0,5)
		for i:=0;i<4;i++{
			if i==0{
				newPokers =append(newPokers,handCards.Pokers[pointer+i])
				continue
			}
			if newPokers[i-1].Face == handCards.Pokers[pointer+i].Face{
				pointer++
				i--
			}else{
				newPokers =append(newPokers,handCards.Pokers[pointer+i])
			}
		}
		handCards.Pokers[0].Face = 1
		newPokers = append(newPokers,handCards.Pokers[0])
		handCards.Pokers = newPokers
		handCards.Level = 5
		handCards.Finish = true
	}

}


