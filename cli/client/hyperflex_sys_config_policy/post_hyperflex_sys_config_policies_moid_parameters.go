// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_sys_config_policy

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

// NewPostHyperflexSysConfigPoliciesMoidParams creates a new PostHyperflexSysConfigPoliciesMoidParams object
// with the default values initialized.
func NewPostHyperflexSysConfigPoliciesMoidParams() *PostHyperflexSysConfigPoliciesMoidParams {
	var ()
	return &PostHyperflexSysConfigPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexSysConfigPoliciesMoidParamsWithTimeout creates a new PostHyperflexSysConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexSysConfigPoliciesMoidParamsWithTimeout(timeout time.Duration) *PostHyperflexSysConfigPoliciesMoidParams {
	var ()
	return &PostHyperflexSysConfigPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostHyperflexSysConfigPoliciesMoidParamsWithContext creates a new PostHyperflexSysConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexSysConfigPoliciesMoidParamsWithContext(ctx context.Context) *PostHyperflexSysConfigPoliciesMoidParams {
	var ()
	return &PostHyperflexSysConfigPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostHyperflexSysConfigPoliciesMoidParamsWithHTTPClient creates a new PostHyperflexSysConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexSysConfigPoliciesMoidParamsWithHTTPClient(client *http.Client) *PostHyperflexSysConfigPoliciesMoidParams {
	var ()
	return &PostHyperflexSysConfigPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostHyperflexSysConfigPoliciesMoidParams contains all the parameters to send to the API endpoint
for the post hyperflex sys config policies moid operation typically these are written to a http.Request
*/
type PostHyperflexSysConfigPoliciesMoidParams struct {

	/*Body
	  hyperflexSysConfigPolicy to update

	*/
	Body *models.HyperflexSysConfigPolicy
	/*Moid
	  The moid of the hyperflexSysConfigPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) WithTimeout(timeout time.Duration) *PostHyperflexSysConfigPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) WithContext(ctx context.Context) *PostHyperflexSysConfigPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) WithHTTPClient(client *http.Client) *PostHyperflexSysConfigPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) WithBody(body *models.HyperflexSysConfigPolicy) *PostHyperflexSysConfigPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) SetBody(body *models.HyperflexSysConfigPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) WithMoid(moid string) *PostHyperflexSysConfigPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post hyperflex sys config policies moid params
func (o *PostHyperflexSysConfigPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexSysConfigPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
