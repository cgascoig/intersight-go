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
)

// NewGetKvmPoliciesMoidParams creates a new GetKvmPoliciesMoidParams object
// with the default values initialized.
func NewGetKvmPoliciesMoidParams() *GetKvmPoliciesMoidParams {
	var ()
	return &GetKvmPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKvmPoliciesMoidParamsWithTimeout creates a new GetKvmPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKvmPoliciesMoidParamsWithTimeout(timeout time.Duration) *GetKvmPoliciesMoidParams {
	var ()
	return &GetKvmPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewGetKvmPoliciesMoidParamsWithContext creates a new GetKvmPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKvmPoliciesMoidParamsWithContext(ctx context.Context) *GetKvmPoliciesMoidParams {
	var ()
	return &GetKvmPoliciesMoidParams{

		Context: ctx,
	}
}

// NewGetKvmPoliciesMoidParamsWithHTTPClient creates a new GetKvmPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKvmPoliciesMoidParamsWithHTTPClient(client *http.Client) *GetKvmPoliciesMoidParams {
	var ()
	return &GetKvmPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*GetKvmPoliciesMoidParams contains all the parameters to send to the API endpoint
for the get kvm policies moid operation typically these are written to a http.Request
*/
type GetKvmPoliciesMoidParams struct {

	/*Moid
	  The moid of the kvmPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kvm policies moid params
func (o *GetKvmPoliciesMoidParams) WithTimeout(timeout time.Duration) *GetKvmPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kvm policies moid params
func (o *GetKvmPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kvm policies moid params
func (o *GetKvmPoliciesMoidParams) WithContext(ctx context.Context) *GetKvmPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kvm policies moid params
func (o *GetKvmPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kvm policies moid params
func (o *GetKvmPoliciesMoidParams) WithHTTPClient(client *http.Client) *GetKvmPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kvm policies moid params
func (o *GetKvmPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get kvm policies moid params
func (o *GetKvmPoliciesMoidParams) WithMoid(moid string) *GetKvmPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get kvm policies moid params
func (o *GetKvmPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetKvmPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
