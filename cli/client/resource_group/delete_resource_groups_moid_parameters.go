// Code generated by go-swagger; DO NOT EDIT.

package resource_group

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

// NewDeleteResourceGroupsMoidParams creates a new DeleteResourceGroupsMoidParams object
// with the default values initialized.
func NewDeleteResourceGroupsMoidParams() *DeleteResourceGroupsMoidParams {
	var ()
	return &DeleteResourceGroupsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResourceGroupsMoidParamsWithTimeout creates a new DeleteResourceGroupsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteResourceGroupsMoidParamsWithTimeout(timeout time.Duration) *DeleteResourceGroupsMoidParams {
	var ()
	return &DeleteResourceGroupsMoidParams{

		timeout: timeout,
	}
}

// NewDeleteResourceGroupsMoidParamsWithContext creates a new DeleteResourceGroupsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteResourceGroupsMoidParamsWithContext(ctx context.Context) *DeleteResourceGroupsMoidParams {
	var ()
	return &DeleteResourceGroupsMoidParams{

		Context: ctx,
	}
}

// NewDeleteResourceGroupsMoidParamsWithHTTPClient creates a new DeleteResourceGroupsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteResourceGroupsMoidParamsWithHTTPClient(client *http.Client) *DeleteResourceGroupsMoidParams {
	var ()
	return &DeleteResourceGroupsMoidParams{
		HTTPClient: client,
	}
}

/*DeleteResourceGroupsMoidParams contains all the parameters to send to the API endpoint
for the delete resource groups moid operation typically these are written to a http.Request
*/
type DeleteResourceGroupsMoidParams struct {

	/*Moid
	  The moid of the resourceGroup instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete resource groups moid params
func (o *DeleteResourceGroupsMoidParams) WithTimeout(timeout time.Duration) *DeleteResourceGroupsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete resource groups moid params
func (o *DeleteResourceGroupsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete resource groups moid params
func (o *DeleteResourceGroupsMoidParams) WithContext(ctx context.Context) *DeleteResourceGroupsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete resource groups moid params
func (o *DeleteResourceGroupsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete resource groups moid params
func (o *DeleteResourceGroupsMoidParams) WithHTTPClient(client *http.Client) *DeleteResourceGroupsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete resource groups moid params
func (o *DeleteResourceGroupsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete resource groups moid params
func (o *DeleteResourceGroupsMoidParams) WithMoid(moid string) *DeleteResourceGroupsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete resource groups moid params
func (o *DeleteResourceGroupsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResourceGroupsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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