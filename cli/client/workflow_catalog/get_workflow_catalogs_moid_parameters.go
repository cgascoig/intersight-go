// Code generated by go-swagger; DO NOT EDIT.

package workflow_catalog

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
)

// NewGetWorkflowCatalogsMoidParams creates a new GetWorkflowCatalogsMoidParams object
// with the default values initialized.
func NewGetWorkflowCatalogsMoidParams() *GetWorkflowCatalogsMoidParams {
	var ()
	return &GetWorkflowCatalogsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowCatalogsMoidParamsWithTimeout creates a new GetWorkflowCatalogsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkflowCatalogsMoidParamsWithTimeout(timeout time.Duration) *GetWorkflowCatalogsMoidParams {
	var ()
	return &GetWorkflowCatalogsMoidParams{

		timeout: timeout,
	}
}

// NewGetWorkflowCatalogsMoidParamsWithContext creates a new GetWorkflowCatalogsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkflowCatalogsMoidParamsWithContext(ctx context.Context) *GetWorkflowCatalogsMoidParams {
	var ()
	return &GetWorkflowCatalogsMoidParams{

		Context: ctx,
	}
}

// NewGetWorkflowCatalogsMoidParamsWithHTTPClient creates a new GetWorkflowCatalogsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkflowCatalogsMoidParamsWithHTTPClient(client *http.Client) *GetWorkflowCatalogsMoidParams {
	var ()
	return &GetWorkflowCatalogsMoidParams{
		HTTPClient: client,
	}
}

/*GetWorkflowCatalogsMoidParams contains all the parameters to send to the API endpoint
for the get workflow catalogs moid operation typically these are written to a http.Request
*/
type GetWorkflowCatalogsMoidParams struct {

	/*Moid
	  The moid of the workflowCatalog instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workflow catalogs moid params
func (o *GetWorkflowCatalogsMoidParams) WithTimeout(timeout time.Duration) *GetWorkflowCatalogsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow catalogs moid params
func (o *GetWorkflowCatalogsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow catalogs moid params
func (o *GetWorkflowCatalogsMoidParams) WithContext(ctx context.Context) *GetWorkflowCatalogsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow catalogs moid params
func (o *GetWorkflowCatalogsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow catalogs moid params
func (o *GetWorkflowCatalogsMoidParams) WithHTTPClient(client *http.Client) *GetWorkflowCatalogsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow catalogs moid params
func (o *GetWorkflowCatalogsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get workflow catalogs moid params
func (o *GetWorkflowCatalogsMoidParams) WithMoid(moid string) *GetWorkflowCatalogsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get workflow catalogs moid params
func (o *GetWorkflowCatalogsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowCatalogsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param moid
	if err := r.SetPathParam("moid", o.Moid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
