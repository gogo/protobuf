regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogoslick
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:. timefail.proto 
