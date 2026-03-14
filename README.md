# routine-tracker

gRPC-based

## Structure

```
.
├── proto/                  # Protobuf source definitions
│   ├── habit.proto
│   └── service.proto
├── api/                    # Generated Go code (do not edit)
│   ├── habit.pb.go
│   └── service.pb.go
├── internal/
│   ├── habit/              # Domain types and business logic
│   ├── repository/         # In-memory storage
│   └── server/             # gRPC handler wiring
├── cmd/
│   └── tracker/            # Binary entrypoint
│       └── main.go
├── go.mod
└── go.sum
```

## Regenerate proto

```bash
protoc \
  --proto_path=proto \
  --go_out=api \
  --go_opt=paths=source_relative \
  --go-grpc_out=api \
  --go-grpc_opt=paths=source_relative \
  habit.proto service.proto
```

