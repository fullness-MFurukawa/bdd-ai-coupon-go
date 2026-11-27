package model

// CouponApplicationResult は適用結果を表す。
// クーポンを使えないことは業務上の正常な結果であるため、エラーではなくこの型で返す。
type CouponApplicationResult struct {
	successful     bool   // 成否
	errorMessage   string // エラーメッセージ（成功時は空文字列）
	discountAmount Money  // 割引額
	paymentAmount  Money  // 支払金額
}

// NewSuccessfulCouponApplicationResult は成功した適用結果を生成する。
// discountAmount は割引額、paymentAmount は支払金額を表す。
// 戻り値は成功した適用結果であり、エラーメッセージは空文字列となる。
func NewSuccessfulCouponApplicationResult(discountAmount Money, paymentAmount Money) CouponApplicationResult {
	panic("未実装")
}

// NewFailedCouponApplicationResult は失敗した適用結果を生成する。
// errorMessage はエラーメッセージ、paymentAmount は支払金額を表す。
// 戻り値は失敗した適用結果であり、割引額は 0 円となる。
func NewFailedCouponApplicationResult(errorMessage string, paymentAmount Money) CouponApplicationResult {
	panic("未実装")
}

// Successful は成否を返す。
// 戻り値はクーポンの適用に成功した場合に true である。
func (r CouponApplicationResult) Successful() bool {
	panic("未実装")
}

// ErrorMessage はエラーメッセージを返す。
// 戻り値はエラーメッセージであり、成功時は空文字列である。
func (r CouponApplicationResult) ErrorMessage() string {
	panic("未実装")
}

// DiscountAmount は割引額を返す。
// 戻り値は割引額である。
func (r CouponApplicationResult) DiscountAmount() Money {
	panic("未実装")
}

// PaymentAmount は支払金額を返す。
// 戻り値は支払金額である。
func (r CouponApplicationResult) PaymentAmount() Money {
	panic("未実装")
}
