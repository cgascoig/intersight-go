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

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPostTestcryptReadOnlyUsersMoidParams creates a new PostTestcryptReadOnlyUsersMoidParams object
// with the default values initialized.
func NewPostTestcryptReadOnlyUsersMoidParams() *PostTestcryptReadOnlyUsersMoidParams {
	var ()
	return &PostTestcryptReadOnlyUsersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTestcryptReadOnlyUsersMoidParamsWithTimeout creates a new PostTestcryptReadOnlyUsersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTestcryptReadOnlyUsersMoidParamsWithTimeout(timeout time.Duration) *PostTestcryptReadOnlyUsersMoidParams {
	var ()
	return &PostTestcryptReadOnlyUsersMoidParams{

		timeout: timeout,
	}
}

// NewPostTestcryptReadOnlyUsersMoidParamsWithContext creates a new PostTestcryptReadOnlyUsersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTestcryptReadOnlyUsersMoidParamsWithContext(ctx context.Context) *PostTestcryptReadOnlyUsersMoidParams {
	var ()
	return &PostTestcryptReadOnlyUsersMoidParams{

		Context: ctx,
	}
}

// NewPostTestcryptReadOnlyUsersMoidParamsWithHTTPClient creates a new PostTestcryptReadOnlyUsersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTestcryptReadOnlyUsersMoidParamsWithHTTPClient(client *http.Client) *PostTestcryptReadOnlyUsersMoidParams {
	var ()
	return &PostTestcryptReadOnlyUsersMoidParams{
		HTTPClient: client,
	}
}

/*PostTestcryptReadOnlyUsersMoidParams contains all the parameters to send to the API endpoint
for the post testcrypt read only users moid operation typically these are written to a http.Request
*/
type PostTestcryptReadOnlyUsersMoidParams struct {

	/*Body
	  testcryptReadOnlyUser to update

	*/
	Body *models.TestcryptReadOnlyUser
	/*Moid
	  The moid of the testcryptReadOnlyUser instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) WithTimeout(timeout time.Duration) *PostTestcryptReadOnlyUsersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) WithContext(ctx context.Context) *PostTestcryptReadOnlyUsersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) WithHTTPClient(client *http.Client) *PostTestcryptReadOnlyUsersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) WithBody(body *models.TestcryptReadOnlyUser) *PostTestcryptReadOnlyUsersMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) SetBody(body *models.TestcryptReadOnlyUser) {
	o.Body = body
}

// WithMoid adds the moid to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) WithMoid(moid string) *PostTestcryptReadOnlyUsersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post testcrypt read only users moid params
func (o *PostTestcryptReadOnlyUsersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostTestcryptReadOnlyUsersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
