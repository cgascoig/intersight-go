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

// OnpremImagePackage Onprem:Image Package
//
// ImagePackage encapsulates a software image package. ImagePackage can be
// a docker image, a UI web image, an endpoint (e.g. UCSM) image, a device
// connector image or an ansible scripts package.
//
// swagger:model onpremImagePackage
type OnpremImagePackage struct {

	// Optional file path of the image package.
	//
	// Read Only: true
	FilePath string `json:"FilePath,omitempty"`

	// Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.
	//
	// Read Only: true
	FileSha string `json:"FileSha,omitempty"`

	// Image file size in bytes.
	//
	// Read Only: true
	FileSize int64 `json:"FileSize,omitempty"`

	// Image file's last modified date and time.
	//
	// Read Only: true
	// Format: date-time
	FileTime strfmt.DateTime `json:"FileTime,omitempty"`

	// Filename of the image package.
	//
	// Read Only: true
	Filename string `json:"Filename,omitempty"`

	// Name of the software image package.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// Image package type (e.g. service, system etc.).
	//
	// Read Only: true
	PackageType string `json:"PackageType,omitempty"`

	// Image package version string.
	//
	// Read Only: true
	Version string `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OnpremImagePackage) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		FilePath string `json:"FilePath,omitempty"`

		FileSha string `json:"FileSha,omitempty"`

		FileSize int64 `json:"FileSize,omitempty"`

		FileTime strfmt.DateTime `json:"FileTime,omitempty"`

		Filename string `json:"Filename,omitempty"`

		Name string `json:"Name,omitempty"`

		PackageType string `json:"PackageType,omitempty"`

		Version string `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.FilePath = dataAO0.FilePath

	m.FileSha = dataAO0.FileSha

	m.FileSize = dataAO0.FileSize

	m.FileTime = dataAO0.FileTime

	m.Filename = dataAO0.Filename

	m.Name = dataAO0.Name

	m.PackageType = dataAO0.PackageType

	m.Version = dataAO0.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OnpremImagePackage) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		FilePath string `json:"FilePath,omitempty"`

		FileSha string `json:"FileSha,omitempty"`

		FileSize int64 `json:"FileSize,omitempty"`

		FileTime strfmt.DateTime `json:"FileTime,omitempty"`

		Filename string `json:"Filename,omitempty"`

		Name string `json:"Name,omitempty"`

		PackageType string `json:"PackageType,omitempty"`

		Version string `json:"Version,omitempty"`
	}

	dataAO0.FilePath = m.FilePath

	dataAO0.FileSha = m.FileSha

	dataAO0.FileSize = m.FileSize

	dataAO0.FileTime = m.FileTime

	dataAO0.Filename = m.Filename

	dataAO0.Name = m.Name

	dataAO0.PackageType = m.PackageType

	dataAO0.Version = m.Version

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this onprem image package
func (m *OnpremImagePackage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnpremImagePackage) validateFileTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FileTime) { // not required
		return nil
	}

	if err := validate.FormatOf("FileTime", "body", "date-time", m.FileTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnpremImagePackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnpremImagePackage) UnmarshalBinary(b []byte) error {
	var res OnpremImagePackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}