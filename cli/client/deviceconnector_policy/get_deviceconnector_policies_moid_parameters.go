// Code generated by go-swagger; DO NOT EDIT.

package deviceconnector_policy

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

// NewGetDeviceconnectorPoliciesMoidParams creates a new GetDeviceconnectorPoliciesMoidParams object
// with the default values initialized.
func NewGetDeviceconnectorPoliciesMoidParams() *GetDeviceconnectorPoliciesMoidParams {
	var ()
	return &GetDeviceconnectorPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceconnectorPoliciesMoidParamsWithTimeout creates a new GetDeviceconnectorPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceconnectorPoliciesMoidParamsWithTimeout(timeout time.Duration) *GetDeviceconnectorPoliciesMoidParams {
	var ()
	return &GetDeviceconnectorPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewGetDeviceconnectorPoliciesMoidParamsWithContext creates a new GetDeviceconnectorPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceconnectorPoliciesMoidParamsWithContext(ctx context.Context) *GetDeviceconnectorPoliciesMoidParams {
	var ()
	return &GetDeviceconnectorPoliciesMoidParams{

		Context: ctx,
	}
}

// NewGetDeviceconnectorPoliciesMoidParamsWithHTTPClient creates a new GetDeviceconnectorPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceconnectorPoliciesMoidParamsWithHTTPClient(client *http.Client) *GetDeviceconnectorPoliciesMoidParams {
	var ()
	return &GetDeviceconnectorPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*GetDeviceconnectorPoliciesMoidParams contains all the parameters to send to the API endpoint
for the get deviceconnector policies moid operation typically these are written to a http.Request
*/
type GetDeviceconnectorPoliciesMoidParams struct {

	/*Moid
	  The moid of the deviceconnectorPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deviceconnector policies moid params
func (o *GetDeviceconnectorPoliciesMoidParams) WithTimeout(timeout time.Duration) *GetDeviceconnectorPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deviceconnector policies moid params
func (o *GetDeviceconnectorPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deviceconnector policies moid params
func (o *GetDeviceconnectorPoliciesMoidParams) WithContext(ctx context.Context) *GetDeviceconnectorPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deviceconnector policies moid params
func (o *GetDeviceconnectorPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deviceconnector policies moid params
func (o *GetDeviceconnectorPoliciesMoidParams) WithHTTPClient(client *http.Client) *GetDeviceconnectorPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deviceconnector policies moid params
func (o *GetDeviceconnectorPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get deviceconnector policies moid params
func (o *GetDeviceconnectorPoliciesMoidParams) WithMoid(moid string) *GetDeviceconnectorPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get deviceconnector policies moid params
func (o *GetDeviceconnectorPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceconnectorPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
