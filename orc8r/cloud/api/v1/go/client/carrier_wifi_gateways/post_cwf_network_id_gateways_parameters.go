// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

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

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPostCwfNetworkIDGatewaysParams creates a new PostCwfNetworkIDGatewaysParams object
// with the default values initialized.
func NewPostCwfNetworkIDGatewaysParams() *PostCwfNetworkIDGatewaysParams {
	var ()
	return &PostCwfNetworkIDGatewaysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCwfNetworkIDGatewaysParamsWithTimeout creates a new PostCwfNetworkIDGatewaysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCwfNetworkIDGatewaysParamsWithTimeout(timeout time.Duration) *PostCwfNetworkIDGatewaysParams {
	var ()
	return &PostCwfNetworkIDGatewaysParams{

		timeout: timeout,
	}
}

// NewPostCwfNetworkIDGatewaysParamsWithContext creates a new PostCwfNetworkIDGatewaysParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCwfNetworkIDGatewaysParamsWithContext(ctx context.Context) *PostCwfNetworkIDGatewaysParams {
	var ()
	return &PostCwfNetworkIDGatewaysParams{

		Context: ctx,
	}
}

// NewPostCwfNetworkIDGatewaysParamsWithHTTPClient creates a new PostCwfNetworkIDGatewaysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCwfNetworkIDGatewaysParamsWithHTTPClient(client *http.Client) *PostCwfNetworkIDGatewaysParams {
	var ()
	return &PostCwfNetworkIDGatewaysParams{
		HTTPClient: client,
	}
}

/*PostCwfNetworkIDGatewaysParams contains all the parameters to send to the API endpoint
for the post cwf network ID gateways operation typically these are written to a http.Request
*/
type PostCwfNetworkIDGatewaysParams struct {

	/*Gateway
	  Full desired configuration of the gateway

	*/
	Gateway *models.MutableCwfGateway
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) WithTimeout(timeout time.Duration) *PostCwfNetworkIDGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) WithContext(ctx context.Context) *PostCwfNetworkIDGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) WithHTTPClient(client *http.Client) *PostCwfNetworkIDGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGateway adds the gateway to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) WithGateway(gateway *models.MutableCwfGateway) *PostCwfNetworkIDGatewaysParams {
	o.SetGateway(gateway)
	return o
}

// SetGateway adds the gateway to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) SetGateway(gateway *models.MutableCwfGateway) {
	o.Gateway = gateway
}

// WithNetworkID adds the networkID to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) WithNetworkID(networkID string) *PostCwfNetworkIDGatewaysParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post cwf network ID gateways params
func (o *PostCwfNetworkIDGatewaysParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostCwfNetworkIDGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Gateway != nil {
		if err := r.SetBodyParam(o.Gateway); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
