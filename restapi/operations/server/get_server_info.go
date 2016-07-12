package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GetServerInfoHandlerFunc turns a function with the right signature into a get server info handler
type GetServerInfoHandlerFunc func(GetServerInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetServerInfoHandlerFunc) Handle(params GetServerInfoParams) middleware.Responder {
	return fn(params)
}

// GetServerInfoHandler interface for that can handle valid get server info params
type GetServerInfoHandler interface {
	Handle(GetServerInfoParams) middleware.Responder
}

// NewGetServerInfo creates a new http.Handler for the get server info operation
func NewGetServerInfo(ctx *middleware.Context, handler GetServerInfoHandler) *GetServerInfo {
	return &GetServerInfo{Context: ctx, Handler: handler}
}

/*GetServerInfo swagger:route GET / server getServerInfo

Returns information about the server version

*/
type GetServerInfo struct {
	Context *middleware.Context
	Handler GetServerInfoHandler
}

func (o *GetServerInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetServerInfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

/*GetServerInfoOKBodyApplication get server info o k body application

swagger:model GetServerInfoOKBodyApplication
*/
type GetServerInfoOKBodyApplication struct {

	/* Application name

	Required: true
	*/
	Name *string `json:"name"`

	/* Application version

	Required: true
	*/
	Version *string `json:"version"`
}

// Validate validates this get server info o k body application
func (o *GetServerInfoOKBodyApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerInfoOKBodyApplication) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getServerInfoOK"+"."+"application"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetServerInfoOKBodyApplication) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("getServerInfoOK"+"."+"application"+"."+"version", "body", o.Version); err != nil {
		return err
	}

	return nil
}

/*GetServerInfoOKBodyBody get server info o k body body

swagger:model GetServerInfoOKBodyBody
*/
type GetServerInfoOKBodyBody struct {

	/* application

	Required: true
	*/
	Application *GetServerInfoOKBodyApplication `json:"application"`
}

// Validate validates this get server info o k body body
func (o *GetServerInfoOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateApplication(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerInfoOKBodyBody) validateApplication(formats strfmt.Registry) error {

	if err := validate.Required("getServerInfoOK"+"."+"application", "body", o.Application); err != nil {
		return err
	}

	if o.Application != nil {

		if err := o.Application.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
