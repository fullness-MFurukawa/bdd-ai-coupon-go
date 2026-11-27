package model

import "time"

// ExpirationDate は有効期限を表す。当日まで有効とする。
type ExpirationDate struct {
	date time.Time // 有効期限の日付（時刻は 0 時、タイムゾーンは time.UTC）
}

// NewExpirationDate は有効期限を生成する。
// date は有効期限の日付を表す。
// 戻り値は生成した有効期限である。
func NewExpirationDate(date time.Time) ExpirationDate {
	panic("未実装")
}

// Date は有効期限の日付を返す。
// 戻り値は有効期限の日付（時刻は 0 時、タイムゾーンは time.UTC）である。
func (e ExpirationDate) Date() time.Time {
	panic("未実装")
}
