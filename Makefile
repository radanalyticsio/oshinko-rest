clean:
	rm -r _output

build:
	godep go build -o _output/oshinko-rest-server ./cmd/oshinko-rest-server

install:
	godep go install ./cmd/oshinko-rest-server

deps:
	godep save ./...

validate-api:
	swagger validate api/swagger.yaml
