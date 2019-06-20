regenerate:
	go install github.com/frankee/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/frankee/protobuf/protobuf/ \
	--gogo_out=. events.proto 
