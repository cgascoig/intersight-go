// Code generated by go-swagger; DO NOT EDIT.

package iaas_connector_pack

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

// NewGetIaasConnectorPacksMoidParams creates a new GetIaasConnectorPacksMoidParams object
// with the default values initialized.
func NewGetIaasConnectorPacksMoidParams() *GetIaasConnectorPacksMoidParams {
	var ()
	return &GetIaasConnectorPacksMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIaasConnectorPacksMoidParamsWithTimeout creates a new GetIaasConnectorPacksMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIaasConnectorPacksMoidParamsWithTimeout(timeout time.Duration) *GetIaasConnectorPacksMoidParams {
	var ()
	return &GetIaasConnectorPacksMoidParams{

		timeout: timeout,
	}
}

// NewGetIaasConnectorPacksMoidParamsWithContext creates a new GetIaasConnectorPacksMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIaasConnectorPacksMoidParamsWithContext(ctx context.Context) *GetIaasConnectorPacksMoidParams {
	var ()
	return &GetIaasConnectorPacksMoidParams{

		Context: ctx,
	}
}

// NewGetIaasConnectorPacksMoidParamsWithHTTPClient creates a new GetIaasConnectorPacksMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIaasConnectorPacksMoidParamsWithHTTPClient(client *http.Client) *GetIaasConnectorPacksMoidParams {
	var ()
	return &GetIaasConnectorPacksMoidParams{
		HTTPClient: client,
	}
}

/*GetIaasConnectorPacksMoidParams contains all the parameters to send to the API endpoint
for the get iaas connector packs moid operation typically these are written to a http.Request
*/
type GetIaasConnectorPacksMoidParams struct {

	/*Moid
	  The moid of the iaasConnectorPack instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iaas connector packs moid params
func (o *GetIaasConnectorPacksMoidParams) WithTimeout(timeout time.Duration) *GetIaasConnectorPacksMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iaas connector packs moid params
func (o *GetIaasConnectorPacksMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iaas connector packs moid params
func (o *GetIaasConnectorPacksMoidParams) WithContext(ctx context.Context) *GetIaasConnectorPacksMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iaas connector packs moid params
func (o *GetIaasConnectorPacksMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iaas connector packs moid params
func (o *GetIaasConnectorPacksMoidParams) WithHTTPClient(client *http.Client) *GetIaasConnectorPacksMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iaas connector packs moid params
func (o *GetIaasConnectorPacksMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get iaas connector packs moid params
func (o *GetIaasConnectorPacksMoidParams) WithMoid(moid string) *GetIaasConnectorPacksMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get iaas connector packs moid params
func (o *GetIaasConnectorPacksMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetIaasConnectorPacksMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
