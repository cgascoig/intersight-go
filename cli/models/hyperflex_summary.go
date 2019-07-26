// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HyperflexSummary Hyperflex:Summary
// swagger:model hyperflexSummary
type HyperflexSummary struct {

	// active nodes
	// Read Only: true
	ActiveNodes string `json:"ActiveNodes,omitempty"`

	// address
	// Read Only: true
	Address string `json:"Address,omitempty"`

	// boottime
	// Read Only: true
	Boottime int64 `json:"Boottime,omitempty"`

	// cluster access policy
	// Read Only: true
	ClusterAccessPolicy string `json:"ClusterAccessPolicy,omitempty"`

	// compression savings
	// Read Only: true
	CompressionSavings float64 `json:"CompressionSavings,omitempty"`

	// data replication compliance
	// Read Only: true
	DataReplicationCompliance string `json:"DataReplicationCompliance,omitempty"`

	// data replication factor
	// Read Only: true
	DataReplicationFactor string `json:"DataReplicationFactor,omitempty"`

	// deduplication savings
	// Read Only: true
	DeduplicationSavings float64 `json:"DeduplicationSavings,omitempty"`

	// downtime
	// Read Only: true
	Downtime string `json:"Downtime,omitempty"`

	// free capacity
	// Read Only: true
	FreeCapacity int64 `json:"FreeCapacity,omitempty"`

	// healing info
	// Read Only: true
	HealingInfo *HyperflexStPlatformClusterHealingInfo `json:"HealingInfo,omitempty"`

	// name
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// resiliency details
	// Read Only: true
	ResiliencyDetails interface{} `json:"ResiliencyDetails,omitempty"`

	// resiliency details size
	// Read Only: true
	ResiliencyDetailsSize int64 `json:"ResiliencyDetailsSize,omitempty"`

	// resiliency info
	// Read Only: true
	ResiliencyInfo *HyperflexStPlatformClusterResiliencyInfo `json:"ResiliencyInfo,omitempty"`

	// space status
	// Read Only: true
	SpaceStatus string `json:"SpaceStatus,omitempty"`

	// state
	// Read Only: true
	State string `json:"State,omitempty"`

	// total capacity
	// Read Only: true
	TotalCapacity int64 `json:"TotalCapacity,omitempty"`

	// total savings
	// Read Only: true
	TotalSavings float64 `json:"TotalSavings,omitempty"`

	// uptime
	// Read Only: true
	Uptime string `json:"Uptime,omitempty"`

	// used capacity
	// Read Only: true
	UsedCapacity int64 `json:"UsedCapacity,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexSummary) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ActiveNodes string `json:"ActiveNodes,omitempty"`

		Address string `json:"Address,omitempty"`

		Boottime int64 `json:"Boottime,omitempty"`

		ClusterAccessPolicy string `json:"ClusterAccessPolicy,omitempty"`

		CompressionSavings float64 `json:"CompressionSavings,omitempty"`

		DataReplicationCompliance string `json:"DataReplicationCompliance,omitempty"`

		DataReplicationFactor string `json:"DataReplicationFactor,omitempty"`

		DeduplicationSavings float64 `json:"DeduplicationSavings,omitempty"`

		Downtime string `json:"Downtime,omitempty"`

		FreeCapacity int64 `json:"FreeCapacity,omitempty"`

		HealingInfo *HyperflexStPlatformClusterHealingInfo `json:"HealingInfo,omitempty"`

		Name string `json:"Name,omitempty"`

		ResiliencyDetails interface{} `json:"ResiliencyDetails,omitempty"`

		ResiliencyDetailsSize int64 `json:"ResiliencyDetailsSize,omitempty"`

		ResiliencyInfo *HyperflexStPlatformClusterResiliencyInfo `json:"ResiliencyInfo,omitempty"`

		SpaceStatus string `json:"SpaceStatus,omitempty"`

		State string `json:"State,omitempty"`

		TotalCapacity int64 `json:"TotalCapacity,omitempty"`

		TotalSavings float64 `json:"TotalSavings,omitempty"`

		Uptime string `json:"Uptime,omitempty"`

		UsedCapacity int64 `json:"UsedCapacity,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.ActiveNodes = dataAO0.ActiveNodes

	m.Address = dataAO0.Address

	m.Boottime = dataAO0.Boottime

	m.ClusterAccessPolicy = dataAO0.ClusterAccessPolicy

	m.CompressionSavings = dataAO0.CompressionSavings

	m.DataReplicationCompliance = dataAO0.DataReplicationCompliance

	m.DataReplicationFactor = dataAO0.DataReplicationFactor

	m.DeduplicationSavings = dataAO0.DeduplicationSavings

	m.Downtime = dataAO0.Downtime

	m.FreeCapacity = dataAO0.FreeCapacity

	m.HealingInfo = dataAO0.HealingInfo

	m.Name = dataAO0.Name

	m.ResiliencyDetails = dataAO0.ResiliencyDetails

	m.ResiliencyDetailsSize = dataAO0.ResiliencyDetailsSize

	m.ResiliencyInfo = dataAO0.ResiliencyInfo

	m.SpaceStatus = dataAO0.SpaceStatus

	m.State = dataAO0.State

	m.TotalCapacity = dataAO0.TotalCapacity

	m.TotalSavings = dataAO0.TotalSavings

	m.Uptime = dataAO0.Uptime

	m.UsedCapacity = dataAO0.UsedCapacity

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexSummary) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		ActiveNodes string `json:"ActiveNodes,omitempty"`

		Address string `json:"Address,omitempty"`

		Boottime int64 `json:"Boottime,omitempty"`

		ClusterAccessPolicy string `json:"ClusterAccessPolicy,omitempty"`

		CompressionSavings float64 `json:"CompressionSavings,omitempty"`

		DataReplicationCompliance string `json:"DataReplicationCompliance,omitempty"`

		DataReplicationFactor string `json:"DataReplicationFactor,omitempty"`

		DeduplicationSavings float64 `json:"DeduplicationSavings,omitempty"`

		Downtime string `json:"Downtime,omitempty"`

		FreeCapacity int64 `json:"FreeCapacity,omitempty"`

		HealingInfo *HyperflexStPlatformClusterHealingInfo `json:"HealingInfo,omitempty"`

		Name string `json:"Name,omitempty"`

		ResiliencyDetails interface{} `json:"ResiliencyDetails,omitempty"`

		ResiliencyDetailsSize int64 `json:"ResiliencyDetailsSize,omitempty"`

		ResiliencyInfo *HyperflexStPlatformClusterResiliencyInfo `json:"ResiliencyInfo,omitempty"`

		SpaceStatus string `json:"SpaceStatus,omitempty"`

		State string `json:"State,omitempty"`

		TotalCapacity int64 `json:"TotalCapacity,omitempty"`

		TotalSavings float64 `json:"TotalSavings,omitempty"`

		Uptime string `json:"Uptime,omitempty"`

		UsedCapacity int64 `json:"UsedCapacity,omitempty"`
	}

	dataAO0.ActiveNodes = m.ActiveNodes

	dataAO0.Address = m.Address

	dataAO0.Boottime = m.Boottime

	dataAO0.ClusterAccessPolicy = m.ClusterAccessPolicy

	dataAO0.CompressionSavings = m.CompressionSavings

	dataAO0.DataReplicationCompliance = m.DataReplicationCompliance

	dataAO0.DataReplicationFactor = m.DataReplicationFactor

	dataAO0.DeduplicationSavings = m.DeduplicationSavings

	dataAO0.Downtime = m.Downtime

	dataAO0.FreeCapacity = m.FreeCapacity

	dataAO0.HealingInfo = m.HealingInfo

	dataAO0.Name = m.Name

	dataAO0.ResiliencyDetails = m.ResiliencyDetails

	dataAO0.ResiliencyDetailsSize = m.ResiliencyDetailsSize

	dataAO0.ResiliencyInfo = m.ResiliencyInfo

	dataAO0.SpaceStatus = m.SpaceStatus

	dataAO0.State = m.State

	dataAO0.TotalCapacity = m.TotalCapacity

	dataAO0.TotalSavings = m.TotalSavings

	dataAO0.Uptime = m.Uptime

	dataAO0.UsedCapacity = m.UsedCapacity

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex summary
func (m *HyperflexSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealingInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliencyInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexSummary) validateHealingInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.HealingInfo) { // not required
		return nil
	}

	if m.HealingInfo != nil {
		if err := m.HealingInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HealingInfo")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexSummary) validateResiliencyInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ResiliencyInfo) { // not required
		return nil
	}

	if m.ResiliencyInfo != nil {
		if err := m.ResiliencyInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResiliencyInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexSummary) UnmarshalBinary(b []byte) error {
	var res HyperflexSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
