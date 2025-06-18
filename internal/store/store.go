package store

import(
	"context"
)

// We’ll flesh this out later. For now a no-op interface
// keeps the dispatcher generic without breaking compile.

type Store interface {
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, value any) error
}

// NullStore lets you pass “nil behaviour” until Redis is added.

type NullStore struct {}

func (NullStore) Get(context.Context, string)(any, error) {return nil,nil}
func (NullStore) Set(context.Context, string, any) error {return nil}
