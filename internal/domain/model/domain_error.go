package model

// DomainError は業務ルール違反を表す。
// 値オブジェクトやエンティティの制約に反する値を渡された場合に、*DomainError を error として返す。
type DomainError struct {
	message string // エラーのメッセージ（開発側に向けたもの）
	cause   error  // 原因となるエラー
}

// NewDomainError は業務ルール違反を生成する。
// message はエラーのメッセージを表す。
// 戻り値は生成した業務ルール違反である。
func NewDomainError(message string) *DomainError {
	panic("未実装")
}

// WrapDomainError は原因となるエラーを包んだ業務ルール違反を生成する。
// message はエラーのメッセージ、cause は原因となるエラーを表す。
// 戻り値は生成した業務ルール違反である。
func WrapDomainError(message string, cause error) *DomainError {
	panic("未実装")
}

// Error はエラーのメッセージを返す。error インタフェースを実装する。
// 戻り値はエラーのメッセージである。
func (e *DomainError) Error() string {
	panic("未実装")
}

// Unwrap は原因となるエラーを返す。
// 戻り値は原因となるエラーであり、包んでいない場合は nil である。
func (e *DomainError) Unwrap() error {
	panic("未実装")
}
