regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install github.com/gogo/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=plugins=grpc:. --proto_path=../../../../../:../../protobuf/:. issue427.proto