// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_flash_controller

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

// NewGetStorageFlexFlashControllersMoidParams creates a new GetStorageFlexFlashControllersMoidParams object
// with the default values initialized.
func NewGetStorageFlexFlashControllersMoidParams() *GetStorageFlexFlashControllersMoidParams {
	var ()
	return &GetStorageFlexFlashControllersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStorageFlexFlashControllersMoidParamsWithTimeout creates a new GetStorageFlexFlashControllersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStorageFlexFlashControllersMoidParamsWithTimeout(timeout time.Duration) *GetStorageFlexFlashControllersMoidParams {
	var ()
	return &GetStorageFlexFlashControllersMoidParams{

		timeout: timeout,
	}
}

// NewGetStorageFlexFlashControllersMoidParamsWithContext creates a new GetStorageFlexFlashControllersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStorageFlexFlashControllersMoidParamsWithContext(ctx context.Context) *GetStorageFlexFlashControllersMoidParams {
	var ()
	return &GetStorageFlexFlashControllersMoidParams{

		Context: ctx,
	}
}

// NewGetStorageFlexFlashControllersMoidParamsWithHTTPClient creates a new GetStorageFlexFlashControllersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStorageFlexFlashControllersMoidParamsWithHTTPClient(client *http.Client) *GetStorageFlexFlashControllersMoidParams {
	var ()
	return &GetStorageFlexFlashControllersMoidParams{
		HTTPClient: client,
	}
}

/*GetStorageFlexFlashControllersMoidParams contains all the parameters to send to the API endpoint
for the get storage flex flash controllers moid operation typically these are written to a http.Request
*/
type GetStorageFlexFlashControllersMoidParams struct {

	/*Moid
	  The moid of the storageFlexFlashController instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get storage flex flash controllers moid params
func (o *GetStorageFlexFlashControllersMoidParams) WithTimeout(timeout time.Duration) *GetStorageFlexFlashControllersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage flex flash controllers moid params
func (o *GetStorageFlexFlashControllersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage flex flash controllers moid params
func (o *GetStorageFlexFlashControllersMoidParams) WithContext(ctx context.Context) *GetStorageFlexFlashControllersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage flex flash controllers moid params
func (o *GetStorageFlexFlashControllersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage flex flash controllers moid params
func (o *GetStorageFlexFlashControllersMoidParams) WithHTTPClient(client *http.Client) *GetStorageFlexFlashControllersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage flex flash controllers moid params
func (o *GetStorageFlexFlashControllersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get storage flex flash controllers moid params
func (o *GetStorageFlexFlashControllersMoidParams) WithMoid(moid string) *GetStorageFlexFlashControllersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get storage flex flash controllers moid params
func (o *GetStorageFlexFlashControllersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetStorageFlexFlashControllersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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