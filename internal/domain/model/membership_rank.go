package model

// MembershipRank は会員ランクを表す。
// 取りうる値は General・Gold・Platinum の3つに限られる。
type MembershipRank struct {
	displayName string // 会員ランクの名称
}

// General は会員ランクの一般を表す。
var General = MembershipRank{displayName: "一般"}

// Gold は会員ランクのゴールドを表す。
var Gold = MembershipRank{displayName: "ゴールド"}

// Platinum は会員ランクのプラチナを表す。
var Platinum = MembershipRank{displayName: "プラチナ"}

// MembershipRankOf は会員ランクの名称から会員ランクを得る。
// displayName は会員ランクの名称（一般・ゴールド・プラチナ）を表す。
// 戻り値は該当する会員ランクと、該当する会員ランクが無い場合の *DomainError である。
func MembershipRankOf(displayName string) (MembershipRank, error) {
	panic("未実装")
}

// DisplayName は会員ランクの名称を返す。
// 戻り値は会員ランクの名称（一般・ゴールド・プラチナ）である。
func (r MembershipRank) DisplayName() string {
	panic("未実装")
}
