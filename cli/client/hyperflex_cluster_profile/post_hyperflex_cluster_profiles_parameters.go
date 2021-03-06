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

// NewPostHyperflexClusterProfilesParams creates a new PostHyperflexClusterProfilesParams object
// with the default values initialized.
func NewPostHyperflexClusterProfilesParams() *PostHyperflexClusterProfilesParams {
	var ()
	return &PostHyperflexClusterProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexClusterProfilesParamsWithTimeout creates a new PostHyperflexClusterProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexClusterProfilesParamsWithTimeout(timeout time.Duration) *PostHyperflexClusterProfilesParams {
	var ()
	return &PostHyperflexClusterProfilesParams{

		timeout: timeout,
	}
}

// NewPostHyperflexClusterProfilesParamsWithContext creates a new PostHyperflexClusterProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexClusterProfilesParamsWithContext(ctx context.Context) *PostHyperflexClusterProfilesParams {
	var ()
	return &PostHyperflexClusterProfilesParams{

		Context: ctx,
	}
}

// NewPostHyperflexClusterProfilesParamsWithHTTPClient creates a new PostHyperflexClusterProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexClusterProfilesParamsWithHTTPClient(client *http.Client) *PostHyperflexClusterProfilesParams {
	var ()
	return &PostHyperflexClusterProfilesParams{
		HTTPClient: client,
	}
}

/*PostHyperflexClusterProfilesParams contains all the parameters to send to the API endpoint
for the post hyperflex cluster profiles operation typically these are written to a http.Request
*/
type PostHyperflexClusterProfilesParams struct {

	/*Body
	  hyperflexClusterProfile to add

	*/
	Body *models.HyperflexClusterProfile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex cluster profiles params
func (o *PostHyperflexClusterProfilesParams) WithTimeout(timeout time.Duration) *PostHyperflexClusterProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex cluster profiles params
func (o *PostHyperflexClusterProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex cluster profiles params
func (o *PostHyperflexClusterProfilesParams) WithContext(ctx context.Context) *PostHyperflexClusterProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex cluster profiles params
func (o *PostHyperflexClusterProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex cluster profiles params
func (o *PostHyperflexClusterProfilesParams) WithHTTPClient(client *http.Client) *PostHyperflexClusterProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex cluster profiles params
func (o *PostHyperflexClusterProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex cluster profiles params
func (o *PostHyperflexClusterProfilesParams) WithBody(body *models.HyperflexClusterProfile) *PostHyperflexClusterProfilesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex cluster profiles params
func (o *PostHyperflexClusterProfilesParams) SetBody(body *models.HyperflexClusterProfile) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexClusterProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
