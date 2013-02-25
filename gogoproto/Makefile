regenerate:
	echo WARNING! THIS RULE IS PROBABLY NOT RIGHT FOR YOUR INSTALLATION
	protoc --go_out=. --proto_path=../../../../:. *.proto

restore:
	cp gogo.pb.golden gogo.pb.go

preserve:
	cp gogo.pb.go gogo.pb.golden