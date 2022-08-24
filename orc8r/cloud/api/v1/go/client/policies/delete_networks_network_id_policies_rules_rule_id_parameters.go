// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteNetworksNetworkIDPoliciesRulesRuleIDParams creates a new DeleteNetworksNetworkIDPoliciesRulesRuleIDParams object
// with the default values initialized.
func NewDeleteNetworksNetworkIDPoliciesRulesRuleIDParams() *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams {
	var ()
	return &DeleteNetworksNetworkIDPoliciesRulesRuleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworksNetworkIDPoliciesRulesRuleIDParamsWithTimeout creates a new DeleteNetworksNetworkIDPoliciesRulesRuleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNetworksNetworkIDPoliciesRulesRuleIDParamsWithTimeout(timeout time.Duration) *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams {
	var ()
	return &DeleteNetworksNetworkIDPoliciesRulesRuleIDParams{

		timeout: timeout,
	}
}

// NewDeleteNetworksNetworkIDPoliciesRulesRuleIDParamsWithContext creates a new DeleteNetworksNetworkIDPoliciesRulesRuleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNetworksNetworkIDPoliciesRulesRuleIDParamsWithContext(ctx context.Context) *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams {
	var ()
	return &DeleteNetworksNetworkIDPoliciesRulesRuleIDParams{

		Context: ctx,
	}
}

// NewDeleteNetworksNetworkIDPoliciesRulesRuleIDParamsWithHTTPClient creates a new DeleteNetworksNetworkIDPoliciesRulesRuleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNetworksNetworkIDPoliciesRulesRuleIDParamsWithHTTPClient(client *http.Client) *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams {
	var ()
	return &DeleteNetworksNetworkIDPoliciesRulesRuleIDParams{
		HTTPClient: client,
	}
}

/*DeleteNetworksNetworkIDPoliciesRulesRuleIDParams contains all the parameters to send to the API endpoint
for the delete networks network ID policies rules rule ID operation typically these are written to a http.Request
*/
type DeleteNetworksNetworkIDPoliciesRulesRuleIDParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*RuleID
	  Rule Id

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) WithTimeout(timeout time.Duration) *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) WithContext(ctx context.Context) *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) WithHTTPClient(client *http.Client) *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) WithNetworkID(networkID string) *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRuleID adds the ruleID to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) WithRuleID(ruleID string) *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the delete networks network ID policies rules rule ID params
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param rule_id
	if err := r.SetPathParam("rule_id", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
