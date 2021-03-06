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

// NewGetResourceGroupsMoidParams creates a new GetResourceGroupsMoidParams object
// with the default values initialized.
func NewGetResourceGroupsMoidParams() *GetResourceGroupsMoidParams {
	var ()
	return &GetResourceGroupsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceGroupsMoidParamsWithTimeout creates a new GetResourceGroupsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResourceGroupsMoidParamsWithTimeout(timeout time.Duration) *GetResourceGroupsMoidParams {
	var ()
	return &GetResourceGroupsMoidParams{

		timeout: timeout,
	}
}

// NewGetResourceGroupsMoidParamsWithContext creates a new GetResourceGroupsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResourceGroupsMoidParamsWithContext(ctx context.Context) *GetResourceGroupsMoidParams {
	var ()
	return &GetResourceGroupsMoidParams{

		Context: ctx,
	}
}

// NewGetResourceGroupsMoidParamsWithHTTPClient creates a new GetResourceGroupsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResourceGroupsMoidParamsWithHTTPClient(client *http.Client) *GetResourceGroupsMoidParams {
	var ()
	return &GetResourceGroupsMoidParams{
		HTTPClient: client,
	}
}

/*GetResourceGroupsMoidParams contains all the parameters to send to the API endpoint
for the get resource groups moid operation typically these are written to a http.Request
*/
type GetResourceGroupsMoidParams struct {

	/*Moid
	  The moid of the resourceGroup instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get resource groups moid params
func (o *GetResourceGroupsMoidParams) WithTimeout(timeout time.Duration) *GetResourceGroupsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource groups moid params
func (o *GetResourceGroupsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource groups moid params
func (o *GetResourceGroupsMoidParams) WithContext(ctx context.Context) *GetResourceGroupsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource groups moid params
func (o *GetResourceGroupsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource groups moid params
func (o *GetResourceGroupsMoidParams) WithHTTPClient(client *http.Client) *GetResourceGroupsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource groups moid params
func (o *GetResourceGroupsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get resource groups moid params
func (o *GetResourceGroupsMoidParams) WithMoid(moid string) *GetResourceGroupsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get resource groups moid params
func (o *GetResourceGroupsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceGroupsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
