package dispatcher_test

import (
	"context"
	"encoding/json"
	"testing"
	"github.com/pillaiharish/protegoapi-mcp/internal/dispatcher"
	"github.com/pillaiharish/protegoapi-mcp/internal/store"
)

type TestTool struct { name string }

func (t TestTool) Name() string { return t.name }

func (t TestTool) Call(_ context.Context, p json.RawMessage , _ store.Store) (any, error) {
	return "ok", nil
}

// Test for duplications during Registrations
func TestRegisterTool(t *testing.T) {
	err := dispatcher.RegisterTool(TestTool{"dup"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := dispatcher.RegisterTool(TestTool{"dup"}); err == nil {
		t.Fatalf("expected duplicate registration error: %v", err)
	}
}

// Test for known dispatch and bogus dispatch
func TestDispatchTool(t *testing.T){
	p := dispatcher.Packet{ID:"1", Tool:"hi"}
	_ = dispatcher.RegisterTool(TestTool{"hi"})
	
	res, err := dispatcher.Dispatch(context.TODO(), p, store.NullStore{})
	if err != nil || res != "ok" {
		t.Fatalf("Dispatch failed: %v and response %v", err, res)
	}

	p.Tool = "bogus"
	res, err = dispatcher.Dispatch(context.TODO(), p, store.NullStore{})
	if err==nil {
		t.Fatalf("Expected unknown tool error")
	}
}

