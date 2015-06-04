regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogofast
	protoc --gogofast_out=./fast/ vanity.proto
	go install github.com/gogo/protobuf/protoc-gen-gogofaster
	protoc --gogofaster_out=./faster/ vanity.proto
	go install github.com/gogo/protobuf/protoc-gen-gogoslick
	protoc --gogoslick_out=./slick/ vanity.proto
	go install github.com/gogo/protobuf/protoc-gen-gofast
	protoc --gofast_out=./gofast/ vanity.proto
