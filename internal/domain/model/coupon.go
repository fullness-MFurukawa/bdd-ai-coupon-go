package model

// Coupon はクーポンを表す。識別子は券種コードで、ポインタ *Coupon で扱う。
type Coupon struct {
	couponCode            CouponCode     // 券種コード
	minimumPurchaseAmount Money          // 最低購入金額
	maximumDiscountAmount Money          // 割引上限額
	expirationDate        ExpirationDate // 有効期限
}

// NewCoupon はクーポンを生成する。
// couponCode は券種コード、minimumPurchaseAmount は最低購入金額、
// maximumDiscountAmount は割引上限額、expirationDate は有効期限を表す。
// 戻り値は生成したクーポンである。
func NewCoupon(
	couponCode CouponCode,
	minimumPurchaseAmount Money,
	maximumDiscountAmount Money,
	expirationDate ExpirationDate,
) *Coupon {
	panic("未実装")
}

// CouponCode は券種コードを返す。
// 戻り値はクーポンを識別する券種コードである。
func (c *Coupon) CouponCode() CouponCode {
	panic("未実装")
}

// MinimumPurchaseAmount は最低購入金額を返す。
// 戻り値は最低購入金額である。
func (c *Coupon) MinimumPurchaseAmount() Money {
	panic("未実装")
}

// MaximumDiscountAmount は割引上限額を返す。
// 戻り値は割引上限額である。
func (c *Coupon) MaximumDiscountAmount() Money {
	panic("未実装")
}

// ExpirationDate は有効期限を返す。
// 戻り値は有効期限である。
func (c *Coupon) ExpirationDate() ExpirationDate {
	panic("未実装")
}

// Equals は同一のクーポンかを判定する。券種コードが等しければ同一とみなす。
// other は比較するクーポンを表す。
// 戻り値は同一のクーポンである場合に true である。
func (c *Coupon) Equals(other *Coupon) bool {
	panic("未実装")
}
