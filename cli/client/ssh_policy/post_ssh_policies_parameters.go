// Code generated by go-swagger; DO NOT EDIT.

package ssh_policy

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

// NewPostSSHPoliciesParams creates a new PostSSHPoliciesParams object
// with the default values initialized.
func NewPostSSHPoliciesParams() *PostSSHPoliciesParams {
	var ()
	return &PostSSHPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSSHPoliciesParamsWithTimeout creates a new PostSSHPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSSHPoliciesParamsWithTimeout(timeout time.Duration) *PostSSHPoliciesParams {
	var ()
	return &PostSSHPoliciesParams{

		timeout: timeout,
	}
}

// NewPostSSHPoliciesParamsWithContext creates a new PostSSHPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSSHPoliciesParamsWithContext(ctx context.Context) *PostSSHPoliciesParams {
	var ()
	return &PostSSHPoliciesParams{

		Context: ctx,
	}
}

// NewPostSSHPoliciesParamsWithHTTPClient creates a new PostSSHPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSSHPoliciesParamsWithHTTPClient(client *http.Client) *PostSSHPoliciesParams {
	var ()
	return &PostSSHPoliciesParams{
		HTTPClient: client,
	}
}

/*PostSSHPoliciesParams contains all the parameters to send to the API endpoint
for the post SSH policies operation typically these are written to a http.Request
*/
type PostSSHPoliciesParams struct {

	/*Body
	  sshPolicy to add

	*/
	Body *models.SSHPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post SSH policies params
func (o *PostSSHPoliciesParams) WithTimeout(timeout time.Duration) *PostSSHPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post SSH policies params
func (o *PostSSHPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post SSH policies params
func (o *PostSSHPoliciesParams) WithContext(ctx context.Context) *PostSSHPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post SSH policies params
func (o *PostSSHPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post SSH policies params
func (o *PostSSHPoliciesParams) WithHTTPClient(client *http.Client) *PostSSHPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post SSH policies params
func (o *PostSSHPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post SSH policies params
func (o *PostSSHPoliciesParams) WithBody(body *models.SSHPolicy) *PostSSHPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post SSH policies params
func (o *PostSSHPoliciesParams) SetBody(body *models.SSHPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSSHPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
