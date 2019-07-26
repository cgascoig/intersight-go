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

// HyperflexInitiateHxClusterUpgrade Hyperflex:Initiate Hx Cluster Upgrade
//
// Initiate HyperFlex upgrade on a set of nodes, provided as an input.
//
// swagger:model hyperflexInitiateHxClusterUpgrade
type HyperflexInitiateHxClusterUpgrade struct {
	MoBaseMo

	// Relationship to the associated Intersight account.
	//
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// Workflow action to perform to either start or cancel an ongoing HyperFlex cluster upgrade workflow.
	//
	// Enum: [Start Cancel]
	Action *string `json:"Action,omitempty"`

	// Relationship to the associated HyperFlex cluster.
	//
	HyperflexCluster *HyperflexClusterRef `json:"HyperflexCluster,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// Information about the package bundle, which will be used to upgrade a HyperFlex cluster.
	//
	UpgradePackageDetails *HyperflexUpgradePackageDetails `json:"UpgradePackageDetails,omitempty"`

	// File name of the static workflow, which has the task definitions to perform HyperFlex cluster upgrade. The file name has to be the absolute path to the workflow file.
	//
	//
	WorkflowFilename string `json:"WorkflowFilename,omitempty"`

	// Relationship to the associated workflow.
	//
	// Read Only: true
	WorkflowInfo *WorkflowWorkflowInfoRef `json:"WorkflowInfo,omitempty"`

	// Input parameters to start the HyperFlex upgrade workflow on a given set of nodes. At a minimum, the input parameters must contain upgrade type, upgrade action and set of nodes.
	//
	//
	WorkflowInputParams *HyperflexWorkflowInputParams `json:"WorkflowInputParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexInitiateHxClusterUpgrade) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Action *string `json:"Action,omitempty"`

		HyperflexCluster *HyperflexClusterRef `json:"HyperflexCluster,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		UpgradePackageDetails *HyperflexUpgradePackageDetails `json:"UpgradePackageDetails,omitempty"`

		WorkflowFilename string `json:"WorkflowFilename,omitempty"`

		WorkflowInfo *WorkflowWorkflowInfoRef `json:"WorkflowInfo,omitempty"`

		WorkflowInputParams *HyperflexWorkflowInputParams `json:"WorkflowInputParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.Action = dataAO1.Action

	m.HyperflexCluster = dataAO1.HyperflexCluster

	m.Organization = dataAO1.Organization

	m.UpgradePackageDetails = dataAO1.UpgradePackageDetails

	m.WorkflowFilename = dataAO1.WorkflowFilename

	m.WorkflowInfo = dataAO1.WorkflowInfo

	m.WorkflowInputParams = dataAO1.WorkflowInputParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexInitiateHxClusterUpgrade) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Action *string `json:"Action,omitempty"`

		HyperflexCluster *HyperflexClusterRef `json:"HyperflexCluster,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		UpgradePackageDetails *HyperflexUpgradePackageDetails `json:"UpgradePackageDetails,omitempty"`

		WorkflowFilename string `json:"WorkflowFilename,omitempty"`

		WorkflowInfo *WorkflowWorkflowInfoRef `json:"WorkflowInfo,omitempty"`

		WorkflowInputParams *HyperflexWorkflowInputParams `json:"WorkflowInputParams,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.Action = m.Action

	dataAO1.HyperflexCluster = m.HyperflexCluster

	dataAO1.Organization = m.Organization

	dataAO1.UpgradePackageDetails = m.UpgradePackageDetails

	dataAO1.WorkflowFilename = m.WorkflowFilename

	dataAO1.WorkflowInfo = m.WorkflowInfo

	dataAO1.WorkflowInputParams = m.WorkflowInputParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex initiate hx cluster upgrade
func (m *HyperflexInitiateHxClusterUpgrade) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHyperflexCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradePackageDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowInputParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexInitiateHxClusterUpgrade) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

var hyperflexInitiateHxClusterUpgradeTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Start","Cancel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexInitiateHxClusterUpgradeTypeActionPropEnum = append(hyperflexInitiateHxClusterUpgradeTypeActionPropEnum, v)
	}
}

// property enum
func (m *HyperflexInitiateHxClusterUpgrade) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexInitiateHxClusterUpgradeTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexInitiateHxClusterUpgrade) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("Action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *HyperflexInitiateHxClusterUpgrade) validateHyperflexCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.HyperflexCluster) { // not required
		return nil
	}

	if m.HyperflexCluster != nil {
		if err := m.HyperflexCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HyperflexCluster")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexInitiateHxClusterUpgrade) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexInitiateHxClusterUpgrade) validateUpgradePackageDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.UpgradePackageDetails) { // not required
		return nil
	}

	if m.UpgradePackageDetails != nil {
		if err := m.UpgradePackageDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UpgradePackageDetails")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexInitiateHxClusterUpgrade) validateWorkflowInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowInfo) { // not required
		return nil
	}

	if m.WorkflowInfo != nil {
		if err := m.WorkflowInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WorkflowInfo")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexInitiateHxClusterUpgrade) validateWorkflowInputParams(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowInputParams) { // not required
		return nil
	}

	if m.WorkflowInputParams != nil {
		if err := m.WorkflowInputParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WorkflowInputParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexInitiateHxClusterUpgrade) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexInitiateHxClusterUpgrade) UnmarshalBinary(b []byte) error {
	var res HyperflexInitiateHxClusterUpgrade
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
