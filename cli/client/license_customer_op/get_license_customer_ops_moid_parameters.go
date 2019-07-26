// Code generated by go-swagger; DO NOT EDIT.

package license_customer_op

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

// NewGetLicenseCustomerOpsMoidParams creates a new GetLicenseCustomerOpsMoidParams object
// with the default values initialized.
func NewGetLicenseCustomerOpsMoidParams() *GetLicenseCustomerOpsMoidParams {
	var ()
	return &GetLicenseCustomerOpsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicenseCustomerOpsMoidParamsWithTimeout creates a new GetLicenseCustomerOpsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLicenseCustomerOpsMoidParamsWithTimeout(timeout time.Duration) *GetLicenseCustomerOpsMoidParams {
	var ()
	return &GetLicenseCustomerOpsMoidParams{

		timeout: timeout,
	}
}

// NewGetLicenseCustomerOpsMoidParamsWithContext creates a new GetLicenseCustomerOpsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLicenseCustomerOpsMoidParamsWithContext(ctx context.Context) *GetLicenseCustomerOpsMoidParams {
	var ()
	return &GetLicenseCustomerOpsMoidParams{

		Context: ctx,
	}
}

// NewGetLicenseCustomerOpsMoidParamsWithHTTPClient creates a new GetLicenseCustomerOpsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLicenseCustomerOpsMoidParamsWithHTTPClient(client *http.Client) *GetLicenseCustomerOpsMoidParams {
	var ()
	return &GetLicenseCustomerOpsMoidParams{
		HTTPClient: client,
	}
}

/*GetLicenseCustomerOpsMoidParams contains all the parameters to send to the API endpoint
for the get license customer ops moid operation typically these are written to a http.Request
*/
type GetLicenseCustomerOpsMoidParams struct {

	/*Moid
	  The moid of the licenseCustomerOp instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get license customer ops moid params
func (o *GetLicenseCustomerOpsMoidParams) WithTimeout(timeout time.Duration) *GetLicenseCustomerOpsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get license customer ops moid params
func (o *GetLicenseCustomerOpsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get license customer ops moid params
func (o *GetLicenseCustomerOpsMoidParams) WithContext(ctx context.Context) *GetLicenseCustomerOpsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get license customer ops moid params
func (o *GetLicenseCustomerOpsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get license customer ops moid params
func (o *GetLicenseCustomerOpsMoidParams) WithHTTPClient(client *http.Client) *GetLicenseCustomerOpsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get license customer ops moid params
func (o *GetLicenseCustomerOpsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get license customer ops moid params
func (o *GetLicenseCustomerOpsMoidParams) WithMoid(moid string) *GetLicenseCustomerOpsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get license customer ops moid params
func (o *GetLicenseCustomerOpsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicenseCustomerOpsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
