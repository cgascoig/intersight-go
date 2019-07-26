// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_server_model

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

// NewPostHyperflexServerModelsParams creates a new PostHyperflexServerModelsParams object
// with the default values initialized.
func NewPostHyperflexServerModelsParams() *PostHyperflexServerModelsParams {
	var ()
	return &PostHyperflexServerModelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexServerModelsParamsWithTimeout creates a new PostHyperflexServerModelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexServerModelsParamsWithTimeout(timeout time.Duration) *PostHyperflexServerModelsParams {
	var ()
	return &PostHyperflexServerModelsParams{

		timeout: timeout,
	}
}

// NewPostHyperflexServerModelsParamsWithContext creates a new PostHyperflexServerModelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexServerModelsParamsWithContext(ctx context.Context) *PostHyperflexServerModelsParams {
	var ()
	return &PostHyperflexServerModelsParams{

		Context: ctx,
	}
}

// NewPostHyperflexServerModelsParamsWithHTTPClient creates a new PostHyperflexServerModelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexServerModelsParamsWithHTTPClient(client *http.Client) *PostHyperflexServerModelsParams {
	var ()
	return &PostHyperflexServerModelsParams{
		HTTPClient: client,
	}
}

/*PostHyperflexServerModelsParams contains all the parameters to send to the API endpoint
for the post hyperflex server models operation typically these are written to a http.Request
*/
type PostHyperflexServerModelsParams struct {

	/*Body
	  hyperflexServerModel to add

	*/
	Body *models.HyperflexServerModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex server models params
func (o *PostHyperflexServerModelsParams) WithTimeout(timeout time.Duration) *PostHyperflexServerModelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex server models params
func (o *PostHyperflexServerModelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex server models params
func (o *PostHyperflexServerModelsParams) WithContext(ctx context.Context) *PostHyperflexServerModelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex server models params
func (o *PostHyperflexServerModelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex server models params
func (o *PostHyperflexServerModelsParams) WithHTTPClient(client *http.Client) *PostHyperflexServerModelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex server models params
func (o *PostHyperflexServerModelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex server models params
func (o *PostHyperflexServerModelsParams) WithBody(body *models.HyperflexServerModel) *PostHyperflexServerModelsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex server models params
func (o *PostHyperflexServerModelsParams) SetBody(body *models.HyperflexServerModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexServerModelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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