// Code generated by go-swagger; DO NOT EDIT.

package sol_policy

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

// NewPostSolPoliciesMoidParams creates a new PostSolPoliciesMoidParams object
// with the default values initialized.
func NewPostSolPoliciesMoidParams() *PostSolPoliciesMoidParams {
	var ()
	return &PostSolPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSolPoliciesMoidParamsWithTimeout creates a new PostSolPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSolPoliciesMoidParamsWithTimeout(timeout time.Duration) *PostSolPoliciesMoidParams {
	var ()
	return &PostSolPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostSolPoliciesMoidParamsWithContext creates a new PostSolPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSolPoliciesMoidParamsWithContext(ctx context.Context) *PostSolPoliciesMoidParams {
	var ()
	return &PostSolPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostSolPoliciesMoidParamsWithHTTPClient creates a new PostSolPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSolPoliciesMoidParamsWithHTTPClient(client *http.Client) *PostSolPoliciesMoidParams {
	var ()
	return &PostSolPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostSolPoliciesMoidParams contains all the parameters to send to the API endpoint
for the post sol policies moid operation typically these are written to a http.Request
*/
type PostSolPoliciesMoidParams struct {

	/*Body
	  solPolicy to update

	*/
	Body *models.SolPolicy
	/*Moid
	  The moid of the solPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) WithTimeout(timeout time.Duration) *PostSolPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) WithContext(ctx context.Context) *PostSolPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) WithHTTPClient(client *http.Client) *PostSolPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) WithBody(body *models.SolPolicy) *PostSolPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) SetBody(body *models.SolPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) WithMoid(moid string) *PostSolPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post sol policies moid params
func (o *PostSolPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostSolPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
