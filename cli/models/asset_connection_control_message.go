// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AssetConnectionControlMessage Asset:Connection Control Message
//
// Control message used to update the context of a devices connection. When a device registration is modified (e.g. by a user modifying the claim status) the process managing the connection must be notified of the change. On changes aurora will attempt to send the change to the current owner of the connection.
//
// swagger:model assetConnectionControlMessage
type AssetConnectionControlMessage struct {

	// The account id to which the device belongs.
	//
	Account string `json:"Account,omitempty"`

	// The version of the device connector currently running on the platform. Deprecated by newer connectors that will report this directly to the device connector gateway in a websocket header, but included to continue to support older versions which report any version change after connect.
	//
	ConnectorVersion string `json:"ConnectorVersion,omitempty"`

	// The Moid of the device under change. Used to route the message to a devices connection.
	//
	DeviceID string `json:"DeviceId,omitempty"`

	// The domain group id to which the device belongs.
	//
	DomainGroup string `json:"DomainGroup,omitempty"`

	// Flag to force any open connections to be evicted. Used in case device has been deleted or blacklisted.
	//
	Evict bool `json:"Evict,omitempty"`

	// The current leadership of a device cluster member.
	//
	// Enum: [Unknown Primary Secondary]
	Leadership *string `json:"Leadership,omitempty"`

	// The new identity assigned to a device on ownership change (claim/unclaim).
	//
	NewIdentity string `json:"NewIdentity,omitempty"`

	// The partition the device was last connected to, used to address the control message to the device connector gateway instance holding the devices connection.
	//
	Partition int64 `json:"Partition,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetConnectionControlMessage) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Account string `json:"Account,omitempty"`

		ConnectorVersion string `json:"ConnectorVersion,omitempty"`

		DeviceID string `json:"DeviceId,omitempty"`

		DomainGroup string `json:"DomainGroup,omitempty"`

		Evict bool `json:"Evict,omitempty"`

		Leadership *string `json:"Leadership,omitempty"`

		NewIdentity string `json:"NewIdentity,omitempty"`

		Partition int64 `json:"Partition,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Account = dataAO0.Account

	m.ConnectorVersion = dataAO0.ConnectorVersion

	m.DeviceID = dataAO0.DeviceID

	m.DomainGroup = dataAO0.DomainGroup

	m.Evict = dataAO0.Evict

	m.Leadership = dataAO0.Leadership

	m.NewIdentity = dataAO0.NewIdentity

	m.Partition = dataAO0.Partition

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetConnectionControlMessage) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Account string `json:"Account,omitempty"`

		ConnectorVersion string `json:"ConnectorVersion,omitempty"`

		DeviceID string `json:"DeviceId,omitempty"`

		DomainGroup string `json:"DomainGroup,omitempty"`

		Evict bool `json:"Evict,omitempty"`

		Leadership *string `json:"Leadership,omitempty"`

		NewIdentity string `json:"NewIdentity,omitempty"`

		Partition int64 `json:"Partition,omitempty"`
	}

	dataAO0.Account = m.Account

	dataAO0.ConnectorVersion = m.ConnectorVersion

	dataAO0.DeviceID = m.DeviceID

	dataAO0.DomainGroup = m.DomainGroup

	dataAO0.Evict = m.Evict

	dataAO0.Leadership = m.Leadership

	dataAO0.NewIdentity = m.NewIdentity

	dataAO0.Partition = m.Partition

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset connection control message
func (m *AssetConnectionControlMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLeadership(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var assetConnectionControlMessageTypeLeadershipPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Primary","Secondary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetConnectionControlMessageTypeLeadershipPropEnum = append(assetConnectionControlMessageTypeLeadershipPropEnum, v)
	}
}

// property enum
func (m *AssetConnectionControlMessage) validateLeadershipEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assetConnectionControlMessageTypeLeadershipPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssetConnectionControlMessage) validateLeadership(formats strfmt.Registry) error {

	if swag.IsZero(m.Leadership) { // not required
		return nil
	}

	// value enum
	if err := m.validateLeadershipEnum("Leadership", "body", *m.Leadership); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetConnectionControlMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetConnectionControlMessage) UnmarshalBinary(b []byte) error {
	var res AssetConnectionControlMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}