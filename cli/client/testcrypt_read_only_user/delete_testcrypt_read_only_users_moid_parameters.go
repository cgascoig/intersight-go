// Code generated by go-swagger; DO NOT EDIT.

package testcrypt_read_only_user

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

// NewDeleteTestcryptReadOnlyUsersMoidParams creates a new DeleteTestcryptReadOnlyUsersMoidParams object
// with the default values initialized.
func NewDeleteTestcryptReadOnlyUsersMoidParams() *DeleteTestcryptReadOnlyUsersMoidParams {
	var ()
	return &DeleteTestcryptReadOnlyUsersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTestcryptReadOnlyUsersMoidParamsWithTimeout creates a new DeleteTestcryptReadOnlyUsersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTestcryptReadOnlyUsersMoidParamsWithTimeout(timeout time.Duration) *DeleteTestcryptReadOnlyUsersMoidParams {
	var ()
	return &DeleteTestcryptReadOnlyUsersMoidParams{

		timeout: timeout,
	}
}

// NewDeleteTestcryptReadOnlyUsersMoidParamsWithContext creates a new DeleteTestcryptReadOnlyUsersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTestcryptReadOnlyUsersMoidParamsWithContext(ctx context.Context) *DeleteTestcryptReadOnlyUsersMoidParams {
	var ()
	return &DeleteTestcryptReadOnlyUsersMoidParams{

		Context: ctx,
	}
}

// NewDeleteTestcryptReadOnlyUsersMoidParamsWithHTTPClient creates a new DeleteTestcryptReadOnlyUsersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTestcryptReadOnlyUsersMoidParamsWithHTTPClient(client *http.Client) *DeleteTestcryptReadOnlyUsersMoidParams {
	var ()
	return &DeleteTestcryptReadOnlyUsersMoidParams{
		HTTPClient: client,
	}
}

/*DeleteTestcryptReadOnlyUsersMoidParams contains all the parameters to send to the API endpoint
for the delete testcrypt read only users moid operation typically these are written to a http.Request
*/
type DeleteTestcryptReadOnlyUsersMoidParams struct {

	/*Moid
	  The moid of the testcryptReadOnlyUser instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete testcrypt read only users moid params
func (o *DeleteTestcryptReadOnlyUsersMoidParams) WithTimeout(timeout time.Duration) *DeleteTestcryptReadOnlyUsersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete testcrypt read only users moid params
func (o *DeleteTestcryptReadOnlyUsersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete testcrypt read only users moid params
func (o *DeleteTestcryptReadOnlyUsersMoidParams) WithContext(ctx context.Context) *DeleteTestcryptReadOnlyUsersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete testcrypt read only users moid params
func (o *DeleteTestcryptReadOnlyUsersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete testcrypt read only users moid params
func (o *DeleteTestcryptReadOnlyUsersMoidParams) WithHTTPClient(client *http.Client) *DeleteTestcryptReadOnlyUsersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete testcrypt read only users moid params
func (o *DeleteTestcryptReadOnlyUsersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete testcrypt read only users moid params
func (o *DeleteTestcryptReadOnlyUsersMoidParams) WithMoid(moid string) *DeleteTestcryptReadOnlyUsersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete testcrypt read only users moid params
func (o *DeleteTestcryptReadOnlyUsersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTestcryptReadOnlyUsersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
