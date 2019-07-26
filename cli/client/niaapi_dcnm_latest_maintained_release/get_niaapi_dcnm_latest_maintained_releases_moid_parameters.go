// Code generated by go-swagger; DO NOT EDIT.

package niaapi_dcnm_latest_maintained_release

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

// NewGetNiaapiDcnmLatestMaintainedReleasesMoidParams creates a new GetNiaapiDcnmLatestMaintainedReleasesMoidParams object
// with the default values initialized.
func NewGetNiaapiDcnmLatestMaintainedReleasesMoidParams() *GetNiaapiDcnmLatestMaintainedReleasesMoidParams {
	var ()
	return &GetNiaapiDcnmLatestMaintainedReleasesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNiaapiDcnmLatestMaintainedReleasesMoidParamsWithTimeout creates a new GetNiaapiDcnmLatestMaintainedReleasesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNiaapiDcnmLatestMaintainedReleasesMoidParamsWithTimeout(timeout time.Duration) *GetNiaapiDcnmLatestMaintainedReleasesMoidParams {
	var ()
	return &GetNiaapiDcnmLatestMaintainedReleasesMoidParams{

		timeout: timeout,
	}
}

// NewGetNiaapiDcnmLatestMaintainedReleasesMoidParamsWithContext creates a new GetNiaapiDcnmLatestMaintainedReleasesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNiaapiDcnmLatestMaintainedReleasesMoidParamsWithContext(ctx context.Context) *GetNiaapiDcnmLatestMaintainedReleasesMoidParams {
	var ()
	return &GetNiaapiDcnmLatestMaintainedReleasesMoidParams{

		Context: ctx,
	}
}

// NewGetNiaapiDcnmLatestMaintainedReleasesMoidParamsWithHTTPClient creates a new GetNiaapiDcnmLatestMaintainedReleasesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNiaapiDcnmLatestMaintainedReleasesMoidParamsWithHTTPClient(client *http.Client) *GetNiaapiDcnmLatestMaintainedReleasesMoidParams {
	var ()
	return &GetNiaapiDcnmLatestMaintainedReleasesMoidParams{
		HTTPClient: client,
	}
}

/*GetNiaapiDcnmLatestMaintainedReleasesMoidParams contains all the parameters to send to the API endpoint
for the get niaapi dcnm latest maintained releases moid operation typically these are written to a http.Request
*/
type GetNiaapiDcnmLatestMaintainedReleasesMoidParams struct {

	/*Moid
	  The moid of the niaapiDcnmLatestMaintainedRelease instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get niaapi dcnm latest maintained releases moid params
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) WithTimeout(timeout time.Duration) *GetNiaapiDcnmLatestMaintainedReleasesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get niaapi dcnm latest maintained releases moid params
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get niaapi dcnm latest maintained releases moid params
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) WithContext(ctx context.Context) *GetNiaapiDcnmLatestMaintainedReleasesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get niaapi dcnm latest maintained releases moid params
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get niaapi dcnm latest maintained releases moid params
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) WithHTTPClient(client *http.Client) *GetNiaapiDcnmLatestMaintainedReleasesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get niaapi dcnm latest maintained releases moid params
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get niaapi dcnm latest maintained releases moid params
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) WithMoid(moid string) *GetNiaapiDcnmLatestMaintainedReleasesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get niaapi dcnm latest maintained releases moid params
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
