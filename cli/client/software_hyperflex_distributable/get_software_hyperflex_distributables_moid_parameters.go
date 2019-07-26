// Code generated by go-swagger; DO NOT EDIT.

package software_hyperflex_distributable

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

// NewGetSoftwareHyperflexDistributablesMoidParams creates a new GetSoftwareHyperflexDistributablesMoidParams object
// with the default values initialized.
func NewGetSoftwareHyperflexDistributablesMoidParams() *GetSoftwareHyperflexDistributablesMoidParams {
	var ()
	return &GetSoftwareHyperflexDistributablesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSoftwareHyperflexDistributablesMoidParamsWithTimeout creates a new GetSoftwareHyperflexDistributablesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSoftwareHyperflexDistributablesMoidParamsWithTimeout(timeout time.Duration) *GetSoftwareHyperflexDistributablesMoidParams {
	var ()
	return &GetSoftwareHyperflexDistributablesMoidParams{

		timeout: timeout,
	}
}

// NewGetSoftwareHyperflexDistributablesMoidParamsWithContext creates a new GetSoftwareHyperflexDistributablesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSoftwareHyperflexDistributablesMoidParamsWithContext(ctx context.Context) *GetSoftwareHyperflexDistributablesMoidParams {
	var ()
	return &GetSoftwareHyperflexDistributablesMoidParams{

		Context: ctx,
	}
}

// NewGetSoftwareHyperflexDistributablesMoidParamsWithHTTPClient creates a new GetSoftwareHyperflexDistributablesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSoftwareHyperflexDistributablesMoidParamsWithHTTPClient(client *http.Client) *GetSoftwareHyperflexDistributablesMoidParams {
	var ()
	return &GetSoftwareHyperflexDistributablesMoidParams{
		HTTPClient: client,
	}
}

/*GetSoftwareHyperflexDistributablesMoidParams contains all the parameters to send to the API endpoint
for the get software hyperflex distributables moid operation typically these are written to a http.Request
*/
type GetSoftwareHyperflexDistributablesMoidParams struct {

	/*Moid
	  The moid of the softwareHyperflexDistributable instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get software hyperflex distributables moid params
func (o *GetSoftwareHyperflexDistributablesMoidParams) WithTimeout(timeout time.Duration) *GetSoftwareHyperflexDistributablesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get software hyperflex distributables moid params
func (o *GetSoftwareHyperflexDistributablesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get software hyperflex distributables moid params
func (o *GetSoftwareHyperflexDistributablesMoidParams) WithContext(ctx context.Context) *GetSoftwareHyperflexDistributablesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get software hyperflex distributables moid params
func (o *GetSoftwareHyperflexDistributablesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get software hyperflex distributables moid params
func (o *GetSoftwareHyperflexDistributablesMoidParams) WithHTTPClient(client *http.Client) *GetSoftwareHyperflexDistributablesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get software hyperflex distributables moid params
func (o *GetSoftwareHyperflexDistributablesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get software hyperflex distributables moid params
func (o *GetSoftwareHyperflexDistributablesMoidParams) WithMoid(moid string) *GetSoftwareHyperflexDistributablesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get software hyperflex distributables moid params
func (o *GetSoftwareHyperflexDistributablesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetSoftwareHyperflexDistributablesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
