// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CellularGatewayPoolConfigs Configuration for gateway pool
// swagger:model cellular_gateway_pool_configs
type CellularGatewayPoolConfigs struct {

	// mme group id
	// Required: true
	// Maximum: 255
	// Minimum: 0
	MmeGroupID uint32 `json:"mme_group_id"`
}

// Validate validates this cellular gateway pool configs
func (m *CellularGatewayPoolConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMmeGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CellularGatewayPoolConfigs) validateMmeGroupID(formats strfmt.Registry) error {

	if err := validate.Required("mme_group_id", "body", uint32(m.MmeGroupID)); err != nil {
		return err
	}

	if err := validate.MinimumInt("mme_group_id", "body", int64(m.MmeGroupID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mme_group_id", "body", int64(m.MmeGroupID), 255, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CellularGatewayPoolConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CellularGatewayPoolConfigs) UnmarshalBinary(b []byte) error {
	var res CellularGatewayPoolConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
