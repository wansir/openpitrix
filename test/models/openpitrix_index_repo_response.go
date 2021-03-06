// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixIndexRepoResponse openpitrix index repo response
// swagger:model openpitrixIndexRepoResponse
type OpenpitrixIndexRepoResponse struct {

	// repo task
	RepoTask *OpenpitrixRepoTask `json:"repo_task,omitempty"`
}

// Validate validates this openpitrix index repo response
func (m *OpenpitrixIndexRepoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepoTask(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixIndexRepoResponse) validateRepoTask(formats strfmt.Registry) error {

	if swag.IsZero(m.RepoTask) { // not required
		return nil
	}

	if m.RepoTask != nil {

		if err := m.RepoTask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repo_task")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixIndexRepoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixIndexRepoResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixIndexRepoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
