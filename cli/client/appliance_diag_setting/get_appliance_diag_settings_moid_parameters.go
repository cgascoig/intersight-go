// Code generated by go-swagger; DO NOT EDIT.

package appliance_diag_setting

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

// NewGetApplianceDiagSettingsMoidParams creates a new GetApplianceDiagSettingsMoidParams object
// with the default values initialized.
func NewGetApplianceDiagSettingsMoidParams() *GetApplianceDiagSettingsMoidParams {
	var ()
	return &GetApplianceDiagSettingsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplianceDiagSettingsMoidParamsWithTimeout creates a new GetApplianceDiagSettingsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplianceDiagSettingsMoidParamsWithTimeout(timeout time.Duration) *GetApplianceDiagSettingsMoidParams {
	var ()
	return &GetApplianceDiagSettingsMoidParams{

		timeout: timeout,
	}
}

// NewGetApplianceDiagSettingsMoidParamsWithContext creates a new GetApplianceDiagSettingsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplianceDiagSettingsMoidParamsWithContext(ctx context.Context) *GetApplianceDiagSettingsMoidParams {
	var ()
	return &GetApplianceDiagSettingsMoidParams{

		Context: ctx,
	}
}

// NewGetApplianceDiagSettingsMoidParamsWithHTTPClient creates a new GetApplianceDiagSettingsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplianceDiagSettingsMoidParamsWithHTTPClient(client *http.Client) *GetApplianceDiagSettingsMoidParams {
	var ()
	return &GetApplianceDiagSettingsMoidParams{
		HTTPClient: client,
	}
}

/*GetApplianceDiagSettingsMoidParams contains all the parameters to send to the API endpoint
for the get appliance diag settings moid operation typically these are written to a http.Request
*/
type GetApplianceDiagSettingsMoidParams struct {

	/*Moid
	  The moid of the applianceDiagSetting instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get appliance diag settings moid params
func (o *GetApplianceDiagSettingsMoidParams) WithTimeout(timeout time.Duration) *GetApplianceDiagSettingsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appliance diag settings moid params
func (o *GetApplianceDiagSettingsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appliance diag settings moid params
func (o *GetApplianceDiagSettingsMoidParams) WithContext(ctx context.Context) *GetApplianceDiagSettingsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appliance diag settings moid params
func (o *GetApplianceDiagSettingsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appliance diag settings moid params
func (o *GetApplianceDiagSettingsMoidParams) WithHTTPClient(client *http.Client) *GetApplianceDiagSettingsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appliance diag settings moid params
func (o *GetApplianceDiagSettingsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get appliance diag settings moid params
func (o *GetApplianceDiagSettingsMoidParams) WithMoid(moid string) *GetApplianceDiagSettingsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get appliance diag settings moid params
func (o *GetApplianceDiagSettingsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplianceDiagSettingsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
