package userdomain

import (
	"errors"

	shareerrs "example-ch7_8/internal/share/domain/errs"
)

type userEmail struct {
	string
}

func (e userEmail) isFormattedEmail() {
}
func (e userEmail) isUniqueEmail() {}

func (e userEmail) Value() string {
	return e.string
}
func ReconstructEmail(e string) Email {
	return userEmail{e}
}

var toValidateEmailImpl ToValidateEmail = func(
	tfef toFormattedEmail,
	tuef toUniqueEmail,
	ext ExternalEmailData,
) func(string) (Email, shareerrs.DomainValidationResult) {
	return func(email string) (Email, shareerrs.DomainValidationResult) {
		var requests []shareerrs.ExternalDataRequest
		var errs shareerrs.ValidationErrors
		formattedEmail, err := tfef(email)
		if err != nil {
			errs = errs.Add("email", err)
			return nil, shareerrs.NewDomainValidationResult(
				nil,
				errs,
			)
		}
		if ext.IsTaken == nil {
			requests = append(requests, CheckIsEmailTakenRequest{formattedEmail})
			return nil, shareerrs.NewDomainValidationResult(
				requests,
				errs,
			)
		}

		uniqueEmail, err := tuef(formattedEmail, *ext.IsTaken)
		if err != nil {
			return nil, shareerrs.NewDomainValidationResult(
				nil,
				errs.Add("email", err),
			)
		}

		return userEmail{uniqueEmail.Value()}, nil
	}
}

type CheckIsEmailTakenRequest struct {
	FormattedEmail
}

func (CheckIsEmailTakenRequest) Key() string {
	return "email"
}

func (CheckIsEmailTakenRequest) Description() string {
	return "メールアドレスの重複チェック"
}

type formattedEmail struct {
	string
}

func (e formattedEmail) Value() string {
	return e.string
}

func (e formattedEmail) isFormattedEmail() {
}

var toFormattedEmailImpl toFormattedEmail = func(email string) (FormattedEmail, error) {
	if email == "" {
		return nil, errors.New("メールアドレスは必須です")
	}
	return formattedEmail{email}, nil
}

type uniqueEmail struct {
	string
}

func (e uniqueEmail) Value() string {
	return e.string
}

func (e uniqueEmail) isUniqueEmail() {
}

var toUniqueEmailImpl toUniqueEmail = func(email FormattedEmail, isTaken bool) (UniqueEmail, error) {
	if isTaken {
		return nil, errors.New("メールアドレスは既に使用されています")
	}
	return uniqueEmail{email.Value()}, nil
}
