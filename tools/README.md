# tools for use with oshinko-rest

## Sample script to deploy oshinko

The `oshinko-deploy.sh` script can deploy the oshinko suite into an existing
OpenShift deployment or it can start an all-in-one docker OpenShift on the
host. It will pull the latest upstream images from the radanalyticsio
organization. It can also be configured to use alternate images, for more
information see the script help text.

**Example all-in-one deployment**

    $ ./oshinko-deploy.sh -d

This will start an OpenShift all-in-one cluster with the `oc cluster up`
command on the host, then it will deploy the oshinko suite into the
`myproject` project as user `developer`. It will apply the default route
url specified by OpenShift.

**Example deployment on remote cluster**

    $ ./oshinko-deploy.sh -c https://10.0.1.100:8443 \
                          -u bob \
                          -p bobsproject \
                          -o bobsoshinko.10.0.1.100.xip.io

This will deploy oshinko into the OpenShift cluster on the 10.0.1.100 host,
in the `bobsproject` project as user `bob`. It will apply the route url
`bobsoshinko.10.0.1.100.xip.io` to the oshinko web console.

### A note on permissions

The all-in-one deployment requires that the user running the script has
permission to issue docker commands. If docker is not configured to
allow non-root access, you will need to invoke this script using `sudo`
or as the `root` user for an all-in-one deployment.