// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StoragePureReplicationSchedule Storage:Pure Replication Schedule
//
// Pure snapshot replication schedule entity.
//
// swagger:model storagePureReplicationSchedule
type StoragePureReplicationSchedule struct {
	StorageReplicationSchedule

	// Total number of snapshots per day to be available on target above and over the specified retention time. PureStorage
	// FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired.
	// In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
	//
	//
	// Read Only: true
	DailyLimit int64 `json:"DailyLimit,omitempty"`

	// Device registration managed object that represents this storage array connection to Intersight.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// Blackout intervals for replication operation in PureStorage FlashArray. System disables replication during these intervals. Blackout periods only apply to scheduled replications. On-demand replications do not observe the blackout period.
	//
	// Read Only: true
	ReplicationBlackoutIntervals []*StorageReplicationBlackout `json:"ReplicationBlackoutIntervals"`

	// Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period.
	//
	// Read Only: true
	SnapshotExpiryTime string `json:"SnapshotExpiryTime,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StoragePureReplicationSchedule) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageReplicationSchedule
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageReplicationSchedule = aO0

	// AO1
	var dataAO1 struct {
		DailyLimit int64 `json:"DailyLimit,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ReplicationBlackoutIntervals []*StorageReplicationBlackout `json:"ReplicationBlackoutIntervals"`

		SnapshotExpiryTime string `json:"SnapshotExpiryTime,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DailyLimit = dataAO1.DailyLimit

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.ReplicationBlackoutIntervals = dataAO1.ReplicationBlackoutIntervals

	m.SnapshotExpiryTime = dataAO1.SnapshotExpiryTime

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StoragePureReplicationSchedule) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageReplicationSchedule)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		DailyLimit int64 `json:"DailyLimit,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ReplicationBlackoutIntervals []*StorageReplicationBlackout `json:"ReplicationBlackoutIntervals"`

		SnapshotExpiryTime string `json:"SnapshotExpiryTime,omitempty"`
	}

	dataAO1.DailyLimit = m.DailyLimit

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.ReplicationBlackoutIntervals = m.ReplicationBlackoutIntervals

	dataAO1.SnapshotExpiryTime = m.SnapshotExpiryTime

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage pure replication schedule
func (m *StoragePureReplicationSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageReplicationSchedule
	if err := m.StorageReplicationSchedule.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicationBlackoutIntervals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePureReplicationSchedule) validateRegisteredDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredDevice) { // not required
		return nil
	}

	if m.RegisteredDevice != nil {
		if err := m.RegisteredDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegisteredDevice")
			}
			return err
		}
	}

	return nil
}

func (m *StoragePureReplicationSchedule) validateReplicationBlackoutIntervals(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplicationBlackoutIntervals) { // not required
		return nil
	}

	for i := 0; i < len(m.ReplicationBlackoutIntervals); i++ {
		if swag.IsZero(m.ReplicationBlackoutIntervals[i]) { // not required
			continue
		}

		if m.ReplicationBlackoutIntervals[i] != nil {
			if err := m.ReplicationBlackoutIntervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ReplicationBlackoutIntervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragePureReplicationSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePureReplicationSchedule) UnmarshalBinary(b []byte) error {
	var res StoragePureReplicationSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
