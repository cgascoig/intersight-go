// Code generated by go-swagger; DO NOT EDIT.

package softwarerepository_authorization

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

// NewPostSoftwarerepositoryAuthorizationsParams creates a new PostSoftwarerepositoryAuthorizationsParams object
// with the default values initialized.
func NewPostSoftwarerepositoryAuthorizationsParams() *PostSoftwarerepositoryAuthorizationsParams {
	var ()
	return &PostSoftwarerepositoryAuthorizationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSoftwarerepositoryAuthorizationsParamsWithTimeout creates a new PostSoftwarerepositoryAuthorizationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSoftwarerepositoryAuthorizationsParamsWithTimeout(timeout time.Duration) *PostSoftwarerepositoryAuthorizationsParams {
	var ()
	return &PostSoftwarerepositoryAuthorizationsParams{

		timeout: timeout,
	}
}

// NewPostSoftwarerepositoryAuthorizationsParamsWithContext creates a new PostSoftwarerepositoryAuthorizationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSoftwarerepositoryAuthorizationsParamsWithContext(ctx context.Context) *PostSoftwarerepositoryAuthorizationsParams {
	var ()
	return &PostSoftwarerepositoryAuthorizationsParams{

		Context: ctx,
	}
}

// NewPostSoftwarerepositoryAuthorizationsParamsWithHTTPClient creates a new PostSoftwarerepositoryAuthorizationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSoftwarerepositoryAuthorizationsParamsWithHTTPClient(client *http.Client) *PostSoftwarerepositoryAuthorizationsParams {
	var ()
	return &PostSoftwarerepositoryAuthorizationsParams{
		HTTPClient: client,
	}
}

/*PostSoftwarerepositoryAuthorizationsParams contains all the parameters to send to the API endpoint
for the post softwarerepository authorizations operation typically these are written to a http.Request
*/
type PostSoftwarerepositoryAuthorizationsParams struct {

	/*Body
	  softwarerepositoryAuthorization to add

	*/
	Body *models.SoftwarerepositoryAuthorization

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post softwarerepository authorizations params
func (o *PostSoftwarerepositoryAuthorizationsParams) WithTimeout(timeout time.Duration) *PostSoftwarerepositoryAuthorizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post softwarerepository authorizations params
func (o *PostSoftwarerepositoryAuthorizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post softwarerepository authorizations params
func (o *PostSoftwarerepositoryAuthorizationsParams) WithContext(ctx context.Context) *PostSoftwarerepositoryAuthorizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post softwarerepository authorizations params
func (o *PostSoftwarerepositoryAuthorizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post softwarerepository authorizations params
func (o *PostSoftwarerepositoryAuthorizationsParams) WithHTTPClient(client *http.Client) *PostSoftwarerepositoryAuthorizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post softwarerepository authorizations params
func (o *PostSoftwarerepositoryAuthorizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post softwarerepository authorizations params
func (o *PostSoftwarerepositoryAuthorizationsParams) WithBody(body *models.SoftwarerepositoryAuthorization) *PostSoftwarerepositoryAuthorizationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post softwarerepository authorizations params
func (o *PostSoftwarerepositoryAuthorizationsParams) SetBody(body *models.SoftwarerepositoryAuthorization) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSoftwarerepositoryAuthorizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
