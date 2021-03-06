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

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPostHyperflexUcsmConfigPoliciesMoidParams creates a new PostHyperflexUcsmConfigPoliciesMoidParams object
// with the default values initialized.
func NewPostHyperflexUcsmConfigPoliciesMoidParams() *PostHyperflexUcsmConfigPoliciesMoidParams {
	var ()
	return &PostHyperflexUcsmConfigPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexUcsmConfigPoliciesMoidParamsWithTimeout creates a new PostHyperflexUcsmConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexUcsmConfigPoliciesMoidParamsWithTimeout(timeout time.Duration) *PostHyperflexUcsmConfigPoliciesMoidParams {
	var ()
	return &PostHyperflexUcsmConfigPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostHyperflexUcsmConfigPoliciesMoidParamsWithContext creates a new PostHyperflexUcsmConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexUcsmConfigPoliciesMoidParamsWithContext(ctx context.Context) *PostHyperflexUcsmConfigPoliciesMoidParams {
	var ()
	return &PostHyperflexUcsmConfigPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostHyperflexUcsmConfigPoliciesMoidParamsWithHTTPClient creates a new PostHyperflexUcsmConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexUcsmConfigPoliciesMoidParamsWithHTTPClient(client *http.Client) *PostHyperflexUcsmConfigPoliciesMoidParams {
	var ()
	return &PostHyperflexUcsmConfigPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostHyperflexUcsmConfigPoliciesMoidParams contains all the parameters to send to the API endpoint
for the post hyperflex ucsm config policies moid operation typically these are written to a http.Request
*/
type PostHyperflexUcsmConfigPoliciesMoidParams struct {

	/*Body
	  hyperflexUcsmConfigPolicy to update

	*/
	Body *models.HyperflexUcsmConfigPolicy
	/*Moid
	  The moid of the hyperflexUcsmConfigPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) WithTimeout(timeout time.Duration) *PostHyperflexUcsmConfigPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) WithContext(ctx context.Context) *PostHyperflexUcsmConfigPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) WithHTTPClient(client *http.Client) *PostHyperflexUcsmConfigPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) WithBody(body *models.HyperflexUcsmConfigPolicy) *PostHyperflexUcsmConfigPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) SetBody(body *models.HyperflexUcsmConfigPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) WithMoid(moid string) *PostHyperflexUcsmConfigPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post hyperflex ucsm config policies moid params
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexUcsmConfigPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param moid
	if err := r.SetPathParam("moid", o.Moid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
