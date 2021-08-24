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

// HandCards 每位选手的手牌及其特征
type HandCards struct {
	Pokers []Poker
	Deck []Deck
	Level int			//从9阶同花顺到1阶单牌
	Finish bool			//被高阶处理后避免再判断低阶以节约运算
}

// Deck 牌组 相同牌为一组
type Deck struct {
	Face int			//相同牌的牌面
	Quantity int		//相同牌的数量
	Pinter int			//该牌组在所有牌中的位置
}

// Series 系列 连续的牌为一个系列
type Series struct {
	Length int			//该连续序列的长度
	Pinter int			//该连续序列的起始位置
}