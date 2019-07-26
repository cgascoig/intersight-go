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

// StorageInitiator Storage:Initiator
//
// An initiator is the consumer of storage, typically a server with an adapter card in it called a Host Bus Adapter (HBA). The initiator "initiates" a connection over the fabric to one or more ports on storage system target ports.
//
// swagger:model storageInitiator
type StorageInitiator struct {

	// IQN (iSCSI qualified name). Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique name.
	//
	// Read Only: true
	Iqn string `json:"Iqn,omitempty"`

	// Name of the initiator represented in storage array.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// Initiator type, it can be FC or iSCSI.
	//
	// Read Only: true
	// Enum: [FC iSCSI]
	Type string `json:"Type,omitempty"`

	// World wide name, 128 bit address represented in hexa decimal notation. (51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01).
	//
	// Read Only: true
	Wwn string `json:"Wwn,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageInitiator) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Iqn string `json:"Iqn,omitempty"`

		Name string `json:"Name,omitempty"`

		Type string `json:"Type,omitempty"`

		Wwn string `json:"Wwn,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Iqn = dataAO0.Iqn

	m.Name = dataAO0.Name

	m.Type = dataAO0.Type

	m.Wwn = dataAO0.Wwn

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageInitiator) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Iqn string `json:"Iqn,omitempty"`

		Name string `json:"Name,omitempty"`

		Type string `json:"Type,omitempty"`

		Wwn string `json:"Wwn,omitempty"`
	}

	dataAO0.Iqn = m.Iqn

	dataAO0.Name = m.Name

	dataAO0.Type = m.Type

	dataAO0.Wwn = m.Wwn

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage initiator
func (m *StorageInitiator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storageInitiatorTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FC","iSCSI"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageInitiatorTypeTypePropEnum = append(storageInitiatorTypeTypePropEnum, v)
	}
}

// property enum
func (m *StorageInitiator) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageInitiatorTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageInitiator) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageInitiator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageInitiator) UnmarshalBinary(b []byte) error {
	var res StorageInitiator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
