// Code generated by go-swagger; DO NOT EDIT.

package firmware_distributable

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

// NewDeleteFirmwareDistributablesMoidParams creates a new DeleteFirmwareDistributablesMoidParams object
// with the default values initialized.
func NewDeleteFirmwareDistributablesMoidParams() *DeleteFirmwareDistributablesMoidParams {
	var ()
	return &DeleteFirmwareDistributablesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFirmwareDistributablesMoidParamsWithTimeout creates a new DeleteFirmwareDistributablesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFirmwareDistributablesMoidParamsWithTimeout(timeout time.Duration) *DeleteFirmwareDistributablesMoidParams {
	var ()
	return &DeleteFirmwareDistributablesMoidParams{

		timeout: timeout,
	}
}

// NewDeleteFirmwareDistributablesMoidParamsWithContext creates a new DeleteFirmwareDistributablesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFirmwareDistributablesMoidParamsWithContext(ctx context.Context) *DeleteFirmwareDistributablesMoidParams {
	var ()
	return &DeleteFirmwareDistributablesMoidParams{

		Context: ctx,
	}
}

// NewDeleteFirmwareDistributablesMoidParamsWithHTTPClient creates a new DeleteFirmwareDistributablesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFirmwareDistributablesMoidParamsWithHTTPClient(client *http.Client) *DeleteFirmwareDistributablesMoidParams {
	var ()
	return &DeleteFirmwareDistributablesMoidParams{
		HTTPClient: client,
	}
}

/*DeleteFirmwareDistributablesMoidParams contains all the parameters to send to the API endpoint
for the delete firmware distributables moid operation typically these are written to a http.Request
*/
type DeleteFirmwareDistributablesMoidParams struct {

	/*Moid
	  The moid of the firmwareDistributable instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete firmware distributables moid params
func (o *DeleteFirmwareDistributablesMoidParams) WithTimeout(timeout time.Duration) *DeleteFirmwareDistributablesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete firmware distributables moid params
func (o *DeleteFirmwareDistributablesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete firmware distributables moid params
func (o *DeleteFirmwareDistributablesMoidParams) WithContext(ctx context.Context) *DeleteFirmwareDistributablesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete firmware distributables moid params
func (o *DeleteFirmwareDistributablesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete firmware distributables moid params
func (o *DeleteFirmwareDistributablesMoidParams) WithHTTPClient(client *http.Client) *DeleteFirmwareDistributablesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete firmware distributables moid params
func (o *DeleteFirmwareDistributablesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete firmware distributables moid params
func (o *DeleteFirmwareDistributablesMoidParams) WithMoid(moid string) *DeleteFirmwareDistributablesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete firmware distributables moid params
func (o *DeleteFirmwareDistributablesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFirmwareDistributablesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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