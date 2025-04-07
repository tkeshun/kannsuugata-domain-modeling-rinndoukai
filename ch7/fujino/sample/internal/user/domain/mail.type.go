package userdomain

import shareerrs "example-ch7_8/internal/share/domain/errs"

type Email interface {
	UniqueEmail
	FormattedEmail
}

// 適切な形式のメールアドレス
type FormattedEmail interface {
	isFormattedEmail()
	Value() string
}

// 重複していないメールアドレス
type UniqueEmail interface {
	isUniqueEmail()
	Value() string
}

// メールアドレスを適切な形式であると保証する関数
type toFormattedEmail func(string) (FormattedEmail, error)

// メールアドレスの重複を検証する関数
type toUniqueEmail func(email FormattedEmail, isTaken bool) (UniqueEmail, error)

type ExternalEmailData struct {
	// 使用されているかどうか
	IsTaken *bool
}

type ToValidateEmail func(toFormattedEmail, toUniqueEmail, ExternalEmailData) func(string) (Email, shareerrs.DomainValidationResult)
