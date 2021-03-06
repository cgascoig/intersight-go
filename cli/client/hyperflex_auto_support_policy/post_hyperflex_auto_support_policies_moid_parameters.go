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

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPostHyperflexAutoSupportPoliciesMoidParams creates a new PostHyperflexAutoSupportPoliciesMoidParams object
// with the default values initialized.
func NewPostHyperflexAutoSupportPoliciesMoidParams() *PostHyperflexAutoSupportPoliciesMoidParams {
	var ()
	return &PostHyperflexAutoSupportPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexAutoSupportPoliciesMoidParamsWithTimeout creates a new PostHyperflexAutoSupportPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexAutoSupportPoliciesMoidParamsWithTimeout(timeout time.Duration) *PostHyperflexAutoSupportPoliciesMoidParams {
	var ()
	return &PostHyperflexAutoSupportPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostHyperflexAutoSupportPoliciesMoidParamsWithContext creates a new PostHyperflexAutoSupportPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexAutoSupportPoliciesMoidParamsWithContext(ctx context.Context) *PostHyperflexAutoSupportPoliciesMoidParams {
	var ()
	return &PostHyperflexAutoSupportPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostHyperflexAutoSupportPoliciesMoidParamsWithHTTPClient creates a new PostHyperflexAutoSupportPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexAutoSupportPoliciesMoidParamsWithHTTPClient(client *http.Client) *PostHyperflexAutoSupportPoliciesMoidParams {
	var ()
	return &PostHyperflexAutoSupportPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostHyperflexAutoSupportPoliciesMoidParams contains all the parameters to send to the API endpoint
for the post hyperflex auto support policies moid operation typically these are written to a http.Request
*/
type PostHyperflexAutoSupportPoliciesMoidParams struct {

	/*Body
	  hyperflexAutoSupportPolicy to update

	*/
	Body *models.HyperflexAutoSupportPolicy
	/*Moid
	  The moid of the hyperflexAutoSupportPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) WithTimeout(timeout time.Duration) *PostHyperflexAutoSupportPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) WithContext(ctx context.Context) *PostHyperflexAutoSupportPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) WithHTTPClient(client *http.Client) *PostHyperflexAutoSupportPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) WithBody(body *models.HyperflexAutoSupportPolicy) *PostHyperflexAutoSupportPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) SetBody(body *models.HyperflexAutoSupportPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) WithMoid(moid string) *PostHyperflexAutoSupportPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post hyperflex auto support policies moid params
func (o *PostHyperflexAutoSupportPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexAutoSupportPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
