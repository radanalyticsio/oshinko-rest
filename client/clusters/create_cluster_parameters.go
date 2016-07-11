package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/redhatanalytics/oshinko-rest/models"
)

// NewCreateClusterParams creates a new CreateClusterParams object
// with the default values initialized.
func NewCreateClusterParams() *CreateClusterParams {
	var ()
	return &CreateClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterParamsWithTimeout creates a new CreateClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterParamsWithTimeout(timeout time.Duration) *CreateClusterParams {
	var ()
	return &CreateClusterParams{

		timeout: timeout,
	}
}

/*CreateClusterParams contains all the parameters to send to the API endpoint
for the create cluster operation typically these are written to a http.Request
*/
type CreateClusterParams struct {

	/*Cluster
	  Cluster to create

	*/
	Cluster *models.NewCluster

	timeout time.Duration
}

// WithCluster adds the cluster to the create cluster params
func (o *CreateClusterParams) WithCluster(Cluster *models.NewCluster) *CreateClusterParams {
	o.Cluster = Cluster
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Cluster == nil {
		o.Cluster = new(models.NewCluster)
	}

	if err := r.SetBodyParam(o.Cluster); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
