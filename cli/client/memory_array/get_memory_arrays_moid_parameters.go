// Code generated by go-swagger; DO NOT EDIT.

package memory_array

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

// NewGetMemoryArraysMoidParams creates a new GetMemoryArraysMoidParams object
// with the default values initialized.
func NewGetMemoryArraysMoidParams() *GetMemoryArraysMoidParams {
	var ()
	return &GetMemoryArraysMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMemoryArraysMoidParamsWithTimeout creates a new GetMemoryArraysMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMemoryArraysMoidParamsWithTimeout(timeout time.Duration) *GetMemoryArraysMoidParams {
	var ()
	return &GetMemoryArraysMoidParams{

		timeout: timeout,
	}
}

// NewGetMemoryArraysMoidParamsWithContext creates a new GetMemoryArraysMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMemoryArraysMoidParamsWithContext(ctx context.Context) *GetMemoryArraysMoidParams {
	var ()
	return &GetMemoryArraysMoidParams{

		Context: ctx,
	}
}

// NewGetMemoryArraysMoidParamsWithHTTPClient creates a new GetMemoryArraysMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMemoryArraysMoidParamsWithHTTPClient(client *http.Client) *GetMemoryArraysMoidParams {
	var ()
	return &GetMemoryArraysMoidParams{
		HTTPClient: client,
	}
}

/*GetMemoryArraysMoidParams contains all the parameters to send to the API endpoint
for the get memory arrays moid operation typically these are written to a http.Request
*/
type GetMemoryArraysMoidParams struct {

	/*Moid
	  The moid of the memoryArray instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get memory arrays moid params
func (o *GetMemoryArraysMoidParams) WithTimeout(timeout time.Duration) *GetMemoryArraysMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get memory arrays moid params
func (o *GetMemoryArraysMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get memory arrays moid params
func (o *GetMemoryArraysMoidParams) WithContext(ctx context.Context) *GetMemoryArraysMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get memory arrays moid params
func (o *GetMemoryArraysMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get memory arrays moid params
func (o *GetMemoryArraysMoidParams) WithHTTPClient(client *http.Client) *GetMemoryArraysMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get memory arrays moid params
func (o *GetMemoryArraysMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get memory arrays moid params
func (o *GetMemoryArraysMoidParams) WithMoid(moid string) *GetMemoryArraysMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get memory arrays moid params
func (o *GetMemoryArraysMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetMemoryArraysMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
