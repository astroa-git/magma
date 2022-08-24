// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutFegLTENetworkIDSubscriberConfigRuleNamesReader is a Reader for the PutFegLTENetworkIDSubscriberConfigRuleNames structure.
type PutFegLTENetworkIDSubscriberConfigRuleNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFegLTENetworkIDSubscriberConfigRuleNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutFegLTENetworkIDSubscriberConfigRuleNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutFegLTENetworkIDSubscriberConfigRuleNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutFegLTENetworkIDSubscriberConfigRuleNamesNoContent creates a PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent with default headers values
func NewPutFegLTENetworkIDSubscriberConfigRuleNamesNoContent() *PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent {
	return &PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent{}
}

/*PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent handles this case with default header values.

Success
*/
type PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent struct {
}

func (o *PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent) Error() string {
	return fmt.Sprintf("[PUT /feg_lte/{network_id}/subscriber_config/rule_names][%d] putFegLteNetworkIdSubscriberConfigRuleNamesNoContent ", 204)
}

func (o *PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFegLTENetworkIDSubscriberConfigRuleNamesDefault creates a PutFegLTENetworkIDSubscriberConfigRuleNamesDefault with default headers values
func NewPutFegLTENetworkIDSubscriberConfigRuleNamesDefault(code int) *PutFegLTENetworkIDSubscriberConfigRuleNamesDefault {
	return &PutFegLTENetworkIDSubscriberConfigRuleNamesDefault{
		_statusCode: code,
	}
}

/*PutFegLTENetworkIDSubscriberConfigRuleNamesDefault handles this case with default header values.

Unexpected Error
*/
type PutFegLTENetworkIDSubscriberConfigRuleNamesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put feg LTE network ID subscriber config rule names default response
func (o *PutFegLTENetworkIDSubscriberConfigRuleNamesDefault) Code() int {
	return o._statusCode
}

func (o *PutFegLTENetworkIDSubscriberConfigRuleNamesDefault) Error() string {
	return fmt.Sprintf("[PUT /feg_lte/{network_id}/subscriber_config/rule_names][%d] PutFegLTENetworkIDSubscriberConfigRuleNames default  %+v", o._statusCode, o.Payload)
}

func (o *PutFegLTENetworkIDSubscriberConfigRuleNamesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutFegLTENetworkIDSubscriberConfigRuleNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
