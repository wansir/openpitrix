// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteRuntimeEnvCredentialResponse openpitrix delete runtime env credential response
// swagger:model openpitrixDeleteRuntimeEnvCredentialResponse
type OpenpitrixDeleteRuntimeEnvCredentialResponse struct {

	// runtime env credential
	RuntimeEnvCredential *OpenpitrixRuntimeEnvCredential `json:"runtime_env_credential,omitempty"`
}

// Validate validates this openpitrix delete runtime env credential response
func (m *OpenpitrixDeleteRuntimeEnvCredentialResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuntimeEnvCredential(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteRuntimeEnvCredentialResponse) validateRuntimeEnvCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.RuntimeEnvCredential) { // not required
		return nil
	}

	if m.RuntimeEnvCredential != nil {

		if err := m.RuntimeEnvCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runtime_env_credential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteRuntimeEnvCredentialResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteRuntimeEnvCredentialResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteRuntimeEnvCredentialResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
