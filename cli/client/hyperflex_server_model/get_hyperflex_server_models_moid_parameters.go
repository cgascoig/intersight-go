// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_server_model

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

// NewGetHyperflexServerModelsMoidParams creates a new GetHyperflexServerModelsMoidParams object
// with the default values initialized.
func NewGetHyperflexServerModelsMoidParams() *GetHyperflexServerModelsMoidParams {
	var ()
	return &GetHyperflexServerModelsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHyperflexServerModelsMoidParamsWithTimeout creates a new GetHyperflexServerModelsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHyperflexServerModelsMoidParamsWithTimeout(timeout time.Duration) *GetHyperflexServerModelsMoidParams {
	var ()
	return &GetHyperflexServerModelsMoidParams{

		timeout: timeout,
	}
}

// NewGetHyperflexServerModelsMoidParamsWithContext creates a new GetHyperflexServerModelsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHyperflexServerModelsMoidParamsWithContext(ctx context.Context) *GetHyperflexServerModelsMoidParams {
	var ()
	return &GetHyperflexServerModelsMoidParams{

		Context: ctx,
	}
}

// NewGetHyperflexServerModelsMoidParamsWithHTTPClient creates a new GetHyperflexServerModelsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHyperflexServerModelsMoidParamsWithHTTPClient(client *http.Client) *GetHyperflexServerModelsMoidParams {
	var ()
	return &GetHyperflexServerModelsMoidParams{
		HTTPClient: client,
	}
}

/*GetHyperflexServerModelsMoidParams contains all the parameters to send to the API endpoint
for the get hyperflex server models moid operation typically these are written to a http.Request
*/
type GetHyperflexServerModelsMoidParams struct {

	/*Moid
	  The moid of the hyperflexServerModel instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hyperflex server models moid params
func (o *GetHyperflexServerModelsMoidParams) WithTimeout(timeout time.Duration) *GetHyperflexServerModelsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hyperflex server models moid params
func (o *GetHyperflexServerModelsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hyperflex server models moid params
func (o *GetHyperflexServerModelsMoidParams) WithContext(ctx context.Context) *GetHyperflexServerModelsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hyperflex server models moid params
func (o *GetHyperflexServerModelsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hyperflex server models moid params
func (o *GetHyperflexServerModelsMoidParams) WithHTTPClient(client *http.Client) *GetHyperflexServerModelsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hyperflex server models moid params
func (o *GetHyperflexServerModelsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hyperflex server models moid params
func (o *GetHyperflexServerModelsMoidParams) WithMoid(moid string) *GetHyperflexServerModelsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hyperflex server models moid params
func (o *GetHyperflexServerModelsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHyperflexServerModelsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
