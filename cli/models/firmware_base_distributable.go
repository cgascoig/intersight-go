// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FirmwareBaseDistributable Firmware:Base Distributable
//
// An image distributed by Cisco.
//
// swagger:model firmwareBaseDistributable
type FirmwareBaseDistributable struct {
	SoftwarerepositoryFile

	// The bundle type of the image, as published on cisco.com.
	//
	// Read Only: true
	BundleType string `json:"BundleType,omitempty"`

	// The unique identifier for an image in a Cisco repository.
	//
	// Read Only: true
	GUID string `json:"Guid,omitempty"`

	// The mdfid of the image provided by cisco.com.
	//
	Mdfid string `json:"Mdfid,omitempty"`

	// The endpoint model for which this firmware image is applicable.
	//
	// Read Only: true
	Model string `json:"Model,omitempty"`

	// The platform type of the image.
	//
	// Read Only: true
	PlatformType string `json:"PlatformType,omitempty"`

	// The build which is recommended by Cisco.
	//
	RecommendedBuild string `json:"RecommendedBuild,omitempty"`

	// The url for the release notes of this image.
	//
	ReleaseNotesURL string `json:"ReleaseNotesUrl,omitempty"`

	// The software type id provided by cisco.com.
	//
	// Read Only: true
	SoftwareTypeID string `json:"SoftwareTypeId,omitempty"`

	// The server models for which this image is applicable.
	//
	SupportedModels []string `json:"SupportedModels"`

	// The vendor or publisher of this file.
	//
	Vendor string `json:"Vendor,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FirmwareBaseDistributable) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SoftwarerepositoryFile
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SoftwarerepositoryFile = aO0

	// AO1
	var dataAO1 struct {
		BundleType string `json:"BundleType,omitempty"`

		GUID string `json:"Guid,omitempty"`

		Mdfid string `json:"Mdfid,omitempty"`

		Model string `json:"Model,omitempty"`

		PlatformType string `json:"PlatformType,omitempty"`

		RecommendedBuild string `json:"RecommendedBuild,omitempty"`

		ReleaseNotesURL string `json:"ReleaseNotesUrl,omitempty"`

		SoftwareTypeID string `json:"SoftwareTypeId,omitempty"`

		SupportedModels []string `json:"SupportedModels"`

		Vendor string `json:"Vendor,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BundleType = dataAO1.BundleType

	m.GUID = dataAO1.GUID

	m.Mdfid = dataAO1.Mdfid

	m.Model = dataAO1.Model

	m.PlatformType = dataAO1.PlatformType

	m.RecommendedBuild = dataAO1.RecommendedBuild

	m.ReleaseNotesURL = dataAO1.ReleaseNotesURL

	m.SoftwareTypeID = dataAO1.SoftwareTypeID

	m.SupportedModels = dataAO1.SupportedModels

	m.Vendor = dataAO1.Vendor

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FirmwareBaseDistributable) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SoftwarerepositoryFile)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		BundleType string `json:"BundleType,omitempty"`

		GUID string `json:"Guid,omitempty"`

		Mdfid string `json:"Mdfid,omitempty"`

		Model string `json:"Model,omitempty"`

		PlatformType string `json:"PlatformType,omitempty"`

		RecommendedBuild string `json:"RecommendedBuild,omitempty"`

		ReleaseNotesURL string `json:"ReleaseNotesUrl,omitempty"`

		SoftwareTypeID string `json:"SoftwareTypeId,omitempty"`

		SupportedModels []string `json:"SupportedModels"`

		Vendor string `json:"Vendor,omitempty"`
	}

	dataAO1.BundleType = m.BundleType

	dataAO1.GUID = m.GUID

	dataAO1.Mdfid = m.Mdfid

	dataAO1.Model = m.Model

	dataAO1.PlatformType = m.PlatformType

	dataAO1.RecommendedBuild = m.RecommendedBuild

	dataAO1.ReleaseNotesURL = m.ReleaseNotesURL

	dataAO1.SoftwareTypeID = m.SoftwareTypeID

	dataAO1.SupportedModels = m.SupportedModels

	dataAO1.Vendor = m.Vendor

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this firmware base distributable
func (m *FirmwareBaseDistributable) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SoftwarerepositoryFile
	if err := m.SoftwarerepositoryFile.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareBaseDistributable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareBaseDistributable) UnmarshalBinary(b []byte) error {
	var res FirmwareBaseDistributable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
