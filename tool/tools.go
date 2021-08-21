package tool

import (
	"Seven_pokers/model"
	"strconv"
)

// ChangeFaceToNumber 把牌面从string转换为int
// 牌面为花牌则按 T:10;J:11;Q:12,K:13;A:14来进行转换
// 返回牌面的数值int型
func ChangeFaceToNumber(f string)int{
	switch f {
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	case "X":
		return 0
	default:
		int,_:=strconv.Atoi(f)
		return int
	}
}

// PutCardIntoHand 将字符串信息转换写入model.poker并放到当前回合里
// 返回model.Turn
func PutCardIntoHand(data *model.Data) (turn model.Turn) {
	length := len(data.Alice)/2
	a := make([]model.Poker, length, length)
	b := make([]model.Poker, length, length)
	for i:=0;i< length;i++{
		a[i].Face = ChangeFaceToNumber(string(data.Alice[i*2]))
		a[i].Color = string(data.Alice[i*2+1])
		b[i].Face = ChangeFaceToNumber(string(data.Bob[i*2]))
		b[i].Color = string(data.Bob[i*2+1])
	}
	turn.Alice.Pokers = a
	turn.Bob.Pokers = b
	return
}

// Sort 把牌按Face从大到小排列
func Sort(pokers []model.Poker)[]model.Poker{
	var tmp model.Poker
	length := len(pokers)

	var target int
	var biggest int

	for i :=range pokers{
		biggest = pokers[i].Face
		target=i
		for t:=i+1;t< length;t++{
			if biggest < pokers[t].Face{
				biggest = pokers[t].Face
				target = t
			}
		}
		if target !=i {
			tmp = pokers[i]
			pokers[i] = pokers[target]
			pokers[target] = tmp
		}
	}
	return pokers
}

func CheckFlush(pokers []model.Poker)(isFlush bool, color string,length int){
	isFlush = false
	hasCardZero := false
	a := [4]int{0}
	t := 0
	for i := range pokers{
		if pokers[i].Color == "s"{
			a[0]++
		}
		if pokers[i].Color == "h"{
			a[1]++
		}
		if pokers[i].Color == "d"{
			a[2]++
		}
		if pokers[i].Color == "c"{
			a[3]++
		}
		if pokers[i].Color == "n"{
			hasCardZero = true
		}
	}
	quantity:=a[0]
	for i:=1;i<4;i++{
		if quantity<a[i]{
			quantity = a[i]
			t = i
		}
	}
	if hasCardZero{
		a[t]++
	}
	if a[t]>4{
		length = a[t]
		isFlush = true
		if t == 0{color = "s"; return}
		if t == 1{color = "h"; return}
		if t == 2{color = "d"; return}
		if t == 3{color = "c"; return}
		panic("Error color!")
	}
	return false,"",0

}

func CheckContinueLength(pokers []model.Poker)(bool,int,int){
	length := len(pokers)
	result := 1
	pointer := 0
	jumpOvertimes := 0
	t:=1

	if pokers[length-1].Face == 0 {}
	for i:=0;i< length-1;i++{
		if pokers[i].Face-1 == pokers[i+1].Face {
			t++
		}else if pokers[i].Face == pokers[i+1].Face{
			jumpOvertimes++
			continue
		}else {
				jumpOvertimes = 0
				t = 1
		}
		if t>result{
			result = t
			pointer = i-t-jumpOvertimes+2
		}

	}
	if result > 3{
		return true,pointer,result
	}
	return false,pointer,result
}

func GetDeck (pokers []model.Poker)(decks []model.Deck){
	length := len(pokers)
	quantity := 1
	face := 0
	pointer := 0
	decks = make([]model.Deck,0,3)
	for i:=0;i<length-1;i++{
		if pokers[i].Face == pokers[i+1].Face {
			quantity = quantity+1
		}
		if pokers[i].Face != pokers[i+1].Face{
			face = pokers[i].Face
			pointer = i-quantity+1
			deck := model.Deck{Face:face,Quantity:quantity,Pinter:pointer}
			decks = append(decks,deck)
			quantity = 1
		}
		if i == length -2{
			face = pokers[i+1].Face
			pointer = i-quantity+2
			deck := model.Deck{Face:face,Quantity:quantity,Pinter:pointer}
			decks = append(decks,deck)

		}
	}
	return
}

func SortDeck (decks []model.Deck)[]model.Deck{
	newDecks := make([]model.Deck,0,7)
	length := len(decks)
	longest := decks[0].Quantity
	for i:=1;i<length;i++{
		if longest<decks[i].Quantity{
			longest = decks[i].Quantity
		}
	}
	for i:=longest;i>0;i--{
		for t:=0;t<length;t++{
			if i == decks[t].Quantity{
				newDecks = append(newDecks,decks[t])
			}
		}
	}
	return newDecks
}

func CompareEachCard(cardsA,cardsB []model.Poker)int{
	for i := range cardsA{
		if i == 5{
			break
		}
		if cardsA[i].Face > cardsB[i].Face{
			return 1
		}
		if cardsA[i].Face < cardsB[i].Face{
			return 2
		}
	}
	return 0
}
