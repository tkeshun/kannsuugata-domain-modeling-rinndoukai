package userusecases

import (
	"context"

	identitytypes "example-ch7_8/internal/identity/domain/types"
	userdomain "example-ch7_8/internal/user/domain"
	userworkflows "example-ch7_8/internal/user/workflows"
)

// ユーザー情報取得ユースケース
type GetUserInfoUsecase func(LoadhUserAggregate, userworkflows.GetUserInfoWorkflow) func(context.Context,
	struct{ ID string }) (*userworkflows.UserInfo, error)

// ユーザー集約から取得
type LoadhUserAggregate func(
	userID identitytypes.IdentityID,
) (userdomain.User, error)
