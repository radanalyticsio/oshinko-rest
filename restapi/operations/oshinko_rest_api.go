package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/redhatanalytics/oshinko-rest/restapi/operations/clusters"
	"github.com/redhatanalytics/oshinko-rest/restapi/operations/server"
)

// NewOshinkoRestAPI creates a new OshinkoRest instance
func NewOshinkoRestAPI(spec *loads.Document) *OshinkoRestAPI {
	o := &OshinkoRestAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*OshinkoRestAPI The REST API server for the Oshinko suite of applications */
type OshinkoRestAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ClustersCreateClusterHandler sets the operation handler for the create cluster operation
	ClustersCreateClusterHandler clusters.CreateClusterHandler
	// ClustersDeleteSingleClusterHandler sets the operation handler for the delete single cluster operation
	ClustersDeleteSingleClusterHandler clusters.DeleteSingleClusterHandler
	// ClustersFindClustersHandler sets the operation handler for the find clusters operation
	ClustersFindClustersHandler clusters.FindClustersHandler
	// ClustersFindSingleClusterHandler sets the operation handler for the find single cluster operation
	ClustersFindSingleClusterHandler clusters.FindSingleClusterHandler
	// ServerGetServerInfoHandler sets the operation handler for the get server info operation
	ServerGetServerInfoHandler server.GetServerInfoHandler
	// ClustersUpdateSingleClusterHandler sets the operation handler for the update single cluster operation
	ClustersUpdateSingleClusterHandler clusters.UpdateSingleClusterHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *OshinkoRestAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *OshinkoRestAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *OshinkoRestAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *OshinkoRestAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *OshinkoRestAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *OshinkoRestAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the OshinkoRestAPI
func (o *OshinkoRestAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ClustersCreateClusterHandler == nil {
		unregistered = append(unregistered, "clusters.CreateClusterHandler")
	}

	if o.ClustersDeleteSingleClusterHandler == nil {
		unregistered = append(unregistered, "clusters.DeleteSingleClusterHandler")
	}

	if o.ClustersFindClustersHandler == nil {
		unregistered = append(unregistered, "clusters.FindClustersHandler")
	}

	if o.ClustersFindSingleClusterHandler == nil {
		unregistered = append(unregistered, "clusters.FindSingleClusterHandler")
	}

	if o.ServerGetServerInfoHandler == nil {
		unregistered = append(unregistered, "server.GetServerInfoHandler")
	}

	if o.ClustersUpdateSingleClusterHandler == nil {
		unregistered = append(unregistered, "clusters.UpdateSingleClusterHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *OshinkoRestAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *OshinkoRestAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *OshinkoRestAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *OshinkoRestAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *OshinkoRestAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *OshinkoRestAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters"] = clusters.NewCreateCluster(o.context, o.ClustersCreateClusterHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/clusters/{name}"] = clusters.NewDeleteSingleCluster(o.context, o.ClustersDeleteSingleClusterHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters"] = clusters.NewFindClusters(o.context, o.ClustersFindClustersHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{name}"] = clusters.NewFindSingleCluster(o.context, o.ClustersFindSingleClusterHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/"] = server.NewGetServerInfo(o.context, o.ServerGetServerInfoHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/clusters/{name}"] = clusters.NewUpdateSingleCluster(o.context, o.ClustersUpdateSingleClusterHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *OshinkoRestAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
