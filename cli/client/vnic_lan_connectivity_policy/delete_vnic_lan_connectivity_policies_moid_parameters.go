// Code generated by go-swagger; DO NOT EDIT.

package vnic_lan_connectivity_policy

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

// NewDeleteVnicLanConnectivityPoliciesMoidParams creates a new DeleteVnicLanConnectivityPoliciesMoidParams object
// with the default values initialized.
func NewDeleteVnicLanConnectivityPoliciesMoidParams() *DeleteVnicLanConnectivityPoliciesMoidParams {
	var ()
	return &DeleteVnicLanConnectivityPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVnicLanConnectivityPoliciesMoidParamsWithTimeout creates a new DeleteVnicLanConnectivityPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVnicLanConnectivityPoliciesMoidParamsWithTimeout(timeout time.Duration) *DeleteVnicLanConnectivityPoliciesMoidParams {
	var ()
	return &DeleteVnicLanConnectivityPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewDeleteVnicLanConnectivityPoliciesMoidParamsWithContext creates a new DeleteVnicLanConnectivityPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVnicLanConnectivityPoliciesMoidParamsWithContext(ctx context.Context) *DeleteVnicLanConnectivityPoliciesMoidParams {
	var ()
	return &DeleteVnicLanConnectivityPoliciesMoidParams{

		Context: ctx,
	}
}

// NewDeleteVnicLanConnectivityPoliciesMoidParamsWithHTTPClient creates a new DeleteVnicLanConnectivityPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVnicLanConnectivityPoliciesMoidParamsWithHTTPClient(client *http.Client) *DeleteVnicLanConnectivityPoliciesMoidParams {
	var ()
	return &DeleteVnicLanConnectivityPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*DeleteVnicLanConnectivityPoliciesMoidParams contains all the parameters to send to the API endpoint
for the delete vnic lan connectivity policies moid operation typically these are written to a http.Request
*/
type DeleteVnicLanConnectivityPoliciesMoidParams struct {

	/*Moid
	  The moid of the vnicLanConnectivityPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete vnic lan connectivity policies moid params
func (o *DeleteVnicLanConnectivityPoliciesMoidParams) WithTimeout(timeout time.Duration) *DeleteVnicLanConnectivityPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vnic lan connectivity policies moid params
func (o *DeleteVnicLanConnectivityPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vnic lan connectivity policies moid params
func (o *DeleteVnicLanConnectivityPoliciesMoidParams) WithContext(ctx context.Context) *DeleteVnicLanConnectivityPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vnic lan connectivity policies moid params
func (o *DeleteVnicLanConnectivityPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vnic lan connectivity policies moid params
func (o *DeleteVnicLanConnectivityPoliciesMoidParams) WithHTTPClient(client *http.Client) *DeleteVnicLanConnectivityPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vnic lan connectivity policies moid params
func (o *DeleteVnicLanConnectivityPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete vnic lan connectivity policies moid params
func (o *DeleteVnicLanConnectivityPoliciesMoidParams) WithMoid(moid string) *DeleteVnicLanConnectivityPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete vnic lan connectivity policies moid params
func (o *DeleteVnicLanConnectivityPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVnicLanConnectivityPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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