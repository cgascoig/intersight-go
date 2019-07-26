// Code generated by go-swagger; DO NOT EDIT.

package vnic_eth_adapter_policy

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

// NewDeleteVnicEthAdapterPoliciesMoidParams creates a new DeleteVnicEthAdapterPoliciesMoidParams object
// with the default values initialized.
func NewDeleteVnicEthAdapterPoliciesMoidParams() *DeleteVnicEthAdapterPoliciesMoidParams {
	var ()
	return &DeleteVnicEthAdapterPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVnicEthAdapterPoliciesMoidParamsWithTimeout creates a new DeleteVnicEthAdapterPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVnicEthAdapterPoliciesMoidParamsWithTimeout(timeout time.Duration) *DeleteVnicEthAdapterPoliciesMoidParams {
	var ()
	return &DeleteVnicEthAdapterPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewDeleteVnicEthAdapterPoliciesMoidParamsWithContext creates a new DeleteVnicEthAdapterPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVnicEthAdapterPoliciesMoidParamsWithContext(ctx context.Context) *DeleteVnicEthAdapterPoliciesMoidParams {
	var ()
	return &DeleteVnicEthAdapterPoliciesMoidParams{

		Context: ctx,
	}
}

// NewDeleteVnicEthAdapterPoliciesMoidParamsWithHTTPClient creates a new DeleteVnicEthAdapterPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVnicEthAdapterPoliciesMoidParamsWithHTTPClient(client *http.Client) *DeleteVnicEthAdapterPoliciesMoidParams {
	var ()
	return &DeleteVnicEthAdapterPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*DeleteVnicEthAdapterPoliciesMoidParams contains all the parameters to send to the API endpoint
for the delete vnic eth adapter policies moid operation typically these are written to a http.Request
*/
type DeleteVnicEthAdapterPoliciesMoidParams struct {

	/*Moid
	  The moid of the vnicEthAdapterPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete vnic eth adapter policies moid params
func (o *DeleteVnicEthAdapterPoliciesMoidParams) WithTimeout(timeout time.Duration) *DeleteVnicEthAdapterPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vnic eth adapter policies moid params
func (o *DeleteVnicEthAdapterPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vnic eth adapter policies moid params
func (o *DeleteVnicEthAdapterPoliciesMoidParams) WithContext(ctx context.Context) *DeleteVnicEthAdapterPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vnic eth adapter policies moid params
func (o *DeleteVnicEthAdapterPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vnic eth adapter policies moid params
func (o *DeleteVnicEthAdapterPoliciesMoidParams) WithHTTPClient(client *http.Client) *DeleteVnicEthAdapterPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vnic eth adapter policies moid params
func (o *DeleteVnicEthAdapterPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete vnic eth adapter policies moid params
func (o *DeleteVnicEthAdapterPoliciesMoidParams) WithMoid(moid string) *DeleteVnicEthAdapterPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete vnic eth adapter policies moid params
func (o *DeleteVnicEthAdapterPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVnicEthAdapterPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
