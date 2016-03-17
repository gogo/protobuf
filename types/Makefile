regenerate:
	protoc --gogo_out=. any/*.proto
	protoc --gogo_out=. duration/*.proto
	protoc --gogo_out=. empty/*.proto
	protoc --gogo_out=. struct/*.proto
	protoc --gogo_out=. timestamp/*.proto
	protoc --gogo_out=. wrappers/*.proto

fetch:
	./regen.sh