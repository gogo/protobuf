build: clean
	(cd gogodocker-gen && make)
	(cd generated-dockerfiles && make)

clean:
	go clean .
	rm -rf generated-dockerfiles

run: build
	(cd generated-dockerfiles && make run)
