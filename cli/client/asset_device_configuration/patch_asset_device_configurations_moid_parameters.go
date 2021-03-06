// Code generated by go-swagger; DO NOT EDIT.

package asset_device_configuration

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

// NewPatchAssetDeviceConfigurationsMoidParams creates a new PatchAssetDeviceConfigurationsMoidParams object
// with the default values initialized.
func NewPatchAssetDeviceConfigurationsMoidParams() *PatchAssetDeviceConfigurationsMoidParams {
	var ()
	return &PatchAssetDeviceConfigurationsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAssetDeviceConfigurationsMoidParamsWithTimeout creates a new PatchAssetDeviceConfigurationsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAssetDeviceConfigurationsMoidParamsWithTimeout(timeout time.Duration) *PatchAssetDeviceConfigurationsMoidParams {
	var ()
	return &PatchAssetDeviceConfigurationsMoidParams{

		timeout: timeout,
	}
}

// NewPatchAssetDeviceConfigurationsMoidParamsWithContext creates a new PatchAssetDeviceConfigurationsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAssetDeviceConfigurationsMoidParamsWithContext(ctx context.Context) *PatchAssetDeviceConfigurationsMoidParams {
	var ()
	return &PatchAssetDeviceConfigurationsMoidParams{

		Context: ctx,
	}
}

// NewPatchAssetDeviceConfigurationsMoidParamsWithHTTPClient creates a new PatchAssetDeviceConfigurationsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAssetDeviceConfigurationsMoidParamsWithHTTPClient(client *http.Client) *PatchAssetDeviceConfigurationsMoidParams {
	var ()
	return &PatchAssetDeviceConfigurationsMoidParams{
		HTTPClient: client,
	}
}

/*PatchAssetDeviceConfigurationsMoidParams contains all the parameters to send to the API endpoint
for the patch asset device configurations moid operation typically these are written to a http.Request
*/
type PatchAssetDeviceConfigurationsMoidParams struct {

	/*Body
	  assetDeviceConfiguration to update

	*/
	Body *models.AssetDeviceConfiguration
	/*Moid
	  The moid of the assetDeviceConfiguration instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) WithTimeout(timeout time.Duration) *PatchAssetDeviceConfigurationsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) WithContext(ctx context.Context) *PatchAssetDeviceConfigurationsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) WithHTTPClient(client *http.Client) *PatchAssetDeviceConfigurationsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) WithBody(body *models.AssetDeviceConfiguration) *PatchAssetDeviceConfigurationsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) SetBody(body *models.AssetDeviceConfiguration) {
	o.Body = body
}

// WithMoid adds the moid to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) WithMoid(moid string) *PatchAssetDeviceConfigurationsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch asset device configurations moid params
func (o *PatchAssetDeviceConfigurationsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAssetDeviceConfigurationsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
