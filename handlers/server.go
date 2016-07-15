package handlers

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/redhatanalytics/oshinko-rest/restapi/operations/server"
	"github.com/redhatanalytics/oshinko-rest/version"
	"github.com/redhatanalytics/oshinko-rest/helpers/info"
	osa "github.com/redhatanalytics/oshinko-rest/helpers/authentication"
)

// ServerResponse respond to the server info request
func ServerResponse(params server.GetServerInfoParams) middleware.Responder {
	vers := version.GetVersion()
	name := version.GetAppName()
	webname := info.GetWebServiceName()
	weburl := GetWebServiceURL()
	payload := server.GetServerInfoOKBodyBody{
		Application: &server.GetServerInfoOKBodyApplication{
			Name: &name, Version: &vers,
		        WebServiceName: &webname, WebURL: &weburl}}
	return server.NewGetServerInfoOK().WithPayload(payload)
}

// Look up routes for current namespace and find the one used by oshinko-web
// Will return empty string if no route can be found
func GetWebServiceURL() string {
	weburl := ""
	osclient, err := osa.GetOpenShiftClient()
	if err != nil {
		return ""
	}
	namespace, _ := info.GetNamespace()
	route, err := osclient.Routes(namespace).Get(info.GetWebServiceName())
	if err != nil || len(route.Status.Ingress) == 0 {
		return ""
	}
	weburl = route.Status.Ingress[0].Host
	return weburl
}