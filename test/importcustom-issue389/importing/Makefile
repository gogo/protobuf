regenerate:
	go install github.com/buptbill220/protobuf/protoc-gen-gogo
	go install github.com/buptbill220/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --proto_path=../../../../../../:../../../protobuf/:. --gogo_out=. c.proto
