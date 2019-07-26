// Code generated by go-swagger; DO NOT EDIT.

package equipment_device_summary

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

// NewGetEquipmentDeviceSummariesMoidParams creates a new GetEquipmentDeviceSummariesMoidParams object
// with the default values initialized.
func NewGetEquipmentDeviceSummariesMoidParams() *GetEquipmentDeviceSummariesMoidParams {
	var ()
	return &GetEquipmentDeviceSummariesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEquipmentDeviceSummariesMoidParamsWithTimeout creates a new GetEquipmentDeviceSummariesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEquipmentDeviceSummariesMoidParamsWithTimeout(timeout time.Duration) *GetEquipmentDeviceSummariesMoidParams {
	var ()
	return &GetEquipmentDeviceSummariesMoidParams{

		timeout: timeout,
	}
}

// NewGetEquipmentDeviceSummariesMoidParamsWithContext creates a new GetEquipmentDeviceSummariesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEquipmentDeviceSummariesMoidParamsWithContext(ctx context.Context) *GetEquipmentDeviceSummariesMoidParams {
	var ()
	return &GetEquipmentDeviceSummariesMoidParams{

		Context: ctx,
	}
}

// NewGetEquipmentDeviceSummariesMoidParamsWithHTTPClient creates a new GetEquipmentDeviceSummariesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEquipmentDeviceSummariesMoidParamsWithHTTPClient(client *http.Client) *GetEquipmentDeviceSummariesMoidParams {
	var ()
	return &GetEquipmentDeviceSummariesMoidParams{
		HTTPClient: client,
	}
}

/*GetEquipmentDeviceSummariesMoidParams contains all the parameters to send to the API endpoint
for the get equipment device summaries moid operation typically these are written to a http.Request
*/
type GetEquipmentDeviceSummariesMoidParams struct {

	/*Moid
	  The moid of the equipmentDeviceSummary instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get equipment device summaries moid params
func (o *GetEquipmentDeviceSummariesMoidParams) WithTimeout(timeout time.Duration) *GetEquipmentDeviceSummariesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get equipment device summaries moid params
func (o *GetEquipmentDeviceSummariesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get equipment device summaries moid params
func (o *GetEquipmentDeviceSummariesMoidParams) WithContext(ctx context.Context) *GetEquipmentDeviceSummariesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get equipment device summaries moid params
func (o *GetEquipmentDeviceSummariesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get equipment device summaries moid params
func (o *GetEquipmentDeviceSummariesMoidParams) WithHTTPClient(client *http.Client) *GetEquipmentDeviceSummariesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get equipment device summaries moid params
func (o *GetEquipmentDeviceSummariesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get equipment device summaries moid params
func (o *GetEquipmentDeviceSummariesMoidParams) WithMoid(moid string) *GetEquipmentDeviceSummariesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get equipment device summaries moid params
func (o *GetEquipmentDeviceSummariesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetEquipmentDeviceSummariesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
