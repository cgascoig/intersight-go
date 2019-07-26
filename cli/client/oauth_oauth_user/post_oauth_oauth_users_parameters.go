// Code generated by go-swagger; DO NOT EDIT.

package oauth_oauth_user

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

// NewPostOauthOauthUsersParams creates a new PostOauthOauthUsersParams object
// with the default values initialized.
func NewPostOauthOauthUsersParams() *PostOauthOauthUsersParams {
	var ()
	return &PostOauthOauthUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOauthOauthUsersParamsWithTimeout creates a new PostOauthOauthUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOauthOauthUsersParamsWithTimeout(timeout time.Duration) *PostOauthOauthUsersParams {
	var ()
	return &PostOauthOauthUsersParams{

		timeout: timeout,
	}
}

// NewPostOauthOauthUsersParamsWithContext creates a new PostOauthOauthUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOauthOauthUsersParamsWithContext(ctx context.Context) *PostOauthOauthUsersParams {
	var ()
	return &PostOauthOauthUsersParams{

		Context: ctx,
	}
}

// NewPostOauthOauthUsersParamsWithHTTPClient creates a new PostOauthOauthUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOauthOauthUsersParamsWithHTTPClient(client *http.Client) *PostOauthOauthUsersParams {
	var ()
	return &PostOauthOauthUsersParams{
		HTTPClient: client,
	}
}

/*PostOauthOauthUsersParams contains all the parameters to send to the API endpoint
for the post oauth oauth users operation typically these are written to a http.Request
*/
type PostOauthOauthUsersParams struct {

	/*Body
	  oauthOauthUser to add

	*/
	Body *models.OauthOauthUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post oauth oauth users params
func (o *PostOauthOauthUsersParams) WithTimeout(timeout time.Duration) *PostOauthOauthUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post oauth oauth users params
func (o *PostOauthOauthUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post oauth oauth users params
func (o *PostOauthOauthUsersParams) WithContext(ctx context.Context) *PostOauthOauthUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post oauth oauth users params
func (o *PostOauthOauthUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post oauth oauth users params
func (o *PostOauthOauthUsersParams) WithHTTPClient(client *http.Client) *PostOauthOauthUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post oauth oauth users params
func (o *PostOauthOauthUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post oauth oauth users params
func (o *PostOauthOauthUsersParams) WithBody(body *models.OauthOauthUser) *PostOauthOauthUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post oauth oauth users params
func (o *PostOauthOauthUsersParams) SetBody(body *models.OauthOauthUser) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostOauthOauthUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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