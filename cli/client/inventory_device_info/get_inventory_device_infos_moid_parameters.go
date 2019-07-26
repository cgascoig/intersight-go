// Code generated by go-swagger; DO NOT EDIT.

package inventory_device_info

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

// NewGetInventoryDeviceInfosMoidParams creates a new GetInventoryDeviceInfosMoidParams object
// with the default values initialized.
func NewGetInventoryDeviceInfosMoidParams() *GetInventoryDeviceInfosMoidParams {
	var ()
	return &GetInventoryDeviceInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryDeviceInfosMoidParamsWithTimeout creates a new GetInventoryDeviceInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInventoryDeviceInfosMoidParamsWithTimeout(timeout time.Duration) *GetInventoryDeviceInfosMoidParams {
	var ()
	return &GetInventoryDeviceInfosMoidParams{

		timeout: timeout,
	}
}

// NewGetInventoryDeviceInfosMoidParamsWithContext creates a new GetInventoryDeviceInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInventoryDeviceInfosMoidParamsWithContext(ctx context.Context) *GetInventoryDeviceInfosMoidParams {
	var ()
	return &GetInventoryDeviceInfosMoidParams{

		Context: ctx,
	}
}

// NewGetInventoryDeviceInfosMoidParamsWithHTTPClient creates a new GetInventoryDeviceInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInventoryDeviceInfosMoidParamsWithHTTPClient(client *http.Client) *GetInventoryDeviceInfosMoidParams {
	var ()
	return &GetInventoryDeviceInfosMoidParams{
		HTTPClient: client,
	}
}

/*GetInventoryDeviceInfosMoidParams contains all the parameters to send to the API endpoint
for the get inventory device infos moid operation typically these are written to a http.Request
*/
type GetInventoryDeviceInfosMoidParams struct {

	/*Moid
	  The moid of the inventoryDeviceInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get inventory device infos moid params
func (o *GetInventoryDeviceInfosMoidParams) WithTimeout(timeout time.Duration) *GetInventoryDeviceInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory device infos moid params
func (o *GetInventoryDeviceInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory device infos moid params
func (o *GetInventoryDeviceInfosMoidParams) WithContext(ctx context.Context) *GetInventoryDeviceInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory device infos moid params
func (o *GetInventoryDeviceInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory device infos moid params
func (o *GetInventoryDeviceInfosMoidParams) WithHTTPClient(client *http.Client) *GetInventoryDeviceInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory device infos moid params
func (o *GetInventoryDeviceInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get inventory device infos moid params
func (o *GetInventoryDeviceInfosMoidParams) WithMoid(moid string) *GetInventoryDeviceInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get inventory device infos moid params
func (o *GetInventoryDeviceInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryDeviceInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
