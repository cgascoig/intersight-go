// Code generated by go-swagger; DO NOT EDIT.

package license_license_info

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

// NewGetLicenseLicenseInfosMoidParams creates a new GetLicenseLicenseInfosMoidParams object
// with the default values initialized.
func NewGetLicenseLicenseInfosMoidParams() *GetLicenseLicenseInfosMoidParams {
	var ()
	return &GetLicenseLicenseInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicenseLicenseInfosMoidParamsWithTimeout creates a new GetLicenseLicenseInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLicenseLicenseInfosMoidParamsWithTimeout(timeout time.Duration) *GetLicenseLicenseInfosMoidParams {
	var ()
	return &GetLicenseLicenseInfosMoidParams{

		timeout: timeout,
	}
}

// NewGetLicenseLicenseInfosMoidParamsWithContext creates a new GetLicenseLicenseInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLicenseLicenseInfosMoidParamsWithContext(ctx context.Context) *GetLicenseLicenseInfosMoidParams {
	var ()
	return &GetLicenseLicenseInfosMoidParams{

		Context: ctx,
	}
}

// NewGetLicenseLicenseInfosMoidParamsWithHTTPClient creates a new GetLicenseLicenseInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLicenseLicenseInfosMoidParamsWithHTTPClient(client *http.Client) *GetLicenseLicenseInfosMoidParams {
	var ()
	return &GetLicenseLicenseInfosMoidParams{
		HTTPClient: client,
	}
}

/*GetLicenseLicenseInfosMoidParams contains all the parameters to send to the API endpoint
for the get license license infos moid operation typically these are written to a http.Request
*/
type GetLicenseLicenseInfosMoidParams struct {

	/*Moid
	  The moid of the licenseLicenseInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get license license infos moid params
func (o *GetLicenseLicenseInfosMoidParams) WithTimeout(timeout time.Duration) *GetLicenseLicenseInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get license license infos moid params
func (o *GetLicenseLicenseInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get license license infos moid params
func (o *GetLicenseLicenseInfosMoidParams) WithContext(ctx context.Context) *GetLicenseLicenseInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get license license infos moid params
func (o *GetLicenseLicenseInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get license license infos moid params
func (o *GetLicenseLicenseInfosMoidParams) WithHTTPClient(client *http.Client) *GetLicenseLicenseInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get license license infos moid params
func (o *GetLicenseLicenseInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get license license infos moid params
func (o *GetLicenseLicenseInfosMoidParams) WithMoid(moid string) *GetLicenseLicenseInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get license license infos moid params
func (o *GetLicenseLicenseInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicenseLicenseInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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