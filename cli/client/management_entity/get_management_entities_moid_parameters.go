// Code generated by go-swagger; DO NOT EDIT.

package management_entity

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

// NewGetManagementEntitiesMoidParams creates a new GetManagementEntitiesMoidParams object
// with the default values initialized.
func NewGetManagementEntitiesMoidParams() *GetManagementEntitiesMoidParams {
	var ()
	return &GetManagementEntitiesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetManagementEntitiesMoidParamsWithTimeout creates a new GetManagementEntitiesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetManagementEntitiesMoidParamsWithTimeout(timeout time.Duration) *GetManagementEntitiesMoidParams {
	var ()
	return &GetManagementEntitiesMoidParams{

		timeout: timeout,
	}
}

// NewGetManagementEntitiesMoidParamsWithContext creates a new GetManagementEntitiesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetManagementEntitiesMoidParamsWithContext(ctx context.Context) *GetManagementEntitiesMoidParams {
	var ()
	return &GetManagementEntitiesMoidParams{

		Context: ctx,
	}
}

// NewGetManagementEntitiesMoidParamsWithHTTPClient creates a new GetManagementEntitiesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetManagementEntitiesMoidParamsWithHTTPClient(client *http.Client) *GetManagementEntitiesMoidParams {
	var ()
	return &GetManagementEntitiesMoidParams{
		HTTPClient: client,
	}
}

/*GetManagementEntitiesMoidParams contains all the parameters to send to the API endpoint
for the get management entities moid operation typically these are written to a http.Request
*/
type GetManagementEntitiesMoidParams struct {

	/*Moid
	  The moid of the managementEntity instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get management entities moid params
func (o *GetManagementEntitiesMoidParams) WithTimeout(timeout time.Duration) *GetManagementEntitiesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get management entities moid params
func (o *GetManagementEntitiesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get management entities moid params
func (o *GetManagementEntitiesMoidParams) WithContext(ctx context.Context) *GetManagementEntitiesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get management entities moid params
func (o *GetManagementEntitiesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get management entities moid params
func (o *GetManagementEntitiesMoidParams) WithHTTPClient(client *http.Client) *GetManagementEntitiesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get management entities moid params
func (o *GetManagementEntitiesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get management entities moid params
func (o *GetManagementEntitiesMoidParams) WithMoid(moid string) *GetManagementEntitiesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get management entities moid params
func (o *GetManagementEntitiesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetManagementEntitiesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
