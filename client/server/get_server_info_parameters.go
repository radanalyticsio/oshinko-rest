package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServerInfoParams creates a new GetServerInfoParams object
// with the default values initialized.
func NewGetServerInfoParams() *GetServerInfoParams {

	return &GetServerInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerInfoParamsWithTimeout creates a new GetServerInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServerInfoParamsWithTimeout(timeout time.Duration) *GetServerInfoParams {

	return &GetServerInfoParams{

		timeout: timeout,
	}
}

/*GetServerInfoParams contains all the parameters to send to the API endpoint
for the get server info operation typically these are written to a http.Request
*/
type GetServerInfoParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
