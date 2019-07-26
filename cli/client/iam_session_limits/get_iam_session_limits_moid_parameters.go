// Code generated by go-swagger; DO NOT EDIT.

package iam_session_limits

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

// NewGetIamSessionLimitsMoidParams creates a new GetIamSessionLimitsMoidParams object
// with the default values initialized.
func NewGetIamSessionLimitsMoidParams() *GetIamSessionLimitsMoidParams {
	var ()
	return &GetIamSessionLimitsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIamSessionLimitsMoidParamsWithTimeout creates a new GetIamSessionLimitsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIamSessionLimitsMoidParamsWithTimeout(timeout time.Duration) *GetIamSessionLimitsMoidParams {
	var ()
	return &GetIamSessionLimitsMoidParams{

		timeout: timeout,
	}
}

// NewGetIamSessionLimitsMoidParamsWithContext creates a new GetIamSessionLimitsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIamSessionLimitsMoidParamsWithContext(ctx context.Context) *GetIamSessionLimitsMoidParams {
	var ()
	return &GetIamSessionLimitsMoidParams{

		Context: ctx,
	}
}

// NewGetIamSessionLimitsMoidParamsWithHTTPClient creates a new GetIamSessionLimitsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIamSessionLimitsMoidParamsWithHTTPClient(client *http.Client) *GetIamSessionLimitsMoidParams {
	var ()
	return &GetIamSessionLimitsMoidParams{
		HTTPClient: client,
	}
}

/*GetIamSessionLimitsMoidParams contains all the parameters to send to the API endpoint
for the get iam session limits moid operation typically these are written to a http.Request
*/
type GetIamSessionLimitsMoidParams struct {

	/*Moid
	  The moid of the iamSessionLimits instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iam session limits moid params
func (o *GetIamSessionLimitsMoidParams) WithTimeout(timeout time.Duration) *GetIamSessionLimitsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iam session limits moid params
func (o *GetIamSessionLimitsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iam session limits moid params
func (o *GetIamSessionLimitsMoidParams) WithContext(ctx context.Context) *GetIamSessionLimitsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iam session limits moid params
func (o *GetIamSessionLimitsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iam session limits moid params
func (o *GetIamSessionLimitsMoidParams) WithHTTPClient(client *http.Client) *GetIamSessionLimitsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iam session limits moid params
func (o *GetIamSessionLimitsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get iam session limits moid params
func (o *GetIamSessionLimitsMoidParams) WithMoid(moid string) *GetIamSessionLimitsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get iam session limits moid params
func (o *GetIamSessionLimitsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetIamSessionLimitsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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