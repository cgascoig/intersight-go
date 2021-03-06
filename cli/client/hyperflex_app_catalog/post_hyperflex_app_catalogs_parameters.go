// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_app_catalog

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

// NewPostHyperflexAppCatalogsParams creates a new PostHyperflexAppCatalogsParams object
// with the default values initialized.
func NewPostHyperflexAppCatalogsParams() *PostHyperflexAppCatalogsParams {
	var ()
	return &PostHyperflexAppCatalogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexAppCatalogsParamsWithTimeout creates a new PostHyperflexAppCatalogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexAppCatalogsParamsWithTimeout(timeout time.Duration) *PostHyperflexAppCatalogsParams {
	var ()
	return &PostHyperflexAppCatalogsParams{

		timeout: timeout,
	}
}

// NewPostHyperflexAppCatalogsParamsWithContext creates a new PostHyperflexAppCatalogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexAppCatalogsParamsWithContext(ctx context.Context) *PostHyperflexAppCatalogsParams {
	var ()
	return &PostHyperflexAppCatalogsParams{

		Context: ctx,
	}
}

// NewPostHyperflexAppCatalogsParamsWithHTTPClient creates a new PostHyperflexAppCatalogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexAppCatalogsParamsWithHTTPClient(client *http.Client) *PostHyperflexAppCatalogsParams {
	var ()
	return &PostHyperflexAppCatalogsParams{
		HTTPClient: client,
	}
}

/*PostHyperflexAppCatalogsParams contains all the parameters to send to the API endpoint
for the post hyperflex app catalogs operation typically these are written to a http.Request
*/
type PostHyperflexAppCatalogsParams struct {

	/*Body
	  hyperflexAppCatalog to add

	*/
	Body *models.HyperflexAppCatalog

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex app catalogs params
func (o *PostHyperflexAppCatalogsParams) WithTimeout(timeout time.Duration) *PostHyperflexAppCatalogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex app catalogs params
func (o *PostHyperflexAppCatalogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex app catalogs params
func (o *PostHyperflexAppCatalogsParams) WithContext(ctx context.Context) *PostHyperflexAppCatalogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex app catalogs params
func (o *PostHyperflexAppCatalogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex app catalogs params
func (o *PostHyperflexAppCatalogsParams) WithHTTPClient(client *http.Client) *PostHyperflexAppCatalogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex app catalogs params
func (o *PostHyperflexAppCatalogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex app catalogs params
func (o *PostHyperflexAppCatalogsParams) WithBody(body *models.HyperflexAppCatalog) *PostHyperflexAppCatalogsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex app catalogs params
func (o *PostHyperflexAppCatalogsParams) SetBody(body *models.HyperflexAppCatalog) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexAppCatalogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
