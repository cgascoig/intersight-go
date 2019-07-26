// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FirmwareNfsServer Firmware:Nfs Server
//
// Network share file server.
//
// swagger:model firmwareNfsServer
type FirmwareNfsServer struct {

	// Mount option as configured on the NFS Server. Example:nolock.
	//
	MountOptions string `json:"MountOptions,omitempty"`

	// Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso.
	//
	RemoteFile string `json:"RemoteFile,omitempty"`

	// NFS Server Hostname or IP Address. Example:NFS-server-hostname or 10.10.8.7.
	//
	RemoteIP string `json:"RemoteIp,omitempty"`

	// Directory where the image is stored. Example:/share/subfolder.
	//
	RemoteShare string `json:"RemoteShare,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FirmwareNfsServer) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		MountOptions string `json:"MountOptions,omitempty"`

		RemoteFile string `json:"RemoteFile,omitempty"`

		RemoteIP string `json:"RemoteIp,omitempty"`

		RemoteShare string `json:"RemoteShare,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.MountOptions = dataAO0.MountOptions

	m.RemoteFile = dataAO0.RemoteFile

	m.RemoteIP = dataAO0.RemoteIP

	m.RemoteShare = dataAO0.RemoteShare

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FirmwareNfsServer) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		MountOptions string `json:"MountOptions,omitempty"`

		RemoteFile string `json:"RemoteFile,omitempty"`

		RemoteIP string `json:"RemoteIp,omitempty"`

		RemoteShare string `json:"RemoteShare,omitempty"`
	}

	dataAO0.MountOptions = m.MountOptions

	dataAO0.RemoteFile = m.RemoteFile

	dataAO0.RemoteIP = m.RemoteIP

	dataAO0.RemoteShare = m.RemoteShare

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this firmware nfs server
func (m *FirmwareNfsServer) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareNfsServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareNfsServer) UnmarshalBinary(b []byte) error {
	var res FirmwareNfsServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}