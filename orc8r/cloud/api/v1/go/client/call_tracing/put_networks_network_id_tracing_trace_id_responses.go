// Code generated by go-swagger; DO NOT EDIT.

package call_tracing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutNetworksNetworkIDTracingTraceIDReader is a Reader for the PutNetworksNetworkIDTracingTraceID structure.
type PutNetworksNetworkIDTracingTraceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDTracingTraceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDTracingTraceIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDTracingTraceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDTracingTraceIDNoContent creates a PutNetworksNetworkIDTracingTraceIDNoContent with default headers values
func NewPutNetworksNetworkIDTracingTraceIDNoContent() *PutNetworksNetworkIDTracingTraceIDNoContent {
	return &PutNetworksNetworkIDTracingTraceIDNoContent{}
}

/*PutNetworksNetworkIDTracingTraceIDNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDTracingTraceIDNoContent struct {
}

func (o *PutNetworksNetworkIDTracingTraceIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tracing/{trace_id}][%d] putNetworksNetworkIdTracingTraceIdNoContent ", 204)
}

func (o *PutNetworksNetworkIDTracingTraceIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDTracingTraceIDDefault creates a PutNetworksNetworkIDTracingTraceIDDefault with default headers values
func NewPutNetworksNetworkIDTracingTraceIDDefault(code int) *PutNetworksNetworkIDTracingTraceIDDefault {
	return &PutNetworksNetworkIDTracingTraceIDDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDTracingTraceIDDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDTracingTraceIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID tracing trace ID default response
func (o *PutNetworksNetworkIDTracingTraceIDDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDTracingTraceIDDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tracing/{trace_id}][%d] PutNetworksNetworkIDTracingTraceID default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDTracingTraceIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDTracingTraceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
