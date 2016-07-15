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
    $ _output/oshinko-rest-server --port 42000
    2016/07/14 16:41:00 Serving oshinko rest at http://127.0.0.1:42000
```

* in a second terminal run a small curl command against the server

```
    $ curl http://localhost:42000/
    {"application":{"name":"oshinko-rest-server","version":"0.1.0"}}
```

*The return value may be different depending on the version of the
server you have built*

## Further reading

Please see the HACKING and CONTRIBUTING docs for more information about
working with this codebase.
