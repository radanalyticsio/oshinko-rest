clean:
	rm -rf _output

build:
	tools/build.sh build

install:
	tools/build.sh install

deps:
	godep save ./...

validate-api:
	swagger validate api/swagger.yaml

image:
	docker build -t oshinko-rest-server .
