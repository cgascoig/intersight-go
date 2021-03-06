// Code generated by go-swagger; DO NOT EDIT.

package port_sub_group

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

// NewGetPortSubGroupsMoidParams creates a new GetPortSubGroupsMoidParams object
// with the default values initialized.
func NewGetPortSubGroupsMoidParams() *GetPortSubGroupsMoidParams {
	var ()
	return &GetPortSubGroupsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPortSubGroupsMoidParamsWithTimeout creates a new GetPortSubGroupsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPortSubGroupsMoidParamsWithTimeout(timeout time.Duration) *GetPortSubGroupsMoidParams {
	var ()
	return &GetPortSubGroupsMoidParams{

		timeout: timeout,
	}
}

// NewGetPortSubGroupsMoidParamsWithContext creates a new GetPortSubGroupsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPortSubGroupsMoidParamsWithContext(ctx context.Context) *GetPortSubGroupsMoidParams {
	var ()
	return &GetPortSubGroupsMoidParams{

		Context: ctx,
	}
}

// NewGetPortSubGroupsMoidParamsWithHTTPClient creates a new GetPortSubGroupsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPortSubGroupsMoidParamsWithHTTPClient(client *http.Client) *GetPortSubGroupsMoidParams {
	var ()
	return &GetPortSubGroupsMoidParams{
		HTTPClient: client,
	}
}

/*GetPortSubGroupsMoidParams contains all the parameters to send to the API endpoint
for the get port sub groups moid operation typically these are written to a http.Request
*/
type GetPortSubGroupsMoidParams struct {

	/*Moid
	  The moid of the portSubGroup instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get port sub groups moid params
func (o *GetPortSubGroupsMoidParams) WithTimeout(timeout time.Duration) *GetPortSubGroupsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get port sub groups moid params
func (o *GetPortSubGroupsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get port sub groups moid params
func (o *GetPortSubGroupsMoidParams) WithContext(ctx context.Context) *GetPortSubGroupsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get port sub groups moid params
func (o *GetPortSubGroupsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get port sub groups moid params
func (o *GetPortSubGroupsMoidParams) WithHTTPClient(client *http.Client) *GetPortSubGroupsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get port sub groups moid params
func (o *GetPortSubGroupsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get port sub groups moid params
func (o *GetPortSubGroupsMoidParams) WithMoid(moid string) *GetPortSubGroupsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get port sub groups moid params
func (o *GetPortSubGroupsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPortSubGroupsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
