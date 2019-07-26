// Code generated by go-swagger; DO NOT EDIT.

package vnic_eth_network_policy

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

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPatchVnicEthNetworkPoliciesMoidParams creates a new PatchVnicEthNetworkPoliciesMoidParams object
// with the default values initialized.
func NewPatchVnicEthNetworkPoliciesMoidParams() *PatchVnicEthNetworkPoliciesMoidParams {
	var ()
	return &PatchVnicEthNetworkPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchVnicEthNetworkPoliciesMoidParamsWithTimeout creates a new PatchVnicEthNetworkPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchVnicEthNetworkPoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchVnicEthNetworkPoliciesMoidParams {
	var ()
	return &PatchVnicEthNetworkPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchVnicEthNetworkPoliciesMoidParamsWithContext creates a new PatchVnicEthNetworkPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchVnicEthNetworkPoliciesMoidParamsWithContext(ctx context.Context) *PatchVnicEthNetworkPoliciesMoidParams {
	var ()
	return &PatchVnicEthNetworkPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchVnicEthNetworkPoliciesMoidParamsWithHTTPClient creates a new PatchVnicEthNetworkPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchVnicEthNetworkPoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchVnicEthNetworkPoliciesMoidParams {
	var ()
	return &PatchVnicEthNetworkPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchVnicEthNetworkPoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch vnic eth network policies moid operation typically these are written to a http.Request
*/
type PatchVnicEthNetworkPoliciesMoidParams struct {

	/*Body
	  vnicEthNetworkPolicy to update

	*/
	Body *models.VnicEthNetworkPolicy
	/*Moid
	  The moid of the vnicEthNetworkPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchVnicEthNetworkPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) WithContext(ctx context.Context) *PatchVnicEthNetworkPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchVnicEthNetworkPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) WithBody(body *models.VnicEthNetworkPolicy) *PatchVnicEthNetworkPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) SetBody(body *models.VnicEthNetworkPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) WithMoid(moid string) *PatchVnicEthNetworkPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch vnic eth network policies moid params
func (o *PatchVnicEthNetworkPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchVnicEthNetworkPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param moid
	if err := r.SetPathParam("moid", o.Moid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}