// Code generated by go-swagger; DO NOT EDIT.

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// DeleteNetworksNetworkIDTiersTierIDReader is a Reader for the DeleteNetworksNetworkIDTiersTierID structure.
type DeleteNetworksNetworkIDTiersTierIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDTiersTierIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworksNetworkIDTiersTierIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDTiersTierIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDTiersTierIDNoContent creates a DeleteNetworksNetworkIDTiersTierIDNoContent with default headers values
func NewDeleteNetworksNetworkIDTiersTierIDNoContent() *DeleteNetworksNetworkIDTiersTierIDNoContent {
	return &DeleteNetworksNetworkIDTiersTierIDNoContent{}
}

/*DeleteNetworksNetworkIDTiersTierIDNoContent handles this case with default header values.

Success
*/
type DeleteNetworksNetworkIDTiersTierIDNoContent struct {
}

func (o *DeleteNetworksNetworkIDTiersTierIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/tiers/{tier_id}][%d] deleteNetworksNetworkIdTiersTierIdNoContent ", 204)
}

func (o *DeleteNetworksNetworkIDTiersTierIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworksNetworkIDTiersTierIDDefault creates a DeleteNetworksNetworkIDTiersTierIDDefault with default headers values
func NewDeleteNetworksNetworkIDTiersTierIDDefault(code int) *DeleteNetworksNetworkIDTiersTierIDDefault {
	return &DeleteNetworksNetworkIDTiersTierIDDefault{
		_statusCode: code,
	}
}

/*DeleteNetworksNetworkIDTiersTierIDDefault handles this case with default header values.

Unexpected Error
*/
type DeleteNetworksNetworkIDTiersTierIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete networks network ID tiers tier ID default response
func (o *DeleteNetworksNetworkIDTiersTierIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDTiersTierIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/tiers/{tier_id}][%d] DeleteNetworksNetworkIDTiersTierID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNetworksNetworkIDTiersTierIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDTiersTierIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
