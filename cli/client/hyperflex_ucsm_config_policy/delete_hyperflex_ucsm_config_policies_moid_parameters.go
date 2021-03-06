// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_ucsm_config_policy

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

// NewDeleteHyperflexUcsmConfigPoliciesMoidParams creates a new DeleteHyperflexUcsmConfigPoliciesMoidParams object
// with the default values initialized.
func NewDeleteHyperflexUcsmConfigPoliciesMoidParams() *DeleteHyperflexUcsmConfigPoliciesMoidParams {
	var ()
	return &DeleteHyperflexUcsmConfigPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHyperflexUcsmConfigPoliciesMoidParamsWithTimeout creates a new DeleteHyperflexUcsmConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteHyperflexUcsmConfigPoliciesMoidParamsWithTimeout(timeout time.Duration) *DeleteHyperflexUcsmConfigPoliciesMoidParams {
	var ()
	return &DeleteHyperflexUcsmConfigPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewDeleteHyperflexUcsmConfigPoliciesMoidParamsWithContext creates a new DeleteHyperflexUcsmConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteHyperflexUcsmConfigPoliciesMoidParamsWithContext(ctx context.Context) *DeleteHyperflexUcsmConfigPoliciesMoidParams {
	var ()
	return &DeleteHyperflexUcsmConfigPoliciesMoidParams{

		Context: ctx,
	}
}

// NewDeleteHyperflexUcsmConfigPoliciesMoidParamsWithHTTPClient creates a new DeleteHyperflexUcsmConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteHyperflexUcsmConfigPoliciesMoidParamsWithHTTPClient(client *http.Client) *DeleteHyperflexUcsmConfigPoliciesMoidParams {
	var ()
	return &DeleteHyperflexUcsmConfigPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*DeleteHyperflexUcsmConfigPoliciesMoidParams contains all the parameters to send to the API endpoint
for the delete hyperflex ucsm config policies moid operation typically these are written to a http.Request
*/
type DeleteHyperflexUcsmConfigPoliciesMoidParams struct {

	/*Moid
	  The moid of the hyperflexUcsmConfigPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete hyperflex ucsm config policies moid params
func (o *DeleteHyperflexUcsmConfigPoliciesMoidParams) WithTimeout(timeout time.Duration) *DeleteHyperflexUcsmConfigPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete hyperflex ucsm config policies moid params
func (o *DeleteHyperflexUcsmConfigPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete hyperflex ucsm config policies moid params
func (o *DeleteHyperflexUcsmConfigPoliciesMoidParams) WithContext(ctx context.Context) *DeleteHyperflexUcsmConfigPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete hyperflex ucsm config policies moid params
func (o *DeleteHyperflexUcsmConfigPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete hyperflex ucsm config policies moid params
func (o *DeleteHyperflexUcsmConfigPoliciesMoidParams) WithHTTPClient(client *http.Client) *DeleteHyperflexUcsmConfigPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete hyperflex ucsm config policies moid params
func (o *DeleteHyperflexUcsmConfigPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete hyperflex ucsm config policies moid params
func (o *DeleteHyperflexUcsmConfigPoliciesMoidParams) WithMoid(moid string) *DeleteHyperflexUcsmConfigPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete hyperflex ucsm config policies moid params
func (o *DeleteHyperflexUcsmConfigPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHyperflexUcsmConfigPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
