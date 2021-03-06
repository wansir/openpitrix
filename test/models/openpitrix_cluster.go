// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCluster openpitrix cluster
// swagger:model openpitrixCluster
type OpenpitrixCluster struct {

	// advanced actions
	AdvancedActions map[string]string `json:"advanced_actions,omitempty"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// cluster node set
	ClusterNodeSet OpenpitrixClusterClusterNodeSet `json:"cluster_node_set"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// endpoints
	Endpoints string `json:"endpoints,omitempty"`

	// frontgate id
	FrontgateID string `json:"frontgate_id,omitempty"`

	// global uuid
	GlobalUUID string `json:"global_uuid,omitempty"`

	// links
	Links map[string]string `json:"links,omitempty"`

	// metadata root access
	MetadataRootAccess bool `json:"metadata_root_access,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// node count
	NodeCount *ProtobufUint32Value `json:"node_count,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// roles
	Roles map[string]string `json:"roles,omitempty"`

	// runtime env id
	RuntimeEnvID string `json:"runtime_env_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// transition status
	TransitionStatus string `json:"transition_status,omitempty"`

	// upgrade status
	UpgradeStatus string `json:"upgrade_status,omitempty"`

	// version id
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix cluster
func (m *OpenpitrixCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixCluster) validateNodeCount(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeCount) { // not required
		return nil
	}

	if m.NodeCount != nil {

		if err := m.NodeCount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_count")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCluster) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
