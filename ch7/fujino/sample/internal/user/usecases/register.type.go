package userusecases

import (
	"context"

	userdomain "example-ch7_8/internal/user/domain"
	userworkflows "example-ch7_8/internal/user/workflows"
)

// ユーザー登録ユースケース
type RegisterUserUsecase func(
	IsEmailTaken,
) func(context.Context, RegisterUser) ([]userworkflows.RegisterUserEvent, error)

type IsEmailTaken func(userdomain.FormattedEmail) (bool, error)

type RegisterUser struct {
	// ユーザーの性
	FirstName string `json:"first_name"`
	// ユーザーの名
	LastName string `json:"last_name"`
	// ユーザーのメールアドレス
	Email string `json:"email"`
	// ユーザーのパスワード
	Password string `json:"password"`
	// ユーザーの郵便番号
	Zipcode string `json:"zipcode"`
	// ユーザーの都道府県
	Prefecture string `json:"prefecture"`
	// ユーザーの市区町村
	Municipalities string `json:"municipalities"`
	// ユーザーの住所
	Address string `json:"address"`
	// ユーザーの電話番号
	Telephone string `json:"telephone"`
}
