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

// NewPostVnicEthNetworkPoliciesParams creates a new PostVnicEthNetworkPoliciesParams object
// with the default values initialized.
func NewPostVnicEthNetworkPoliciesParams() *PostVnicEthNetworkPoliciesParams {
	var ()
	return &PostVnicEthNetworkPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVnicEthNetworkPoliciesParamsWithTimeout creates a new PostVnicEthNetworkPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVnicEthNetworkPoliciesParamsWithTimeout(timeout time.Duration) *PostVnicEthNetworkPoliciesParams {
	var ()
	return &PostVnicEthNetworkPoliciesParams{

		timeout: timeout,
	}
}

// NewPostVnicEthNetworkPoliciesParamsWithContext creates a new PostVnicEthNetworkPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVnicEthNetworkPoliciesParamsWithContext(ctx context.Context) *PostVnicEthNetworkPoliciesParams {
	var ()
	return &PostVnicEthNetworkPoliciesParams{

		Context: ctx,
	}
}

// NewPostVnicEthNetworkPoliciesParamsWithHTTPClient creates a new PostVnicEthNetworkPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVnicEthNetworkPoliciesParamsWithHTTPClient(client *http.Client) *PostVnicEthNetworkPoliciesParams {
	var ()
	return &PostVnicEthNetworkPoliciesParams{
		HTTPClient: client,
	}
}

/*PostVnicEthNetworkPoliciesParams contains all the parameters to send to the API endpoint
for the post vnic eth network policies operation typically these are written to a http.Request
*/
type PostVnicEthNetworkPoliciesParams struct {

	/*Body
	  vnicEthNetworkPolicy to add

	*/
	Body *models.VnicEthNetworkPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vnic eth network policies params
func (o *PostVnicEthNetworkPoliciesParams) WithTimeout(timeout time.Duration) *PostVnicEthNetworkPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vnic eth network policies params
func (o *PostVnicEthNetworkPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vnic eth network policies params
func (o *PostVnicEthNetworkPoliciesParams) WithContext(ctx context.Context) *PostVnicEthNetworkPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vnic eth network policies params
func (o *PostVnicEthNetworkPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vnic eth network policies params
func (o *PostVnicEthNetworkPoliciesParams) WithHTTPClient(client *http.Client) *PostVnicEthNetworkPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vnic eth network policies params
func (o *PostVnicEthNetworkPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vnic eth network policies params
func (o *PostVnicEthNetworkPoliciesParams) WithBody(body *models.VnicEthNetworkPolicy) *PostVnicEthNetworkPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vnic eth network policies params
func (o *PostVnicEthNetworkPoliciesParams) SetBody(body *models.VnicEthNetworkPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVnicEthNetworkPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
