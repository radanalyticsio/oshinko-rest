# package the oshinko-rest app into a container
# by default this image will run the server listening on port 8080
FROM golang

ADD . /go/src/github.com/redhatanalytics/oshinko-rest

RUN go get github.com/tools/godep

WORKDIR /go/src/github.com/redhatanalytics/oshinko-rest
RUN make install

ENTRYPOINT /go/bin/oshinko-rest-server --host 0.0.0.0 --port 8080

EXPOSE 8080
