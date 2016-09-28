# oshinko-rest

REST based API server for an Apache Spark on OpenShift cluster management
application.

## Building oshinko-rest

To build the project simply run the `build` or `install` target in the
makefile.

**Example**

    $ make build

Assuming a successful build, the output will be stored in the `_output`
directory. For an `install` target, the binary will be placed in your
`$GOPATH/bin`.

## Running oshinko-rest

For most functionality an OpenShift cluster will be needed, but the
application can be tested for basic operation without one.

**Example**

After building the binary a basic test can be performed as follows:

* start the server in a terminal

```
    $ _output/oshinko-rest-server --port 42000 --scheme http
    2016/07/14 16:41:00 Serving oshinko rest at http://127.0.0.1:42000
```

* in a second terminal run a small curl command against the server

```
    $ curl http://localhost:42000/
    {"application":{"name":"oshinko-rest-server","version":"0.1.0"}}
```

*The return value may be different depending on the version of the
server you have built*

### TLS

To start the server with TLS enabled, you will need to supply a certificate
file and a key file for the server to use. Once these files are created you
can start the server as follows to enable HTTPS access:

```
    $ _output/oshinko-rest-server --port 42000 --tls-port 42443 --tls-key keyfile.key --tls-certificate certificatefile.cert
    2016/09/28 12:10:47 Serving oshinko rest at http://127.0.0.1:42000
    2016/09/28 12:10:47 Serving oshinko rest at https://127.0.0.1:42443
```

At this point the server is ready to accept both HTTP and HTTPS requests. If
would like to resitrct access to **only** use TLS, add the `--scheme https`
flag to the command line as follows:

```
    $ _output/oshinko-rest-server --scheme https --tls-port 42443 --tls-key keyfile.key --tls-certificate certificatefile.cert
    2016/09/28 12:10:47 Serving oshinko rest at https://127.0.0.1:42443
```

## Further reading

Please see the HACKING and CONTRIBUTING docs for more information about
working with this codebase.
