// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutNetworksNetworkIDGatewaysGatewayIDDescriptionReader is a Reader for the PutNetworksNetworkIDGatewaysGatewayIDDescription structure.
type PutNetworksNetworkIDGatewaysGatewayIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent creates a PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent() *PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent {
	return &PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent{}
}

/*PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent struct {
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}/description][%d] putNetworksNetworkIdGatewaysGatewayIdDescriptionNoContent ", 204)
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault creates a PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault(code int) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault {
	return &PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID gateways gateway ID description default response
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}/description][%d] PutNetworksNetworkIDGatewaysGatewayIDDescription default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
