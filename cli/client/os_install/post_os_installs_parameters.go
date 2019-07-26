// Code generated by go-swagger; DO NOT EDIT.

package os_install

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

// NewPostOsInstallsParams creates a new PostOsInstallsParams object
// with the default values initialized.
func NewPostOsInstallsParams() *PostOsInstallsParams {
	var ()
	return &PostOsInstallsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOsInstallsParamsWithTimeout creates a new PostOsInstallsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOsInstallsParamsWithTimeout(timeout time.Duration) *PostOsInstallsParams {
	var ()
	return &PostOsInstallsParams{

		timeout: timeout,
	}
}

// NewPostOsInstallsParamsWithContext creates a new PostOsInstallsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOsInstallsParamsWithContext(ctx context.Context) *PostOsInstallsParams {
	var ()
	return &PostOsInstallsParams{

		Context: ctx,
	}
}

// NewPostOsInstallsParamsWithHTTPClient creates a new PostOsInstallsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOsInstallsParamsWithHTTPClient(client *http.Client) *PostOsInstallsParams {
	var ()
	return &PostOsInstallsParams{
		HTTPClient: client,
	}
}

/*PostOsInstallsParams contains all the parameters to send to the API endpoint
for the post os installs operation typically these are written to a http.Request
*/
type PostOsInstallsParams struct {

	/*Body
	  osInstall to add

	*/
	Body *models.OsInstall

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post os installs params
func (o *PostOsInstallsParams) WithTimeout(timeout time.Duration) *PostOsInstallsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post os installs params
func (o *PostOsInstallsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post os installs params
func (o *PostOsInstallsParams) WithContext(ctx context.Context) *PostOsInstallsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post os installs params
func (o *PostOsInstallsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post os installs params
func (o *PostOsInstallsParams) WithHTTPClient(client *http.Client) *PostOsInstallsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post os installs params
func (o *PostOsInstallsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post os installs params
func (o *PostOsInstallsParams) WithBody(body *models.OsInstall) *PostOsInstallsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post os installs params
func (o *PostOsInstallsParams) SetBody(body *models.OsInstall) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostOsInstallsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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