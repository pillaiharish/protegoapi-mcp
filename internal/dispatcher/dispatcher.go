package dispatcher

import (
	"fmt"
	"encoding/json"
	"context"
	"github.com/pillaiharish/protegoapi-mcp/internal/store"
)

type Packet struct {
	ID interface{}			`json:"id"`
	Tool string				`json:"tool"`
	Params json.RawMessage	`json:"params"`
}

type Tool interface {
	Name() string
	Call(ctx context.Context, params json.RawMessage, s store.Store) (any, error)
}

var registry = map[string]Tool{}

func RegisterTool(t Tool)error{
	if _, exists := registry[t.Name()]; exists{
		return fmt.Errorf("Tool %s already registered", t.Name())
	}
	registry[t.Name()] = t
	return nil
}

func Dispatch(ctx context.Context, p Packet, s store.Store) (any, error) {
	t, ok := registry[p.Tool]
	if !ok {
		return nil, fmt.Errorf("Unknown tool %s", p.Tool)
	}

	return t.Call(ctx, p.Params, s)
}
