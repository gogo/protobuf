regenerate:
	go install github.com/buptbill220/protobuf/protoc-gen-gogo
	go install github.com/buptbill220/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mgoogle/protobuf/any.proto=github.com/buptbill220/protobuf/types,\
	Mgoogle/protobuf/api.proto=github.com/buptbill220/protobuf/types,\
	Mgoogle/protobuf/duration.proto=github.com/buptbill220/protobuf/types,\
	Mgoogle/protobuf/empty.proto=github.com/buptbill220/protobuf/types,\
	Mgoogle/protobuf/field_mask.proto=github.com/buptbill220/protobuf/types,\
	Mgoogle/protobuf/source_context.proto=github.com/buptbill220/protobuf/types,\
	Mgoogle/protobuf/struct.proto=github.com/buptbill220/protobuf/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/buptbill220/protobuf/types,\
	Mgoogle/protobuf/type.proto=github.com/buptbill220/protobuf/types,\
	Mgoogle/protobuf/wrappers.proto=github.com/buptbill220/protobuf/types,\
	:. --proto_path=../../../../../:../../protobuf/:. issue438.proto
