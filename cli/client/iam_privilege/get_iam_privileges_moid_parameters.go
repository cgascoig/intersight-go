// Code generated by go-swagger; DO NOT EDIT.

package iam_privilege

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

// NewGetIamPrivilegesMoidParams creates a new GetIamPrivilegesMoidParams object
// with the default values initialized.
func NewGetIamPrivilegesMoidParams() *GetIamPrivilegesMoidParams {
	var ()
	return &GetIamPrivilegesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIamPrivilegesMoidParamsWithTimeout creates a new GetIamPrivilegesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIamPrivilegesMoidParamsWithTimeout(timeout time.Duration) *GetIamPrivilegesMoidParams {
	var ()
	return &GetIamPrivilegesMoidParams{

		timeout: timeout,
	}
}

// NewGetIamPrivilegesMoidParamsWithContext creates a new GetIamPrivilegesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIamPrivilegesMoidParamsWithContext(ctx context.Context) *GetIamPrivilegesMoidParams {
	var ()
	return &GetIamPrivilegesMoidParams{

		Context: ctx,
	}
}

// NewGetIamPrivilegesMoidParamsWithHTTPClient creates a new GetIamPrivilegesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIamPrivilegesMoidParamsWithHTTPClient(client *http.Client) *GetIamPrivilegesMoidParams {
	var ()
	return &GetIamPrivilegesMoidParams{
		HTTPClient: client,
	}
}

/*GetIamPrivilegesMoidParams contains all the parameters to send to the API endpoint
for the get iam privileges moid operation typically these are written to a http.Request
*/
type GetIamPrivilegesMoidParams struct {

	/*Moid
	  The moid of the iamPrivilege instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iam privileges moid params
func (o *GetIamPrivilegesMoidParams) WithTimeout(timeout time.Duration) *GetIamPrivilegesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iam privileges moid params
func (o *GetIamPrivilegesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iam privileges moid params
func (o *GetIamPrivilegesMoidParams) WithContext(ctx context.Context) *GetIamPrivilegesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iam privileges moid params
func (o *GetIamPrivilegesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iam privileges moid params
func (o *GetIamPrivilegesMoidParams) WithHTTPClient(client *http.Client) *GetIamPrivilegesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iam privileges moid params
func (o *GetIamPrivilegesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get iam privileges moid params
func (o *GetIamPrivilegesMoidParams) WithMoid(moid string) *GetIamPrivilegesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get iam privileges moid params
func (o *GetIamPrivilegesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetIamPrivilegesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
