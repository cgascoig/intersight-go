// Code generated by go-swagger; DO NOT EDIT.

package iam_user

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

// NewPostIamUsersParams creates a new PostIamUsersParams object
// with the default values initialized.
func NewPostIamUsersParams() *PostIamUsersParams {
	var ()
	return &PostIamUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIamUsersParamsWithTimeout creates a new PostIamUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIamUsersParamsWithTimeout(timeout time.Duration) *PostIamUsersParams {
	var ()
	return &PostIamUsersParams{

		timeout: timeout,
	}
}

// NewPostIamUsersParamsWithContext creates a new PostIamUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIamUsersParamsWithContext(ctx context.Context) *PostIamUsersParams {
	var ()
	return &PostIamUsersParams{

		Context: ctx,
	}
}

// NewPostIamUsersParamsWithHTTPClient creates a new PostIamUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIamUsersParamsWithHTTPClient(client *http.Client) *PostIamUsersParams {
	var ()
	return &PostIamUsersParams{
		HTTPClient: client,
	}
}

/*PostIamUsersParams contains all the parameters to send to the API endpoint
for the post iam users operation typically these are written to a http.Request
*/
type PostIamUsersParams struct {

	/*Body
	  iamUser to add

	*/
	Body *models.IamUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post iam users params
func (o *PostIamUsersParams) WithTimeout(timeout time.Duration) *PostIamUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post iam users params
func (o *PostIamUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post iam users params
func (o *PostIamUsersParams) WithContext(ctx context.Context) *PostIamUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post iam users params
func (o *PostIamUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post iam users params
func (o *PostIamUsersParams) WithHTTPClient(client *http.Client) *PostIamUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post iam users params
func (o *PostIamUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post iam users params
func (o *PostIamUsersParams) WithBody(body *models.IamUser) *PostIamUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post iam users params
func (o *PostIamUsersParams) SetBody(body *models.IamUser) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostIamUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
