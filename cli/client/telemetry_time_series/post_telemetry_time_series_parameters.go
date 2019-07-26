// Code generated by go-swagger; DO NOT EDIT.

package telemetry_time_series

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

// NewPostTelemetryTimeSeriesParams creates a new PostTelemetryTimeSeriesParams object
// with the default values initialized.
func NewPostTelemetryTimeSeriesParams() *PostTelemetryTimeSeriesParams {
	var ()
	return &PostTelemetryTimeSeriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTelemetryTimeSeriesParamsWithTimeout creates a new PostTelemetryTimeSeriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTelemetryTimeSeriesParamsWithTimeout(timeout time.Duration) *PostTelemetryTimeSeriesParams {
	var ()
	return &PostTelemetryTimeSeriesParams{

		timeout: timeout,
	}
}

// NewPostTelemetryTimeSeriesParamsWithContext creates a new PostTelemetryTimeSeriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTelemetryTimeSeriesParamsWithContext(ctx context.Context) *PostTelemetryTimeSeriesParams {
	var ()
	return &PostTelemetryTimeSeriesParams{

		Context: ctx,
	}
}

// NewPostTelemetryTimeSeriesParamsWithHTTPClient creates a new PostTelemetryTimeSeriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTelemetryTimeSeriesParamsWithHTTPClient(client *http.Client) *PostTelemetryTimeSeriesParams {
	var ()
	return &PostTelemetryTimeSeriesParams{
		HTTPClient: client,
	}
}

/*PostTelemetryTimeSeriesParams contains all the parameters to send to the API endpoint
for the post telemetry time series operation typically these are written to a http.Request
*/
type PostTelemetryTimeSeriesParams struct {

	/*Body
	  telemetryTimeSeries to add

	*/
	Body *models.TelemetryTimeSeries

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post telemetry time series params
func (o *PostTelemetryTimeSeriesParams) WithTimeout(timeout time.Duration) *PostTelemetryTimeSeriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post telemetry time series params
func (o *PostTelemetryTimeSeriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post telemetry time series params
func (o *PostTelemetryTimeSeriesParams) WithContext(ctx context.Context) *PostTelemetryTimeSeriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post telemetry time series params
func (o *PostTelemetryTimeSeriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post telemetry time series params
func (o *PostTelemetryTimeSeriesParams) WithHTTPClient(client *http.Client) *PostTelemetryTimeSeriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post telemetry time series params
func (o *PostTelemetryTimeSeriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post telemetry time series params
func (o *PostTelemetryTimeSeriesParams) WithBody(body *models.TelemetryTimeSeries) *PostTelemetryTimeSeriesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post telemetry time series params
func (o *PostTelemetryTimeSeriesParams) SetBody(body *models.TelemetryTimeSeries) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTelemetryTimeSeriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}