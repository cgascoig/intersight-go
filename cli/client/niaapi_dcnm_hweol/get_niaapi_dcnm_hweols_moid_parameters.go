// Code generated by go-swagger; DO NOT EDIT.

package niaapi_dcnm_hweol

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

// NewGetNiaapiDcnmHweolsMoidParams creates a new GetNiaapiDcnmHweolsMoidParams object
// with the default values initialized.
func NewGetNiaapiDcnmHweolsMoidParams() *GetNiaapiDcnmHweolsMoidParams {
	var ()
	return &GetNiaapiDcnmHweolsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNiaapiDcnmHweolsMoidParamsWithTimeout creates a new GetNiaapiDcnmHweolsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNiaapiDcnmHweolsMoidParamsWithTimeout(timeout time.Duration) *GetNiaapiDcnmHweolsMoidParams {
	var ()
	return &GetNiaapiDcnmHweolsMoidParams{

		timeout: timeout,
	}
}

// NewGetNiaapiDcnmHweolsMoidParamsWithContext creates a new GetNiaapiDcnmHweolsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNiaapiDcnmHweolsMoidParamsWithContext(ctx context.Context) *GetNiaapiDcnmHweolsMoidParams {
	var ()
	return &GetNiaapiDcnmHweolsMoidParams{

		Context: ctx,
	}
}

// NewGetNiaapiDcnmHweolsMoidParamsWithHTTPClient creates a new GetNiaapiDcnmHweolsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNiaapiDcnmHweolsMoidParamsWithHTTPClient(client *http.Client) *GetNiaapiDcnmHweolsMoidParams {
	var ()
	return &GetNiaapiDcnmHweolsMoidParams{
		HTTPClient: client,
	}
}

/*GetNiaapiDcnmHweolsMoidParams contains all the parameters to send to the API endpoint
for the get niaapi dcnm hweols moid operation typically these are written to a http.Request
*/
type GetNiaapiDcnmHweolsMoidParams struct {

	/*Moid
	  The moid of the niaapiDcnmHweol instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get niaapi dcnm hweols moid params
func (o *GetNiaapiDcnmHweolsMoidParams) WithTimeout(timeout time.Duration) *GetNiaapiDcnmHweolsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get niaapi dcnm hweols moid params
func (o *GetNiaapiDcnmHweolsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get niaapi dcnm hweols moid params
func (o *GetNiaapiDcnmHweolsMoidParams) WithContext(ctx context.Context) *GetNiaapiDcnmHweolsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get niaapi dcnm hweols moid params
func (o *GetNiaapiDcnmHweolsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get niaapi dcnm hweols moid params
func (o *GetNiaapiDcnmHweolsMoidParams) WithHTTPClient(client *http.Client) *GetNiaapiDcnmHweolsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get niaapi dcnm hweols moid params
func (o *GetNiaapiDcnmHweolsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get niaapi dcnm hweols moid params
func (o *GetNiaapiDcnmHweolsMoidParams) WithMoid(moid string) *GetNiaapiDcnmHweolsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get niaapi dcnm hweols moid params
func (o *GetNiaapiDcnmHweolsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetNiaapiDcnmHweolsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
