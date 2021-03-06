// Code generated by go-swagger; DO NOT EDIT.

package equipment_fan_module

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPatchEquipmentFanModulesMoidParams creates a new PatchEquipmentFanModulesMoidParams object
// with the default values initialized.
func NewPatchEquipmentFanModulesMoidParams() *PatchEquipmentFanModulesMoidParams {
	var ()
	return &PatchEquipmentFanModulesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEquipmentFanModulesMoidParamsWithTimeout creates a new PatchEquipmentFanModulesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchEquipmentFanModulesMoidParamsWithTimeout(timeout time.Duration) *PatchEquipmentFanModulesMoidParams {
	var ()
	return &PatchEquipmentFanModulesMoidParams{

		timeout: timeout,
	}
}

// NewPatchEquipmentFanModulesMoidParamsWithContext creates a new PatchEquipmentFanModulesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchEquipmentFanModulesMoidParamsWithContext(ctx context.Context) *PatchEquipmentFanModulesMoidParams {
	var ()
	return &PatchEquipmentFanModulesMoidParams{

		Context: ctx,
	}
}

// NewPatchEquipmentFanModulesMoidParamsWithHTTPClient creates a new PatchEquipmentFanModulesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchEquipmentFanModulesMoidParamsWithHTTPClient(client *http.Client) *PatchEquipmentFanModulesMoidParams {
	var ()
	return &PatchEquipmentFanModulesMoidParams{
		HTTPClient: client,
	}
}

/*PatchEquipmentFanModulesMoidParams contains all the parameters to send to the API endpoint
for the patch equipment fan modules moid operation typically these are written to a http.Request
*/
type PatchEquipmentFanModulesMoidParams struct {

	/*Body
	  equipmentFanModule to update

	*/
	Body *models.EquipmentFanModule
	/*Moid
	  The moid of the equipmentFanModule instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) WithTimeout(timeout time.Duration) *PatchEquipmentFanModulesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) WithContext(ctx context.Context) *PatchEquipmentFanModulesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) WithHTTPClient(client *http.Client) *PatchEquipmentFanModulesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) WithBody(body *models.EquipmentFanModule) *PatchEquipmentFanModulesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) SetBody(body *models.EquipmentFanModule) {
	o.Body = body
}

// WithMoid adds the moid to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) WithMoid(moid string) *PatchEquipmentFanModulesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch equipment fan modules moid params
func (o *PatchEquipmentFanModulesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEquipmentFanModulesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param moid
	if err := r.SetPathParam("moid", o.Moid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
