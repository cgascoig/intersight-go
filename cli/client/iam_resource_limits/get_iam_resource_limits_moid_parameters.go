// Code generated by go-swagger; DO NOT EDIT.

package iam_resource_limits

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

// NewGetIamResourceLimitsMoidParams creates a new GetIamResourceLimitsMoidParams object
// with the default values initialized.
func NewGetIamResourceLimitsMoidParams() *GetIamResourceLimitsMoidParams {
	var ()
	return &GetIamResourceLimitsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIamResourceLimitsMoidParamsWithTimeout creates a new GetIamResourceLimitsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIamResourceLimitsMoidParamsWithTimeout(timeout time.Duration) *GetIamResourceLimitsMoidParams {
	var ()
	return &GetIamResourceLimitsMoidParams{

		timeout: timeout,
	}
}

// NewGetIamResourceLimitsMoidParamsWithContext creates a new GetIamResourceLimitsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIamResourceLimitsMoidParamsWithContext(ctx context.Context) *GetIamResourceLimitsMoidParams {
	var ()
	return &GetIamResourceLimitsMoidParams{

		Context: ctx,
	}
}

// NewGetIamResourceLimitsMoidParamsWithHTTPClient creates a new GetIamResourceLimitsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIamResourceLimitsMoidParamsWithHTTPClient(client *http.Client) *GetIamResourceLimitsMoidParams {
	var ()
	return &GetIamResourceLimitsMoidParams{
		HTTPClient: client,
	}
}

/*GetIamResourceLimitsMoidParams contains all the parameters to send to the API endpoint
for the get iam resource limits moid operation typically these are written to a http.Request
*/
type GetIamResourceLimitsMoidParams struct {

	/*Moid
	  The moid of the iamResourceLimits instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iam resource limits moid params
func (o *GetIamResourceLimitsMoidParams) WithTimeout(timeout time.Duration) *GetIamResourceLimitsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iam resource limits moid params
func (o *GetIamResourceLimitsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iam resource limits moid params
func (o *GetIamResourceLimitsMoidParams) WithContext(ctx context.Context) *GetIamResourceLimitsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iam resource limits moid params
func (o *GetIamResourceLimitsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iam resource limits moid params
func (o *GetIamResourceLimitsMoidParams) WithHTTPClient(client *http.Client) *GetIamResourceLimitsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iam resource limits moid params
func (o *GetIamResourceLimitsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get iam resource limits moid params
func (o *GetIamResourceLimitsMoidParams) WithMoid(moid string) *GetIamResourceLimitsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get iam resource limits moid params
func (o *GetIamResourceLimitsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetIamResourceLimitsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
