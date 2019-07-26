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

// NewPostHyperflexUcsmConfigPoliciesParams creates a new PostHyperflexUcsmConfigPoliciesParams object
// with the default values initialized.
func NewPostHyperflexUcsmConfigPoliciesParams() *PostHyperflexUcsmConfigPoliciesParams {
	var ()
	return &PostHyperflexUcsmConfigPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexUcsmConfigPoliciesParamsWithTimeout creates a new PostHyperflexUcsmConfigPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexUcsmConfigPoliciesParamsWithTimeout(timeout time.Duration) *PostHyperflexUcsmConfigPoliciesParams {
	var ()
	return &PostHyperflexUcsmConfigPoliciesParams{

		timeout: timeout,
	}
}

// NewPostHyperflexUcsmConfigPoliciesParamsWithContext creates a new PostHyperflexUcsmConfigPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexUcsmConfigPoliciesParamsWithContext(ctx context.Context) *PostHyperflexUcsmConfigPoliciesParams {
	var ()
	return &PostHyperflexUcsmConfigPoliciesParams{

		Context: ctx,
	}
}

// NewPostHyperflexUcsmConfigPoliciesParamsWithHTTPClient creates a new PostHyperflexUcsmConfigPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexUcsmConfigPoliciesParamsWithHTTPClient(client *http.Client) *PostHyperflexUcsmConfigPoliciesParams {
	var ()
	return &PostHyperflexUcsmConfigPoliciesParams{
		HTTPClient: client,
	}
}

/*PostHyperflexUcsmConfigPoliciesParams contains all the parameters to send to the API endpoint
for the post hyperflex ucsm config policies operation typically these are written to a http.Request
*/
type PostHyperflexUcsmConfigPoliciesParams struct {

	/*Body
	  hyperflexUcsmConfigPolicy to add

	*/
	Body *models.HyperflexUcsmConfigPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex ucsm config policies params
func (o *PostHyperflexUcsmConfigPoliciesParams) WithTimeout(timeout time.Duration) *PostHyperflexUcsmConfigPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex ucsm config policies params
func (o *PostHyperflexUcsmConfigPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex ucsm config policies params
func (o *PostHyperflexUcsmConfigPoliciesParams) WithContext(ctx context.Context) *PostHyperflexUcsmConfigPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex ucsm config policies params
func (o *PostHyperflexUcsmConfigPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex ucsm config policies params
func (o *PostHyperflexUcsmConfigPoliciesParams) WithHTTPClient(client *http.Client) *PostHyperflexUcsmConfigPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex ucsm config policies params
func (o *PostHyperflexUcsmConfigPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex ucsm config policies params
func (o *PostHyperflexUcsmConfigPoliciesParams) WithBody(body *models.HyperflexUcsmConfigPolicy) *PostHyperflexUcsmConfigPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex ucsm config policies params
func (o *PostHyperflexUcsmConfigPoliciesParams) SetBody(body *models.HyperflexUcsmConfigPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexUcsmConfigPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
