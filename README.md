# Protego MCP Protocol v0.1

## Envelope
(id, ts, auth, method, tool, params, ctx)

## Error codes
1000 InvalidJSON
1001 UnkknownTool
1002 Unauthorized
...

# Transport
- HTTP POST /mcp
- Websocket sub-protocol: mcpv1

# Run and Test
``` bash
harish $ go run ./cmd/server
harish $ curl -X POST localhost:8991/mcp -d '{"id":1,"tool":"echo","params":{"msg":"namaste"}}'
{"error":null,"id":1,"result":{"msg":"namaste"}}

``` 
