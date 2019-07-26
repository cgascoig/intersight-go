// Code generated by go-swagger; DO NOT EDIT.

package security_unit

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

// NewGetSecurityUnitsMoidParams creates a new GetSecurityUnitsMoidParams object
// with the default values initialized.
func NewGetSecurityUnitsMoidParams() *GetSecurityUnitsMoidParams {
	var ()
	return &GetSecurityUnitsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecurityUnitsMoidParamsWithTimeout creates a new GetSecurityUnitsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSecurityUnitsMoidParamsWithTimeout(timeout time.Duration) *GetSecurityUnitsMoidParams {
	var ()
	return &GetSecurityUnitsMoidParams{

		timeout: timeout,
	}
}

// NewGetSecurityUnitsMoidParamsWithContext creates a new GetSecurityUnitsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSecurityUnitsMoidParamsWithContext(ctx context.Context) *GetSecurityUnitsMoidParams {
	var ()
	return &GetSecurityUnitsMoidParams{

		Context: ctx,
	}
}

// NewGetSecurityUnitsMoidParamsWithHTTPClient creates a new GetSecurityUnitsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSecurityUnitsMoidParamsWithHTTPClient(client *http.Client) *GetSecurityUnitsMoidParams {
	var ()
	return &GetSecurityUnitsMoidParams{
		HTTPClient: client,
	}
}

/*GetSecurityUnitsMoidParams contains all the parameters to send to the API endpoint
for the get security units moid operation typically these are written to a http.Request
*/
type GetSecurityUnitsMoidParams struct {

	/*Moid
	  The moid of the securityUnit instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get security units moid params
func (o *GetSecurityUnitsMoidParams) WithTimeout(timeout time.Duration) *GetSecurityUnitsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get security units moid params
func (o *GetSecurityUnitsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get security units moid params
func (o *GetSecurityUnitsMoidParams) WithContext(ctx context.Context) *GetSecurityUnitsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get security units moid params
func (o *GetSecurityUnitsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get security units moid params
func (o *GetSecurityUnitsMoidParams) WithHTTPClient(client *http.Client) *GetSecurityUnitsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get security units moid params
func (o *GetSecurityUnitsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get security units moid params
func (o *GetSecurityUnitsMoidParams) WithMoid(moid string) *GetSecurityUnitsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get security units moid params
func (o *GetSecurityUnitsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecurityUnitsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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