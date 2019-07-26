// Code generated by go-swagger; DO NOT EDIT.

package kvm_policy

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

// NewPostKvmPoliciesParams creates a new PostKvmPoliciesParams object
// with the default values initialized.
func NewPostKvmPoliciesParams() *PostKvmPoliciesParams {
	var ()
	return &PostKvmPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostKvmPoliciesParamsWithTimeout creates a new PostKvmPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostKvmPoliciesParamsWithTimeout(timeout time.Duration) *PostKvmPoliciesParams {
	var ()
	return &PostKvmPoliciesParams{

		timeout: timeout,
	}
}

// NewPostKvmPoliciesParamsWithContext creates a new PostKvmPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostKvmPoliciesParamsWithContext(ctx context.Context) *PostKvmPoliciesParams {
	var ()
	return &PostKvmPoliciesParams{

		Context: ctx,
	}
}

// NewPostKvmPoliciesParamsWithHTTPClient creates a new PostKvmPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostKvmPoliciesParamsWithHTTPClient(client *http.Client) *PostKvmPoliciesParams {
	var ()
	return &PostKvmPoliciesParams{
		HTTPClient: client,
	}
}

/*PostKvmPoliciesParams contains all the parameters to send to the API endpoint
for the post kvm policies operation typically these are written to a http.Request
*/
type PostKvmPoliciesParams struct {

	/*Body
	  kvmPolicy to add

	*/
	Body *models.KvmPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post kvm policies params
func (o *PostKvmPoliciesParams) WithTimeout(timeout time.Duration) *PostKvmPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post kvm policies params
func (o *PostKvmPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post kvm policies params
func (o *PostKvmPoliciesParams) WithContext(ctx context.Context) *PostKvmPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post kvm policies params
func (o *PostKvmPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post kvm policies params
func (o *PostKvmPoliciesParams) WithHTTPClient(client *http.Client) *PostKvmPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post kvm policies params
func (o *PostKvmPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post kvm policies params
func (o *PostKvmPoliciesParams) WithBody(body *models.KvmPolicy) *PostKvmPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post kvm policies params
func (o *PostKvmPoliciesParams) SetBody(body *models.KvmPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostKvmPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
