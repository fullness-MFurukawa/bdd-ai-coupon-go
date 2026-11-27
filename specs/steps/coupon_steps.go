package steps

import (
	"github.com/cucumber/godog"
)

// CouponSteps はクーポンによる割引の適用のステップ定義を表す。
type CouponSteps struct{}

// InitializeScenario は、ステップ定義を godog に登録する。
// ctx はシナリオごとに godog から渡される登録先を表す。
func InitializeScenario(ctx *godog.ScenarioContext) {
	s := &CouponSteps{}

	ctx.Given(`^クーポン "([^"]*)" が存在する$`, s.CouponExists)
	ctx.Given(`^クーポン "([^"]*)" の最低購入金額は (\d+) 円である$`, s.CouponHasMinimumPurchaseAmount)
	ctx.Given(`^クーポン "([^"]*)" の割引上限額は (\d+) 円である$`, s.CouponHasMaximumDiscountAmount)
	ctx.Given(`^クーポン "([^"]*)" の有効期限は "([^"]*)" である$`, s.CouponExpiresOn)
	ctx.Given(`^割引額の端数は切り捨てる$`, s.DiscountFractionIsTruncated)
	ctx.Given(`^ポイントは (\d+) ポイントを (\d+) 円として計算する$`, s.PointRateIs)
	ctx.Given(`^現在日時は "([^"]*)" である$`, s.CurrentDateTimeIs)
	ctx.Given(`^ポイントを利用しない$`, s.PointsAreNotUsed)
	ctx.Given(`^会員ランクは "([^"]*)" である$`, s.MembershipRankIs)
	ctx.Given(`^カートの購入金額は (\d+) 円である$`, s.CartTotalAmountIs)

	ctx.When(`^クーポン "([^"]*)" を適用する$`, s.ApplyCoupon)

	ctx.Then(`^クーポンの適用は失敗する$`, s.CouponApplicationFails)
	ctx.Then(`^エラーメッセージ "([^"]*)" が表示される$`, s.ErrorMessageIsShown)
	ctx.Then(`^割引額は (\d+) 円である$`, s.DiscountAmountIs)
	ctx.Then(`^支払金額は (\d+) 円である$`, s.PaymentAmountIs)
}

// CouponExists は、クーポンが存在する状況を準備する。
// couponCode は券種コードを表す。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) CouponExists(couponCode string) error {
	return godog.ErrPending
}

// CouponHasMinimumPurchaseAmount は、クーポンの最低購入金額を準備する。
// couponCode は券種コード、minimumPurchaseAmount は最低購入金額（円）を表す。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) CouponHasMinimumPurchaseAmount(couponCode string, minimumPurchaseAmount int) error {
	return godog.ErrPending
}

// CouponHasMaximumDiscountAmount は、クーポンの割引上限額を準備する。
// couponCode は券種コード、maximumDiscountAmount は割引上限額（円）を表す。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) CouponHasMaximumDiscountAmount(couponCode string, maximumDiscountAmount int) error {
	return godog.ErrPending
}

// CouponExpiresOn は、クーポンの有効期限を準備する。
// couponCode は券種コード、expirationDate は有効期限の日付の文字列を表す。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) CouponExpiresOn(couponCode string, expirationDate string) error {
	return godog.ErrPending
}

// DiscountFractionIsTruncated は、割引額の端数を切り捨てる状況を準備する。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) DiscountFractionIsTruncated() error {
	return godog.ErrPending
}

// PointRateIs は、ポイントの換算率を準備する。
// points はポイント数、yen はそのポイント数に相当する金額（円）を表す。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) PointRateIs(points int, yen int) error {
	return godog.ErrPending
}

// CurrentDateTimeIs は、現在日時を準備する。
// now は現在日時の文字列を表す。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) CurrentDateTimeIs(now string) error {
	return godog.ErrPending
}

// PointsAreNotUsed は、ポイントを利用しない状況を準備する。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) PointsAreNotUsed() error {
	return godog.ErrPending
}

// MembershipRankIs は、会員ランクを準備する。
// membershipRank は会員ランクの名称を表す。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) MembershipRankIs(membershipRank string) error {
	return godog.ErrPending
}

// CartTotalAmountIs は、買い物かごの購入金額を準備する。
// totalAmount は購入金額（円）を表す。
// 戻り値は準備に失敗した場合のエラーである。
func (s *CouponSteps) CartTotalAmountIs(totalAmount int) error {
	return godog.ErrPending
}

// ApplyCoupon は、クーポンを適用する。
// couponCode は適用するクーポンの券種コードを表す。
// 戻り値は操作に失敗した場合のエラーである。
func (s *CouponSteps) ApplyCoupon(couponCode string) error {
	return godog.ErrPending
}

// CouponApplicationFails は、クーポンの適用が失敗したことを検証する。
// 戻り値は検証に失敗した場合の、期待した値と実際の値を含むエラーである。
func (s *CouponSteps) CouponApplicationFails() error {
	return godog.ErrPending
}

// ErrorMessageIsShown は、エラーメッセージが表示されることを検証する。
// errorMessage は期待するエラーメッセージを表す。
// 戻り値は検証に失敗した場合の、期待した値と実際の値を含むエラーである。
func (s *CouponSteps) ErrorMessageIsShown(errorMessage string) error {
	return godog.ErrPending
}

// DiscountAmountIs は、割引額を検証する。
// discountAmount は期待する割引額（円）を表す。
// 戻り値は検証に失敗した場合の、期待した値と実際の値を含むエラーである。
func (s *CouponSteps) DiscountAmountIs(discountAmount int) error {
	return godog.ErrPending
}

// PaymentAmountIs は、支払金額を検証する。
// paymentAmount は期待する支払金額（円）を表す。
// 戻り値は検証に失敗した場合の、期待した値と実際の値を含むエラーである。
func (s *CouponSteps) PaymentAmountIs(paymentAmount int) error {
	return godog.ErrPending
}
