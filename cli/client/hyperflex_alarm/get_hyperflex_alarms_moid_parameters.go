// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_alarm

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

// NewGetHyperflexAlarmsMoidParams creates a new GetHyperflexAlarmsMoidParams object
// with the default values initialized.
func NewGetHyperflexAlarmsMoidParams() *GetHyperflexAlarmsMoidParams {
	var ()
	return &GetHyperflexAlarmsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHyperflexAlarmsMoidParamsWithTimeout creates a new GetHyperflexAlarmsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHyperflexAlarmsMoidParamsWithTimeout(timeout time.Duration) *GetHyperflexAlarmsMoidParams {
	var ()
	return &GetHyperflexAlarmsMoidParams{

		timeout: timeout,
	}
}

// NewGetHyperflexAlarmsMoidParamsWithContext creates a new GetHyperflexAlarmsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHyperflexAlarmsMoidParamsWithContext(ctx context.Context) *GetHyperflexAlarmsMoidParams {
	var ()
	return &GetHyperflexAlarmsMoidParams{

		Context: ctx,
	}
}

// NewGetHyperflexAlarmsMoidParamsWithHTTPClient creates a new GetHyperflexAlarmsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHyperflexAlarmsMoidParamsWithHTTPClient(client *http.Client) *GetHyperflexAlarmsMoidParams {
	var ()
	return &GetHyperflexAlarmsMoidParams{
		HTTPClient: client,
	}
}

/*GetHyperflexAlarmsMoidParams contains all the parameters to send to the API endpoint
for the get hyperflex alarms moid operation typically these are written to a http.Request
*/
type GetHyperflexAlarmsMoidParams struct {

	/*Moid
	  The moid of the hyperflexAlarm instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hyperflex alarms moid params
func (o *GetHyperflexAlarmsMoidParams) WithTimeout(timeout time.Duration) *GetHyperflexAlarmsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hyperflex alarms moid params
func (o *GetHyperflexAlarmsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hyperflex alarms moid params
func (o *GetHyperflexAlarmsMoidParams) WithContext(ctx context.Context) *GetHyperflexAlarmsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hyperflex alarms moid params
func (o *GetHyperflexAlarmsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hyperflex alarms moid params
func (o *GetHyperflexAlarmsMoidParams) WithHTTPClient(client *http.Client) *GetHyperflexAlarmsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hyperflex alarms moid params
func (o *GetHyperflexAlarmsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hyperflex alarms moid params
func (o *GetHyperflexAlarmsMoidParams) WithMoid(moid string) *GetHyperflexAlarmsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hyperflex alarms moid params
func (o *GetHyperflexAlarmsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHyperflexAlarmsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
