package model

// Money は金額を表す。負の値を許さない。
type Money struct {
	amount int64 // 金額の値（円単位）
}

// NewMoney は金額を生成する。
// amount は金額の値（円単位）を表す。
// 戻り値は生成した金額と、負の値が渡された場合の *DomainError である。
func NewMoney(amount int64) (Money, error) {
	panic("未実装")
}

// Amount は金額の値を返す。
// 戻り値は円単位の金額の値である。
func (m Money) Amount() int64 {
	panic("未実装")
}
