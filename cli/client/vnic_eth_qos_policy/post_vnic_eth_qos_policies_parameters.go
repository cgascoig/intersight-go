// Code generated by go-swagger; DO NOT EDIT.

package vnic_eth_qos_policy

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

// NewPostVnicEthQosPoliciesParams creates a new PostVnicEthQosPoliciesParams object
// with the default values initialized.
func NewPostVnicEthQosPoliciesParams() *PostVnicEthQosPoliciesParams {
	var ()
	return &PostVnicEthQosPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVnicEthQosPoliciesParamsWithTimeout creates a new PostVnicEthQosPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVnicEthQosPoliciesParamsWithTimeout(timeout time.Duration) *PostVnicEthQosPoliciesParams {
	var ()
	return &PostVnicEthQosPoliciesParams{

		timeout: timeout,
	}
}

// NewPostVnicEthQosPoliciesParamsWithContext creates a new PostVnicEthQosPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVnicEthQosPoliciesParamsWithContext(ctx context.Context) *PostVnicEthQosPoliciesParams {
	var ()
	return &PostVnicEthQosPoliciesParams{

		Context: ctx,
	}
}

// NewPostVnicEthQosPoliciesParamsWithHTTPClient creates a new PostVnicEthQosPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVnicEthQosPoliciesParamsWithHTTPClient(client *http.Client) *PostVnicEthQosPoliciesParams {
	var ()
	return &PostVnicEthQosPoliciesParams{
		HTTPClient: client,
	}
}

/*PostVnicEthQosPoliciesParams contains all the parameters to send to the API endpoint
for the post vnic eth qos policies operation typically these are written to a http.Request
*/
type PostVnicEthQosPoliciesParams struct {

	/*Body
	  vnicEthQosPolicy to add

	*/
	Body *models.VnicEthQosPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vnic eth qos policies params
func (o *PostVnicEthQosPoliciesParams) WithTimeout(timeout time.Duration) *PostVnicEthQosPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vnic eth qos policies params
func (o *PostVnicEthQosPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vnic eth qos policies params
func (o *PostVnicEthQosPoliciesParams) WithContext(ctx context.Context) *PostVnicEthQosPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vnic eth qos policies params
func (o *PostVnicEthQosPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vnic eth qos policies params
func (o *PostVnicEthQosPoliciesParams) WithHTTPClient(client *http.Client) *PostVnicEthQosPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vnic eth qos policies params
func (o *PostVnicEthQosPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vnic eth qos policies params
func (o *PostVnicEthQosPoliciesParams) WithBody(body *models.VnicEthQosPolicy) *PostVnicEthQosPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vnic eth qos policies params
func (o *PostVnicEthQosPoliciesParams) SetBody(body *models.VnicEthQosPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVnicEthQosPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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