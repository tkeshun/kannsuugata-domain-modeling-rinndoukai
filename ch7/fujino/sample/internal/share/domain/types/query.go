package sharetypes

import "context"

type Query[T any] struct {
	Context context.Context
	Data    T
}
