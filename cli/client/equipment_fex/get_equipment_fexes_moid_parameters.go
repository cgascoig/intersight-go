// Code generated by go-swagger; DO NOT EDIT.

package equipment_fex

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

// NewGetEquipmentFexesMoidParams creates a new GetEquipmentFexesMoidParams object
// with the default values initialized.
func NewGetEquipmentFexesMoidParams() *GetEquipmentFexesMoidParams {
	var ()
	return &GetEquipmentFexesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEquipmentFexesMoidParamsWithTimeout creates a new GetEquipmentFexesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEquipmentFexesMoidParamsWithTimeout(timeout time.Duration) *GetEquipmentFexesMoidParams {
	var ()
	return &GetEquipmentFexesMoidParams{

		timeout: timeout,
	}
}

// NewGetEquipmentFexesMoidParamsWithContext creates a new GetEquipmentFexesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEquipmentFexesMoidParamsWithContext(ctx context.Context) *GetEquipmentFexesMoidParams {
	var ()
	return &GetEquipmentFexesMoidParams{

		Context: ctx,
	}
}

// NewGetEquipmentFexesMoidParamsWithHTTPClient creates a new GetEquipmentFexesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEquipmentFexesMoidParamsWithHTTPClient(client *http.Client) *GetEquipmentFexesMoidParams {
	var ()
	return &GetEquipmentFexesMoidParams{
		HTTPClient: client,
	}
}

/*GetEquipmentFexesMoidParams contains all the parameters to send to the API endpoint
for the get equipment fexes moid operation typically these are written to a http.Request
*/
type GetEquipmentFexesMoidParams struct {

	/*Moid
	  The moid of the equipmentFex instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get equipment fexes moid params
func (o *GetEquipmentFexesMoidParams) WithTimeout(timeout time.Duration) *GetEquipmentFexesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get equipment fexes moid params
func (o *GetEquipmentFexesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get equipment fexes moid params
func (o *GetEquipmentFexesMoidParams) WithContext(ctx context.Context) *GetEquipmentFexesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get equipment fexes moid params
func (o *GetEquipmentFexesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get equipment fexes moid params
func (o *GetEquipmentFexesMoidParams) WithHTTPClient(client *http.Client) *GetEquipmentFexesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get equipment fexes moid params
func (o *GetEquipmentFexesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get equipment fexes moid params
func (o *GetEquipmentFexesMoidParams) WithMoid(moid string) *GetEquipmentFexesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get equipment fexes moid params
func (o *GetEquipmentFexesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetEquipmentFexesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
