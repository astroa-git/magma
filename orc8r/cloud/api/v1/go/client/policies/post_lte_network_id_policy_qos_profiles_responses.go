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

// PostLTENetworkIDPolicyQosProfilesReader is a Reader for the PostLTENetworkIDPolicyQosProfiles structure.
type PostLTENetworkIDPolicyQosProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLTENetworkIDPolicyQosProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLTENetworkIDPolicyQosProfilesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostLTENetworkIDPolicyQosProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLTENetworkIDPolicyQosProfilesCreated creates a PostLTENetworkIDPolicyQosProfilesCreated with default headers values
func NewPostLTENetworkIDPolicyQosProfilesCreated() *PostLTENetworkIDPolicyQosProfilesCreated {
	return &PostLTENetworkIDPolicyQosProfilesCreated{}
}

/*PostLTENetworkIDPolicyQosProfilesCreated handles this case with default header values.

Success
*/
type PostLTENetworkIDPolicyQosProfilesCreated struct {
}

func (o *PostLTENetworkIDPolicyQosProfilesCreated) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/policy_qos_profiles][%d] postLteNetworkIdPolicyQosProfilesCreated ", 201)
}

func (o *PostLTENetworkIDPolicyQosProfilesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLTENetworkIDPolicyQosProfilesDefault creates a PostLTENetworkIDPolicyQosProfilesDefault with default headers values
func NewPostLTENetworkIDPolicyQosProfilesDefault(code int) *PostLTENetworkIDPolicyQosProfilesDefault {
	return &PostLTENetworkIDPolicyQosProfilesDefault{
		_statusCode: code,
	}
}

/*PostLTENetworkIDPolicyQosProfilesDefault handles this case with default header values.

Unexpected Error
*/
type PostLTENetworkIDPolicyQosProfilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post LTE network ID policy qos profiles default response
func (o *PostLTENetworkIDPolicyQosProfilesDefault) Code() int {
	return o._statusCode
}

func (o *PostLTENetworkIDPolicyQosProfilesDefault) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/policy_qos_profiles][%d] PostLTENetworkIDPolicyQosProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *PostLTENetworkIDPolicyQosProfilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostLTENetworkIDPolicyQosProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
