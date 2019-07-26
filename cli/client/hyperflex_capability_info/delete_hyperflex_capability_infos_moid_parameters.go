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

// NewDeleteHyperflexCapabilityInfosMoidParams creates a new DeleteHyperflexCapabilityInfosMoidParams object
// with the default values initialized.
func NewDeleteHyperflexCapabilityInfosMoidParams() *DeleteHyperflexCapabilityInfosMoidParams {
	var ()
	return &DeleteHyperflexCapabilityInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHyperflexCapabilityInfosMoidParamsWithTimeout creates a new DeleteHyperflexCapabilityInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteHyperflexCapabilityInfosMoidParamsWithTimeout(timeout time.Duration) *DeleteHyperflexCapabilityInfosMoidParams {
	var ()
	return &DeleteHyperflexCapabilityInfosMoidParams{

		timeout: timeout,
	}
}

// NewDeleteHyperflexCapabilityInfosMoidParamsWithContext creates a new DeleteHyperflexCapabilityInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteHyperflexCapabilityInfosMoidParamsWithContext(ctx context.Context) *DeleteHyperflexCapabilityInfosMoidParams {
	var ()
	return &DeleteHyperflexCapabilityInfosMoidParams{

		Context: ctx,
	}
}

// NewDeleteHyperflexCapabilityInfosMoidParamsWithHTTPClient creates a new DeleteHyperflexCapabilityInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteHyperflexCapabilityInfosMoidParamsWithHTTPClient(client *http.Client) *DeleteHyperflexCapabilityInfosMoidParams {
	var ()
	return &DeleteHyperflexCapabilityInfosMoidParams{
		HTTPClient: client,
	}
}

/*DeleteHyperflexCapabilityInfosMoidParams contains all the parameters to send to the API endpoint
for the delete hyperflex capability infos moid operation typically these are written to a http.Request
*/
type DeleteHyperflexCapabilityInfosMoidParams struct {

	/*Moid
	  The moid of the hyperflexCapabilityInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete hyperflex capability infos moid params
func (o *DeleteHyperflexCapabilityInfosMoidParams) WithTimeout(timeout time.Duration) *DeleteHyperflexCapabilityInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete hyperflex capability infos moid params
func (o *DeleteHyperflexCapabilityInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete hyperflex capability infos moid params
func (o *DeleteHyperflexCapabilityInfosMoidParams) WithContext(ctx context.Context) *DeleteHyperflexCapabilityInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete hyperflex capability infos moid params
func (o *DeleteHyperflexCapabilityInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete hyperflex capability infos moid params
func (o *DeleteHyperflexCapabilityInfosMoidParams) WithHTTPClient(client *http.Client) *DeleteHyperflexCapabilityInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete hyperflex capability infos moid params
func (o *DeleteHyperflexCapabilityInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete hyperflex capability infos moid params
func (o *DeleteHyperflexCapabilityInfosMoidParams) WithMoid(moid string) *DeleteHyperflexCapabilityInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete hyperflex capability infos moid params
func (o *DeleteHyperflexCapabilityInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHyperflexCapabilityInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
