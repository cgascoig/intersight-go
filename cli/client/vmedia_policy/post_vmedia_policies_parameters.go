// Code generated by go-swagger; DO NOT EDIT.

package vmedia_policy

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

// NewPostVmediaPoliciesParams creates a new PostVmediaPoliciesParams object
// with the default values initialized.
func NewPostVmediaPoliciesParams() *PostVmediaPoliciesParams {
	var ()
	return &PostVmediaPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVmediaPoliciesParamsWithTimeout creates a new PostVmediaPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVmediaPoliciesParamsWithTimeout(timeout time.Duration) *PostVmediaPoliciesParams {
	var ()
	return &PostVmediaPoliciesParams{

		timeout: timeout,
	}
}

// NewPostVmediaPoliciesParamsWithContext creates a new PostVmediaPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVmediaPoliciesParamsWithContext(ctx context.Context) *PostVmediaPoliciesParams {
	var ()
	return &PostVmediaPoliciesParams{

		Context: ctx,
	}
}

// NewPostVmediaPoliciesParamsWithHTTPClient creates a new PostVmediaPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVmediaPoliciesParamsWithHTTPClient(client *http.Client) *PostVmediaPoliciesParams {
	var ()
	return &PostVmediaPoliciesParams{
		HTTPClient: client,
	}
}

/*PostVmediaPoliciesParams contains all the parameters to send to the API endpoint
for the post vmedia policies operation typically these are written to a http.Request
*/
type PostVmediaPoliciesParams struct {

	/*Body
	  vmediaPolicy to add

	*/
	Body *models.VmediaPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vmedia policies params
func (o *PostVmediaPoliciesParams) WithTimeout(timeout time.Duration) *PostVmediaPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vmedia policies params
func (o *PostVmediaPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vmedia policies params
func (o *PostVmediaPoliciesParams) WithContext(ctx context.Context) *PostVmediaPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vmedia policies params
func (o *PostVmediaPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vmedia policies params
func (o *PostVmediaPoliciesParams) WithHTTPClient(client *http.Client) *PostVmediaPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vmedia policies params
func (o *PostVmediaPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vmedia policies params
func (o *PostVmediaPoliciesParams) WithBody(body *models.VmediaPolicy) *PostVmediaPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vmedia policies params
func (o *PostVmediaPoliciesParams) SetBody(body *models.VmediaPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVmediaPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
