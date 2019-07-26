// Code generated by go-swagger; DO NOT EDIT.

package snmp_policy

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

// NewPostSnmpPoliciesParams creates a new PostSnmpPoliciesParams object
// with the default values initialized.
func NewPostSnmpPoliciesParams() *PostSnmpPoliciesParams {
	var ()
	return &PostSnmpPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSnmpPoliciesParamsWithTimeout creates a new PostSnmpPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSnmpPoliciesParamsWithTimeout(timeout time.Duration) *PostSnmpPoliciesParams {
	var ()
	return &PostSnmpPoliciesParams{

		timeout: timeout,
	}
}

// NewPostSnmpPoliciesParamsWithContext creates a new PostSnmpPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSnmpPoliciesParamsWithContext(ctx context.Context) *PostSnmpPoliciesParams {
	var ()
	return &PostSnmpPoliciesParams{

		Context: ctx,
	}
}

// NewPostSnmpPoliciesParamsWithHTTPClient creates a new PostSnmpPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSnmpPoliciesParamsWithHTTPClient(client *http.Client) *PostSnmpPoliciesParams {
	var ()
	return &PostSnmpPoliciesParams{
		HTTPClient: client,
	}
}

/*PostSnmpPoliciesParams contains all the parameters to send to the API endpoint
for the post snmp policies operation typically these are written to a http.Request
*/
type PostSnmpPoliciesParams struct {

	/*Body
	  snmpPolicy to add

	*/
	Body *models.SnmpPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post snmp policies params
func (o *PostSnmpPoliciesParams) WithTimeout(timeout time.Duration) *PostSnmpPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post snmp policies params
func (o *PostSnmpPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post snmp policies params
func (o *PostSnmpPoliciesParams) WithContext(ctx context.Context) *PostSnmpPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post snmp policies params
func (o *PostSnmpPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post snmp policies params
func (o *PostSnmpPoliciesParams) WithHTTPClient(client *http.Client) *PostSnmpPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post snmp policies params
func (o *PostSnmpPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post snmp policies params
func (o *PostSnmpPoliciesParams) WithBody(body *models.SnmpPolicy) *PostSnmpPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post snmp policies params
func (o *PostSnmpPoliciesParams) SetBody(body *models.SnmpPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSnmpPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
