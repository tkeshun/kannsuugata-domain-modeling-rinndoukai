package sharetypes

import "time"

type AuditInfo struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 新しいAuditInfoを作成するヘルパー関数
func NewAuditInfo() AuditInfo {
	now := time.Now()
	return AuditInfo{
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// 更新時にUpdateAtを更新するメソッド
func (a AuditInfo) Update() AuditInfo {
	a.UpdatedAt = time.Now()
	return a
}
