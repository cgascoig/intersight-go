// Code generated by go-swagger; DO NOT EDIT.

package equipment_psu

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

// NewPostEquipmentPsusMoidParams creates a new PostEquipmentPsusMoidParams object
// with the default values initialized.
func NewPostEquipmentPsusMoidParams() *PostEquipmentPsusMoidParams {
	var ()
	return &PostEquipmentPsusMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEquipmentPsusMoidParamsWithTimeout creates a new PostEquipmentPsusMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEquipmentPsusMoidParamsWithTimeout(timeout time.Duration) *PostEquipmentPsusMoidParams {
	var ()
	return &PostEquipmentPsusMoidParams{

		timeout: timeout,
	}
}

// NewPostEquipmentPsusMoidParamsWithContext creates a new PostEquipmentPsusMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEquipmentPsusMoidParamsWithContext(ctx context.Context) *PostEquipmentPsusMoidParams {
	var ()
	return &PostEquipmentPsusMoidParams{

		Context: ctx,
	}
}

// NewPostEquipmentPsusMoidParamsWithHTTPClient creates a new PostEquipmentPsusMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEquipmentPsusMoidParamsWithHTTPClient(client *http.Client) *PostEquipmentPsusMoidParams {
	var ()
	return &PostEquipmentPsusMoidParams{
		HTTPClient: client,
	}
}

/*PostEquipmentPsusMoidParams contains all the parameters to send to the API endpoint
for the post equipment psus moid operation typically these are written to a http.Request
*/
type PostEquipmentPsusMoidParams struct {

	/*Body
	  equipmentPsu to update

	*/
	Body *models.EquipmentPsu
	/*Moid
	  The moid of the equipmentPsu instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) WithTimeout(timeout time.Duration) *PostEquipmentPsusMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) WithContext(ctx context.Context) *PostEquipmentPsusMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) WithHTTPClient(client *http.Client) *PostEquipmentPsusMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) WithBody(body *models.EquipmentPsu) *PostEquipmentPsusMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) SetBody(body *models.EquipmentPsu) {
	o.Body = body
}

// WithMoid adds the moid to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) WithMoid(moid string) *PostEquipmentPsusMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post equipment psus moid params
func (o *PostEquipmentPsusMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostEquipmentPsusMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
