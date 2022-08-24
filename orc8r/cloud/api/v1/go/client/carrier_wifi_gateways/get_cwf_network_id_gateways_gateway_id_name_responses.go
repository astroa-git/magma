// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetCwfNetworkIDGatewaysGatewayIDNameReader is a Reader for the GetCwfNetworkIDGatewaysGatewayIDName structure.
type GetCwfNetworkIDGatewaysGatewayIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDGatewaysGatewayIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDGatewaysGatewayIDNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDGatewaysGatewayIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDNameOK creates a GetCwfNetworkIDGatewaysGatewayIDNameOK with default headers values
func NewGetCwfNetworkIDGatewaysGatewayIDNameOK() *GetCwfNetworkIDGatewaysGatewayIDNameOK {
	return &GetCwfNetworkIDGatewaysGatewayIDNameOK{}
}

/*GetCwfNetworkIDGatewaysGatewayIDNameOK handles this case with default header values.

The name of the gateway
*/
type GetCwfNetworkIDGatewaysGatewayIDNameOK struct {
	Payload models.GatewayName
}

func (o *GetCwfNetworkIDGatewaysGatewayIDNameOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways/{gateway_id}/name][%d] getCwfNetworkIdGatewaysGatewayIdNameOK  %+v", 200, o.Payload)
}

func (o *GetCwfNetworkIDGatewaysGatewayIDNameOK) GetPayload() models.GatewayName {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysGatewayIDNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDGatewaysGatewayIDNameDefault creates a GetCwfNetworkIDGatewaysGatewayIDNameDefault with default headers values
func NewGetCwfNetworkIDGatewaysGatewayIDNameDefault(code int) *GetCwfNetworkIDGatewaysGatewayIDNameDefault {
	return &GetCwfNetworkIDGatewaysGatewayIDNameDefault{
		_statusCode: code,
	}
}

/*GetCwfNetworkIDGatewaysGatewayIDNameDefault handles this case with default header values.

Unexpected Error
*/
type GetCwfNetworkIDGatewaysGatewayIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID gateways gateway ID name default response
func (o *GetCwfNetworkIDGatewaysGatewayIDNameDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDGatewaysGatewayIDNameDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways/{gateway_id}/name][%d] GetCwfNetworkIDGatewaysGatewayIDName default  %+v", o._statusCode, o.Payload)
}

func (o *GetCwfNetworkIDGatewaysGatewayIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysGatewayIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
