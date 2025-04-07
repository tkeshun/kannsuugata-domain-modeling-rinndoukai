package userusecases

import (
	"context"
	"fmt"

	userdomain "example-ch7_8/internal/user/domain"
	userworkflows "example-ch7_8/internal/user/workflows"
)

var NewRegisterUserUsecase RegisterUserUsecase = func(IsEmailTaken IsEmailTaken) func(context.Context, RegisterUser) ([]userworkflows.RegisterUserEvent, error) {
	flow := userworkflows.NewRegisterUserWorkflow(
		userworkflows.ValidateUserImpl,
		userworkflows.RegistUserImpl,
	)
	return func(ctx context.Context, ru RegisterUser) ([]userworkflows.RegisterUserEvent, error) {
		var ext userdomain.ExternalUserData

		cmd := userworkflows.RegisterUserCommand{
			Context: ctx,
			Data: userworkflows.RegisterUserCommandData{
				UnvalidatedUser:  ru.toUnvalidatedUser(),
				ExternalUserData: ext,
			},
		}
		events, result, err := flow(cmd)
		if err != nil {
			return nil, fmt.Errorf("ユーザー登録に失敗しました: %w", err)
		}

		// 必要な外部データリクエストがある場合は、リクエストを処理する
		for result.HasRequest() {
			for _, req := range result.ExternalDataRequests() {
				switch v := req.(type) {
				case userdomain.CheckIsEmailTakenRequest:
					{
						isTaken, err := IsEmailTaken(v.FormattedEmail)
						if err != nil {
							return nil, fmt.Errorf("メールの重複確認に失敗しました: %w", err)
						}
						ext = userdomain.ExternalUserData{
							ExternalEmailData: userdomain.ExternalEmailData{
								IsTaken: &isTaken,
							},
						}
					}
				}
			}
			// 再度、ユーザー登録コマンドを実行
			cmd = userworkflows.RegisterUserCommand{
				Context: ctx,
				Data: userworkflows.RegisterUserCommandData{
					UnvalidatedUser:  ru.toUnvalidatedUser(),
					ExternalUserData: ext,
				},
			}
			events, result, err = flow(cmd)
			if err != nil {
				return nil, fmt.Errorf("ユーザー登録に失敗しました: %w", err)
			}
		}
		return events, result.ValidationErrors()
	}
}

func (ru RegisterUser) toUnvalidatedUser() userdomain.UnvalidatedUser {
	return userdomain.UnvalidatedUser{
		FirstName:      ru.FirstName,
		LastName:       ru.LastName,
		Email:          ru.Email,
		Password:       ru.Password,
		Zipcode:        ru.Zipcode,
		Prefecture:     ru.Prefecture,
		Municipalities: ru.Municipalities,
		Address:        ru.Address,
		Telephone:      ru.Telephone,
	}
}
