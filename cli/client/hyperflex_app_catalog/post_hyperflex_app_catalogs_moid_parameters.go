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

// NewPostHyperflexAppCatalogsMoidParams creates a new PostHyperflexAppCatalogsMoidParams object
// with the default values initialized.
func NewPostHyperflexAppCatalogsMoidParams() *PostHyperflexAppCatalogsMoidParams {
	var ()
	return &PostHyperflexAppCatalogsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexAppCatalogsMoidParamsWithTimeout creates a new PostHyperflexAppCatalogsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexAppCatalogsMoidParamsWithTimeout(timeout time.Duration) *PostHyperflexAppCatalogsMoidParams {
	var ()
	return &PostHyperflexAppCatalogsMoidParams{

		timeout: timeout,
	}
}

// NewPostHyperflexAppCatalogsMoidParamsWithContext creates a new PostHyperflexAppCatalogsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexAppCatalogsMoidParamsWithContext(ctx context.Context) *PostHyperflexAppCatalogsMoidParams {
	var ()
	return &PostHyperflexAppCatalogsMoidParams{

		Context: ctx,
	}
}

// NewPostHyperflexAppCatalogsMoidParamsWithHTTPClient creates a new PostHyperflexAppCatalogsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexAppCatalogsMoidParamsWithHTTPClient(client *http.Client) *PostHyperflexAppCatalogsMoidParams {
	var ()
	return &PostHyperflexAppCatalogsMoidParams{
		HTTPClient: client,
	}
}

/*PostHyperflexAppCatalogsMoidParams contains all the parameters to send to the API endpoint
for the post hyperflex app catalogs moid operation typically these are written to a http.Request
*/
type PostHyperflexAppCatalogsMoidParams struct {

	/*Body
	  hyperflexAppCatalog to update

	*/
	Body *models.HyperflexAppCatalog
	/*Moid
	  The moid of the hyperflexAppCatalog instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) WithTimeout(timeout time.Duration) *PostHyperflexAppCatalogsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) WithContext(ctx context.Context) *PostHyperflexAppCatalogsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) WithHTTPClient(client *http.Client) *PostHyperflexAppCatalogsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) WithBody(body *models.HyperflexAppCatalog) *PostHyperflexAppCatalogsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) SetBody(body *models.HyperflexAppCatalog) {
	o.Body = body
}

// WithMoid adds the moid to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) WithMoid(moid string) *PostHyperflexAppCatalogsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post hyperflex app catalogs moid params
func (o *PostHyperflexAppCatalogsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexAppCatalogsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
