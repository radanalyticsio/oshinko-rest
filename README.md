# oshinko application

The oshinko application manages Apache Spark clusters on OpenShift.
The application consists of a REST server (oshinko-rest) and a web UI
and is designed to run in an OpenShift project.

This repository contains tools to launch the oshinko application
along with the source code for the oshinko REST server. The source
code for the web UI is located in a different repository.

# Launching the oshinko application with public images

This section describes how to simply launch the oshinko components
in an OpenShift project using public images and defaults.

## Create an oshinko service account

The oshinko-rest image uses a service account to perform OpenShift
operations. This service account must be created and given the admin role in
*each* project which will use oshinko, for example:

Create a new project *myproject*:

    $ oc new-project myproject

Create the oshinko service account:

    $ oc create sa oshinko -n myproject

Assign the admin role:

    $ oc policy add-role-to-user admin system:serviceaccount:myproject:oshinko -n myproject

## Process the server-ui-template.yaml

This template deploys the oshinko application in an OpenShift project.
By default it uses several images available from radanalyticsio on the
docker hub (https://hub.docker.com/u/radanalyticsio/)

To deploy oshinko from the command line:
 
    $ oc process -f tools/server-ui-template.yaml | oc create -f -

To load the template for use from the console with "Add to project":

    $ oc create -f tools/server-ui-template.yaml

# Building and running oshinko-rest

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
you would like to restrict access to **only** use TLS, add the
`--scheme https` flag to the command line as follows:

```
    $ _output/oshinko-rest-server --scheme https --tls-port 42443 --tls-key keyfile.key --tls-certificate certificatefile.cert
    2016/09/28 12:10:47 Serving oshinko rest at https://127.0.0.1:42443
```

# Further reading

Please see the CONTRIBUTING and HACKING docs for more information about
working with this codebase and the docs directory for more general information on usage.
