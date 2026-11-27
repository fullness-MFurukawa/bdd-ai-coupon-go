package service

import (
	"time"

	"github.com/fullness-MFurukawa/bdd-ai-coupon-go/internal/domain/model"
)

// DiscountCalculationService は割引計算を表す。ポインタ *DiscountCalculationService で扱う。
type DiscountCalculationService struct{}

// NewDiscountCalculationService は割引計算を生成する。
// 戻り値は生成した割引計算である。
func NewDiscountCalculationService() *DiscountCalculationService {
	panic("未実装")
}

// Apply はクーポンを適用する。
// 現在日時を引数で受け取るのは、背景で指定された日時で有効期限を判定するためである。
// coupon はクーポン、cart は買い物かご、now は現在日時を表す。
// 戻り値は適用結果である。
func (s *DiscountCalculationService) Apply(
	coupon *model.Coupon,
	cart *model.Cart,
	now time.Time,
) model.CouponApplicationResult {
	panic("未実装")
}

// CanApply はクーポンが使えるかを判定する。
// coupon はクーポン、cart は買い物かご、now は現在日時を表す。
// 戻り値はクーポンが使える場合に true である。
func (s *DiscountCalculationService) CanApply(coupon *model.Coupon, cart *model.Cart, now time.Time) bool {
	panic("未実装")
}
