package userusecases

import (
	"context"

	userdomain "example-ch7_8/internal/user/domain"
	userworkflows "example-ch7_8/internal/user/workflows"
)

var NewGetUserInfoUsecase GetUserInfoUsecase = func(lua LoadhUserAggregate, guiw userworkflows.GetUserInfoWorkflow) func(context.Context, struct{ ID string }) (*userworkflows.UserInfo, error) {
	return func(ctx context.Context, prm struct{ ID string }) (*userworkflows.UserInfo, error) {
		user, err := lua(prm.ID)
		if err != nil {
			return nil, err
		}
		query := userworkflows.GetUserInfoQuery{
			Context: ctx,
			Data:    user,
		}
		return guiw(query)
	}

}

func toUserInfo(user userdomain.User) userworkflows.UserInfo {
	return userworkflows.UserInfo{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Zipcode:        user.Zipcode,
		Prefecture:     user.Prefecture,
		Municipalities: user.Municipalities,
		Address:        user.Address,
		Telephone:      user.Telephone,
		AuditInfo:      user.AuditInfo,
	}
}
