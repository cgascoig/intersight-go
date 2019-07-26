// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HyperflexHxdpVersion Hyperflex:Hxdp Version
//
// A HyperFlex Data Platform version.
//
// swagger:model hyperflexHxdpVersion
type HyperflexHxdpVersion struct {
	MoBaseMo

	// A collection of references to the [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) MO unsets its reference to this deleted MO.
	//
	AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

	// Corresponding installer image for the HyperFlex Data Platform version.
	//
	InstallerImage *HyperflexInstallerImageRef `json:"InstallerImage,omitempty"`

	// The HyperFlex Data Platform version.
	//
	Version string `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexHxdpVersion) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

		InstallerImage *HyperflexInstallerImageRef `json:"InstallerImage,omitempty"`

		Version string `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AppCatalog = dataAO1.AppCatalog

	m.InstallerImage = dataAO1.InstallerImage

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexHxdpVersion) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

		InstallerImage *HyperflexInstallerImageRef `json:"InstallerImage,omitempty"`

		Version string `json:"Version,omitempty"`
	}

	dataAO1.AppCatalog = m.AppCatalog

	dataAO1.InstallerImage = m.InstallerImage

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex hxdp version
func (m *HyperflexHxdpVersion) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallerImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexHxdpVersion) validateAppCatalog(formats strfmt.Registry) error {

	if swag.IsZero(m.AppCatalog) { // not required
		return nil
	}

	if m.AppCatalog != nil {
		if err := m.AppCatalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AppCatalog")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexHxdpVersion) validateInstallerImage(formats strfmt.Registry) error {

	if swag.IsZero(m.InstallerImage) { // not required
		return nil
	}

	if m.InstallerImage != nil {
		if err := m.InstallerImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InstallerImage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxdpVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxdpVersion) UnmarshalBinary(b []byte) error {
	var res HyperflexHxdpVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}