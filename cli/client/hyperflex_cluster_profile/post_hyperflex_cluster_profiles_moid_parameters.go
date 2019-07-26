// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster_profile

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

// NewPostHyperflexClusterProfilesMoidParams creates a new PostHyperflexClusterProfilesMoidParams object
// with the default values initialized.
func NewPostHyperflexClusterProfilesMoidParams() *PostHyperflexClusterProfilesMoidParams {
	var ()
	return &PostHyperflexClusterProfilesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexClusterProfilesMoidParamsWithTimeout creates a new PostHyperflexClusterProfilesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexClusterProfilesMoidParamsWithTimeout(timeout time.Duration) *PostHyperflexClusterProfilesMoidParams {
	var ()
	return &PostHyperflexClusterProfilesMoidParams{

		timeout: timeout,
	}
}

// NewPostHyperflexClusterProfilesMoidParamsWithContext creates a new PostHyperflexClusterProfilesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexClusterProfilesMoidParamsWithContext(ctx context.Context) *PostHyperflexClusterProfilesMoidParams {
	var ()
	return &PostHyperflexClusterProfilesMoidParams{

		Context: ctx,
	}
}

// NewPostHyperflexClusterProfilesMoidParamsWithHTTPClient creates a new PostHyperflexClusterProfilesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexClusterProfilesMoidParamsWithHTTPClient(client *http.Client) *PostHyperflexClusterProfilesMoidParams {
	var ()
	return &PostHyperflexClusterProfilesMoidParams{
		HTTPClient: client,
	}
}

/*PostHyperflexClusterProfilesMoidParams contains all the parameters to send to the API endpoint
for the post hyperflex cluster profiles moid operation typically these are written to a http.Request
*/
type PostHyperflexClusterProfilesMoidParams struct {

	/*Body
	  hyperflexClusterProfile to update

	*/
	Body *models.HyperflexClusterProfile
	/*Moid
	  The moid of the hyperflexClusterProfile instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) WithTimeout(timeout time.Duration) *PostHyperflexClusterProfilesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) WithContext(ctx context.Context) *PostHyperflexClusterProfilesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) WithHTTPClient(client *http.Client) *PostHyperflexClusterProfilesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) WithBody(body *models.HyperflexClusterProfile) *PostHyperflexClusterProfilesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) SetBody(body *models.HyperflexClusterProfile) {
	o.Body = body
}

// WithMoid adds the moid to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) WithMoid(moid string) *PostHyperflexClusterProfilesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post hyperflex cluster profiles moid params
func (o *PostHyperflexClusterProfilesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexClusterProfilesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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