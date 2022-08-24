// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDGatewaysGatewayIDVpnReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDVpn structure.
type GetLTENetworkIDGatewaysGatewayIDVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysGatewayIDVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDVpnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDVpnOK creates a GetLTENetworkIDGatewaysGatewayIDVpnOK with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDVpnOK() *GetLTENetworkIDGatewaysGatewayIDVpnOK {
	return &GetLTENetworkIDGatewaysGatewayIDVpnOK{}
}

/*GetLTENetworkIDGatewaysGatewayIDVpnOK handles this case with default header values.

The current VPN gateway config
*/
type GetLTENetworkIDGatewaysGatewayIDVpnOK struct {
	Payload *models.GatewayVpnConfigs
}

func (o *GetLTENetworkIDGatewaysGatewayIDVpnOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/vpn][%d] getLteNetworkIdGatewaysGatewayIdVpnOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDVpnOK) GetPayload() *models.GatewayVpnConfigs {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayVpnConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDVpnDefault creates a GetLTENetworkIDGatewaysGatewayIDVpnDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDVpnDefault(code int) *GetLTENetworkIDGatewaysGatewayIDVpnDefault {
	return &GetLTENetworkIDGatewaysGatewayIDVpnDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDVpnDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDVpnDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID vpn default response
func (o *GetLTENetworkIDGatewaysGatewayIDVpnDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDVpnDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/vpn][%d] GetLTENetworkIDGatewaysGatewayIDVpn default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDVpnDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDVpnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
