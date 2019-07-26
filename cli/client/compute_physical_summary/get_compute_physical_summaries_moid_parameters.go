// Code generated by go-swagger; DO NOT EDIT.

package compute_physical_summary

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

// NewGetComputePhysicalSummariesMoidParams creates a new GetComputePhysicalSummariesMoidParams object
// with the default values initialized.
func NewGetComputePhysicalSummariesMoidParams() *GetComputePhysicalSummariesMoidParams {
	var ()
	return &GetComputePhysicalSummariesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetComputePhysicalSummariesMoidParamsWithTimeout creates a new GetComputePhysicalSummariesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetComputePhysicalSummariesMoidParamsWithTimeout(timeout time.Duration) *GetComputePhysicalSummariesMoidParams {
	var ()
	return &GetComputePhysicalSummariesMoidParams{

		timeout: timeout,
	}
}

// NewGetComputePhysicalSummariesMoidParamsWithContext creates a new GetComputePhysicalSummariesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetComputePhysicalSummariesMoidParamsWithContext(ctx context.Context) *GetComputePhysicalSummariesMoidParams {
	var ()
	return &GetComputePhysicalSummariesMoidParams{

		Context: ctx,
	}
}

// NewGetComputePhysicalSummariesMoidParamsWithHTTPClient creates a new GetComputePhysicalSummariesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetComputePhysicalSummariesMoidParamsWithHTTPClient(client *http.Client) *GetComputePhysicalSummariesMoidParams {
	var ()
	return &GetComputePhysicalSummariesMoidParams{
		HTTPClient: client,
	}
}

/*GetComputePhysicalSummariesMoidParams contains all the parameters to send to the API endpoint
for the get compute physical summaries moid operation typically these are written to a http.Request
*/
type GetComputePhysicalSummariesMoidParams struct {

	/*Moid
	  The moid of the computePhysicalSummary instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get compute physical summaries moid params
func (o *GetComputePhysicalSummariesMoidParams) WithTimeout(timeout time.Duration) *GetComputePhysicalSummariesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compute physical summaries moid params
func (o *GetComputePhysicalSummariesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compute physical summaries moid params
func (o *GetComputePhysicalSummariesMoidParams) WithContext(ctx context.Context) *GetComputePhysicalSummariesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compute physical summaries moid params
func (o *GetComputePhysicalSummariesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compute physical summaries moid params
func (o *GetComputePhysicalSummariesMoidParams) WithHTTPClient(client *http.Client) *GetComputePhysicalSummariesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compute physical summaries moid params
func (o *GetComputePhysicalSummariesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get compute physical summaries moid params
func (o *GetComputePhysicalSummariesMoidParams) WithMoid(moid string) *GetComputePhysicalSummariesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get compute physical summaries moid params
func (o *GetComputePhysicalSummariesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetComputePhysicalSummariesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
