// Code generated by go-swagger; DO NOT EDIT.

package firmware_driver_distributable

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

// NewPostFirmwareDriverDistributablesParams creates a new PostFirmwareDriverDistributablesParams object
// with the default values initialized.
func NewPostFirmwareDriverDistributablesParams() *PostFirmwareDriverDistributablesParams {
	var ()
	return &PostFirmwareDriverDistributablesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFirmwareDriverDistributablesParamsWithTimeout creates a new PostFirmwareDriverDistributablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFirmwareDriverDistributablesParamsWithTimeout(timeout time.Duration) *PostFirmwareDriverDistributablesParams {
	var ()
	return &PostFirmwareDriverDistributablesParams{

		timeout: timeout,
	}
}

// NewPostFirmwareDriverDistributablesParamsWithContext creates a new PostFirmwareDriverDistributablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFirmwareDriverDistributablesParamsWithContext(ctx context.Context) *PostFirmwareDriverDistributablesParams {
	var ()
	return &PostFirmwareDriverDistributablesParams{

		Context: ctx,
	}
}

// NewPostFirmwareDriverDistributablesParamsWithHTTPClient creates a new PostFirmwareDriverDistributablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFirmwareDriverDistributablesParamsWithHTTPClient(client *http.Client) *PostFirmwareDriverDistributablesParams {
	var ()
	return &PostFirmwareDriverDistributablesParams{
		HTTPClient: client,
	}
}

/*PostFirmwareDriverDistributablesParams contains all the parameters to send to the API endpoint
for the post firmware driver distributables operation typically these are written to a http.Request
*/
type PostFirmwareDriverDistributablesParams struct {

	/*Body
	  firmwareDriverDistributable to add

	*/
	Body *models.FirmwareDriverDistributable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post firmware driver distributables params
func (o *PostFirmwareDriverDistributablesParams) WithTimeout(timeout time.Duration) *PostFirmwareDriverDistributablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post firmware driver distributables params
func (o *PostFirmwareDriverDistributablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post firmware driver distributables params
func (o *PostFirmwareDriverDistributablesParams) WithContext(ctx context.Context) *PostFirmwareDriverDistributablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post firmware driver distributables params
func (o *PostFirmwareDriverDistributablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post firmware driver distributables params
func (o *PostFirmwareDriverDistributablesParams) WithHTTPClient(client *http.Client) *PostFirmwareDriverDistributablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post firmware driver distributables params
func (o *PostFirmwareDriverDistributablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post firmware driver distributables params
func (o *PostFirmwareDriverDistributablesParams) WithBody(body *models.FirmwareDriverDistributable) *PostFirmwareDriverDistributablesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post firmware driver distributables params
func (o *PostFirmwareDriverDistributablesParams) SetBody(body *models.FirmwareDriverDistributable) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostFirmwareDriverDistributablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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