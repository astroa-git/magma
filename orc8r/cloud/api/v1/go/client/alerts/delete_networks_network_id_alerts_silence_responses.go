// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteNetworksNetworkIDAlertsSilenceReader is a Reader for the DeleteNetworksNetworkIDAlertsSilence structure.
type DeleteNetworksNetworkIDAlertsSilenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDAlertsSilenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworksNetworkIDAlertsSilenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDAlertsSilenceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDAlertsSilenceOK creates a DeleteNetworksNetworkIDAlertsSilenceOK with default headers values
func NewDeleteNetworksNetworkIDAlertsSilenceOK() *DeleteNetworksNetworkIDAlertsSilenceOK {
	return &DeleteNetworksNetworkIDAlertsSilenceOK{}
}

/*DeleteNetworksNetworkIDAlertsSilenceOK handles this case with default header values.

Successfully deleted silencer
*/
type DeleteNetworksNetworkIDAlertsSilenceOK struct {
	Payload string
}

func (o *DeleteNetworksNetworkIDAlertsSilenceOK) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/alerts/silence][%d] deleteNetworksNetworkIdAlertsSilenceOK  %+v", 200, o.Payload)
}

func (o *DeleteNetworksNetworkIDAlertsSilenceOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDAlertsSilenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworksNetworkIDAlertsSilenceDefault creates a DeleteNetworksNetworkIDAlertsSilenceDefault with default headers values
func NewDeleteNetworksNetworkIDAlertsSilenceDefault(code int) *DeleteNetworksNetworkIDAlertsSilenceDefault {
	return &DeleteNetworksNetworkIDAlertsSilenceDefault{
		_statusCode: code,
	}
}

/*DeleteNetworksNetworkIDAlertsSilenceDefault handles this case with default header values.

Network ID
*/
type DeleteNetworksNetworkIDAlertsSilenceDefault struct {
	_statusCode int
}

// Code gets the status code for the delete networks network ID alerts silence default response
func (o *DeleteNetworksNetworkIDAlertsSilenceDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDAlertsSilenceDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/alerts/silence][%d] DeleteNetworksNetworkIDAlertsSilence default ", o._statusCode)
}

func (o *DeleteNetworksNetworkIDAlertsSilenceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
