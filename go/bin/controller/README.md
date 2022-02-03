```
protoc --go_out=./grpc --go_opt=paths=source_relative \
    --go-proto_path=../proto/servers/Control/Control.proto \
    --go-grpc_out=./grpc --go-grpc_opt=paths=source_relative \
    ../proto/servers/Control/Control.proto
```