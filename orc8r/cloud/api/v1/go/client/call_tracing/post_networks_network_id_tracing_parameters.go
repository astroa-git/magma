// Code generated by go-swagger; DO NOT EDIT.

package call_tracing

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

// NewPostNetworksNetworkIDTracingParams creates a new PostNetworksNetworkIDTracingParams object
// with the default values initialized.
func NewPostNetworksNetworkIDTracingParams() *PostNetworksNetworkIDTracingParams {
	var ()
	return &PostNetworksNetworkIDTracingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksNetworkIDTracingParamsWithTimeout creates a new PostNetworksNetworkIDTracingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNetworksNetworkIDTracingParamsWithTimeout(timeout time.Duration) *PostNetworksNetworkIDTracingParams {
	var ()
	return &PostNetworksNetworkIDTracingParams{

		timeout: timeout,
	}
}

// NewPostNetworksNetworkIDTracingParamsWithContext creates a new PostNetworksNetworkIDTracingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNetworksNetworkIDTracingParamsWithContext(ctx context.Context) *PostNetworksNetworkIDTracingParams {
	var ()
	return &PostNetworksNetworkIDTracingParams{

		Context: ctx,
	}
}

// NewPostNetworksNetworkIDTracingParamsWithHTTPClient creates a new PostNetworksNetworkIDTracingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNetworksNetworkIDTracingParamsWithHTTPClient(client *http.Client) *PostNetworksNetworkIDTracingParams {
	var ()
	return &PostNetworksNetworkIDTracingParams{
		HTTPClient: client,
	}
}

/*PostNetworksNetworkIDTracingParams contains all the parameters to send to the API endpoint
for the post networks network ID tracing operation typically these are written to a http.Request
*/
type PostNetworksNetworkIDTracingParams struct {

	/*CallTraceConfiguration
	  Configuration of call trace to start

	*/
	CallTraceConfiguration *models.CallTraceConfig
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) WithTimeout(timeout time.Duration) *PostNetworksNetworkIDTracingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) WithContext(ctx context.Context) *PostNetworksNetworkIDTracingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) WithHTTPClient(client *http.Client) *PostNetworksNetworkIDTracingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallTraceConfiguration adds the callTraceConfiguration to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) WithCallTraceConfiguration(callTraceConfiguration *models.CallTraceConfig) *PostNetworksNetworkIDTracingParams {
	o.SetCallTraceConfiguration(callTraceConfiguration)
	return o
}

// SetCallTraceConfiguration adds the callTraceConfiguration to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) SetCallTraceConfiguration(callTraceConfiguration *models.CallTraceConfig) {
	o.CallTraceConfiguration = callTraceConfiguration
}

// WithNetworkID adds the networkID to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) WithNetworkID(networkID string) *PostNetworksNetworkIDTracingParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post networks network ID tracing params
func (o *PostNetworksNetworkIDTracingParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksNetworkIDTracingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CallTraceConfiguration != nil {
		if err := r.SetBodyParam(o.CallTraceConfiguration); err != nil {
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
