// Code generated by go-swagger; DO NOT EDIT.

package appliance_system_info

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

// NewGetApplianceSystemInfosMoidParams creates a new GetApplianceSystemInfosMoidParams object
// with the default values initialized.
func NewGetApplianceSystemInfosMoidParams() *GetApplianceSystemInfosMoidParams {
	var ()
	return &GetApplianceSystemInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplianceSystemInfosMoidParamsWithTimeout creates a new GetApplianceSystemInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplianceSystemInfosMoidParamsWithTimeout(timeout time.Duration) *GetApplianceSystemInfosMoidParams {
	var ()
	return &GetApplianceSystemInfosMoidParams{

		timeout: timeout,
	}
}

// NewGetApplianceSystemInfosMoidParamsWithContext creates a new GetApplianceSystemInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplianceSystemInfosMoidParamsWithContext(ctx context.Context) *GetApplianceSystemInfosMoidParams {
	var ()
	return &GetApplianceSystemInfosMoidParams{

		Context: ctx,
	}
}

// NewGetApplianceSystemInfosMoidParamsWithHTTPClient creates a new GetApplianceSystemInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplianceSystemInfosMoidParamsWithHTTPClient(client *http.Client) *GetApplianceSystemInfosMoidParams {
	var ()
	return &GetApplianceSystemInfosMoidParams{
		HTTPClient: client,
	}
}

/*GetApplianceSystemInfosMoidParams contains all the parameters to send to the API endpoint
for the get appliance system infos moid operation typically these are written to a http.Request
*/
type GetApplianceSystemInfosMoidParams struct {

	/*Moid
	  The moid of the applianceSystemInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get appliance system infos moid params
func (o *GetApplianceSystemInfosMoidParams) WithTimeout(timeout time.Duration) *GetApplianceSystemInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appliance system infos moid params
func (o *GetApplianceSystemInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appliance system infos moid params
func (o *GetApplianceSystemInfosMoidParams) WithContext(ctx context.Context) *GetApplianceSystemInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appliance system infos moid params
func (o *GetApplianceSystemInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appliance system infos moid params
func (o *GetApplianceSystemInfosMoidParams) WithHTTPClient(client *http.Client) *GetApplianceSystemInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appliance system infos moid params
func (o *GetApplianceSystemInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get appliance system infos moid params
func (o *GetApplianceSystemInfosMoidParams) WithMoid(moid string) *GetApplianceSystemInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get appliance system infos moid params
func (o *GetApplianceSystemInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplianceSystemInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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