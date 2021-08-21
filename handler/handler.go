package handler

import (
	"Seven_pokers/model"
	"Seven_pokers/tool"
)

func CreateTurn(data *model.Data)model.Turn{
	turn := tool.PutCardIntoHand(data)
	turn.Alice.Pokers = tool.Sort(turn.Alice.Pokers)
	turn.Bob.Pokers = tool.Sort(turn.Bob.Pokers)
	return turn
}

func Analyse(turn *model.Turn){

	analyseDecks(&turn.Alice)
	analyseDecks(&turn.Bob)

	getLevelByDeck(&turn.Alice)
	getLevelByDeck(&turn.Bob)

	if turn.Alice.Pokers[6].Face == 0{

	}

	if !turn.Alice.Finish{
		analyseFlush(&turn.Alice)
	}

	if !turn.Alice.Finish{
		analyseContinue(&turn.Alice)
	}

	if !turn.Bob.Finish{
		analyseFlush(&turn.Bob)
	}

	if !turn.Bob.Finish{
		analyseContinue(&turn.Bob)
	}

	if !turn.Alice.Finish{
		sortByDeck(&turn.Alice)
	}
	if !turn.Bob.Finish{
		sortByDeck(&turn.Bob)
	}




}

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

func getLevelByDeck(handCards *model.HandCards){
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
		//handCards.Finish = true
		return
	}
	if three == 2{
		handCards.Level = 7
		//handCards.Finish = true
		return
	}
	if three == 1 && two > 0{
		handCards.Level = 7
		//handCards.Finish =true
		return
	}
	if three == 1{
		handCards.Level = 4
		return
	}
	if two > 1{
		handCards.Level = 3
		return
	}
	if two == 1{
		handCards.Level = 2
		return
	}
	handCards.Level = 1
	return

}

// sortByDeck 按Deck选择牌
func sortByDeck(handCardsCards *model.HandCards) {
	newPokers := make([]model.Poker,0,5)
	decks := tool.SortDeck(handCardsCards.Deck)
	num := 0
	if handCardsCards.Level == 3{
		newPokers = append(newPokers,handCardsCards.Pokers[decks[0].Pinter])
		newPokers = append(newPokers,handCardsCards.Pokers[decks[0].Pinter+1])
		newPokers = append(newPokers,handCardsCards.Pokers[decks[1].Pinter])
		newPokers = append(newPokers,handCardsCards.Pokers[decks[1].Pinter+1])
		if decks[2].Face > decks[3].Face{
			newPokers = append(newPokers,handCardsCards.Pokers[decks[2].Pinter])
		}else{
			newPokers = append(newPokers,handCardsCards.Pokers[decks[3].Pinter])
		}

	}else if handCardsCards.Level == 8{
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
	if handCards.Pokers[0].Face == 15{handCards.Pokers[0].Color = color}
	if flush {
		newPokers := make([]model.Poker, 0, 7)
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

