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

// NewPatchHyperflexAppCatalogsMoidParams creates a new PatchHyperflexAppCatalogsMoidParams object
// with the default values initialized.
func NewPatchHyperflexAppCatalogsMoidParams() *PatchHyperflexAppCatalogsMoidParams {
	var ()
	return &PatchHyperflexAppCatalogsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHyperflexAppCatalogsMoidParamsWithTimeout creates a new PatchHyperflexAppCatalogsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchHyperflexAppCatalogsMoidParamsWithTimeout(timeout time.Duration) *PatchHyperflexAppCatalogsMoidParams {
	var ()
	return &PatchHyperflexAppCatalogsMoidParams{

		timeout: timeout,
	}
}

// NewPatchHyperflexAppCatalogsMoidParamsWithContext creates a new PatchHyperflexAppCatalogsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchHyperflexAppCatalogsMoidParamsWithContext(ctx context.Context) *PatchHyperflexAppCatalogsMoidParams {
	var ()
	return &PatchHyperflexAppCatalogsMoidParams{

		Context: ctx,
	}
}

// NewPatchHyperflexAppCatalogsMoidParamsWithHTTPClient creates a new PatchHyperflexAppCatalogsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchHyperflexAppCatalogsMoidParamsWithHTTPClient(client *http.Client) *PatchHyperflexAppCatalogsMoidParams {
	var ()
	return &PatchHyperflexAppCatalogsMoidParams{
		HTTPClient: client,
	}
}

/*PatchHyperflexAppCatalogsMoidParams contains all the parameters to send to the API endpoint
for the patch hyperflex app catalogs moid operation typically these are written to a http.Request
*/
type PatchHyperflexAppCatalogsMoidParams struct {

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

// WithTimeout adds the timeout to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) WithTimeout(timeout time.Duration) *PatchHyperflexAppCatalogsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) WithContext(ctx context.Context) *PatchHyperflexAppCatalogsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) WithHTTPClient(client *http.Client) *PatchHyperflexAppCatalogsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) WithBody(body *models.HyperflexAppCatalog) *PatchHyperflexAppCatalogsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) SetBody(body *models.HyperflexAppCatalog) {
	o.Body = body
}

// WithMoid adds the moid to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) WithMoid(moid string) *PatchHyperflexAppCatalogsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch hyperflex app catalogs moid params
func (o *PatchHyperflexAppCatalogsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHyperflexAppCatalogsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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