// Code generated by go-swagger; DO NOT EDIT.

package pci_switch

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

// NewGetPciSwitchesMoidParams creates a new GetPciSwitchesMoidParams object
// with the default values initialized.
func NewGetPciSwitchesMoidParams() *GetPciSwitchesMoidParams {
	var ()
	return &GetPciSwitchesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPciSwitchesMoidParamsWithTimeout creates a new GetPciSwitchesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPciSwitchesMoidParamsWithTimeout(timeout time.Duration) *GetPciSwitchesMoidParams {
	var ()
	return &GetPciSwitchesMoidParams{

		timeout: timeout,
	}
}

// NewGetPciSwitchesMoidParamsWithContext creates a new GetPciSwitchesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPciSwitchesMoidParamsWithContext(ctx context.Context) *GetPciSwitchesMoidParams {
	var ()
	return &GetPciSwitchesMoidParams{

		Context: ctx,
	}
}

// NewGetPciSwitchesMoidParamsWithHTTPClient creates a new GetPciSwitchesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPciSwitchesMoidParamsWithHTTPClient(client *http.Client) *GetPciSwitchesMoidParams {
	var ()
	return &GetPciSwitchesMoidParams{
		HTTPClient: client,
	}
}

/*GetPciSwitchesMoidParams contains all the parameters to send to the API endpoint
for the get pci switches moid operation typically these are written to a http.Request
*/
type GetPciSwitchesMoidParams struct {

	/*Moid
	  The moid of the pciSwitch instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pci switches moid params
func (o *GetPciSwitchesMoidParams) WithTimeout(timeout time.Duration) *GetPciSwitchesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pci switches moid params
func (o *GetPciSwitchesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pci switches moid params
func (o *GetPciSwitchesMoidParams) WithContext(ctx context.Context) *GetPciSwitchesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pci switches moid params
func (o *GetPciSwitchesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pci switches moid params
func (o *GetPciSwitchesMoidParams) WithHTTPClient(client *http.Client) *GetPciSwitchesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pci switches moid params
func (o *GetPciSwitchesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get pci switches moid params
func (o *GetPciSwitchesMoidParams) WithMoid(moid string) *GetPciSwitchesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get pci switches moid params
func (o *GetPciSwitchesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPciSwitchesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
