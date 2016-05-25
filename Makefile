build:
	tools/build.sh build

clean:
	rm -rf _output

install:
	tools/build.sh install

test:
	tools/build.sh test

deps:
	export GO15VENDOREXPERIMENT=1 ; godep save ./...

validate-api:
	swagger validate api/swagger.yaml

image:
	docker build -t oshinko-rest-server .

generate-server:
	swagger generate server -f api/swagger.yaml -A oshinko-rest

generate-client:
	swagger generate client -f api/swagger.yaml -A oshinko-rest
