// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_ext_fc_storage_policy

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

// NewPostHyperflexExtFcStoragePoliciesMoidParams creates a new PostHyperflexExtFcStoragePoliciesMoidParams object
// with the default values initialized.
func NewPostHyperflexExtFcStoragePoliciesMoidParams() *PostHyperflexExtFcStoragePoliciesMoidParams {
	var ()
	return &PostHyperflexExtFcStoragePoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexExtFcStoragePoliciesMoidParamsWithTimeout creates a new PostHyperflexExtFcStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexExtFcStoragePoliciesMoidParamsWithTimeout(timeout time.Duration) *PostHyperflexExtFcStoragePoliciesMoidParams {
	var ()
	return &PostHyperflexExtFcStoragePoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostHyperflexExtFcStoragePoliciesMoidParamsWithContext creates a new PostHyperflexExtFcStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexExtFcStoragePoliciesMoidParamsWithContext(ctx context.Context) *PostHyperflexExtFcStoragePoliciesMoidParams {
	var ()
	return &PostHyperflexExtFcStoragePoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostHyperflexExtFcStoragePoliciesMoidParamsWithHTTPClient creates a new PostHyperflexExtFcStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexExtFcStoragePoliciesMoidParamsWithHTTPClient(client *http.Client) *PostHyperflexExtFcStoragePoliciesMoidParams {
	var ()
	return &PostHyperflexExtFcStoragePoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostHyperflexExtFcStoragePoliciesMoidParams contains all the parameters to send to the API endpoint
for the post hyperflex ext fc storage policies moid operation typically these are written to a http.Request
*/
type PostHyperflexExtFcStoragePoliciesMoidParams struct {

	/*Body
	  hyperflexExtFcStoragePolicy to update

	*/
	Body *models.HyperflexExtFcStoragePolicy
	/*Moid
	  The moid of the hyperflexExtFcStoragePolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) WithTimeout(timeout time.Duration) *PostHyperflexExtFcStoragePoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) WithContext(ctx context.Context) *PostHyperflexExtFcStoragePoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) WithHTTPClient(client *http.Client) *PostHyperflexExtFcStoragePoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) WithBody(body *models.HyperflexExtFcStoragePolicy) *PostHyperflexExtFcStoragePoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) SetBody(body *models.HyperflexExtFcStoragePolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) WithMoid(moid string) *PostHyperflexExtFcStoragePoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post hyperflex ext fc storage policies moid params
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexExtFcStoragePoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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