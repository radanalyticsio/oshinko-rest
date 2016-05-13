clean:
	rm -rf _output

build:
	tools/build.sh

install:
	godep go install ./cmd/oshinko-rest-server

deps:
	godep save ./...

validate-api:
	swagger validate api/swagger.yaml

image:
	docker build -t oshinko-rest-server .
