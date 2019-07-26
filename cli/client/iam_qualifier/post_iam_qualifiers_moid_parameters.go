// Code generated by go-swagger; DO NOT EDIT.

package iam_qualifier

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

// NewPostIamQualifiersMoidParams creates a new PostIamQualifiersMoidParams object
// with the default values initialized.
func NewPostIamQualifiersMoidParams() *PostIamQualifiersMoidParams {
	var ()
	return &PostIamQualifiersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIamQualifiersMoidParamsWithTimeout creates a new PostIamQualifiersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIamQualifiersMoidParamsWithTimeout(timeout time.Duration) *PostIamQualifiersMoidParams {
	var ()
	return &PostIamQualifiersMoidParams{

		timeout: timeout,
	}
}

// NewPostIamQualifiersMoidParamsWithContext creates a new PostIamQualifiersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIamQualifiersMoidParamsWithContext(ctx context.Context) *PostIamQualifiersMoidParams {
	var ()
	return &PostIamQualifiersMoidParams{

		Context: ctx,
	}
}

// NewPostIamQualifiersMoidParamsWithHTTPClient creates a new PostIamQualifiersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIamQualifiersMoidParamsWithHTTPClient(client *http.Client) *PostIamQualifiersMoidParams {
	var ()
	return &PostIamQualifiersMoidParams{
		HTTPClient: client,
	}
}

/*PostIamQualifiersMoidParams contains all the parameters to send to the API endpoint
for the post iam qualifiers moid operation typically these are written to a http.Request
*/
type PostIamQualifiersMoidParams struct {

	/*Body
	  iamQualifier to update

	*/
	Body *models.IamQualifier
	/*Moid
	  The moid of the iamQualifier instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) WithTimeout(timeout time.Duration) *PostIamQualifiersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) WithContext(ctx context.Context) *PostIamQualifiersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) WithHTTPClient(client *http.Client) *PostIamQualifiersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) WithBody(body *models.IamQualifier) *PostIamQualifiersMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) SetBody(body *models.IamQualifier) {
	o.Body = body
}

// WithMoid adds the moid to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) WithMoid(moid string) *PostIamQualifiersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post iam qualifiers moid params
func (o *PostIamQualifiersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostIamQualifiersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
