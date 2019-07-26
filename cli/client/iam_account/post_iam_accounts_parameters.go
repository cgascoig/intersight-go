// Code generated by go-swagger; DO NOT EDIT.

package iam_account

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

// NewPostIamAccountsParams creates a new PostIamAccountsParams object
// with the default values initialized.
func NewPostIamAccountsParams() *PostIamAccountsParams {
	var ()
	return &PostIamAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIamAccountsParamsWithTimeout creates a new PostIamAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIamAccountsParamsWithTimeout(timeout time.Duration) *PostIamAccountsParams {
	var ()
	return &PostIamAccountsParams{

		timeout: timeout,
	}
}

// NewPostIamAccountsParamsWithContext creates a new PostIamAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIamAccountsParamsWithContext(ctx context.Context) *PostIamAccountsParams {
	var ()
	return &PostIamAccountsParams{

		Context: ctx,
	}
}

// NewPostIamAccountsParamsWithHTTPClient creates a new PostIamAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIamAccountsParamsWithHTTPClient(client *http.Client) *PostIamAccountsParams {
	var ()
	return &PostIamAccountsParams{
		HTTPClient: client,
	}
}

/*PostIamAccountsParams contains all the parameters to send to the API endpoint
for the post iam accounts operation typically these are written to a http.Request
*/
type PostIamAccountsParams struct {

	/*Body
	  iamAccount to add

	*/
	Body *models.IamAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post iam accounts params
func (o *PostIamAccountsParams) WithTimeout(timeout time.Duration) *PostIamAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post iam accounts params
func (o *PostIamAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post iam accounts params
func (o *PostIamAccountsParams) WithContext(ctx context.Context) *PostIamAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post iam accounts params
func (o *PostIamAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post iam accounts params
func (o *PostIamAccountsParams) WithHTTPClient(client *http.Client) *PostIamAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post iam accounts params
func (o *PostIamAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post iam accounts params
func (o *PostIamAccountsParams) WithBody(body *models.IamAccount) *PostIamAccountsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post iam accounts params
func (o *PostIamAccountsParams) SetBody(body *models.IamAccount) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostIamAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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