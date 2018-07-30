test:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install github.com/gogo/protobuf/protoc-min-version
	go get -u golang.org/x/net/context
	go get -u google.golang.org/grpc
	protoc-min-version --version="3.0.0" --gogo_out=plugins=grpc,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:. --proto_path=../../../../../:../../protobuf/:. issue438.proto
	go test ./...