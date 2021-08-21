package model

// Poker 每张牌
type Poker struct {
	Face int
	Color string
}

// Data 输入的Data结构
type Data struct {
	Alice string
	Bob string
	Result int
}

// InputData 从json获取的数据结构
type InputData struct {
	Matches []Data
}

// Turn 每一局的所以数据
type Turn struct {
	Alice HandCards
	Bob HandCards
	Winner int
}
// Feature 每一组牌的特征值
type Feature struct {
	Continue int
	SameCards int
	Flush bool
	Finish bool
}

type HandCards struct {
	Pokers []Poker
	Deck []Deck
	Level int
	Finish bool
}

type Deck struct {
	Face int
	Quantity int
	Pinter int
}