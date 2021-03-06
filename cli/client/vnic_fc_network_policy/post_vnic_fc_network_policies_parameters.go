// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_network_policy

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

// NewPostVnicFcNetworkPoliciesParams creates a new PostVnicFcNetworkPoliciesParams object
// with the default values initialized.
func NewPostVnicFcNetworkPoliciesParams() *PostVnicFcNetworkPoliciesParams {
	var ()
	return &PostVnicFcNetworkPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVnicFcNetworkPoliciesParamsWithTimeout creates a new PostVnicFcNetworkPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVnicFcNetworkPoliciesParamsWithTimeout(timeout time.Duration) *PostVnicFcNetworkPoliciesParams {
	var ()
	return &PostVnicFcNetworkPoliciesParams{

		timeout: timeout,
	}
}

// NewPostVnicFcNetworkPoliciesParamsWithContext creates a new PostVnicFcNetworkPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVnicFcNetworkPoliciesParamsWithContext(ctx context.Context) *PostVnicFcNetworkPoliciesParams {
	var ()
	return &PostVnicFcNetworkPoliciesParams{

		Context: ctx,
	}
}

// NewPostVnicFcNetworkPoliciesParamsWithHTTPClient creates a new PostVnicFcNetworkPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVnicFcNetworkPoliciesParamsWithHTTPClient(client *http.Client) *PostVnicFcNetworkPoliciesParams {
	var ()
	return &PostVnicFcNetworkPoliciesParams{
		HTTPClient: client,
	}
}

/*PostVnicFcNetworkPoliciesParams contains all the parameters to send to the API endpoint
for the post vnic fc network policies operation typically these are written to a http.Request
*/
type PostVnicFcNetworkPoliciesParams struct {

	/*Body
	  vnicFcNetworkPolicy to add

	*/
	Body *models.VnicFcNetworkPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vnic fc network policies params
func (o *PostVnicFcNetworkPoliciesParams) WithTimeout(timeout time.Duration) *PostVnicFcNetworkPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vnic fc network policies params
func (o *PostVnicFcNetworkPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vnic fc network policies params
func (o *PostVnicFcNetworkPoliciesParams) WithContext(ctx context.Context) *PostVnicFcNetworkPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vnic fc network policies params
func (o *PostVnicFcNetworkPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vnic fc network policies params
func (o *PostVnicFcNetworkPoliciesParams) WithHTTPClient(client *http.Client) *PostVnicFcNetworkPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vnic fc network policies params
func (o *PostVnicFcNetworkPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vnic fc network policies params
func (o *PostVnicFcNetworkPoliciesParams) WithBody(body *models.VnicFcNetworkPolicy) *PostVnicFcNetworkPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vnic fc network policies params
func (o *PostVnicFcNetworkPoliciesParams) SetBody(body *models.VnicFcNetworkPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVnicFcNetworkPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
