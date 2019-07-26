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
)

// NewDeleteDeviceconnectorPoliciesMoidParams creates a new DeleteDeviceconnectorPoliciesMoidParams object
// with the default values initialized.
func NewDeleteDeviceconnectorPoliciesMoidParams() *DeleteDeviceconnectorPoliciesMoidParams {
	var ()
	return &DeleteDeviceconnectorPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceconnectorPoliciesMoidParamsWithTimeout creates a new DeleteDeviceconnectorPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeviceconnectorPoliciesMoidParamsWithTimeout(timeout time.Duration) *DeleteDeviceconnectorPoliciesMoidParams {
	var ()
	return &DeleteDeviceconnectorPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewDeleteDeviceconnectorPoliciesMoidParamsWithContext creates a new DeleteDeviceconnectorPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeviceconnectorPoliciesMoidParamsWithContext(ctx context.Context) *DeleteDeviceconnectorPoliciesMoidParams {
	var ()
	return &DeleteDeviceconnectorPoliciesMoidParams{

		Context: ctx,
	}
}

// NewDeleteDeviceconnectorPoliciesMoidParamsWithHTTPClient creates a new DeleteDeviceconnectorPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeviceconnectorPoliciesMoidParamsWithHTTPClient(client *http.Client) *DeleteDeviceconnectorPoliciesMoidParams {
	var ()
	return &DeleteDeviceconnectorPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*DeleteDeviceconnectorPoliciesMoidParams contains all the parameters to send to the API endpoint
for the delete deviceconnector policies moid operation typically these are written to a http.Request
*/
type DeleteDeviceconnectorPoliciesMoidParams struct {

	/*Moid
	  The moid of the deviceconnectorPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete deviceconnector policies moid params
func (o *DeleteDeviceconnectorPoliciesMoidParams) WithTimeout(timeout time.Duration) *DeleteDeviceconnectorPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deviceconnector policies moid params
func (o *DeleteDeviceconnectorPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deviceconnector policies moid params
func (o *DeleteDeviceconnectorPoliciesMoidParams) WithContext(ctx context.Context) *DeleteDeviceconnectorPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deviceconnector policies moid params
func (o *DeleteDeviceconnectorPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deviceconnector policies moid params
func (o *DeleteDeviceconnectorPoliciesMoidParams) WithHTTPClient(client *http.Client) *DeleteDeviceconnectorPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deviceconnector policies moid params
func (o *DeleteDeviceconnectorPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete deviceconnector policies moid params
func (o *DeleteDeviceconnectorPoliciesMoidParams) WithMoid(moid string) *DeleteDeviceconnectorPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete deviceconnector policies moid params
func (o *DeleteDeviceconnectorPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceconnectorPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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