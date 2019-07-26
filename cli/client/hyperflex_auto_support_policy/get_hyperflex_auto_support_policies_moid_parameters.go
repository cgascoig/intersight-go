// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_auto_support_policy

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

// NewGetHyperflexAutoSupportPoliciesMoidParams creates a new GetHyperflexAutoSupportPoliciesMoidParams object
// with the default values initialized.
func NewGetHyperflexAutoSupportPoliciesMoidParams() *GetHyperflexAutoSupportPoliciesMoidParams {
	var ()
	return &GetHyperflexAutoSupportPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHyperflexAutoSupportPoliciesMoidParamsWithTimeout creates a new GetHyperflexAutoSupportPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHyperflexAutoSupportPoliciesMoidParamsWithTimeout(timeout time.Duration) *GetHyperflexAutoSupportPoliciesMoidParams {
	var ()
	return &GetHyperflexAutoSupportPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewGetHyperflexAutoSupportPoliciesMoidParamsWithContext creates a new GetHyperflexAutoSupportPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHyperflexAutoSupportPoliciesMoidParamsWithContext(ctx context.Context) *GetHyperflexAutoSupportPoliciesMoidParams {
	var ()
	return &GetHyperflexAutoSupportPoliciesMoidParams{

		Context: ctx,
	}
}

// NewGetHyperflexAutoSupportPoliciesMoidParamsWithHTTPClient creates a new GetHyperflexAutoSupportPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHyperflexAutoSupportPoliciesMoidParamsWithHTTPClient(client *http.Client) *GetHyperflexAutoSupportPoliciesMoidParams {
	var ()
	return &GetHyperflexAutoSupportPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*GetHyperflexAutoSupportPoliciesMoidParams contains all the parameters to send to the API endpoint
for the get hyperflex auto support policies moid operation typically these are written to a http.Request
*/
type GetHyperflexAutoSupportPoliciesMoidParams struct {

	/*Moid
	  The moid of the hyperflexAutoSupportPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hyperflex auto support policies moid params
func (o *GetHyperflexAutoSupportPoliciesMoidParams) WithTimeout(timeout time.Duration) *GetHyperflexAutoSupportPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hyperflex auto support policies moid params
func (o *GetHyperflexAutoSupportPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hyperflex auto support policies moid params
func (o *GetHyperflexAutoSupportPoliciesMoidParams) WithContext(ctx context.Context) *GetHyperflexAutoSupportPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hyperflex auto support policies moid params
func (o *GetHyperflexAutoSupportPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hyperflex auto support policies moid params
func (o *GetHyperflexAutoSupportPoliciesMoidParams) WithHTTPClient(client *http.Client) *GetHyperflexAutoSupportPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hyperflex auto support policies moid params
func (o *GetHyperflexAutoSupportPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hyperflex auto support policies moid params
func (o *GetHyperflexAutoSupportPoliciesMoidParams) WithMoid(moid string) *GetHyperflexAutoSupportPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hyperflex auto support policies moid params
func (o *GetHyperflexAutoSupportPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHyperflexAutoSupportPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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