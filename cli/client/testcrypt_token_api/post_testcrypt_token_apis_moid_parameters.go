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

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPostTestcryptTokenApisMoidParams creates a new PostTestcryptTokenApisMoidParams object
// with the default values initialized.
func NewPostTestcryptTokenApisMoidParams() *PostTestcryptTokenApisMoidParams {
	var ()
	return &PostTestcryptTokenApisMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTestcryptTokenApisMoidParamsWithTimeout creates a new PostTestcryptTokenApisMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTestcryptTokenApisMoidParamsWithTimeout(timeout time.Duration) *PostTestcryptTokenApisMoidParams {
	var ()
	return &PostTestcryptTokenApisMoidParams{

		timeout: timeout,
	}
}

// NewPostTestcryptTokenApisMoidParamsWithContext creates a new PostTestcryptTokenApisMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTestcryptTokenApisMoidParamsWithContext(ctx context.Context) *PostTestcryptTokenApisMoidParams {
	var ()
	return &PostTestcryptTokenApisMoidParams{

		Context: ctx,
	}
}

// NewPostTestcryptTokenApisMoidParamsWithHTTPClient creates a new PostTestcryptTokenApisMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTestcryptTokenApisMoidParamsWithHTTPClient(client *http.Client) *PostTestcryptTokenApisMoidParams {
	var ()
	return &PostTestcryptTokenApisMoidParams{
		HTTPClient: client,
	}
}

/*PostTestcryptTokenApisMoidParams contains all the parameters to send to the API endpoint
for the post testcrypt token apis moid operation typically these are written to a http.Request
*/
type PostTestcryptTokenApisMoidParams struct {

	/*Body
	  testcryptTokenApi to update

	*/
	Body *models.TestcryptTokenAPI
	/*Moid
	  The moid of the testcryptTokenApi instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) WithTimeout(timeout time.Duration) *PostTestcryptTokenApisMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) WithContext(ctx context.Context) *PostTestcryptTokenApisMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) WithHTTPClient(client *http.Client) *PostTestcryptTokenApisMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) WithBody(body *models.TestcryptTokenAPI) *PostTestcryptTokenApisMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) SetBody(body *models.TestcryptTokenAPI) {
	o.Body = body
}

// WithMoid adds the moid to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) WithMoid(moid string) *PostTestcryptTokenApisMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post testcrypt token apis moid params
func (o *PostTestcryptTokenApisMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostTestcryptTokenApisMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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