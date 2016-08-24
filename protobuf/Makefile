URL="https://raw.githubusercontent.com/google/protobuf/master/src/google/protobuf/"

update:
	(cd ./google/protobuf && rm descriptor.proto; wget ${URL}/descriptor.proto)
	(cd ./google/protobuf/compiler && rm plugin.proto; wget ${URL}/compiler/plugin.proto)
	(cd ./google/protobuf && rm any.proto; wget ${URL}/any.proto)
	(cd ./google/protobuf && rm timestamp.proto; wget ${URL}/timestamp.proto)
	(cd ./google/protobuf && rm duration.proto; wget ${URL}/duration.proto)
	(cd ./google/protobuf && rm struct.proto; wget ${URL}/struct.proto)
	(cd ./google/protobuf && rm wrappers.proto; wget ${URL}/wrappers.proto)
