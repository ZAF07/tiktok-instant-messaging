start-http:
	go run http-server/main.go

start-rpc:
	go run rpc-server/main.go

gen-proto:
	protoc --go_out=./message-service/internal/core/domain/ --go_opt=paths=source_relative ./message-service/proto/message.proto