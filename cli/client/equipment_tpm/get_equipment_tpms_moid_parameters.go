// Code generated by go-swagger; DO NOT EDIT.

package equipment_tpm

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
)

// NewGetEquipmentTpmsMoidParams creates a new GetEquipmentTpmsMoidParams object
// with the default values initialized.
func NewGetEquipmentTpmsMoidParams() *GetEquipmentTpmsMoidParams {
	var ()
	return &GetEquipmentTpmsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEquipmentTpmsMoidParamsWithTimeout creates a new GetEquipmentTpmsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEquipmentTpmsMoidParamsWithTimeout(timeout time.Duration) *GetEquipmentTpmsMoidParams {
	var ()
	return &GetEquipmentTpmsMoidParams{

		timeout: timeout,
	}
}

// NewGetEquipmentTpmsMoidParamsWithContext creates a new GetEquipmentTpmsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEquipmentTpmsMoidParamsWithContext(ctx context.Context) *GetEquipmentTpmsMoidParams {
	var ()
	return &GetEquipmentTpmsMoidParams{

		Context: ctx,
	}
}

// NewGetEquipmentTpmsMoidParamsWithHTTPClient creates a new GetEquipmentTpmsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEquipmentTpmsMoidParamsWithHTTPClient(client *http.Client) *GetEquipmentTpmsMoidParams {
	var ()
	return &GetEquipmentTpmsMoidParams{
		HTTPClient: client,
	}
}

/*GetEquipmentTpmsMoidParams contains all the parameters to send to the API endpoint
for the get equipment tpms moid operation typically these are written to a http.Request
*/
type GetEquipmentTpmsMoidParams struct {

	/*Moid
	  The moid of the equipmentTpm instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get equipment tpms moid params
func (o *GetEquipmentTpmsMoidParams) WithTimeout(timeout time.Duration) *GetEquipmentTpmsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get equipment tpms moid params
func (o *GetEquipmentTpmsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get equipment tpms moid params
func (o *GetEquipmentTpmsMoidParams) WithContext(ctx context.Context) *GetEquipmentTpmsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get equipment tpms moid params
func (o *GetEquipmentTpmsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get equipment tpms moid params
func (o *GetEquipmentTpmsMoidParams) WithHTTPClient(client *http.Client) *GetEquipmentTpmsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get equipment tpms moid params
func (o *GetEquipmentTpmsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get equipment tpms moid params
func (o *GetEquipmentTpmsMoidParams) WithMoid(moid string) *GetEquipmentTpmsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get equipment tpms moid params
func (o *GetEquipmentTpmsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetEquipmentTpmsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param moid
	if err := r.SetPathParam("moid", o.Moid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}