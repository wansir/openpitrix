// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyRuntimeEnvRequest openpitrix modify runtime env request
// swagger:model openpitrixModifyRuntimeEnvRequest
type OpenpitrixModifyRuntimeEnvRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// labels
	Labels string `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// runtime env id
	RuntimeEnvID string `json:"runtime_env_id,omitempty"`
}

// Validate validates this openpitrix modify runtime env request
func (m *OpenpitrixModifyRuntimeEnvRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyRuntimeEnvRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyRuntimeEnvRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyRuntimeEnvRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
