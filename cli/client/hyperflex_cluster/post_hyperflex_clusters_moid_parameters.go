// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster

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

// NewPostHyperflexClustersMoidParams creates a new PostHyperflexClustersMoidParams object
// with the default values initialized.
func NewPostHyperflexClustersMoidParams() *PostHyperflexClustersMoidParams {
	var ()
	return &PostHyperflexClustersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexClustersMoidParamsWithTimeout creates a new PostHyperflexClustersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexClustersMoidParamsWithTimeout(timeout time.Duration) *PostHyperflexClustersMoidParams {
	var ()
	return &PostHyperflexClustersMoidParams{

		timeout: timeout,
	}
}

// NewPostHyperflexClustersMoidParamsWithContext creates a new PostHyperflexClustersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexClustersMoidParamsWithContext(ctx context.Context) *PostHyperflexClustersMoidParams {
	var ()
	return &PostHyperflexClustersMoidParams{

		Context: ctx,
	}
}

// NewPostHyperflexClustersMoidParamsWithHTTPClient creates a new PostHyperflexClustersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexClustersMoidParamsWithHTTPClient(client *http.Client) *PostHyperflexClustersMoidParams {
	var ()
	return &PostHyperflexClustersMoidParams{
		HTTPClient: client,
	}
}

/*PostHyperflexClustersMoidParams contains all the parameters to send to the API endpoint
for the post hyperflex clusters moid operation typically these are written to a http.Request
*/
type PostHyperflexClustersMoidParams struct {

	/*Body
	  hyperflexCluster to update

	*/
	Body *models.HyperflexCluster
	/*Moid
	  The moid of the hyperflexCluster instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) WithTimeout(timeout time.Duration) *PostHyperflexClustersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) WithContext(ctx context.Context) *PostHyperflexClustersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) WithHTTPClient(client *http.Client) *PostHyperflexClustersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) WithBody(body *models.HyperflexCluster) *PostHyperflexClustersMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) SetBody(body *models.HyperflexCluster) {
	o.Body = body
}

// WithMoid adds the moid to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) WithMoid(moid string) *PostHyperflexClustersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post hyperflex clusters moid params
func (o *PostHyperflexClustersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexClustersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
