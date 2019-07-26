// Code generated by go-swagger; DO NOT EDIT.

package compute_board

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

// NewGetComputeBoardsMoidParams creates a new GetComputeBoardsMoidParams object
// with the default values initialized.
func NewGetComputeBoardsMoidParams() *GetComputeBoardsMoidParams {
	var ()
	return &GetComputeBoardsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetComputeBoardsMoidParamsWithTimeout creates a new GetComputeBoardsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetComputeBoardsMoidParamsWithTimeout(timeout time.Duration) *GetComputeBoardsMoidParams {
	var ()
	return &GetComputeBoardsMoidParams{

		timeout: timeout,
	}
}

// NewGetComputeBoardsMoidParamsWithContext creates a new GetComputeBoardsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetComputeBoardsMoidParamsWithContext(ctx context.Context) *GetComputeBoardsMoidParams {
	var ()
	return &GetComputeBoardsMoidParams{

		Context: ctx,
	}
}

// NewGetComputeBoardsMoidParamsWithHTTPClient creates a new GetComputeBoardsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetComputeBoardsMoidParamsWithHTTPClient(client *http.Client) *GetComputeBoardsMoidParams {
	var ()
	return &GetComputeBoardsMoidParams{
		HTTPClient: client,
	}
}

/*GetComputeBoardsMoidParams contains all the parameters to send to the API endpoint
for the get compute boards moid operation typically these are written to a http.Request
*/
type GetComputeBoardsMoidParams struct {

	/*Moid
	  The moid of the computeBoard instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get compute boards moid params
func (o *GetComputeBoardsMoidParams) WithTimeout(timeout time.Duration) *GetComputeBoardsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compute boards moid params
func (o *GetComputeBoardsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compute boards moid params
func (o *GetComputeBoardsMoidParams) WithContext(ctx context.Context) *GetComputeBoardsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compute boards moid params
func (o *GetComputeBoardsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compute boards moid params
func (o *GetComputeBoardsMoidParams) WithHTTPClient(client *http.Client) *GetComputeBoardsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compute boards moid params
func (o *GetComputeBoardsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get compute boards moid params
func (o *GetComputeBoardsMoidParams) WithMoid(moid string) *GetComputeBoardsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get compute boards moid params
func (o *GetComputeBoardsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetComputeBoardsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
