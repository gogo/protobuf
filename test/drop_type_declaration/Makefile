regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install github.com/gogo/protobuf/protoc-gen-gofast
	protoc --gofast_out=. --proto_path=../../../../../:../../protobuf/:. thetest.proto
