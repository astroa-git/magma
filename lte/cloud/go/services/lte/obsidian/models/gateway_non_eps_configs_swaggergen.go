// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayNonEpsConfigs Non-EPS service configuration for a gateway
//
// swagger:model gateway_non_eps_configs
type GatewayNonEpsConfigs struct {

	// arfcn 2g
	Arfcn2g []uint32 `json:"arfcn_2g"`

	// csfb mcc
	// Example: 001
	// Min Length: 1
	// Pattern: ^(\d{3})$
	CsfbMcc string `json:"csfb_mcc,omitempty"`

	// csfb mnc
	// Example: 01
	// Min Length: 1
	// Pattern: ^(\d{2,3})$
	CsfbMnc string `json:"csfb_mnc,omitempty"`

	// csfb rat
	// Enum: [0 1]
	CsfbRat *uint32 `json:"csfb_rat,omitempty"`

	// lac
	// Example: 1
	Lac *uint32 `json:"lac,omitempty"`

	// non eps service control
	// Required: true
	// Enum: [0 1 2 3]
	NonEpsServiceControl *uint32 `json:"non_eps_service_control"`
}

// Validate validates this gateway non eps configs
func (m *GatewayNonEpsConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCsfbMcc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsfbMnc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsfbRat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonEpsServiceControl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayNonEpsConfigs) validateCsfbMcc(formats strfmt.Registry) error {
	if swag.IsZero(m.CsfbMcc) { // not required
		return nil
	}

	if err := validate.MinLength("csfb_mcc", "body", m.CsfbMcc, 1); err != nil {
		return err
	}

	if err := validate.Pattern("csfb_mcc", "body", m.CsfbMcc, `^(\d{3})$`); err != nil {
		return err
	}

	return nil
}

func (m *GatewayNonEpsConfigs) validateCsfbMnc(formats strfmt.Registry) error {
	if swag.IsZero(m.CsfbMnc) { // not required
		return nil
	}

	if err := validate.MinLength("csfb_mnc", "body", m.CsfbMnc, 1); err != nil {
		return err
	}

	if err := validate.Pattern("csfb_mnc", "body", m.CsfbMnc, `^(\d{2,3})$`); err != nil {
		return err
	}

	return nil
}

var gatewayNonEpsConfigsTypeCsfbRatPropEnum []interface{}

func init() {
	var res []uint32
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gatewayNonEpsConfigsTypeCsfbRatPropEnum = append(gatewayNonEpsConfigsTypeCsfbRatPropEnum, v)
	}
}

// prop value enum
func (m *GatewayNonEpsConfigs) validateCsfbRatEnum(path, location string, value uint32) error {
	if err := validate.EnumCase(path, location, value, gatewayNonEpsConfigsTypeCsfbRatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GatewayNonEpsConfigs) validateCsfbRat(formats strfmt.Registry) error {
	if swag.IsZero(m.CsfbRat) { // not required
		return nil
	}

	// value enum
	if err := m.validateCsfbRatEnum("csfb_rat", "body", *m.CsfbRat); err != nil {
		return err
	}

	return nil
}

var gatewayNonEpsConfigsTypeNonEpsServiceControlPropEnum []interface{}

func init() {
	var res []uint32
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gatewayNonEpsConfigsTypeNonEpsServiceControlPropEnum = append(gatewayNonEpsConfigsTypeNonEpsServiceControlPropEnum, v)
	}
}

// prop value enum
func (m *GatewayNonEpsConfigs) validateNonEpsServiceControlEnum(path, location string, value uint32) error {
	if err := validate.EnumCase(path, location, value, gatewayNonEpsConfigsTypeNonEpsServiceControlPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GatewayNonEpsConfigs) validateNonEpsServiceControl(formats strfmt.Registry) error {

	if err := validate.Required("non_eps_service_control", "body", m.NonEpsServiceControl); err != nil {
		return err
	}

	// value enum
	if err := m.validateNonEpsServiceControlEnum("non_eps_service_control", "body", *m.NonEpsServiceControl); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gateway non eps configs based on context it is used
func (m *GatewayNonEpsConfigs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GatewayNonEpsConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayNonEpsConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayNonEpsConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
