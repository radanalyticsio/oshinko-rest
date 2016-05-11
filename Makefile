clean:
	rm -r _output

build:
	go build -o _output/oshinko-rest-server ./cmd/oshinko-rest-server

install:
	go install ./cmd/oshinko-rest-server
