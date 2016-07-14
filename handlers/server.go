package handlers

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/redhatanalytics/oshinko-rest/restapi/operations/server"
	"github.com/redhatanalytics/oshinko-rest/version"
	"github.com/redhatanalytics/oshinko-rest/helpers/info"
)

// ServerResponse respond to the server info request
func ServerResponse(params server.GetServerInfoParams) middleware.Responder {
	vers := version.GetVersion()
	name := version.GetAppName()
	webname := info.GetWebServiceName()
	weburl := info.GetWebServiceURL()
	payload := server.GetServerInfoOKBodyBody{
		Application: &server.GetServerInfoOKBodyApplication{
			Name: &name, Version: &vers,
		        WebServiceName: &webname, WebURL: &weburl}}
	return server.NewGetServerInfoOK().WithPayload(payload)
}
