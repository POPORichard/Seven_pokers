package main

import (
	"Seven_pokers/handler"
	"Seven_pokers/model"
	"fmt"
	"time"
)

func main () {
	start := time.Now()
	data := handler.ReadDataToModel("./seven_cards_with_ghost.result.json")

	for i := range data {
		turn := handler.CreateTurn(&data[i])
		handler.Analyse(&turn)
		handler.JudgeWinner(&turn)
		if turn.Winner != data[i].Result {
			outPutErrors(turn,data[i])
			panic("result wrong")
		}
	}
	cost := time.Since(start)
	fmt.Println("成功！ 耗时：", cost)
}

func outPutErrors(turn model.Turn,data model.Data){
	fmt.Println("-----Error-----")
	fmt.Println(data.Alice)
	fmt.Println("Alice.Pokers:")
	fmt.Println(turn.Alice.Pokers)
	fmt.Println("Alice.Decks:")
	fmt.Println(turn.Alice.Deck)
	fmt.Println("Alice.Level:")
	fmt.Println(turn.Alice.Level,turn.Alice.Finish)
	fmt.Println("-------------------------------------------")
	fmt.Println(data.Bob)
	fmt.Println("Bob.Pokers:")
	fmt.Println(turn.Bob.Pokers)
	fmt.Println("Bob.Decks:")
	fmt.Println(turn.Bob.Deck)
	fmt.Println("Bob.Level:")
	fmt.Println(turn.Bob.Level,turn.Bob.Finish)
	fmt.Println("your winner is :",turn.Winner,"Should be :",data.Result)
}

