package model

// CouponCode は券種コードを表す。クーポンを識別する。
type CouponCode struct {
	value string // 券種コードの値
}

// NewCouponCode は券種コードを生成する。
// value は券種コードの値を表す。
// 戻り値は生成した券種コードと、業務上の制約に反する値が渡された場合の *DomainError である。
func NewCouponCode(value string) (CouponCode, error) {
	panic("未実装")
}

// Value は券種コードの値を返す。
// 戻り値は券種コードの値である。
func (c CouponCode) Value() string {
	panic("未実装")
}
