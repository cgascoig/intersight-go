// Code generated by go-swagger; DO NOT EDIT.

package iam_domain_group

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

// NewGetIamDomainGroupsMoidParams creates a new GetIamDomainGroupsMoidParams object
// with the default values initialized.
func NewGetIamDomainGroupsMoidParams() *GetIamDomainGroupsMoidParams {
	var ()
	return &GetIamDomainGroupsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIamDomainGroupsMoidParamsWithTimeout creates a new GetIamDomainGroupsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIamDomainGroupsMoidParamsWithTimeout(timeout time.Duration) *GetIamDomainGroupsMoidParams {
	var ()
	return &GetIamDomainGroupsMoidParams{

		timeout: timeout,
	}
}

// NewGetIamDomainGroupsMoidParamsWithContext creates a new GetIamDomainGroupsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIamDomainGroupsMoidParamsWithContext(ctx context.Context) *GetIamDomainGroupsMoidParams {
	var ()
	return &GetIamDomainGroupsMoidParams{

		Context: ctx,
	}
}

// NewGetIamDomainGroupsMoidParamsWithHTTPClient creates a new GetIamDomainGroupsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIamDomainGroupsMoidParamsWithHTTPClient(client *http.Client) *GetIamDomainGroupsMoidParams {
	var ()
	return &GetIamDomainGroupsMoidParams{
		HTTPClient: client,
	}
}

/*GetIamDomainGroupsMoidParams contains all the parameters to send to the API endpoint
for the get iam domain groups moid operation typically these are written to a http.Request
*/
type GetIamDomainGroupsMoidParams struct {

	/*Moid
	  The moid of the iamDomainGroup instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iam domain groups moid params
func (o *GetIamDomainGroupsMoidParams) WithTimeout(timeout time.Duration) *GetIamDomainGroupsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iam domain groups moid params
func (o *GetIamDomainGroupsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iam domain groups moid params
func (o *GetIamDomainGroupsMoidParams) WithContext(ctx context.Context) *GetIamDomainGroupsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iam domain groups moid params
func (o *GetIamDomainGroupsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iam domain groups moid params
func (o *GetIamDomainGroupsMoidParams) WithHTTPClient(client *http.Client) *GetIamDomainGroupsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iam domain groups moid params
func (o *GetIamDomainGroupsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get iam domain groups moid params
func (o *GetIamDomainGroupsMoidParams) WithMoid(moid string) *GetIamDomainGroupsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get iam domain groups moid params
func (o *GetIamDomainGroupsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetIamDomainGroupsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
