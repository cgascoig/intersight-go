// Code generated by go-swagger; DO NOT EDIT.

package iam_end_point_role

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

// NewGetIamEndPointRolesMoidParams creates a new GetIamEndPointRolesMoidParams object
// with the default values initialized.
func NewGetIamEndPointRolesMoidParams() *GetIamEndPointRolesMoidParams {
	var ()
	return &GetIamEndPointRolesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIamEndPointRolesMoidParamsWithTimeout creates a new GetIamEndPointRolesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIamEndPointRolesMoidParamsWithTimeout(timeout time.Duration) *GetIamEndPointRolesMoidParams {
	var ()
	return &GetIamEndPointRolesMoidParams{

		timeout: timeout,
	}
}

// NewGetIamEndPointRolesMoidParamsWithContext creates a new GetIamEndPointRolesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIamEndPointRolesMoidParamsWithContext(ctx context.Context) *GetIamEndPointRolesMoidParams {
	var ()
	return &GetIamEndPointRolesMoidParams{

		Context: ctx,
	}
}

// NewGetIamEndPointRolesMoidParamsWithHTTPClient creates a new GetIamEndPointRolesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIamEndPointRolesMoidParamsWithHTTPClient(client *http.Client) *GetIamEndPointRolesMoidParams {
	var ()
	return &GetIamEndPointRolesMoidParams{
		HTTPClient: client,
	}
}

/*GetIamEndPointRolesMoidParams contains all the parameters to send to the API endpoint
for the get iam end point roles moid operation typically these are written to a http.Request
*/
type GetIamEndPointRolesMoidParams struct {

	/*Moid
	  The moid of the iamEndPointRole instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iam end point roles moid params
func (o *GetIamEndPointRolesMoidParams) WithTimeout(timeout time.Duration) *GetIamEndPointRolesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iam end point roles moid params
func (o *GetIamEndPointRolesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iam end point roles moid params
func (o *GetIamEndPointRolesMoidParams) WithContext(ctx context.Context) *GetIamEndPointRolesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iam end point roles moid params
func (o *GetIamEndPointRolesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iam end point roles moid params
func (o *GetIamEndPointRolesMoidParams) WithHTTPClient(client *http.Client) *GetIamEndPointRolesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iam end point roles moid params
func (o *GetIamEndPointRolesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get iam end point roles moid params
func (o *GetIamEndPointRolesMoidParams) WithMoid(moid string) *GetIamEndPointRolesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get iam end point roles moid params
func (o *GetIamEndPointRolesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetIamEndPointRolesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
