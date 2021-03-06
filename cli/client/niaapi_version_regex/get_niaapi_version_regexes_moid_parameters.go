// Code generated by go-swagger; DO NOT EDIT.

package niaapi_version_regex

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

// NewGetNiaapiVersionRegexesMoidParams creates a new GetNiaapiVersionRegexesMoidParams object
// with the default values initialized.
func NewGetNiaapiVersionRegexesMoidParams() *GetNiaapiVersionRegexesMoidParams {
	var ()
	return &GetNiaapiVersionRegexesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNiaapiVersionRegexesMoidParamsWithTimeout creates a new GetNiaapiVersionRegexesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNiaapiVersionRegexesMoidParamsWithTimeout(timeout time.Duration) *GetNiaapiVersionRegexesMoidParams {
	var ()
	return &GetNiaapiVersionRegexesMoidParams{

		timeout: timeout,
	}
}

// NewGetNiaapiVersionRegexesMoidParamsWithContext creates a new GetNiaapiVersionRegexesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNiaapiVersionRegexesMoidParamsWithContext(ctx context.Context) *GetNiaapiVersionRegexesMoidParams {
	var ()
	return &GetNiaapiVersionRegexesMoidParams{

		Context: ctx,
	}
}

// NewGetNiaapiVersionRegexesMoidParamsWithHTTPClient creates a new GetNiaapiVersionRegexesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNiaapiVersionRegexesMoidParamsWithHTTPClient(client *http.Client) *GetNiaapiVersionRegexesMoidParams {
	var ()
	return &GetNiaapiVersionRegexesMoidParams{
		HTTPClient: client,
	}
}

/*GetNiaapiVersionRegexesMoidParams contains all the parameters to send to the API endpoint
for the get niaapi version regexes moid operation typically these are written to a http.Request
*/
type GetNiaapiVersionRegexesMoidParams struct {

	/*Moid
	  The moid of the niaapiVersionRegex instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get niaapi version regexes moid params
func (o *GetNiaapiVersionRegexesMoidParams) WithTimeout(timeout time.Duration) *GetNiaapiVersionRegexesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get niaapi version regexes moid params
func (o *GetNiaapiVersionRegexesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get niaapi version regexes moid params
func (o *GetNiaapiVersionRegexesMoidParams) WithContext(ctx context.Context) *GetNiaapiVersionRegexesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get niaapi version regexes moid params
func (o *GetNiaapiVersionRegexesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get niaapi version regexes moid params
func (o *GetNiaapiVersionRegexesMoidParams) WithHTTPClient(client *http.Client) *GetNiaapiVersionRegexesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get niaapi version regexes moid params
func (o *GetNiaapiVersionRegexesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get niaapi version regexes moid params
func (o *GetNiaapiVersionRegexesMoidParams) WithMoid(moid string) *GetNiaapiVersionRegexesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get niaapi version regexes moid params
func (o *GetNiaapiVersionRegexesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetNiaapiVersionRegexesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
