// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_ext_iscsi_storage_policy

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

// NewPostHyperflexExtIscsiStoragePoliciesParams creates a new PostHyperflexExtIscsiStoragePoliciesParams object
// with the default values initialized.
func NewPostHyperflexExtIscsiStoragePoliciesParams() *PostHyperflexExtIscsiStoragePoliciesParams {
	var ()
	return &PostHyperflexExtIscsiStoragePoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexExtIscsiStoragePoliciesParamsWithTimeout creates a new PostHyperflexExtIscsiStoragePoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexExtIscsiStoragePoliciesParamsWithTimeout(timeout time.Duration) *PostHyperflexExtIscsiStoragePoliciesParams {
	var ()
	return &PostHyperflexExtIscsiStoragePoliciesParams{

		timeout: timeout,
	}
}

// NewPostHyperflexExtIscsiStoragePoliciesParamsWithContext creates a new PostHyperflexExtIscsiStoragePoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexExtIscsiStoragePoliciesParamsWithContext(ctx context.Context) *PostHyperflexExtIscsiStoragePoliciesParams {
	var ()
	return &PostHyperflexExtIscsiStoragePoliciesParams{

		Context: ctx,
	}
}

// NewPostHyperflexExtIscsiStoragePoliciesParamsWithHTTPClient creates a new PostHyperflexExtIscsiStoragePoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexExtIscsiStoragePoliciesParamsWithHTTPClient(client *http.Client) *PostHyperflexExtIscsiStoragePoliciesParams {
	var ()
	return &PostHyperflexExtIscsiStoragePoliciesParams{
		HTTPClient: client,
	}
}

/*PostHyperflexExtIscsiStoragePoliciesParams contains all the parameters to send to the API endpoint
for the post hyperflex ext iscsi storage policies operation typically these are written to a http.Request
*/
type PostHyperflexExtIscsiStoragePoliciesParams struct {

	/*Body
	  hyperflexExtIscsiStoragePolicy to add

	*/
	Body *models.HyperflexExtIscsiStoragePolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex ext iscsi storage policies params
func (o *PostHyperflexExtIscsiStoragePoliciesParams) WithTimeout(timeout time.Duration) *PostHyperflexExtIscsiStoragePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex ext iscsi storage policies params
func (o *PostHyperflexExtIscsiStoragePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex ext iscsi storage policies params
func (o *PostHyperflexExtIscsiStoragePoliciesParams) WithContext(ctx context.Context) *PostHyperflexExtIscsiStoragePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex ext iscsi storage policies params
func (o *PostHyperflexExtIscsiStoragePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex ext iscsi storage policies params
func (o *PostHyperflexExtIscsiStoragePoliciesParams) WithHTTPClient(client *http.Client) *PostHyperflexExtIscsiStoragePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex ext iscsi storage policies params
func (o *PostHyperflexExtIscsiStoragePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex ext iscsi storage policies params
func (o *PostHyperflexExtIscsiStoragePoliciesParams) WithBody(body *models.HyperflexExtIscsiStoragePolicy) *PostHyperflexExtIscsiStoragePoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex ext iscsi storage policies params
func (o *PostHyperflexExtIscsiStoragePoliciesParams) SetBody(body *models.HyperflexExtIscsiStoragePolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexExtIscsiStoragePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
