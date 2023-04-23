protos:  proto/account.proto proto/list.proto proto/login.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/account.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/list.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/login.proto

clean:
	rm proto/account.pb.go	proto/account_grpc.pb.go
	rm proto/list.pb.go	proto/list_grpc.pb.go
	rm proto/login.pb.go	proto/login_grpc.pb.go

wasm: main.go
	GOOS=js GOARCH=wasm go build -o  assets/json.wasm

run: main.go
	go run .


