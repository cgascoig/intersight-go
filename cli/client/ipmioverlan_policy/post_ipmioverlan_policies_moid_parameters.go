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

// NewPostIpmioverlanPoliciesMoidParams creates a new PostIpmioverlanPoliciesMoidParams object
// with the default values initialized.
func NewPostIpmioverlanPoliciesMoidParams() *PostIpmioverlanPoliciesMoidParams {
	var ()
	return &PostIpmioverlanPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIpmioverlanPoliciesMoidParamsWithTimeout creates a new PostIpmioverlanPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIpmioverlanPoliciesMoidParamsWithTimeout(timeout time.Duration) *PostIpmioverlanPoliciesMoidParams {
	var ()
	return &PostIpmioverlanPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostIpmioverlanPoliciesMoidParamsWithContext creates a new PostIpmioverlanPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIpmioverlanPoliciesMoidParamsWithContext(ctx context.Context) *PostIpmioverlanPoliciesMoidParams {
	var ()
	return &PostIpmioverlanPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostIpmioverlanPoliciesMoidParamsWithHTTPClient creates a new PostIpmioverlanPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIpmioverlanPoliciesMoidParamsWithHTTPClient(client *http.Client) *PostIpmioverlanPoliciesMoidParams {
	var ()
	return &PostIpmioverlanPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostIpmioverlanPoliciesMoidParams contains all the parameters to send to the API endpoint
for the post ipmioverlan policies moid operation typically these are written to a http.Request
*/
type PostIpmioverlanPoliciesMoidParams struct {

	/*Body
	  ipmioverlanPolicy to update

	*/
	Body *models.IpmioverlanPolicy
	/*Moid
	  The moid of the ipmioverlanPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) WithTimeout(timeout time.Duration) *PostIpmioverlanPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) WithContext(ctx context.Context) *PostIpmioverlanPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) WithHTTPClient(client *http.Client) *PostIpmioverlanPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) WithBody(body *models.IpmioverlanPolicy) *PostIpmioverlanPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) SetBody(body *models.IpmioverlanPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) WithMoid(moid string) *PostIpmioverlanPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post ipmioverlan policies moid params
func (o *PostIpmioverlanPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostIpmioverlanPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
