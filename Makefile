.PHONY : build clean install test validate-api image generate

build:
	tools/build.sh build

clean:
	rm -rf _output

install:
	tools/build.sh install

test:
	tools/build.sh test

validate-api:
	swagger validate api/swagger.yaml

image:
	docker build -t oshinko-rest-server .

generate:
	swagger generate server -f api/swagger.yaml -A oshinko-rest
	swagger generate client -f api/swagger.yaml -A oshinko-rest
