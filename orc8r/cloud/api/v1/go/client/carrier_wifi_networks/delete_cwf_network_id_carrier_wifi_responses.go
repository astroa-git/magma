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

// DeleteCwfNetworkIDCarrierWifiReader is a Reader for the DeleteCwfNetworkIDCarrierWifi structure.
type DeleteCwfNetworkIDCarrierWifiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCwfNetworkIDCarrierWifiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCwfNetworkIDCarrierWifiNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCwfNetworkIDCarrierWifiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCwfNetworkIDCarrierWifiNoContent creates a DeleteCwfNetworkIDCarrierWifiNoContent with default headers values
func NewDeleteCwfNetworkIDCarrierWifiNoContent() *DeleteCwfNetworkIDCarrierWifiNoContent {
	return &DeleteCwfNetworkIDCarrierWifiNoContent{}
}

/*DeleteCwfNetworkIDCarrierWifiNoContent handles this case with default header values.

Success
*/
type DeleteCwfNetworkIDCarrierWifiNoContent struct {
}

func (o *DeleteCwfNetworkIDCarrierWifiNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cwf/{network_id}/carrier_wifi][%d] deleteCwfNetworkIdCarrierWifiNoContent ", 204)
}

func (o *DeleteCwfNetworkIDCarrierWifiNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCwfNetworkIDCarrierWifiDefault creates a DeleteCwfNetworkIDCarrierWifiDefault with default headers values
func NewDeleteCwfNetworkIDCarrierWifiDefault(code int) *DeleteCwfNetworkIDCarrierWifiDefault {
	return &DeleteCwfNetworkIDCarrierWifiDefault{
		_statusCode: code,
	}
}

/*DeleteCwfNetworkIDCarrierWifiDefault handles this case with default header values.

Unexpected Error
*/
type DeleteCwfNetworkIDCarrierWifiDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete cwf network ID carrier wifi default response
func (o *DeleteCwfNetworkIDCarrierWifiDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCwfNetworkIDCarrierWifiDefault) Error() string {
	return fmt.Sprintf("[DELETE /cwf/{network_id}/carrier_wifi][%d] DeleteCwfNetworkIDCarrierWifi default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCwfNetworkIDCarrierWifiDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCwfNetworkIDCarrierWifiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
