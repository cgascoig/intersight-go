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

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPostFirmwareDistributablesMoidParams creates a new PostFirmwareDistributablesMoidParams object
// with the default values initialized.
func NewPostFirmwareDistributablesMoidParams() *PostFirmwareDistributablesMoidParams {
	var ()
	return &PostFirmwareDistributablesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFirmwareDistributablesMoidParamsWithTimeout creates a new PostFirmwareDistributablesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFirmwareDistributablesMoidParamsWithTimeout(timeout time.Duration) *PostFirmwareDistributablesMoidParams {
	var ()
	return &PostFirmwareDistributablesMoidParams{

		timeout: timeout,
	}
}

// NewPostFirmwareDistributablesMoidParamsWithContext creates a new PostFirmwareDistributablesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFirmwareDistributablesMoidParamsWithContext(ctx context.Context) *PostFirmwareDistributablesMoidParams {
	var ()
	return &PostFirmwareDistributablesMoidParams{

		Context: ctx,
	}
}

// NewPostFirmwareDistributablesMoidParamsWithHTTPClient creates a new PostFirmwareDistributablesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFirmwareDistributablesMoidParamsWithHTTPClient(client *http.Client) *PostFirmwareDistributablesMoidParams {
	var ()
	return &PostFirmwareDistributablesMoidParams{
		HTTPClient: client,
	}
}

/*PostFirmwareDistributablesMoidParams contains all the parameters to send to the API endpoint
for the post firmware distributables moid operation typically these are written to a http.Request
*/
type PostFirmwareDistributablesMoidParams struct {

	/*Body
	  firmwareDistributable to update

	*/
	Body *models.FirmwareDistributable
	/*Moid
	  The moid of the firmwareDistributable instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) WithTimeout(timeout time.Duration) *PostFirmwareDistributablesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) WithContext(ctx context.Context) *PostFirmwareDistributablesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) WithHTTPClient(client *http.Client) *PostFirmwareDistributablesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) WithBody(body *models.FirmwareDistributable) *PostFirmwareDistributablesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) SetBody(body *models.FirmwareDistributable) {
	o.Body = body
}

// WithMoid adds the moid to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) WithMoid(moid string) *PostFirmwareDistributablesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post firmware distributables moid params
func (o *PostFirmwareDistributablesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostFirmwareDistributablesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
