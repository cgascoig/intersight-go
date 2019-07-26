// Code generated by go-swagger; DO NOT EDIT.

package appliance_data_export_policy

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

// NewGetApplianceDataExportPoliciesMoidParams creates a new GetApplianceDataExportPoliciesMoidParams object
// with the default values initialized.
func NewGetApplianceDataExportPoliciesMoidParams() *GetApplianceDataExportPoliciesMoidParams {
	var ()
	return &GetApplianceDataExportPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplianceDataExportPoliciesMoidParamsWithTimeout creates a new GetApplianceDataExportPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplianceDataExportPoliciesMoidParamsWithTimeout(timeout time.Duration) *GetApplianceDataExportPoliciesMoidParams {
	var ()
	return &GetApplianceDataExportPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewGetApplianceDataExportPoliciesMoidParamsWithContext creates a new GetApplianceDataExportPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplianceDataExportPoliciesMoidParamsWithContext(ctx context.Context) *GetApplianceDataExportPoliciesMoidParams {
	var ()
	return &GetApplianceDataExportPoliciesMoidParams{

		Context: ctx,
	}
}

// NewGetApplianceDataExportPoliciesMoidParamsWithHTTPClient creates a new GetApplianceDataExportPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplianceDataExportPoliciesMoidParamsWithHTTPClient(client *http.Client) *GetApplianceDataExportPoliciesMoidParams {
	var ()
	return &GetApplianceDataExportPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*GetApplianceDataExportPoliciesMoidParams contains all the parameters to send to the API endpoint
for the get appliance data export policies moid operation typically these are written to a http.Request
*/
type GetApplianceDataExportPoliciesMoidParams struct {

	/*Moid
	  The moid of the applianceDataExportPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get appliance data export policies moid params
func (o *GetApplianceDataExportPoliciesMoidParams) WithTimeout(timeout time.Duration) *GetApplianceDataExportPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appliance data export policies moid params
func (o *GetApplianceDataExportPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appliance data export policies moid params
func (o *GetApplianceDataExportPoliciesMoidParams) WithContext(ctx context.Context) *GetApplianceDataExportPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appliance data export policies moid params
func (o *GetApplianceDataExportPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appliance data export policies moid params
func (o *GetApplianceDataExportPoliciesMoidParams) WithHTTPClient(client *http.Client) *GetApplianceDataExportPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appliance data export policies moid params
func (o *GetApplianceDataExportPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get appliance data export policies moid params
func (o *GetApplianceDataExportPoliciesMoidParams) WithMoid(moid string) *GetApplianceDataExportPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get appliance data export policies moid params
func (o *GetApplianceDataExportPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplianceDataExportPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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