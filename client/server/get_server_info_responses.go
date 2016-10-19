package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/radanalyticsio/oshinko-rest/models"
)

// GetServerInfoReader is a Reader for the GetServerInfo structure.
type GetServerInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServerInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetServerInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetServerInfoOK creates a GetServerInfoOK with default headers values
func NewGetServerInfoOK() *GetServerInfoOK {
	return &GetServerInfoOK{}
}

/*GetServerInfoOK handles this case with default header values.

Server info response
*/
type GetServerInfoOK struct {
	Payload GetServerInfoOKBodyBody
}

func (o *GetServerInfoOK) Error() string {
	return fmt.Sprintf("[GET /][%d] getServerInfoOK  %+v", 200, o.Payload)
}

func (o *GetServerInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerInfoDefault creates a GetServerInfoDefault with default headers values
func NewGetServerInfoDefault(code int) *GetServerInfoDefault {
	return &GetServerInfoDefault{
		_statusCode: code,
	}
}

/*GetServerInfoDefault handles this case with default header values.

Unexpected error
*/
type GetServerInfoDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get server info default response
func (o *GetServerInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetServerInfoDefault) Error() string {
	return fmt.Sprintf("[GET /][%d] getServerInfo default  %+v", o._statusCode, o.Payload)
}

func (o *GetServerInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
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

	/* Oshinko Web Service Name

	Required: true
	*/
	WebServiceName *string `json:"web-service-name"`

	/* Oshinko Web URL

	Required: true
	*/
	WebURL *string `json:"web-url"`
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

	if err := o.validateWebServiceName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateWebURL(formats); err != nil {
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

func (o *GetServerInfoOKBodyApplication) validateWebServiceName(formats strfmt.Registry) error {

	if err := validate.Required("getServerInfoOK"+"."+"application"+"."+"web-service-name", "body", o.WebServiceName); err != nil {
		return err
	}

	return nil
}

func (o *GetServerInfoOKBodyApplication) validateWebURL(formats strfmt.Registry) error {

	if err := validate.Required("getServerInfoOK"+"."+"application"+"."+"web-url", "body", o.WebURL); err != nil {
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
