package userworkflows

import (
	identitytypes "example-ch7_8/internal/identity/domain/types"
	sharetypes "example-ch7_8/internal/share/domain/types"
	userdomain "example-ch7_8/internal/user/domain"
)

// ユーザー表示情報取得
type GetUserInfo func(UserToInfo) GetUserInfoWorkflow

// ユーザー表示情報取得クエリ
type GetUserInfoQuery sharetypes.Query[userdomain.User]

// ユーザー情報取得処理ワークフロー
type GetUserInfoWorkflow func(
	GetUserInfoQuery,
) (*UserInfo, error)

// ユーザーエンティティをユーザー表示情報への変換ステップ
type UserToInfo func(user userdomain.User) UserInfo

// ユーザー表示情報
type UserInfo struct {
	ID             identitytypes.IdentityID  `json:"id"`
	FirstName      userdomain.FirstName      `json:"first_name"`
	LastName       userdomain.LastName       `json:"last_name"`
	Email          userdomain.Email          `json:"email"`
	Zipcode        userdomain.Zipcode        `json:"zipcode"`
	Prefecture     userdomain.Prefecture     `json:"prefecture"`
	Municipalities userdomain.Municipalities `json:"municipalities"`
	Address        userdomain.Address        `json:"address"`
	Telephone      userdomain.Telephone      `json:"telephone"`
	AuditInfo      sharetypes.AuditInfo      `json:"audit_info"`

	// パスワードは表示しない
	// Password       userdomain.Password       `json:"password"`
}
