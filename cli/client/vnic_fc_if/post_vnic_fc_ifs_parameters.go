// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_if

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

// NewPostVnicFcIfsParams creates a new PostVnicFcIfsParams object
// with the default values initialized.
func NewPostVnicFcIfsParams() *PostVnicFcIfsParams {
	var ()
	return &PostVnicFcIfsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVnicFcIfsParamsWithTimeout creates a new PostVnicFcIfsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVnicFcIfsParamsWithTimeout(timeout time.Duration) *PostVnicFcIfsParams {
	var ()
	return &PostVnicFcIfsParams{

		timeout: timeout,
	}
}

// NewPostVnicFcIfsParamsWithContext creates a new PostVnicFcIfsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVnicFcIfsParamsWithContext(ctx context.Context) *PostVnicFcIfsParams {
	var ()
	return &PostVnicFcIfsParams{

		Context: ctx,
	}
}

// NewPostVnicFcIfsParamsWithHTTPClient creates a new PostVnicFcIfsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVnicFcIfsParamsWithHTTPClient(client *http.Client) *PostVnicFcIfsParams {
	var ()
	return &PostVnicFcIfsParams{
		HTTPClient: client,
	}
}

/*PostVnicFcIfsParams contains all the parameters to send to the API endpoint
for the post vnic fc ifs operation typically these are written to a http.Request
*/
type PostVnicFcIfsParams struct {

	/*Body
	  vnicFcIf to add

	*/
	Body *models.VnicFcIf

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vnic fc ifs params
func (o *PostVnicFcIfsParams) WithTimeout(timeout time.Duration) *PostVnicFcIfsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vnic fc ifs params
func (o *PostVnicFcIfsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vnic fc ifs params
func (o *PostVnicFcIfsParams) WithContext(ctx context.Context) *PostVnicFcIfsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vnic fc ifs params
func (o *PostVnicFcIfsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vnic fc ifs params
func (o *PostVnicFcIfsParams) WithHTTPClient(client *http.Client) *PostVnicFcIfsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vnic fc ifs params
func (o *PostVnicFcIfsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vnic fc ifs params
func (o *PostVnicFcIfsParams) WithBody(body *models.VnicFcIf) *PostVnicFcIfsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vnic fc ifs params
func (o *PostVnicFcIfsParams) SetBody(body *models.VnicFcIf) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVnicFcIfsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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