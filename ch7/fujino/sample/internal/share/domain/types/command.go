package sharetypes

import (
	"context"
	"time"

	identitytypes "example-ch7_8/internal/identity/domain/types"
)

type Command[T any] struct {
	Context    context.Context
	Data       T
	Timestamp  time.Time
	IdentityID identitytypes.IdentityID
}
