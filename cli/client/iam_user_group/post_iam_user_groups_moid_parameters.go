// Code generated by go-swagger; DO NOT EDIT.

package iam_user_group

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

// NewPostIamUserGroupsMoidParams creates a new PostIamUserGroupsMoidParams object
// with the default values initialized.
func NewPostIamUserGroupsMoidParams() *PostIamUserGroupsMoidParams {
	var ()
	return &PostIamUserGroupsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIamUserGroupsMoidParamsWithTimeout creates a new PostIamUserGroupsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIamUserGroupsMoidParamsWithTimeout(timeout time.Duration) *PostIamUserGroupsMoidParams {
	var ()
	return &PostIamUserGroupsMoidParams{

		timeout: timeout,
	}
}

// NewPostIamUserGroupsMoidParamsWithContext creates a new PostIamUserGroupsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIamUserGroupsMoidParamsWithContext(ctx context.Context) *PostIamUserGroupsMoidParams {
	var ()
	return &PostIamUserGroupsMoidParams{

		Context: ctx,
	}
}

// NewPostIamUserGroupsMoidParamsWithHTTPClient creates a new PostIamUserGroupsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIamUserGroupsMoidParamsWithHTTPClient(client *http.Client) *PostIamUserGroupsMoidParams {
	var ()
	return &PostIamUserGroupsMoidParams{
		HTTPClient: client,
	}
}

/*PostIamUserGroupsMoidParams contains all the parameters to send to the API endpoint
for the post iam user groups moid operation typically these are written to a http.Request
*/
type PostIamUserGroupsMoidParams struct {

	/*Body
	  iamUserGroup to update

	*/
	Body *models.IamUserGroup
	/*Moid
	  The moid of the iamUserGroup instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) WithTimeout(timeout time.Duration) *PostIamUserGroupsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) WithContext(ctx context.Context) *PostIamUserGroupsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) WithHTTPClient(client *http.Client) *PostIamUserGroupsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) WithBody(body *models.IamUserGroup) *PostIamUserGroupsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) SetBody(body *models.IamUserGroup) {
	o.Body = body
}

// WithMoid adds the moid to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) WithMoid(moid string) *PostIamUserGroupsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post iam user groups moid params
func (o *PostIamUserGroupsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostIamUserGroupsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
