regenerate:
	go install github.com/gogo/protobuf/vanity/cmd/protoc-gen-gogofast
	protoc --gogofast_out=./fast/ vanity.proto
	go install github.com/gogo/protobuf/vanity/cmd/protoc-gen-gogofaster
	protoc --gogofaster_out=./faster/ vanity.proto
	go install github.com/gogo/protobuf/vanity/cmd/protoc-gen-gogoslick
	protoc --gogoslick_out=./slick/ vanity.proto
