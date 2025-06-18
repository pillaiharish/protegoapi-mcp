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

# Test with Token
``` bash
harish $ TOKEN=$(go run ./cmd/jwtgen)

harish $ curl -v   -H "Authorization: Bearer $TOKEN"   -H "Content-Type: application/json"   -X POST "http://localhost:8991/mcp"   -d '{"id":1,"tool":"echo","params":{"msg":"namaste"}}'
Note: Unnecessary use of -X or --request, POST is already inferred.
* Host localhost:8991 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8991...
* Connected to localhost (::1) port 8991
> POST /mcp HTTP/1.1
> Host: localhost:8991
> User-Agent: curl/8.7.1
> Accept: */*
> Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
> Content-Type: application/json
> Content-Length: 49
>
* upload completely sent off: 49 bytes
< HTTP/1.1 200 OK
< Date: Wed, 18 Jun 2025 19:32:33 GMT
< Content-Length: 49
< Content-Type: text/plain; charset=utf-8
<
{"error":null,"id":1,"result":{"msg":"namaste"}}
* Connection #0 to host localhost left intact
```

# Test without Token
``` bash
harish $ curl -X POST localhost:8991/mcp -d '{"id":1,"tool":"echo","params":{"msg":"namaste"}}'
missing token

```

# Test with Invalid Token
``` bash
harish $ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAyNzQ1MzAsInJvbGVzIjpbImFkbWluIl0sInN1YiI6ImhhcmlzaCJ9.0ggk9J2Ck5Fe8CAH_iyLmWPGP3yzAqqx7yiqn3T2EGs" -X POST http://localhost:8991/mcp -d '{"id":1,"tool":"echo","params":{"msg":"namaste"}}'
invalid token 
```

