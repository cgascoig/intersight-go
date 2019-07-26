// Code generated by go-swagger; DO NOT EDIT.

package boot_precision_policy

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

// NewPostBootPrecisionPoliciesMoidParams creates a new PostBootPrecisionPoliciesMoidParams object
// with the default values initialized.
func NewPostBootPrecisionPoliciesMoidParams() *PostBootPrecisionPoliciesMoidParams {
	var ()
	return &PostBootPrecisionPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostBootPrecisionPoliciesMoidParamsWithTimeout creates a new PostBootPrecisionPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostBootPrecisionPoliciesMoidParamsWithTimeout(timeout time.Duration) *PostBootPrecisionPoliciesMoidParams {
	var ()
	return &PostBootPrecisionPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostBootPrecisionPoliciesMoidParamsWithContext creates a new PostBootPrecisionPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostBootPrecisionPoliciesMoidParamsWithContext(ctx context.Context) *PostBootPrecisionPoliciesMoidParams {
	var ()
	return &PostBootPrecisionPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostBootPrecisionPoliciesMoidParamsWithHTTPClient creates a new PostBootPrecisionPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostBootPrecisionPoliciesMoidParamsWithHTTPClient(client *http.Client) *PostBootPrecisionPoliciesMoidParams {
	var ()
	return &PostBootPrecisionPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostBootPrecisionPoliciesMoidParams contains all the parameters to send to the API endpoint
for the post boot precision policies moid operation typically these are written to a http.Request
*/
type PostBootPrecisionPoliciesMoidParams struct {

	/*Body
	  bootPrecisionPolicy to update

	*/
	Body *models.BootPrecisionPolicy
	/*Moid
	  The moid of the bootPrecisionPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) WithTimeout(timeout time.Duration) *PostBootPrecisionPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) WithContext(ctx context.Context) *PostBootPrecisionPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) WithHTTPClient(client *http.Client) *PostBootPrecisionPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) WithBody(body *models.BootPrecisionPolicy) *PostBootPrecisionPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) SetBody(body *models.BootPrecisionPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) WithMoid(moid string) *PostBootPrecisionPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post boot precision policies moid params
func (o *PostBootPrecisionPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostBootPrecisionPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
