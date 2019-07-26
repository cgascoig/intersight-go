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

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPostVnicFcAdapterPoliciesParams creates a new PostVnicFcAdapterPoliciesParams object
// with the default values initialized.
func NewPostVnicFcAdapterPoliciesParams() *PostVnicFcAdapterPoliciesParams {
	var ()
	return &PostVnicFcAdapterPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVnicFcAdapterPoliciesParamsWithTimeout creates a new PostVnicFcAdapterPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVnicFcAdapterPoliciesParamsWithTimeout(timeout time.Duration) *PostVnicFcAdapterPoliciesParams {
	var ()
	return &PostVnicFcAdapterPoliciesParams{

		timeout: timeout,
	}
}

// NewPostVnicFcAdapterPoliciesParamsWithContext creates a new PostVnicFcAdapterPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVnicFcAdapterPoliciesParamsWithContext(ctx context.Context) *PostVnicFcAdapterPoliciesParams {
	var ()
	return &PostVnicFcAdapterPoliciesParams{

		Context: ctx,
	}
}

// NewPostVnicFcAdapterPoliciesParamsWithHTTPClient creates a new PostVnicFcAdapterPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVnicFcAdapterPoliciesParamsWithHTTPClient(client *http.Client) *PostVnicFcAdapterPoliciesParams {
	var ()
	return &PostVnicFcAdapterPoliciesParams{
		HTTPClient: client,
	}
}

/*PostVnicFcAdapterPoliciesParams contains all the parameters to send to the API endpoint
for the post vnic fc adapter policies operation typically these are written to a http.Request
*/
type PostVnicFcAdapterPoliciesParams struct {

	/*Body
	  vnicFcAdapterPolicy to add

	*/
	Body *models.VnicFcAdapterPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vnic fc adapter policies params
func (o *PostVnicFcAdapterPoliciesParams) WithTimeout(timeout time.Duration) *PostVnicFcAdapterPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vnic fc adapter policies params
func (o *PostVnicFcAdapterPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vnic fc adapter policies params
func (o *PostVnicFcAdapterPoliciesParams) WithContext(ctx context.Context) *PostVnicFcAdapterPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vnic fc adapter policies params
func (o *PostVnicFcAdapterPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vnic fc adapter policies params
func (o *PostVnicFcAdapterPoliciesParams) WithHTTPClient(client *http.Client) *PostVnicFcAdapterPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vnic fc adapter policies params
func (o *PostVnicFcAdapterPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vnic fc adapter policies params
func (o *PostVnicFcAdapterPoliciesParams) WithBody(body *models.VnicFcAdapterPolicy) *PostVnicFcAdapterPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vnic fc adapter policies params
func (o *PostVnicFcAdapterPoliciesParams) SetBody(body *models.VnicFcAdapterPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVnicFcAdapterPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
