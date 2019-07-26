// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_capability_info

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

// NewGetHyperflexCapabilityInfosMoidParams creates a new GetHyperflexCapabilityInfosMoidParams object
// with the default values initialized.
func NewGetHyperflexCapabilityInfosMoidParams() *GetHyperflexCapabilityInfosMoidParams {
	var ()
	return &GetHyperflexCapabilityInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHyperflexCapabilityInfosMoidParamsWithTimeout creates a new GetHyperflexCapabilityInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHyperflexCapabilityInfosMoidParamsWithTimeout(timeout time.Duration) *GetHyperflexCapabilityInfosMoidParams {
	var ()
	return &GetHyperflexCapabilityInfosMoidParams{

		timeout: timeout,
	}
}

// NewGetHyperflexCapabilityInfosMoidParamsWithContext creates a new GetHyperflexCapabilityInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHyperflexCapabilityInfosMoidParamsWithContext(ctx context.Context) *GetHyperflexCapabilityInfosMoidParams {
	var ()
	return &GetHyperflexCapabilityInfosMoidParams{

		Context: ctx,
	}
}

// NewGetHyperflexCapabilityInfosMoidParamsWithHTTPClient creates a new GetHyperflexCapabilityInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHyperflexCapabilityInfosMoidParamsWithHTTPClient(client *http.Client) *GetHyperflexCapabilityInfosMoidParams {
	var ()
	return &GetHyperflexCapabilityInfosMoidParams{
		HTTPClient: client,
	}
}

/*GetHyperflexCapabilityInfosMoidParams contains all the parameters to send to the API endpoint
for the get hyperflex capability infos moid operation typically these are written to a http.Request
*/
type GetHyperflexCapabilityInfosMoidParams struct {

	/*Moid
	  The moid of the hyperflexCapabilityInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hyperflex capability infos moid params
func (o *GetHyperflexCapabilityInfosMoidParams) WithTimeout(timeout time.Duration) *GetHyperflexCapabilityInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hyperflex capability infos moid params
func (o *GetHyperflexCapabilityInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hyperflex capability infos moid params
func (o *GetHyperflexCapabilityInfosMoidParams) WithContext(ctx context.Context) *GetHyperflexCapabilityInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hyperflex capability infos moid params
func (o *GetHyperflexCapabilityInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hyperflex capability infos moid params
func (o *GetHyperflexCapabilityInfosMoidParams) WithHTTPClient(client *http.Client) *GetHyperflexCapabilityInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hyperflex capability infos moid params
func (o *GetHyperflexCapabilityInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hyperflex capability infos moid params
func (o *GetHyperflexCapabilityInfosMoidParams) WithMoid(moid string) *GetHyperflexCapabilityInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hyperflex capability infos moid params
func (o *GetHyperflexCapabilityInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHyperflexCapabilityInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
