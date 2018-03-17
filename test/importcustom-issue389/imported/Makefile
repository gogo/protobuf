regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	(protoc --proto_path=../../../../../../:../../../protobuf/:. --gogo_out=. a.proto)