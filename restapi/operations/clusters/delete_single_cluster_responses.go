package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/radanalyticsio/oshinko-rest/models"
)

/*DeleteSingleClusterNoContent Cluster deletion response

swagger:response deleteSingleClusterNoContent
*/
type DeleteSingleClusterNoContent struct {
}

// NewDeleteSingleClusterNoContent creates DeleteSingleClusterNoContent with default headers values
func NewDeleteSingleClusterNoContent() *DeleteSingleClusterNoContent {
	return &DeleteSingleClusterNoContent{}
}

// WriteResponse to the client
func (o *DeleteSingleClusterNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteSingleClusterDefault Unexpected error

swagger:response deleteSingleClusterDefault
*/
type DeleteSingleClusterDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteSingleClusterDefault creates DeleteSingleClusterDefault with default headers values
func NewDeleteSingleClusterDefault(code int) *DeleteSingleClusterDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSingleClusterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete single cluster default response
func (o *DeleteSingleClusterDefault) WithStatusCode(code int) *DeleteSingleClusterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete single cluster default response
func (o *DeleteSingleClusterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete single cluster default response
func (o *DeleteSingleClusterDefault) WithPayload(payload *models.ErrorResponse) *DeleteSingleClusterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete single cluster default response
func (o *DeleteSingleClusterDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSingleClusterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
