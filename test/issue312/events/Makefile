regenerate:
	go install github.com/buptbill220/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/buptbill220/protobuf/protobuf/ \
	--gogo_out=. events.proto 
