// Code generated by go-swagger; DO NOT EDIT.

package inventory_generic_inventory_holder

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

// NewGetInventoryGenericInventoryHoldersMoidParams creates a new GetInventoryGenericInventoryHoldersMoidParams object
// with the default values initialized.
func NewGetInventoryGenericInventoryHoldersMoidParams() *GetInventoryGenericInventoryHoldersMoidParams {
	var ()
	return &GetInventoryGenericInventoryHoldersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryGenericInventoryHoldersMoidParamsWithTimeout creates a new GetInventoryGenericInventoryHoldersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInventoryGenericInventoryHoldersMoidParamsWithTimeout(timeout time.Duration) *GetInventoryGenericInventoryHoldersMoidParams {
	var ()
	return &GetInventoryGenericInventoryHoldersMoidParams{

		timeout: timeout,
	}
}

// NewGetInventoryGenericInventoryHoldersMoidParamsWithContext creates a new GetInventoryGenericInventoryHoldersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInventoryGenericInventoryHoldersMoidParamsWithContext(ctx context.Context) *GetInventoryGenericInventoryHoldersMoidParams {
	var ()
	return &GetInventoryGenericInventoryHoldersMoidParams{

		Context: ctx,
	}
}

// NewGetInventoryGenericInventoryHoldersMoidParamsWithHTTPClient creates a new GetInventoryGenericInventoryHoldersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInventoryGenericInventoryHoldersMoidParamsWithHTTPClient(client *http.Client) *GetInventoryGenericInventoryHoldersMoidParams {
	var ()
	return &GetInventoryGenericInventoryHoldersMoidParams{
		HTTPClient: client,
	}
}

/*GetInventoryGenericInventoryHoldersMoidParams contains all the parameters to send to the API endpoint
for the get inventory generic inventory holders moid operation typically these are written to a http.Request
*/
type GetInventoryGenericInventoryHoldersMoidParams struct {

	/*Moid
	  The moid of the inventoryGenericInventoryHolder instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get inventory generic inventory holders moid params
func (o *GetInventoryGenericInventoryHoldersMoidParams) WithTimeout(timeout time.Duration) *GetInventoryGenericInventoryHoldersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory generic inventory holders moid params
func (o *GetInventoryGenericInventoryHoldersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory generic inventory holders moid params
func (o *GetInventoryGenericInventoryHoldersMoidParams) WithContext(ctx context.Context) *GetInventoryGenericInventoryHoldersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory generic inventory holders moid params
func (o *GetInventoryGenericInventoryHoldersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory generic inventory holders moid params
func (o *GetInventoryGenericInventoryHoldersMoidParams) WithHTTPClient(client *http.Client) *GetInventoryGenericInventoryHoldersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory generic inventory holders moid params
func (o *GetInventoryGenericInventoryHoldersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get inventory generic inventory holders moid params
func (o *GetInventoryGenericInventoryHoldersMoidParams) WithMoid(moid string) *GetInventoryGenericInventoryHoldersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get inventory generic inventory holders moid params
func (o *GetInventoryGenericInventoryHoldersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryGenericInventoryHoldersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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