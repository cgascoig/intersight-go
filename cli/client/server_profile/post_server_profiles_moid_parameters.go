// Code generated by go-swagger; DO NOT EDIT.

package server_profile

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

// NewPostServerProfilesMoidParams creates a new PostServerProfilesMoidParams object
// with the default values initialized.
func NewPostServerProfilesMoidParams() *PostServerProfilesMoidParams {
	var ()
	return &PostServerProfilesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostServerProfilesMoidParamsWithTimeout creates a new PostServerProfilesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostServerProfilesMoidParamsWithTimeout(timeout time.Duration) *PostServerProfilesMoidParams {
	var ()
	return &PostServerProfilesMoidParams{

		timeout: timeout,
	}
}

// NewPostServerProfilesMoidParamsWithContext creates a new PostServerProfilesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostServerProfilesMoidParamsWithContext(ctx context.Context) *PostServerProfilesMoidParams {
	var ()
	return &PostServerProfilesMoidParams{

		Context: ctx,
	}
}

// NewPostServerProfilesMoidParamsWithHTTPClient creates a new PostServerProfilesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostServerProfilesMoidParamsWithHTTPClient(client *http.Client) *PostServerProfilesMoidParams {
	var ()
	return &PostServerProfilesMoidParams{
		HTTPClient: client,
	}
}

/*PostServerProfilesMoidParams contains all the parameters to send to the API endpoint
for the post server profiles moid operation typically these are written to a http.Request
*/
type PostServerProfilesMoidParams struct {

	/*Body
	  serverProfile to update

	*/
	Body *models.ServerProfile
	/*Moid
	  The moid of the serverProfile instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post server profiles moid params
func (o *PostServerProfilesMoidParams) WithTimeout(timeout time.Duration) *PostServerProfilesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post server profiles moid params
func (o *PostServerProfilesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post server profiles moid params
func (o *PostServerProfilesMoidParams) WithContext(ctx context.Context) *PostServerProfilesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post server profiles moid params
func (o *PostServerProfilesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post server profiles moid params
func (o *PostServerProfilesMoidParams) WithHTTPClient(client *http.Client) *PostServerProfilesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post server profiles moid params
func (o *PostServerProfilesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post server profiles moid params
func (o *PostServerProfilesMoidParams) WithBody(body *models.ServerProfile) *PostServerProfilesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post server profiles moid params
func (o *PostServerProfilesMoidParams) SetBody(body *models.ServerProfile) {
	o.Body = body
}

// WithMoid adds the moid to the post server profiles moid params
func (o *PostServerProfilesMoidParams) WithMoid(moid string) *PostServerProfilesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post server profiles moid params
func (o *PostServerProfilesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostServerProfilesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
