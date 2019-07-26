// Code generated by go-swagger; DO NOT EDIT.

package ipmioverlan_policy

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

// NewPostIpmioverlanPoliciesParams creates a new PostIpmioverlanPoliciesParams object
// with the default values initialized.
func NewPostIpmioverlanPoliciesParams() *PostIpmioverlanPoliciesParams {
	var ()
	return &PostIpmioverlanPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIpmioverlanPoliciesParamsWithTimeout creates a new PostIpmioverlanPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIpmioverlanPoliciesParamsWithTimeout(timeout time.Duration) *PostIpmioverlanPoliciesParams {
	var ()
	return &PostIpmioverlanPoliciesParams{

		timeout: timeout,
	}
}

// NewPostIpmioverlanPoliciesParamsWithContext creates a new PostIpmioverlanPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIpmioverlanPoliciesParamsWithContext(ctx context.Context) *PostIpmioverlanPoliciesParams {
	var ()
	return &PostIpmioverlanPoliciesParams{

		Context: ctx,
	}
}

// NewPostIpmioverlanPoliciesParamsWithHTTPClient creates a new PostIpmioverlanPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIpmioverlanPoliciesParamsWithHTTPClient(client *http.Client) *PostIpmioverlanPoliciesParams {
	var ()
	return &PostIpmioverlanPoliciesParams{
		HTTPClient: client,
	}
}

/*PostIpmioverlanPoliciesParams contains all the parameters to send to the API endpoint
for the post ipmioverlan policies operation typically these are written to a http.Request
*/
type PostIpmioverlanPoliciesParams struct {

	/*Body
	  ipmioverlanPolicy to add

	*/
	Body *models.IpmioverlanPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ipmioverlan policies params
func (o *PostIpmioverlanPoliciesParams) WithTimeout(timeout time.Duration) *PostIpmioverlanPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ipmioverlan policies params
func (o *PostIpmioverlanPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ipmioverlan policies params
func (o *PostIpmioverlanPoliciesParams) WithContext(ctx context.Context) *PostIpmioverlanPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ipmioverlan policies params
func (o *PostIpmioverlanPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ipmioverlan policies params
func (o *PostIpmioverlanPoliciesParams) WithHTTPClient(client *http.Client) *PostIpmioverlanPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ipmioverlan policies params
func (o *PostIpmioverlanPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post ipmioverlan policies params
func (o *PostIpmioverlanPoliciesParams) WithBody(body *models.IpmioverlanPolicy) *PostIpmioverlanPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post ipmioverlan policies params
func (o *PostIpmioverlanPoliciesParams) SetBody(body *models.IpmioverlanPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostIpmioverlanPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
