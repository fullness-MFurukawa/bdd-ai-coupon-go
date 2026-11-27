package model

// Cart は買い物かごを表す。ポインタ *Cart で扱う。
type Cart struct {
	totalAmount    Money          // 購入金額
	membershipRank MembershipRank // 会員ランク
}

// NewCart は買い物かごを生成する。
// totalAmount は購入金額、membershipRank は会員ランクを表す。
// 戻り値は生成した買い物かごである。
func NewCart(totalAmount Money, membershipRank MembershipRank) *Cart {
	panic("未実装")
}

// TotalAmount は購入金額を返す。
// 戻り値は購入金額である。
func (c *Cart) TotalAmount() Money {
	panic("未実装")
}

// MembershipRank は会員ランクを返す。
// 戻り値は会員ランクである。
func (c *Cart) MembershipRank() MembershipRank {
	panic("未実装")
}
