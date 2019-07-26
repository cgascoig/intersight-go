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

// NewDeleteTestcryptTokenApisMoidParams creates a new DeleteTestcryptTokenApisMoidParams object
// with the default values initialized.
func NewDeleteTestcryptTokenApisMoidParams() *DeleteTestcryptTokenApisMoidParams {
	var ()
	return &DeleteTestcryptTokenApisMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTestcryptTokenApisMoidParamsWithTimeout creates a new DeleteTestcryptTokenApisMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTestcryptTokenApisMoidParamsWithTimeout(timeout time.Duration) *DeleteTestcryptTokenApisMoidParams {
	var ()
	return &DeleteTestcryptTokenApisMoidParams{

		timeout: timeout,
	}
}

// NewDeleteTestcryptTokenApisMoidParamsWithContext creates a new DeleteTestcryptTokenApisMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTestcryptTokenApisMoidParamsWithContext(ctx context.Context) *DeleteTestcryptTokenApisMoidParams {
	var ()
	return &DeleteTestcryptTokenApisMoidParams{

		Context: ctx,
	}
}

// NewDeleteTestcryptTokenApisMoidParamsWithHTTPClient creates a new DeleteTestcryptTokenApisMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTestcryptTokenApisMoidParamsWithHTTPClient(client *http.Client) *DeleteTestcryptTokenApisMoidParams {
	var ()
	return &DeleteTestcryptTokenApisMoidParams{
		HTTPClient: client,
	}
}

/*DeleteTestcryptTokenApisMoidParams contains all the parameters to send to the API endpoint
for the delete testcrypt token apis moid operation typically these are written to a http.Request
*/
type DeleteTestcryptTokenApisMoidParams struct {

	/*Moid
	  The moid of the testcryptTokenApi instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete testcrypt token apis moid params
func (o *DeleteTestcryptTokenApisMoidParams) WithTimeout(timeout time.Duration) *DeleteTestcryptTokenApisMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete testcrypt token apis moid params
func (o *DeleteTestcryptTokenApisMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete testcrypt token apis moid params
func (o *DeleteTestcryptTokenApisMoidParams) WithContext(ctx context.Context) *DeleteTestcryptTokenApisMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete testcrypt token apis moid params
func (o *DeleteTestcryptTokenApisMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete testcrypt token apis moid params
func (o *DeleteTestcryptTokenApisMoidParams) WithHTTPClient(client *http.Client) *DeleteTestcryptTokenApisMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete testcrypt token apis moid params
func (o *DeleteTestcryptTokenApisMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete testcrypt token apis moid params
func (o *DeleteTestcryptTokenApisMoidParams) WithMoid(moid string) *DeleteTestcryptTokenApisMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete testcrypt token apis moid params
func (o *DeleteTestcryptTokenApisMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTestcryptTokenApisMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
