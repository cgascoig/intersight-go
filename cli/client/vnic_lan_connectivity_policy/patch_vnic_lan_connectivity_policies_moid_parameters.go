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

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPatchVnicLanConnectivityPoliciesMoidParams creates a new PatchVnicLanConnectivityPoliciesMoidParams object
// with the default values initialized.
func NewPatchVnicLanConnectivityPoliciesMoidParams() *PatchVnicLanConnectivityPoliciesMoidParams {
	var ()
	return &PatchVnicLanConnectivityPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchVnicLanConnectivityPoliciesMoidParamsWithTimeout creates a new PatchVnicLanConnectivityPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchVnicLanConnectivityPoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchVnicLanConnectivityPoliciesMoidParams {
	var ()
	return &PatchVnicLanConnectivityPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchVnicLanConnectivityPoliciesMoidParamsWithContext creates a new PatchVnicLanConnectivityPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchVnicLanConnectivityPoliciesMoidParamsWithContext(ctx context.Context) *PatchVnicLanConnectivityPoliciesMoidParams {
	var ()
	return &PatchVnicLanConnectivityPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchVnicLanConnectivityPoliciesMoidParamsWithHTTPClient creates a new PatchVnicLanConnectivityPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchVnicLanConnectivityPoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchVnicLanConnectivityPoliciesMoidParams {
	var ()
	return &PatchVnicLanConnectivityPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchVnicLanConnectivityPoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch vnic lan connectivity policies moid operation typically these are written to a http.Request
*/
type PatchVnicLanConnectivityPoliciesMoidParams struct {

	/*Body
	  vnicLanConnectivityPolicy to update

	*/
	Body *models.VnicLanConnectivityPolicy
	/*Moid
	  The moid of the vnicLanConnectivityPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchVnicLanConnectivityPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) WithContext(ctx context.Context) *PatchVnicLanConnectivityPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchVnicLanConnectivityPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) WithBody(body *models.VnicLanConnectivityPolicy) *PatchVnicLanConnectivityPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) SetBody(body *models.VnicLanConnectivityPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) WithMoid(moid string) *PatchVnicLanConnectivityPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch vnic lan connectivity policies moid params
func (o *PatchVnicLanConnectivityPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchVnicLanConnectivityPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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