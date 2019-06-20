regenerate:
	go install github.com/frankee/protobuf/protoc-min-version
	go install github.com/frankee/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/frankee/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/frankee/protobuf/types:. timefail.proto
