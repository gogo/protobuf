all:
	protoc --gogofast_out=paths=source_relative:. ./huge.proto

clean:
	rm -rf huge.pb.go gen1.go gen2.go

.PHONY: clean
