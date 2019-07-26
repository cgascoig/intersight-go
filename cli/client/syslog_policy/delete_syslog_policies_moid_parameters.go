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
)

// NewDeleteSyslogPoliciesMoidParams creates a new DeleteSyslogPoliciesMoidParams object
// with the default values initialized.
func NewDeleteSyslogPoliciesMoidParams() *DeleteSyslogPoliciesMoidParams {
	var ()
	return &DeleteSyslogPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSyslogPoliciesMoidParamsWithTimeout creates a new DeleteSyslogPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSyslogPoliciesMoidParamsWithTimeout(timeout time.Duration) *DeleteSyslogPoliciesMoidParams {
	var ()
	return &DeleteSyslogPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewDeleteSyslogPoliciesMoidParamsWithContext creates a new DeleteSyslogPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSyslogPoliciesMoidParamsWithContext(ctx context.Context) *DeleteSyslogPoliciesMoidParams {
	var ()
	return &DeleteSyslogPoliciesMoidParams{

		Context: ctx,
	}
}

// NewDeleteSyslogPoliciesMoidParamsWithHTTPClient creates a new DeleteSyslogPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSyslogPoliciesMoidParamsWithHTTPClient(client *http.Client) *DeleteSyslogPoliciesMoidParams {
	var ()
	return &DeleteSyslogPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*DeleteSyslogPoliciesMoidParams contains all the parameters to send to the API endpoint
for the delete syslog policies moid operation typically these are written to a http.Request
*/
type DeleteSyslogPoliciesMoidParams struct {

	/*Moid
	  The moid of the syslogPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete syslog policies moid params
func (o *DeleteSyslogPoliciesMoidParams) WithTimeout(timeout time.Duration) *DeleteSyslogPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete syslog policies moid params
func (o *DeleteSyslogPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete syslog policies moid params
func (o *DeleteSyslogPoliciesMoidParams) WithContext(ctx context.Context) *DeleteSyslogPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete syslog policies moid params
func (o *DeleteSyslogPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete syslog policies moid params
func (o *DeleteSyslogPoliciesMoidParams) WithHTTPClient(client *http.Client) *DeleteSyslogPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete syslog policies moid params
func (o *DeleteSyslogPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete syslog policies moid params
func (o *DeleteSyslogPoliciesMoidParams) WithMoid(moid string) *DeleteSyslogPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete syslog policies moid params
func (o *DeleteSyslogPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSyslogPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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