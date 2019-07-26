// Code generated by go-swagger; DO NOT EDIT.

package equipment_fan

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

// NewPatchEquipmentFansMoidParams creates a new PatchEquipmentFansMoidParams object
// with the default values initialized.
func NewPatchEquipmentFansMoidParams() *PatchEquipmentFansMoidParams {
	var ()
	return &PatchEquipmentFansMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEquipmentFansMoidParamsWithTimeout creates a new PatchEquipmentFansMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchEquipmentFansMoidParamsWithTimeout(timeout time.Duration) *PatchEquipmentFansMoidParams {
	var ()
	return &PatchEquipmentFansMoidParams{

		timeout: timeout,
	}
}

// NewPatchEquipmentFansMoidParamsWithContext creates a new PatchEquipmentFansMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchEquipmentFansMoidParamsWithContext(ctx context.Context) *PatchEquipmentFansMoidParams {
	var ()
	return &PatchEquipmentFansMoidParams{

		Context: ctx,
	}
}

// NewPatchEquipmentFansMoidParamsWithHTTPClient creates a new PatchEquipmentFansMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchEquipmentFansMoidParamsWithHTTPClient(client *http.Client) *PatchEquipmentFansMoidParams {
	var ()
	return &PatchEquipmentFansMoidParams{
		HTTPClient: client,
	}
}

/*PatchEquipmentFansMoidParams contains all the parameters to send to the API endpoint
for the patch equipment fans moid operation typically these are written to a http.Request
*/
type PatchEquipmentFansMoidParams struct {

	/*Body
	  equipmentFan to update

	*/
	Body *models.EquipmentFan
	/*Moid
	  The moid of the equipmentFan instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) WithTimeout(timeout time.Duration) *PatchEquipmentFansMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) WithContext(ctx context.Context) *PatchEquipmentFansMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) WithHTTPClient(client *http.Client) *PatchEquipmentFansMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) WithBody(body *models.EquipmentFan) *PatchEquipmentFansMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) SetBody(body *models.EquipmentFan) {
	o.Body = body
}

// WithMoid adds the moid to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) WithMoid(moid string) *PatchEquipmentFansMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch equipment fans moid params
func (o *PatchEquipmentFansMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEquipmentFansMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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