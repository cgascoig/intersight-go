// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_feature_limit_internal

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

// NewGetHyperflexFeatureLimitInternalsMoidParams creates a new GetHyperflexFeatureLimitInternalsMoidParams object
// with the default values initialized.
func NewGetHyperflexFeatureLimitInternalsMoidParams() *GetHyperflexFeatureLimitInternalsMoidParams {
	var ()
	return &GetHyperflexFeatureLimitInternalsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHyperflexFeatureLimitInternalsMoidParamsWithTimeout creates a new GetHyperflexFeatureLimitInternalsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHyperflexFeatureLimitInternalsMoidParamsWithTimeout(timeout time.Duration) *GetHyperflexFeatureLimitInternalsMoidParams {
	var ()
	return &GetHyperflexFeatureLimitInternalsMoidParams{

		timeout: timeout,
	}
}

// NewGetHyperflexFeatureLimitInternalsMoidParamsWithContext creates a new GetHyperflexFeatureLimitInternalsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHyperflexFeatureLimitInternalsMoidParamsWithContext(ctx context.Context) *GetHyperflexFeatureLimitInternalsMoidParams {
	var ()
	return &GetHyperflexFeatureLimitInternalsMoidParams{

		Context: ctx,
	}
}

// NewGetHyperflexFeatureLimitInternalsMoidParamsWithHTTPClient creates a new GetHyperflexFeatureLimitInternalsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHyperflexFeatureLimitInternalsMoidParamsWithHTTPClient(client *http.Client) *GetHyperflexFeatureLimitInternalsMoidParams {
	var ()
	return &GetHyperflexFeatureLimitInternalsMoidParams{
		HTTPClient: client,
	}
}

/*GetHyperflexFeatureLimitInternalsMoidParams contains all the parameters to send to the API endpoint
for the get hyperflex feature limit internals moid operation typically these are written to a http.Request
*/
type GetHyperflexFeatureLimitInternalsMoidParams struct {

	/*Moid
	  The moid of the hyperflexFeatureLimitInternal instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hyperflex feature limit internals moid params
func (o *GetHyperflexFeatureLimitInternalsMoidParams) WithTimeout(timeout time.Duration) *GetHyperflexFeatureLimitInternalsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hyperflex feature limit internals moid params
func (o *GetHyperflexFeatureLimitInternalsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hyperflex feature limit internals moid params
func (o *GetHyperflexFeatureLimitInternalsMoidParams) WithContext(ctx context.Context) *GetHyperflexFeatureLimitInternalsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hyperflex feature limit internals moid params
func (o *GetHyperflexFeatureLimitInternalsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hyperflex feature limit internals moid params
func (o *GetHyperflexFeatureLimitInternalsMoidParams) WithHTTPClient(client *http.Client) *GetHyperflexFeatureLimitInternalsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hyperflex feature limit internals moid params
func (o *GetHyperflexFeatureLimitInternalsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hyperflex feature limit internals moid params
func (o *GetHyperflexFeatureLimitInternalsMoidParams) WithMoid(moid string) *GetHyperflexFeatureLimitInternalsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hyperflex feature limit internals moid params
func (o *GetHyperflexFeatureLimitInternalsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHyperflexFeatureLimitInternalsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
