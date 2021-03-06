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
)

// NewGetEquipmentFansMoidParams creates a new GetEquipmentFansMoidParams object
// with the default values initialized.
func NewGetEquipmentFansMoidParams() *GetEquipmentFansMoidParams {
	var ()
	return &GetEquipmentFansMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEquipmentFansMoidParamsWithTimeout creates a new GetEquipmentFansMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEquipmentFansMoidParamsWithTimeout(timeout time.Duration) *GetEquipmentFansMoidParams {
	var ()
	return &GetEquipmentFansMoidParams{

		timeout: timeout,
	}
}

// NewGetEquipmentFansMoidParamsWithContext creates a new GetEquipmentFansMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEquipmentFansMoidParamsWithContext(ctx context.Context) *GetEquipmentFansMoidParams {
	var ()
	return &GetEquipmentFansMoidParams{

		Context: ctx,
	}
}

// NewGetEquipmentFansMoidParamsWithHTTPClient creates a new GetEquipmentFansMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEquipmentFansMoidParamsWithHTTPClient(client *http.Client) *GetEquipmentFansMoidParams {
	var ()
	return &GetEquipmentFansMoidParams{
		HTTPClient: client,
	}
}

/*GetEquipmentFansMoidParams contains all the parameters to send to the API endpoint
for the get equipment fans moid operation typically these are written to a http.Request
*/
type GetEquipmentFansMoidParams struct {

	/*Moid
	  The moid of the equipmentFan instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get equipment fans moid params
func (o *GetEquipmentFansMoidParams) WithTimeout(timeout time.Duration) *GetEquipmentFansMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get equipment fans moid params
func (o *GetEquipmentFansMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get equipment fans moid params
func (o *GetEquipmentFansMoidParams) WithContext(ctx context.Context) *GetEquipmentFansMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get equipment fans moid params
func (o *GetEquipmentFansMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get equipment fans moid params
func (o *GetEquipmentFansMoidParams) WithHTTPClient(client *http.Client) *GetEquipmentFansMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get equipment fans moid params
func (o *GetEquipmentFansMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get equipment fans moid params
func (o *GetEquipmentFansMoidParams) WithMoid(moid string) *GetEquipmentFansMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get equipment fans moid params
func (o *GetEquipmentFansMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetEquipmentFansMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
