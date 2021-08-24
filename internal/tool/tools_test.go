package tool

import (
	"Seven_pokers/internal/model"
	"fmt"
	"testing"
)

var data = "2hTh4dAdAc5s3s"

func createPokers(data string)(turn model.Turn){
	length := len(data)/2
	a := make([]model.Poker, length, length)
	b := make([]model.Poker, length, length)
	for i:=0;i< length;i++{
		a[i].Face = ChangeFaceToNumber(string(data[i*2]))
		a[i].Color = string(data[i*2+1])
		b[i].Face = ChangeFaceToNumber(string(data[i*2]))
		b[i].Color = string(data[i*2+1])
	}
	turn.Alice.Pokers = a
	turn.Bob.Pokers = b
	return
}

func TestCheckContinueLength(t *testing.T) {
	turn := createPokers(data)
	turn.Alice.Pokers = Sort(turn.Alice.Pokers)
	bool,pointer,result:= CheckContinueLength(turn.Alice.Pokers)
	fmt.Println(bool,pointer,result)

}

func TestGetDeck(t *testing.T) {
	turn := createPokers(data)
	decks := GetDeck(turn.Alice.Pokers)
	fmt.Println(decks)

}
