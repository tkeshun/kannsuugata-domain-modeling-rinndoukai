package userpersistence

import (
	"context"
	"time"

	identitytypes "example-ch7_8/internal/identity/domain/types"
	sharetypes "example-ch7_8/internal/share/domain/types"
	userdomain "example-ch7_8/internal/user/domain"
	userusecases "example-ch7_8/internal/user/usecases"
)

func NewLoadUserAggregate(
	ctx context.Context,
) userusecases.LoadhUserAggregate {
	return func(userID identitytypes.IdentityID) (userdomain.User, error) {

		// TODO(char5742): ユーザーを取得する処理を実装
		// if err != nil {
		// 	return userdomain.User{}, fmt.Errorf("ユーザーの取得に失敗: %w", err)
		// }
		// if user == nil {
		// 	return userdomain.User{}, fmt.Errorf("ユーザーが見つかりません")
		// }
		return userdomain.User{
			ID:             userID,
			FirstName:      userdomain.FirstName("太郎"),
			LastName:       userdomain.LastName("山田"),
			Email:          userdomain.ReconstructEmail("taro.yamada@example.com"),
			Zipcode:        userdomain.Zipcode("123-4567"),
			Prefecture:     userdomain.Prefecture("東京都"),
			Municipalities: userdomain.Municipalities("千代田区"),
			Address:        userdomain.Address("丸の内1-1-1"),
			Telephone:      userdomain.Telephone("03-1234-5678"),
			AuditInfo: sharetypes.AuditInfo{
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		}, nil
	}
}

func NewIsEmailTaken(
	ctx context.Context,
) userusecases.IsEmailTaken {
	return func(email userdomain.FormattedEmail) (bool, error) {
		// TODO(char5742): メールアドレスの重複チェックを実装
		return false, nil
	}
}
