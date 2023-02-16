
1. protoc-gen-go 설치
go get google.golang.org/protobuf/cmd/protoc-gen-go
export PATH="$PATH:$(go env GOPATH)/bin"


2. protoc-gen-gogrpc 설치 
grpc plugin을 사용해야 server와 같

go get google.golang.org/grpc/cmd/protoc-gen-go-grpc


protoc chat/chat.proto --go-grpc_out=chat --go_out=chat