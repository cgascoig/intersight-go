// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HclExemptedCatalog Hcl:Exempted Catalog
//
// Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.
//
// swagger:model hclExemptedCatalog
type HclExemptedCatalog struct {
	MoBaseMo

	// The catalog under which the definition is present.
	//
	Catalog *HclCatalogRef `json:"Catalog,omitempty"`

	// Reason for the exemption.
	//
	Comments string `json:"Comments,omitempty"`

	// Specifies whether the exemption is enabled or disabled.
	//
	IsMarkedForDelete bool `json:"IsMarkedForDelete,omitempty"`

	// A unique descriptive name of the exemption.
	//
	Name string `json:"Name,omitempty"`

	// Vendor of the Operating System.
	//
	OsVendor string `json:"OsVendor,omitempty"`

	// Version of the Operating system.
	//
	OsVersion string `json:"OsVersion,omitempty"`

	// Name of the processor supported for the server.
	//
	ProcessorName string `json:"ProcessorName,omitempty"`

	// Models of the product/adapter.
	//
	ProductModels []string `json:"ProductModels"`

	// Type of the product/adapter say PT for Pass Through controllers.
	//
	ProductType string `json:"ProductType,omitempty"`

	// Three part ID representing the server model as returned by UCSM/CIMC XML APIs.
	//
	ServerPid string `json:"ServerPid,omitempty"`

	// Version of the UCS software.
	//
	UcsVersion string `json:"UcsVersion,omitempty"`

	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
	//
	VersionType string `json:"VersionType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HclExemptedCatalog) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Catalog *HclCatalogRef `json:"Catalog,omitempty"`

		Comments string `json:"Comments,omitempty"`

		IsMarkedForDelete bool `json:"IsMarkedForDelete,omitempty"`

		Name string `json:"Name,omitempty"`

		OsVendor string `json:"OsVendor,omitempty"`

		OsVersion string `json:"OsVersion,omitempty"`

		ProcessorName string `json:"ProcessorName,omitempty"`

		ProductModels []string `json:"ProductModels"`

		ProductType string `json:"ProductType,omitempty"`

		ServerPid string `json:"ServerPid,omitempty"`

		UcsVersion string `json:"UcsVersion,omitempty"`

		VersionType string `json:"VersionType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Catalog = dataAO1.Catalog

	m.Comments = dataAO1.Comments

	m.IsMarkedForDelete = dataAO1.IsMarkedForDelete

	m.Name = dataAO1.Name

	m.OsVendor = dataAO1.OsVendor

	m.OsVersion = dataAO1.OsVersion

	m.ProcessorName = dataAO1.ProcessorName

	m.ProductModels = dataAO1.ProductModels

	m.ProductType = dataAO1.ProductType

	m.ServerPid = dataAO1.ServerPid

	m.UcsVersion = dataAO1.UcsVersion

	m.VersionType = dataAO1.VersionType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HclExemptedCatalog) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Catalog *HclCatalogRef `json:"Catalog,omitempty"`

		Comments string `json:"Comments,omitempty"`

		IsMarkedForDelete bool `json:"IsMarkedForDelete,omitempty"`

		Name string `json:"Name,omitempty"`

		OsVendor string `json:"OsVendor,omitempty"`

		OsVersion string `json:"OsVersion,omitempty"`

		ProcessorName string `json:"ProcessorName,omitempty"`

		ProductModels []string `json:"ProductModels"`

		ProductType string `json:"ProductType,omitempty"`

		ServerPid string `json:"ServerPid,omitempty"`

		UcsVersion string `json:"UcsVersion,omitempty"`

		VersionType string `json:"VersionType,omitempty"`
	}

	dataAO1.Catalog = m.Catalog

	dataAO1.Comments = m.Comments

	dataAO1.IsMarkedForDelete = m.IsMarkedForDelete

	dataAO1.Name = m.Name

	dataAO1.OsVendor = m.OsVendor

	dataAO1.OsVersion = m.OsVersion

	dataAO1.ProcessorName = m.ProcessorName

	dataAO1.ProductModels = m.ProductModels

	dataAO1.ProductType = m.ProductType

	dataAO1.ServerPid = m.ServerPid

	dataAO1.UcsVersion = m.UcsVersion

	dataAO1.VersionType = m.VersionType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hcl exempted catalog
func (m *HclExemptedCatalog) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HclExemptedCatalog) validateCatalog(formats strfmt.Registry) error {

	if swag.IsZero(m.Catalog) { // not required
		return nil
	}

	if m.Catalog != nil {
		if err := m.Catalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Catalog")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HclExemptedCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HclExemptedCatalog) UnmarshalBinary(b []byte) error {
	var res HclExemptedCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
