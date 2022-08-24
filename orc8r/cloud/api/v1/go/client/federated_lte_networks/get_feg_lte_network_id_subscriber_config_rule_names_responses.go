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

// GetFegLTENetworkIDSubscriberConfigRuleNamesReader is a Reader for the GetFegLTENetworkIDSubscriberConfigRuleNames structure.
type GetFegLTENetworkIDSubscriberConfigRuleNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegLTENetworkIDSubscriberConfigRuleNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegLTENetworkIDSubscriberConfigRuleNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegLTENetworkIDSubscriberConfigRuleNamesOK creates a GetFegLTENetworkIDSubscriberConfigRuleNamesOK with default headers values
func NewGetFegLTENetworkIDSubscriberConfigRuleNamesOK() *GetFegLTENetworkIDSubscriberConfigRuleNamesOK {
	return &GetFegLTENetworkIDSubscriberConfigRuleNamesOK{}
}

/*GetFegLTENetworkIDSubscriberConfigRuleNamesOK handles this case with default header values.

Subscriber Config
*/
type GetFegLTENetworkIDSubscriberConfigRuleNamesOK struct {
	Payload models.RuleNames
}

func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesOK) Error() string {
	return fmt.Sprintf("[GET /feg_lte/{network_id}/subscriber_config/rule_names][%d] getFegLteNetworkIdSubscriberConfigRuleNamesOK  %+v", 200, o.Payload)
}

func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesOK) GetPayload() models.RuleNames {
	return o.Payload
}

func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegLTENetworkIDSubscriberConfigRuleNamesDefault creates a GetFegLTENetworkIDSubscriberConfigRuleNamesDefault with default headers values
func NewGetFegLTENetworkIDSubscriberConfigRuleNamesDefault(code int) *GetFegLTENetworkIDSubscriberConfigRuleNamesDefault {
	return &GetFegLTENetworkIDSubscriberConfigRuleNamesDefault{
		_statusCode: code,
	}
}

/*GetFegLTENetworkIDSubscriberConfigRuleNamesDefault handles this case with default header values.

Unexpected Error
*/
type GetFegLTENetworkIDSubscriberConfigRuleNamesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg LTE network ID subscriber config rule names default response
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesDefault) Code() int {
	return o._statusCode
}

func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesDefault) Error() string {
	return fmt.Sprintf("[GET /feg_lte/{network_id}/subscriber_config/rule_names][%d] GetFegLTENetworkIDSubscriberConfigRuleNames default  %+v", o._statusCode, o.Payload)
}

func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
