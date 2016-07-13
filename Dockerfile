# package the oshinko-rest app into a container
# by default this image will run the server listening on port 8080
FROM golang

ADD . /go/src/github.com/redhatanalytics/oshinko-rest

WORKDIR /go/src/github.com/redhatanalytics/oshinko-rest
RUN make install

ENTRYPOINT /go/src/github.com/redhatanalytics/oshinko-rest/tools/start_server.sh
