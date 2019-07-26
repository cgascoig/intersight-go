// Code generated by go-swagger; DO NOT EDIT.

package testcrypt_token_api

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

// NewGetTestcryptTokenApisMoidParams creates a new GetTestcryptTokenApisMoidParams object
// with the default values initialized.
func NewGetTestcryptTokenApisMoidParams() *GetTestcryptTokenApisMoidParams {
	var ()
	return &GetTestcryptTokenApisMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestcryptTokenApisMoidParamsWithTimeout creates a new GetTestcryptTokenApisMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestcryptTokenApisMoidParamsWithTimeout(timeout time.Duration) *GetTestcryptTokenApisMoidParams {
	var ()
	return &GetTestcryptTokenApisMoidParams{

		timeout: timeout,
	}
}

// NewGetTestcryptTokenApisMoidParamsWithContext creates a new GetTestcryptTokenApisMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestcryptTokenApisMoidParamsWithContext(ctx context.Context) *GetTestcryptTokenApisMoidParams {
	var ()
	return &GetTestcryptTokenApisMoidParams{

		Context: ctx,
	}
}

// NewGetTestcryptTokenApisMoidParamsWithHTTPClient creates a new GetTestcryptTokenApisMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestcryptTokenApisMoidParamsWithHTTPClient(client *http.Client) *GetTestcryptTokenApisMoidParams {
	var ()
	return &GetTestcryptTokenApisMoidParams{
		HTTPClient: client,
	}
}

/*GetTestcryptTokenApisMoidParams contains all the parameters to send to the API endpoint
for the get testcrypt token apis moid operation typically these are written to a http.Request
*/
type GetTestcryptTokenApisMoidParams struct {

	/*Moid
	  The moid of the testcryptTokenApi instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get testcrypt token apis moid params
func (o *GetTestcryptTokenApisMoidParams) WithTimeout(timeout time.Duration) *GetTestcryptTokenApisMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get testcrypt token apis moid params
func (o *GetTestcryptTokenApisMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get testcrypt token apis moid params
func (o *GetTestcryptTokenApisMoidParams) WithContext(ctx context.Context) *GetTestcryptTokenApisMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get testcrypt token apis moid params
func (o *GetTestcryptTokenApisMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get testcrypt token apis moid params
func (o *GetTestcryptTokenApisMoidParams) WithHTTPClient(client *http.Client) *GetTestcryptTokenApisMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get testcrypt token apis moid params
func (o *GetTestcryptTokenApisMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get testcrypt token apis moid params
func (o *GetTestcryptTokenApisMoidParams) WithMoid(moid string) *GetTestcryptTokenApisMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get testcrypt token apis moid params
func (o *GetTestcryptTokenApisMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestcryptTokenApisMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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