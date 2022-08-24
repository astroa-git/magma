// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDPoliciesBaseNamesBaseNameReader is a Reader for the GetNetworksNetworkIDPoliciesBaseNamesBaseName structure.
type GetNetworksNetworkIDPoliciesBaseNamesBaseNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDPoliciesBaseNamesBaseNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDPoliciesBaseNamesBaseNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDPoliciesBaseNamesBaseNameOK creates a GetNetworksNetworkIDPoliciesBaseNamesBaseNameOK with default headers values
func NewGetNetworksNetworkIDPoliciesBaseNamesBaseNameOK() *GetNetworksNetworkIDPoliciesBaseNamesBaseNameOK {
	return &GetNetworksNetworkIDPoliciesBaseNamesBaseNameOK{}
}

/*GetNetworksNetworkIDPoliciesBaseNamesBaseNameOK handles this case with default header values.

Charging Rule Base Name on success
*/
type GetNetworksNetworkIDPoliciesBaseNamesBaseNameOK struct {
	Payload *models.BaseNameRecord
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesBaseNameOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/policies/base_names/{base_name}][%d] getNetworksNetworkIdPoliciesBaseNamesBaseNameOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesBaseNameOK) GetPayload() *models.BaseNameRecord {
	return o.Payload
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesBaseNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BaseNameRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault creates a GetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault with default headers values
func NewGetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault(code int) *GetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault {
	return &GetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID policies base names base name default response
func (o *GetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/policies/base_names/{base_name}][%d] GetNetworksNetworkIDPoliciesBaseNamesBaseName default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDPoliciesBaseNamesBaseNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
