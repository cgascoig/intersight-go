// Code generated by go-swagger; DO NOT EDIT.

package softwarerepository_operating_system_file

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

// NewPostSoftwarerepositoryOperatingSystemFilesParams creates a new PostSoftwarerepositoryOperatingSystemFilesParams object
// with the default values initialized.
func NewPostSoftwarerepositoryOperatingSystemFilesParams() *PostSoftwarerepositoryOperatingSystemFilesParams {
	var ()
	return &PostSoftwarerepositoryOperatingSystemFilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSoftwarerepositoryOperatingSystemFilesParamsWithTimeout creates a new PostSoftwarerepositoryOperatingSystemFilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSoftwarerepositoryOperatingSystemFilesParamsWithTimeout(timeout time.Duration) *PostSoftwarerepositoryOperatingSystemFilesParams {
	var ()
	return &PostSoftwarerepositoryOperatingSystemFilesParams{

		timeout: timeout,
	}
}

// NewPostSoftwarerepositoryOperatingSystemFilesParamsWithContext creates a new PostSoftwarerepositoryOperatingSystemFilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSoftwarerepositoryOperatingSystemFilesParamsWithContext(ctx context.Context) *PostSoftwarerepositoryOperatingSystemFilesParams {
	var ()
	return &PostSoftwarerepositoryOperatingSystemFilesParams{

		Context: ctx,
	}
}

// NewPostSoftwarerepositoryOperatingSystemFilesParamsWithHTTPClient creates a new PostSoftwarerepositoryOperatingSystemFilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSoftwarerepositoryOperatingSystemFilesParamsWithHTTPClient(client *http.Client) *PostSoftwarerepositoryOperatingSystemFilesParams {
	var ()
	return &PostSoftwarerepositoryOperatingSystemFilesParams{
		HTTPClient: client,
	}
}

/*PostSoftwarerepositoryOperatingSystemFilesParams contains all the parameters to send to the API endpoint
for the post softwarerepository operating system files operation typically these are written to a http.Request
*/
type PostSoftwarerepositoryOperatingSystemFilesParams struct {

	/*Body
	  softwarerepositoryOperatingSystemFile to add

	*/
	Body *models.SoftwarerepositoryOperatingSystemFile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post softwarerepository operating system files params
func (o *PostSoftwarerepositoryOperatingSystemFilesParams) WithTimeout(timeout time.Duration) *PostSoftwarerepositoryOperatingSystemFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post softwarerepository operating system files params
func (o *PostSoftwarerepositoryOperatingSystemFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post softwarerepository operating system files params
func (o *PostSoftwarerepositoryOperatingSystemFilesParams) WithContext(ctx context.Context) *PostSoftwarerepositoryOperatingSystemFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post softwarerepository operating system files params
func (o *PostSoftwarerepositoryOperatingSystemFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post softwarerepository operating system files params
func (o *PostSoftwarerepositoryOperatingSystemFilesParams) WithHTTPClient(client *http.Client) *PostSoftwarerepositoryOperatingSystemFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post softwarerepository operating system files params
func (o *PostSoftwarerepositoryOperatingSystemFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post softwarerepository operating system files params
func (o *PostSoftwarerepositoryOperatingSystemFilesParams) WithBody(body *models.SoftwarerepositoryOperatingSystemFile) *PostSoftwarerepositoryOperatingSystemFilesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post softwarerepository operating system files params
func (o *PostSoftwarerepositoryOperatingSystemFilesParams) SetBody(body *models.SoftwarerepositoryOperatingSystemFile) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSoftwarerepositoryOperatingSystemFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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