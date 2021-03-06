// Code generated by go-swagger; DO NOT EDIT.

package testcrypt_credential

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

// NewDeleteTestcryptCredentialsMoidParams creates a new DeleteTestcryptCredentialsMoidParams object
// with the default values initialized.
func NewDeleteTestcryptCredentialsMoidParams() *DeleteTestcryptCredentialsMoidParams {
	var ()
	return &DeleteTestcryptCredentialsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTestcryptCredentialsMoidParamsWithTimeout creates a new DeleteTestcryptCredentialsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTestcryptCredentialsMoidParamsWithTimeout(timeout time.Duration) *DeleteTestcryptCredentialsMoidParams {
	var ()
	return &DeleteTestcryptCredentialsMoidParams{

		timeout: timeout,
	}
}

// NewDeleteTestcryptCredentialsMoidParamsWithContext creates a new DeleteTestcryptCredentialsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTestcryptCredentialsMoidParamsWithContext(ctx context.Context) *DeleteTestcryptCredentialsMoidParams {
	var ()
	return &DeleteTestcryptCredentialsMoidParams{

		Context: ctx,
	}
}

// NewDeleteTestcryptCredentialsMoidParamsWithHTTPClient creates a new DeleteTestcryptCredentialsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTestcryptCredentialsMoidParamsWithHTTPClient(client *http.Client) *DeleteTestcryptCredentialsMoidParams {
	var ()
	return &DeleteTestcryptCredentialsMoidParams{
		HTTPClient: client,
	}
}

/*DeleteTestcryptCredentialsMoidParams contains all the parameters to send to the API endpoint
for the delete testcrypt credentials moid operation typically these are written to a http.Request
*/
type DeleteTestcryptCredentialsMoidParams struct {

	/*Moid
	  The moid of the testcryptCredential instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete testcrypt credentials moid params
func (o *DeleteTestcryptCredentialsMoidParams) WithTimeout(timeout time.Duration) *DeleteTestcryptCredentialsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete testcrypt credentials moid params
func (o *DeleteTestcryptCredentialsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete testcrypt credentials moid params
func (o *DeleteTestcryptCredentialsMoidParams) WithContext(ctx context.Context) *DeleteTestcryptCredentialsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete testcrypt credentials moid params
func (o *DeleteTestcryptCredentialsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete testcrypt credentials moid params
func (o *DeleteTestcryptCredentialsMoidParams) WithHTTPClient(client *http.Client) *DeleteTestcryptCredentialsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete testcrypt credentials moid params
func (o *DeleteTestcryptCredentialsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete testcrypt credentials moid params
func (o *DeleteTestcryptCredentialsMoidParams) WithMoid(moid string) *DeleteTestcryptCredentialsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete testcrypt credentials moid params
func (o *DeleteTestcryptCredentialsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTestcryptCredentialsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
