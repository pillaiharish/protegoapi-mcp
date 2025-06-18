package main

import (
	"net/http"
	"context"
	"encoding/json"
	"log"
	"github.com/pillaiharish/protegoapi-mcp/internal/dispatcher"
	"github.com/pillaiharish/protegoapi-mcp/internal/store"
)

type echoTool struct{}

func (e echoTool) Name() string {
	return "echo"
}

func (e echoTool) Call(_ context.Context, p json.RawMessage, _ store.Store) (any, error) {
	var v any
	if err := json.Unmarshal(p, &v); err!=nil{
		return nil, err
	}
	return v, nil
}



func main() {
	_ = dispatcher.RegisterTool(echoTool{})
	
	http.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {
		var pkt dispatcher.Packet
		if err := json.NewDecoder(r.Body).Decode(&pkt); err!=nil{
			http.Error(w, "bad json", 400); 
			return
		}
		res, err := dispatcher.Dispatch(r.Context(), pkt, nil)
		resp := map[string]any{"id":pkt.ID, "result": res, "error": err}
		_ = json.NewEncoder(w).Encode(resp)
	})
	log.Println("-> listening 8991")
	log.Fatal(http.ListenAndServe(":8991", nil))
}
