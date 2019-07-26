// Code generated by go-swagger; DO NOT EDIT.

package hcl_operating_system

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

// NewGetHclOperatingSystemsMoidParams creates a new GetHclOperatingSystemsMoidParams object
// with the default values initialized.
func NewGetHclOperatingSystemsMoidParams() *GetHclOperatingSystemsMoidParams {
	var ()
	return &GetHclOperatingSystemsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHclOperatingSystemsMoidParamsWithTimeout creates a new GetHclOperatingSystemsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHclOperatingSystemsMoidParamsWithTimeout(timeout time.Duration) *GetHclOperatingSystemsMoidParams {
	var ()
	return &GetHclOperatingSystemsMoidParams{

		timeout: timeout,
	}
}

// NewGetHclOperatingSystemsMoidParamsWithContext creates a new GetHclOperatingSystemsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHclOperatingSystemsMoidParamsWithContext(ctx context.Context) *GetHclOperatingSystemsMoidParams {
	var ()
	return &GetHclOperatingSystemsMoidParams{

		Context: ctx,
	}
}

// NewGetHclOperatingSystemsMoidParamsWithHTTPClient creates a new GetHclOperatingSystemsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHclOperatingSystemsMoidParamsWithHTTPClient(client *http.Client) *GetHclOperatingSystemsMoidParams {
	var ()
	return &GetHclOperatingSystemsMoidParams{
		HTTPClient: client,
	}
}

/*GetHclOperatingSystemsMoidParams contains all the parameters to send to the API endpoint
for the get hcl operating systems moid operation typically these are written to a http.Request
*/
type GetHclOperatingSystemsMoidParams struct {

	/*Moid
	  The moid of the hclOperatingSystem instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hcl operating systems moid params
func (o *GetHclOperatingSystemsMoidParams) WithTimeout(timeout time.Duration) *GetHclOperatingSystemsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hcl operating systems moid params
func (o *GetHclOperatingSystemsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hcl operating systems moid params
func (o *GetHclOperatingSystemsMoidParams) WithContext(ctx context.Context) *GetHclOperatingSystemsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hcl operating systems moid params
func (o *GetHclOperatingSystemsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hcl operating systems moid params
func (o *GetHclOperatingSystemsMoidParams) WithHTTPClient(client *http.Client) *GetHclOperatingSystemsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hcl operating systems moid params
func (o *GetHclOperatingSystemsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hcl operating systems moid params
func (o *GetHclOperatingSystemsMoidParams) WithMoid(moid string) *GetHclOperatingSystemsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hcl operating systems moid params
func (o *GetHclOperatingSystemsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHclOperatingSystemsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
