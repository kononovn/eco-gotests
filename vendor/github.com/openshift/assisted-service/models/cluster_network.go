// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterNetwork A network from which Pod IPs are allocated. This block must not overlap with existing physical networks. These IP addresses are used for the Pod network, and if you need to access the Pods from an external network, configure load balancers and routers to manage the traffic.
//
// swagger:model cluster_network
type ClusterNetwork struct {

	// The IP block address pool.
	Cidr Subnet `json:"cidr,omitempty" gorm:"primaryKey"`

	// The cluster that this network is associated with.
	// Format: uuid
	ClusterID strfmt.UUID `json:"cluster_id,omitempty" gorm:"primaryKey"`

	// The subnet prefix length to assign to each individual node. For example if is set to 23, then each node is assigned a /23 subnet out of the given CIDR, which allows for 510 (2^(32 - 23) - 2) pod IPs addresses.
	// Maximum: 128
	// Minimum: 1
	HostPrefix int64 `json:"host_prefix,omitempty"`
}

// Validate validates this cluster network
func (m *ClusterNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostPrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNetwork) validateCidr(formats strfmt.Registry) error {
	if swag.IsZero(m.Cidr) { // not required
		return nil
	}

	if err := m.Cidr.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cidr")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cidr")
		}
		return err
	}

	return nil
}

func (m *ClusterNetwork) validateClusterID(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster_id", "body", "uuid", m.ClusterID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNetwork) validateHostPrefix(formats strfmt.Registry) error {
	if swag.IsZero(m.HostPrefix) { // not required
		return nil
	}

	if err := validate.MinimumInt("host_prefix", "body", m.HostPrefix, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("host_prefix", "body", m.HostPrefix, 128, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster network based on the context it is used
func (m *ClusterNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCidr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNetwork) contextValidateCidr(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Cidr.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cidr")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cidr")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNetwork) UnmarshalBinary(b []byte) error {
	var res ClusterNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
