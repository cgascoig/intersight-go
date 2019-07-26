// Code generated by go-swagger; DO NOT EDIT.

package syslog_policy

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

// NewPostSyslogPoliciesParams creates a new PostSyslogPoliciesParams object
// with the default values initialized.
func NewPostSyslogPoliciesParams() *PostSyslogPoliciesParams {
	var ()
	return &PostSyslogPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSyslogPoliciesParamsWithTimeout creates a new PostSyslogPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSyslogPoliciesParamsWithTimeout(timeout time.Duration) *PostSyslogPoliciesParams {
	var ()
	return &PostSyslogPoliciesParams{

		timeout: timeout,
	}
}

// NewPostSyslogPoliciesParamsWithContext creates a new PostSyslogPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSyslogPoliciesParamsWithContext(ctx context.Context) *PostSyslogPoliciesParams {
	var ()
	return &PostSyslogPoliciesParams{

		Context: ctx,
	}
}

// NewPostSyslogPoliciesParamsWithHTTPClient creates a new PostSyslogPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSyslogPoliciesParamsWithHTTPClient(client *http.Client) *PostSyslogPoliciesParams {
	var ()
	return &PostSyslogPoliciesParams{
		HTTPClient: client,
	}
}

/*PostSyslogPoliciesParams contains all the parameters to send to the API endpoint
for the post syslog policies operation typically these are written to a http.Request
*/
type PostSyslogPoliciesParams struct {

	/*Body
	  syslogPolicy to add

	*/
	Body *models.SyslogPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post syslog policies params
func (o *PostSyslogPoliciesParams) WithTimeout(timeout time.Duration) *PostSyslogPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post syslog policies params
func (o *PostSyslogPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post syslog policies params
func (o *PostSyslogPoliciesParams) WithContext(ctx context.Context) *PostSyslogPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post syslog policies params
func (o *PostSyslogPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post syslog policies params
func (o *PostSyslogPoliciesParams) WithHTTPClient(client *http.Client) *PostSyslogPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post syslog policies params
func (o *PostSyslogPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post syslog policies params
func (o *PostSyslogPoliciesParams) WithBody(body *models.SyslogPolicy) *PostSyslogPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post syslog policies params
func (o *PostSyslogPoliciesParams) SetBody(body *models.SyslogPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSyslogPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
