package userworkflows

import (
	shareerrs "example-ch7_8/internal/share/domain/errs"
	sharetypes "example-ch7_8/internal/share/domain/types"
	userdomain "example-ch7_8/internal/user/domain"
)

// ユーザー登録コマンド
type RegisterUserCommand sharetypes.Command[RegisterUserCommandData]
type RegisterUserCommandData struct {
	userdomain.UnvalidatedUser
	userdomain.ExternalUserData
}

// ユーザー登録ワークフロー
type RegisterUserWorkflow func(RegisterUserCommand) ([]RegisterUserEvent, shareerrs.DomainValidationResult, error)

// ユーザー登録
type RegisterUser func(ValidateUser, RegistUser) RegisterUserWorkflow

// ユーザーのローカル検証ステップ
type ValidateUser func(userdomain.UnvalidatedUser, userdomain.ExternalUserData) (*userdomain.ValidatedUser, shareerrs.DomainValidationResult)

// ユーザーの登録ステップ
type RegistUser func(userdomain.ValidatedUser) (*userdomain.RegisteredUser, error)

type RegisterUserEvent interface {
	registerUserEvent()
	IsEvent()
}

// ユーザー登録イベント
type UserRegistered struct {
	RegisteredUser userdomain.RegisteredUser
}

func (UserRegistered) registerUserEvent() {}
func (UserRegistered) IsEvent()           {}
