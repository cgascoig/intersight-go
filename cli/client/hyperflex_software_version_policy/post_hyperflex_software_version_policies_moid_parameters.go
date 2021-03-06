// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_software_version_policy

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

// NewPostHyperflexSoftwareVersionPoliciesMoidParams creates a new PostHyperflexSoftwareVersionPoliciesMoidParams object
// with the default values initialized.
func NewPostHyperflexSoftwareVersionPoliciesMoidParams() *PostHyperflexSoftwareVersionPoliciesMoidParams {
	var ()
	return &PostHyperflexSoftwareVersionPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexSoftwareVersionPoliciesMoidParamsWithTimeout creates a new PostHyperflexSoftwareVersionPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexSoftwareVersionPoliciesMoidParamsWithTimeout(timeout time.Duration) *PostHyperflexSoftwareVersionPoliciesMoidParams {
	var ()
	return &PostHyperflexSoftwareVersionPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostHyperflexSoftwareVersionPoliciesMoidParamsWithContext creates a new PostHyperflexSoftwareVersionPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexSoftwareVersionPoliciesMoidParamsWithContext(ctx context.Context) *PostHyperflexSoftwareVersionPoliciesMoidParams {
	var ()
	return &PostHyperflexSoftwareVersionPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostHyperflexSoftwareVersionPoliciesMoidParamsWithHTTPClient creates a new PostHyperflexSoftwareVersionPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexSoftwareVersionPoliciesMoidParamsWithHTTPClient(client *http.Client) *PostHyperflexSoftwareVersionPoliciesMoidParams {
	var ()
	return &PostHyperflexSoftwareVersionPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostHyperflexSoftwareVersionPoliciesMoidParams contains all the parameters to send to the API endpoint
for the post hyperflex software version policies moid operation typically these are written to a http.Request
*/
type PostHyperflexSoftwareVersionPoliciesMoidParams struct {

	/*Body
	  hyperflexSoftwareVersionPolicy to update

	*/
	Body *models.HyperflexSoftwareVersionPolicy
	/*Moid
	  The moid of the hyperflexSoftwareVersionPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) WithTimeout(timeout time.Duration) *PostHyperflexSoftwareVersionPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) WithContext(ctx context.Context) *PostHyperflexSoftwareVersionPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) WithHTTPClient(client *http.Client) *PostHyperflexSoftwareVersionPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) WithBody(body *models.HyperflexSoftwareVersionPolicy) *PostHyperflexSoftwareVersionPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) SetBody(body *models.HyperflexSoftwareVersionPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) WithMoid(moid string) *PostHyperflexSoftwareVersionPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post hyperflex software version policies moid params
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexSoftwareVersionPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
