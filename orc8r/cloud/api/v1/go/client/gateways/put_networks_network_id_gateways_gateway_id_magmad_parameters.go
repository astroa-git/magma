// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewPutNetworksNetworkIDGatewaysGatewayIDMagmadParams creates a new PutNetworksNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized.
func NewPutNetworksNetworkIDGatewaysGatewayIDMagmadParams() *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDMagmadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDMagmadParamsWithTimeout creates a new PutNetworksNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNetworksNetworkIDGatewaysGatewayIDMagmadParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDMagmadParams{

		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDMagmadParamsWithContext creates a new PutNetworksNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNetworksNetworkIDGatewaysGatewayIDMagmadParamsWithContext(ctx context.Context) *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDMagmadParams{

		Context: ctx,
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDMagmadParamsWithHTTPClient creates a new PutNetworksNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNetworksNetworkIDGatewaysGatewayIDMagmadParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDMagmadParams{
		HTTPClient: client,
	}
}

/*PutNetworksNetworkIDGatewaysGatewayIDMagmadParams contains all the parameters to send to the API endpoint
for the put networks network ID gateways gateway ID magmad operation typically these are written to a http.Request
*/
type PutNetworksNetworkIDGatewaysGatewayIDMagmadParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*Magmad
	  New magmad configuration

	*/
	Magmad *models.MagmadGatewayConfigs
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) WithContext(ctx context.Context) *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) WithGatewayID(gatewayID string) *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithMagmad adds the magmad to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) WithMagmad(magmad *models.MagmadGatewayConfigs) *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetMagmad(magmad)
	return o
}

// SetMagmad adds the magmad to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) SetMagmad(magmad *models.MagmadGatewayConfigs) {
	o.Magmad = magmad
}

// WithNetworkID adds the networkID to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) WithNetworkID(networkID string) *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID gateways gateway ID magmad params
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	if o.Magmad != nil {
		if err := r.SetBodyParam(o.Magmad); err != nil {
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
