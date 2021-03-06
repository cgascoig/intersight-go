// Code generated by go-swagger; DO NOT EDIT.

package deviceconnector_policy

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

// NewPostDeviceconnectorPoliciesParams creates a new PostDeviceconnectorPoliciesParams object
// with the default values initialized.
func NewPostDeviceconnectorPoliciesParams() *PostDeviceconnectorPoliciesParams {
	var ()
	return &PostDeviceconnectorPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDeviceconnectorPoliciesParamsWithTimeout creates a new PostDeviceconnectorPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDeviceconnectorPoliciesParamsWithTimeout(timeout time.Duration) *PostDeviceconnectorPoliciesParams {
	var ()
	return &PostDeviceconnectorPoliciesParams{

		timeout: timeout,
	}
}

// NewPostDeviceconnectorPoliciesParamsWithContext creates a new PostDeviceconnectorPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDeviceconnectorPoliciesParamsWithContext(ctx context.Context) *PostDeviceconnectorPoliciesParams {
	var ()
	return &PostDeviceconnectorPoliciesParams{

		Context: ctx,
	}
}

// NewPostDeviceconnectorPoliciesParamsWithHTTPClient creates a new PostDeviceconnectorPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDeviceconnectorPoliciesParamsWithHTTPClient(client *http.Client) *PostDeviceconnectorPoliciesParams {
	var ()
	return &PostDeviceconnectorPoliciesParams{
		HTTPClient: client,
	}
}

/*PostDeviceconnectorPoliciesParams contains all the parameters to send to the API endpoint
for the post deviceconnector policies operation typically these are written to a http.Request
*/
type PostDeviceconnectorPoliciesParams struct {

	/*Body
	  deviceconnectorPolicy to add

	*/
	Body *models.DeviceconnectorPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post deviceconnector policies params
func (o *PostDeviceconnectorPoliciesParams) WithTimeout(timeout time.Duration) *PostDeviceconnectorPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post deviceconnector policies params
func (o *PostDeviceconnectorPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post deviceconnector policies params
func (o *PostDeviceconnectorPoliciesParams) WithContext(ctx context.Context) *PostDeviceconnectorPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post deviceconnector policies params
func (o *PostDeviceconnectorPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post deviceconnector policies params
func (o *PostDeviceconnectorPoliciesParams) WithHTTPClient(client *http.Client) *PostDeviceconnectorPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post deviceconnector policies params
func (o *PostDeviceconnectorPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post deviceconnector policies params
func (o *PostDeviceconnectorPoliciesParams) WithBody(body *models.DeviceconnectorPolicy) *PostDeviceconnectorPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post deviceconnector policies params
func (o *PostDeviceconnectorPoliciesParams) SetBody(body *models.DeviceconnectorPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDeviceconnectorPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
