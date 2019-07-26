// Code generated by go-swagger; DO NOT EDIT.

package appliance_node_info

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

// NewGetApplianceNodeInfosMoidParams creates a new GetApplianceNodeInfosMoidParams object
// with the default values initialized.
func NewGetApplianceNodeInfosMoidParams() *GetApplianceNodeInfosMoidParams {
	var ()
	return &GetApplianceNodeInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplianceNodeInfosMoidParamsWithTimeout creates a new GetApplianceNodeInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplianceNodeInfosMoidParamsWithTimeout(timeout time.Duration) *GetApplianceNodeInfosMoidParams {
	var ()
	return &GetApplianceNodeInfosMoidParams{

		timeout: timeout,
	}
}

// NewGetApplianceNodeInfosMoidParamsWithContext creates a new GetApplianceNodeInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplianceNodeInfosMoidParamsWithContext(ctx context.Context) *GetApplianceNodeInfosMoidParams {
	var ()
	return &GetApplianceNodeInfosMoidParams{

		Context: ctx,
	}
}

// NewGetApplianceNodeInfosMoidParamsWithHTTPClient creates a new GetApplianceNodeInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplianceNodeInfosMoidParamsWithHTTPClient(client *http.Client) *GetApplianceNodeInfosMoidParams {
	var ()
	return &GetApplianceNodeInfosMoidParams{
		HTTPClient: client,
	}
}

/*GetApplianceNodeInfosMoidParams contains all the parameters to send to the API endpoint
for the get appliance node infos moid operation typically these are written to a http.Request
*/
type GetApplianceNodeInfosMoidParams struct {

	/*Moid
	  The moid of the applianceNodeInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get appliance node infos moid params
func (o *GetApplianceNodeInfosMoidParams) WithTimeout(timeout time.Duration) *GetApplianceNodeInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appliance node infos moid params
func (o *GetApplianceNodeInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appliance node infos moid params
func (o *GetApplianceNodeInfosMoidParams) WithContext(ctx context.Context) *GetApplianceNodeInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appliance node infos moid params
func (o *GetApplianceNodeInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appliance node infos moid params
func (o *GetApplianceNodeInfosMoidParams) WithHTTPClient(client *http.Client) *GetApplianceNodeInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appliance node infos moid params
func (o *GetApplianceNodeInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get appliance node infos moid params
func (o *GetApplianceNodeInfosMoidParams) WithMoid(moid string) *GetApplianceNodeInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get appliance node infos moid params
func (o *GetApplianceNodeInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplianceNodeInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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