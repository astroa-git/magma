// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutCwfNetworkIDLiUesReader is a Reader for the PutCwfNetworkIDLiUes structure.
type PutCwfNetworkIDLiUesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDLiUesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCwfNetworkIDLiUesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDLiUesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDLiUesNoContent creates a PutCwfNetworkIDLiUesNoContent with default headers values
func NewPutCwfNetworkIDLiUesNoContent() *PutCwfNetworkIDLiUesNoContent {
	return &PutCwfNetworkIDLiUesNoContent{}
}

/*PutCwfNetworkIDLiUesNoContent handles this case with default header values.

Success
*/
type PutCwfNetworkIDLiUesNoContent struct {
}

func (o *PutCwfNetworkIDLiUesNoContent) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/li_ues][%d] putCwfNetworkIdLiUesNoContent ", 204)
}

func (o *PutCwfNetworkIDLiUesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDLiUesDefault creates a PutCwfNetworkIDLiUesDefault with default headers values
func NewPutCwfNetworkIDLiUesDefault(code int) *PutCwfNetworkIDLiUesDefault {
	return &PutCwfNetworkIDLiUesDefault{
		_statusCode: code,
	}
}

/*PutCwfNetworkIDLiUesDefault handles this case with default header values.

Unexpected Error
*/
type PutCwfNetworkIDLiUesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID li ues default response
func (o *PutCwfNetworkIDLiUesDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDLiUesDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/li_ues][%d] PutCwfNetworkIDLiUes default  %+v", o._statusCode, o.Payload)
}

func (o *PutCwfNetworkIDLiUesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDLiUesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
