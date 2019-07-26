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

// NewPostDeviceconnectorPoliciesMoidParams creates a new PostDeviceconnectorPoliciesMoidParams object
// with the default values initialized.
func NewPostDeviceconnectorPoliciesMoidParams() *PostDeviceconnectorPoliciesMoidParams {
	var ()
	return &PostDeviceconnectorPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDeviceconnectorPoliciesMoidParamsWithTimeout creates a new PostDeviceconnectorPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDeviceconnectorPoliciesMoidParamsWithTimeout(timeout time.Duration) *PostDeviceconnectorPoliciesMoidParams {
	var ()
	return &PostDeviceconnectorPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostDeviceconnectorPoliciesMoidParamsWithContext creates a new PostDeviceconnectorPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDeviceconnectorPoliciesMoidParamsWithContext(ctx context.Context) *PostDeviceconnectorPoliciesMoidParams {
	var ()
	return &PostDeviceconnectorPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostDeviceconnectorPoliciesMoidParamsWithHTTPClient creates a new PostDeviceconnectorPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDeviceconnectorPoliciesMoidParamsWithHTTPClient(client *http.Client) *PostDeviceconnectorPoliciesMoidParams {
	var ()
	return &PostDeviceconnectorPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostDeviceconnectorPoliciesMoidParams contains all the parameters to send to the API endpoint
for the post deviceconnector policies moid operation typically these are written to a http.Request
*/
type PostDeviceconnectorPoliciesMoidParams struct {

	/*Body
	  deviceconnectorPolicy to update

	*/
	Body *models.DeviceconnectorPolicy
	/*Moid
	  The moid of the deviceconnectorPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) WithTimeout(timeout time.Duration) *PostDeviceconnectorPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) WithContext(ctx context.Context) *PostDeviceconnectorPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) WithHTTPClient(client *http.Client) *PostDeviceconnectorPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) WithBody(body *models.DeviceconnectorPolicy) *PostDeviceconnectorPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) SetBody(body *models.DeviceconnectorPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) WithMoid(moid string) *PostDeviceconnectorPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post deviceconnector policies moid params
func (o *PostDeviceconnectorPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostDeviceconnectorPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
