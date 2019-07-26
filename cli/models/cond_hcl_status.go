// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CondHclStatus Cond:Hcl Status
//
// The HCL status of a managed object after we have validated the managed object components' firmware and drivers against the HCL.
//
// swagger:model condHclStatus
type CondHclStatus struct {
	MoBaseMo

	// The overall status for the components found in the HCL. This will provide the HCL validation status for all the components. It can be one of the following. "Validated" - all the components hardware/software profiles are listed in the HCL. "Not-Listed" - one or more components hardware/software profiles are not listed in the HCL "Incomplete" - the components are not evaluated as the server's software/hardware profiles are not listed in the HCL.
	//
	// Enum: [Incomplete Not-Found Not-Listed Validated Not-Evaluated]
	ComponentStatus *string `json:"ComponentStatus,omitempty"`

	// The collection of all the detailed related components for which we performed HCL validation.
	//
	// Read Only: true
	Details []*CondHclStatusDetailRef `json:"Details"`

	// The server model, processor and firmware are considered as part of the hardware profile for the server. This will provide the HCL validation status for the hardware profile. For the failure reason see the serverReason property. The status can be one of the following "Validated" - The server model, processor and firmware combination is listed in the HCL "Not-Listed" - The server model, processor and firmware combination is not listed in the HCL.
	//
	// Enum: [Incomplete Not-Found Not-Listed Validated Not-Evaluated]
	HardwareStatus *string `json:"HardwareStatus,omitempty"`

	// The current CIMC version for the server normalized for querying HCL data. It is empty if we are missing this information.
	//
	HclFirmwareVersion string `json:"HclFirmwareVersion,omitempty"`

	// The managed object's model to validate normalized for querying HCL data. It is empty if we are missing this information.
	//
	HclModel string `json:"HclModel,omitempty"`

	// The OS Vendor for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.
	//
	HclOsVendor string `json:"HclOsVendor,omitempty"`

	// The OS Version for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.
	//
	HclOsVersion string `json:"HclOsVersion,omitempty"`

	// The managed object's processor to validate if applicable normalized for querying HCL data. It is empty if we are missing this information.
	//
	HclProcessor string `json:"HclProcessor,omitempty"`

	// The current CIMC version for the server as received from inventory. It is empty if we are missing this information.
	//
	InvFirmwareVersion string `json:"InvFirmwareVersion,omitempty"`

	// The managed object's model to validate as received from the inventory. It is empty if we are missing this information.
	//
	InvModel string `json:"InvModel,omitempty"`

	// The OS Vendor for the managed object to validate as received from inventory. It is empty if we are missing this information.
	//
	InvOsVendor string `json:"InvOsVendor,omitempty"`

	// The OS Version for the managed object to validate as received from inventory. It is empty if we are missing this information.
	//
	InvOsVersion string `json:"InvOsVersion,omitempty"`

	// The managed object's processor to validate if applicable as received from inventory. It is empty if we are missing this information.
	//
	InvProcessor string `json:"InvProcessor,omitempty"`

	// The managed object relationship for this HCLStatus.
	//
	// Read Only: true
	ManagedObject *InventoryBaseRef `json:"ManagedObject,omitempty"`

	// The reason for the HCL status. It will be one of the following "Missing-Os-Info" - we are missing os information in the inventory from the device connector "Incompatible-Components" - we have 1 or more components with "Not-Validated" status "Compatible" - all the components have "Validated" status.
	//
	// Enum: [Missing-Os-Info Incompatible-Components Compatible]
	Reason *string `json:"Reason,omitempty"`

	// The Relationship to the registered device. We need this in order to correctly set permissions during device claim.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// The reason generated by the server's HCL validation. For HCL the evaluation can be seen in three logical stages 1. Evaluate the server's hardware status 2. Evaluate the server's software status 3. Evaluate the server's components (each component has its own hardware/software evaluation) The evaluation does not proceed to the next stage until the previous stage is evaluated. Therefore there can be only one validation reason. "Incompatible-Server" - the server model is not listed in the HCL. "Incompatible-Processor" - the server model and processor combination is not listed in HCL. "Incompatible-Firmware" - the server model, processor and server firmware is not listed in HCL. "Missing-Os-Info" - the os vendor and version is not listed in HCL with the HW profile. "Incompatible-Os-Info" - the os vendor and version is not listed in HCL with the HW profile. "Incompatible-Components" - there is one or more components with "Not-Validated" status "Service-Unavailable" - HCL data service is unavailable at the moment (try again later). "Compatible" - the server and all its components are validated.
	//
	// Enum: [Missing-Os-Driver-Info Incompatible-Server Incompatible-Processor Incompatible-Os-Info Incompatible-Firmware Service-Unavailable Service-Error Incompatible-Components Compatible]
	ServerReason *string `json:"ServerReason,omitempty"`

	// The OS vendor and version are considered part of the software profile for the server. This will provide the HCL validation status for the software profile. For the failure reason see the serverReason property. The status can be be one of the following "Validated" - The os vendor/version is listed in the HCL for the server model, processor and firmware "Not-Listed" - The os vendor/version is not listed in the HCL for the server model, processor and firmware "Incomplete" - The inventory is missing os vendor/version and HCL validation was not performed.
	//
	// Enum: [Incomplete Not-Found Not-Listed Validated Not-Evaluated]
	SoftwareStatus *string `json:"SoftwareStatus,omitempty"`

	// The HCL compatibility status of the managed object. The status can be one of the following "Incomplete" - there is no enough information to evaluate against the HCL data "Validated" - all components have been evaluated against the HCL and they all have "Validated" status "Not-Listed" - all components have been evaluated against the HCL and one or more have "Not-Listed" status.
	//
	// Enum: [Incomplete Not-Found Not-Listed Validated Not-Evaluated]
	Status *string `json:"Status,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CondHclStatus) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		ComponentStatus *string `json:"ComponentStatus,omitempty"`

		Details []*CondHclStatusDetailRef `json:"Details"`

		HardwareStatus *string `json:"HardwareStatus,omitempty"`

		HclFirmwareVersion string `json:"HclFirmwareVersion,omitempty"`

		HclModel string `json:"HclModel,omitempty"`

		HclOsVendor string `json:"HclOsVendor,omitempty"`

		HclOsVersion string `json:"HclOsVersion,omitempty"`

		HclProcessor string `json:"HclProcessor,omitempty"`

		InvFirmwareVersion string `json:"InvFirmwareVersion,omitempty"`

		InvModel string `json:"InvModel,omitempty"`

		InvOsVendor string `json:"InvOsVendor,omitempty"`

		InvOsVersion string `json:"InvOsVersion,omitempty"`

		InvProcessor string `json:"InvProcessor,omitempty"`

		ManagedObject *InventoryBaseRef `json:"ManagedObject,omitempty"`

		Reason *string `json:"Reason,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ServerReason *string `json:"ServerReason,omitempty"`

		SoftwareStatus *string `json:"SoftwareStatus,omitempty"`

		Status *string `json:"Status,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ComponentStatus = dataAO1.ComponentStatus

	m.Details = dataAO1.Details

	m.HardwareStatus = dataAO1.HardwareStatus

	m.HclFirmwareVersion = dataAO1.HclFirmwareVersion

	m.HclModel = dataAO1.HclModel

	m.HclOsVendor = dataAO1.HclOsVendor

	m.HclOsVersion = dataAO1.HclOsVersion

	m.HclProcessor = dataAO1.HclProcessor

	m.InvFirmwareVersion = dataAO1.InvFirmwareVersion

	m.InvModel = dataAO1.InvModel

	m.InvOsVendor = dataAO1.InvOsVendor

	m.InvOsVersion = dataAO1.InvOsVersion

	m.InvProcessor = dataAO1.InvProcessor

	m.ManagedObject = dataAO1.ManagedObject

	m.Reason = dataAO1.Reason

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.ServerReason = dataAO1.ServerReason

	m.SoftwareStatus = dataAO1.SoftwareStatus

	m.Status = dataAO1.Status

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CondHclStatus) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ComponentStatus *string `json:"ComponentStatus,omitempty"`

		Details []*CondHclStatusDetailRef `json:"Details"`

		HardwareStatus *string `json:"HardwareStatus,omitempty"`

		HclFirmwareVersion string `json:"HclFirmwareVersion,omitempty"`

		HclModel string `json:"HclModel,omitempty"`

		HclOsVendor string `json:"HclOsVendor,omitempty"`

		HclOsVersion string `json:"HclOsVersion,omitempty"`

		HclProcessor string `json:"HclProcessor,omitempty"`

		InvFirmwareVersion string `json:"InvFirmwareVersion,omitempty"`

		InvModel string `json:"InvModel,omitempty"`

		InvOsVendor string `json:"InvOsVendor,omitempty"`

		InvOsVersion string `json:"InvOsVersion,omitempty"`

		InvProcessor string `json:"InvProcessor,omitempty"`

		ManagedObject *InventoryBaseRef `json:"ManagedObject,omitempty"`

		Reason *string `json:"Reason,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ServerReason *string `json:"ServerReason,omitempty"`

		SoftwareStatus *string `json:"SoftwareStatus,omitempty"`

		Status *string `json:"Status,omitempty"`
	}

	dataAO1.ComponentStatus = m.ComponentStatus

	dataAO1.Details = m.Details

	dataAO1.HardwareStatus = m.HardwareStatus

	dataAO1.HclFirmwareVersion = m.HclFirmwareVersion

	dataAO1.HclModel = m.HclModel

	dataAO1.HclOsVendor = m.HclOsVendor

	dataAO1.HclOsVersion = m.HclOsVersion

	dataAO1.HclProcessor = m.HclProcessor

	dataAO1.InvFirmwareVersion = m.InvFirmwareVersion

	dataAO1.InvModel = m.InvModel

	dataAO1.InvOsVendor = m.InvOsVendor

	dataAO1.InvOsVersion = m.InvOsVersion

	dataAO1.InvProcessor = m.InvProcessor

	dataAO1.ManagedObject = m.ManagedObject

	dataAO1.Reason = m.Reason

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.ServerReason = m.ServerReason

	dataAO1.SoftwareStatus = m.SoftwareStatus

	dataAO1.Status = m.Status

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this cond hcl status
func (m *CondHclStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardwareStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var condHclStatusTypeComponentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Incomplete","Not-Found","Not-Listed","Validated","Not-Evaluated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		condHclStatusTypeComponentStatusPropEnum = append(condHclStatusTypeComponentStatusPropEnum, v)
	}
}

// property enum
func (m *CondHclStatus) validateComponentStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, condHclStatusTypeComponentStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CondHclStatus) validateComponentStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ComponentStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateComponentStatusEnum("ComponentStatus", "body", *m.ComponentStatus); err != nil {
		return err
	}

	return nil
}

func (m *CondHclStatus) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	for i := 0; i < len(m.Details); i++ {
		if swag.IsZero(m.Details[i]) { // not required
			continue
		}

		if m.Details[i] != nil {
			if err := m.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var condHclStatusTypeHardwareStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Incomplete","Not-Found","Not-Listed","Validated","Not-Evaluated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		condHclStatusTypeHardwareStatusPropEnum = append(condHclStatusTypeHardwareStatusPropEnum, v)
	}
}

// property enum
func (m *CondHclStatus) validateHardwareStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, condHclStatusTypeHardwareStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CondHclStatus) validateHardwareStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.HardwareStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateHardwareStatusEnum("HardwareStatus", "body", *m.HardwareStatus); err != nil {
		return err
	}

	return nil
}

func (m *CondHclStatus) validateManagedObject(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagedObject) { // not required
		return nil
	}

	if m.ManagedObject != nil {
		if err := m.ManagedObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ManagedObject")
			}
			return err
		}
	}

	return nil
}

var condHclStatusTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Missing-Os-Info","Incompatible-Components","Compatible"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		condHclStatusTypeReasonPropEnum = append(condHclStatusTypeReasonPropEnum, v)
	}
}

// property enum
func (m *CondHclStatus) validateReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, condHclStatusTypeReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CondHclStatus) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	// value enum
	if err := m.validateReasonEnum("Reason", "body", *m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *CondHclStatus) validateRegisteredDevice(formats strfmt.Registry) error {

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

var condHclStatusTypeServerReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Missing-Os-Driver-Info","Incompatible-Server","Incompatible-Processor","Incompatible-Os-Info","Incompatible-Firmware","Service-Unavailable","Service-Error","Incompatible-Components","Compatible"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		condHclStatusTypeServerReasonPropEnum = append(condHclStatusTypeServerReasonPropEnum, v)
	}
}

// property enum
func (m *CondHclStatus) validateServerReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, condHclStatusTypeServerReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CondHclStatus) validateServerReason(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateServerReasonEnum("ServerReason", "body", *m.ServerReason); err != nil {
		return err
	}

	return nil
}

var condHclStatusTypeSoftwareStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Incomplete","Not-Found","Not-Listed","Validated","Not-Evaluated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		condHclStatusTypeSoftwareStatusPropEnum = append(condHclStatusTypeSoftwareStatusPropEnum, v)
	}
}

// property enum
func (m *CondHclStatus) validateSoftwareStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, condHclStatusTypeSoftwareStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CondHclStatus) validateSoftwareStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SoftwareStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateSoftwareStatusEnum("SoftwareStatus", "body", *m.SoftwareStatus); err != nil {
		return err
	}

	return nil
}

var condHclStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Incomplete","Not-Found","Not-Listed","Validated","Not-Evaluated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		condHclStatusTypeStatusPropEnum = append(condHclStatusTypeStatusPropEnum, v)
	}
}

// property enum
func (m *CondHclStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, condHclStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CondHclStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CondHclStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CondHclStatus) UnmarshalBinary(b []byte) error {
	var res CondHclStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
