// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_adapter_policy

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

// NewGetVnicFcAdapterPoliciesMoidParams creates a new GetVnicFcAdapterPoliciesMoidParams object
// with the default values initialized.
func NewGetVnicFcAdapterPoliciesMoidParams() *GetVnicFcAdapterPoliciesMoidParams {
	var ()
	return &GetVnicFcAdapterPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVnicFcAdapterPoliciesMoidParamsWithTimeout creates a new GetVnicFcAdapterPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVnicFcAdapterPoliciesMoidParamsWithTimeout(timeout time.Duration) *GetVnicFcAdapterPoliciesMoidParams {
	var ()
	return &GetVnicFcAdapterPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewGetVnicFcAdapterPoliciesMoidParamsWithContext creates a new GetVnicFcAdapterPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVnicFcAdapterPoliciesMoidParamsWithContext(ctx context.Context) *GetVnicFcAdapterPoliciesMoidParams {
	var ()
	return &GetVnicFcAdapterPoliciesMoidParams{

		Context: ctx,
	}
}

// NewGetVnicFcAdapterPoliciesMoidParamsWithHTTPClient creates a new GetVnicFcAdapterPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVnicFcAdapterPoliciesMoidParamsWithHTTPClient(client *http.Client) *GetVnicFcAdapterPoliciesMoidParams {
	var ()
	return &GetVnicFcAdapterPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*GetVnicFcAdapterPoliciesMoidParams contains all the parameters to send to the API endpoint
for the get vnic fc adapter policies moid operation typically these are written to a http.Request
*/
type GetVnicFcAdapterPoliciesMoidParams struct {

	/*Moid
	  The moid of the vnicFcAdapterPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vnic fc adapter policies moid params
func (o *GetVnicFcAdapterPoliciesMoidParams) WithTimeout(timeout time.Duration) *GetVnicFcAdapterPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vnic fc adapter policies moid params
func (o *GetVnicFcAdapterPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vnic fc adapter policies moid params
func (o *GetVnicFcAdapterPoliciesMoidParams) WithContext(ctx context.Context) *GetVnicFcAdapterPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vnic fc adapter policies moid params
func (o *GetVnicFcAdapterPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vnic fc adapter policies moid params
func (o *GetVnicFcAdapterPoliciesMoidParams) WithHTTPClient(client *http.Client) *GetVnicFcAdapterPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vnic fc adapter policies moid params
func (o *GetVnicFcAdapterPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get vnic fc adapter policies moid params
func (o *GetVnicFcAdapterPoliciesMoidParams) WithMoid(moid string) *GetVnicFcAdapterPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get vnic fc adapter policies moid params
func (o *GetVnicFcAdapterPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetVnicFcAdapterPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
