// Code generated by go-swagger; DO NOT EDIT.

package pci_device

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

// NewGetPciDevicesMoidParams creates a new GetPciDevicesMoidParams object
// with the default values initialized.
func NewGetPciDevicesMoidParams() *GetPciDevicesMoidParams {
	var ()
	return &GetPciDevicesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPciDevicesMoidParamsWithTimeout creates a new GetPciDevicesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPciDevicesMoidParamsWithTimeout(timeout time.Duration) *GetPciDevicesMoidParams {
	var ()
	return &GetPciDevicesMoidParams{

		timeout: timeout,
	}
}

// NewGetPciDevicesMoidParamsWithContext creates a new GetPciDevicesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPciDevicesMoidParamsWithContext(ctx context.Context) *GetPciDevicesMoidParams {
	var ()
	return &GetPciDevicesMoidParams{

		Context: ctx,
	}
}

// NewGetPciDevicesMoidParamsWithHTTPClient creates a new GetPciDevicesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPciDevicesMoidParamsWithHTTPClient(client *http.Client) *GetPciDevicesMoidParams {
	var ()
	return &GetPciDevicesMoidParams{
		HTTPClient: client,
	}
}

/*GetPciDevicesMoidParams contains all the parameters to send to the API endpoint
for the get pci devices moid operation typically these are written to a http.Request
*/
type GetPciDevicesMoidParams struct {

	/*Moid
	  The moid of the pciDevice instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pci devices moid params
func (o *GetPciDevicesMoidParams) WithTimeout(timeout time.Duration) *GetPciDevicesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pci devices moid params
func (o *GetPciDevicesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pci devices moid params
func (o *GetPciDevicesMoidParams) WithContext(ctx context.Context) *GetPciDevicesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pci devices moid params
func (o *GetPciDevicesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pci devices moid params
func (o *GetPciDevicesMoidParams) WithHTTPClient(client *http.Client) *GetPciDevicesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pci devices moid params
func (o *GetPciDevicesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get pci devices moid params
func (o *GetPciDevicesMoidParams) WithMoid(moid string) *GetPciDevicesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get pci devices moid params
func (o *GetPciDevicesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPciDevicesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
