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
)

// NewDeleteIamUsersMoidParams creates a new DeleteIamUsersMoidParams object
// with the default values initialized.
func NewDeleteIamUsersMoidParams() *DeleteIamUsersMoidParams {
	var ()
	return &DeleteIamUsersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIamUsersMoidParamsWithTimeout creates a new DeleteIamUsersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIamUsersMoidParamsWithTimeout(timeout time.Duration) *DeleteIamUsersMoidParams {
	var ()
	return &DeleteIamUsersMoidParams{

		timeout: timeout,
	}
}

// NewDeleteIamUsersMoidParamsWithContext creates a new DeleteIamUsersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIamUsersMoidParamsWithContext(ctx context.Context) *DeleteIamUsersMoidParams {
	var ()
	return &DeleteIamUsersMoidParams{

		Context: ctx,
	}
}

// NewDeleteIamUsersMoidParamsWithHTTPClient creates a new DeleteIamUsersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIamUsersMoidParamsWithHTTPClient(client *http.Client) *DeleteIamUsersMoidParams {
	var ()
	return &DeleteIamUsersMoidParams{
		HTTPClient: client,
	}
}

/*DeleteIamUsersMoidParams contains all the parameters to send to the API endpoint
for the delete iam users moid operation typically these are written to a http.Request
*/
type DeleteIamUsersMoidParams struct {

	/*Moid
	  The moid of the iamUser instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete iam users moid params
func (o *DeleteIamUsersMoidParams) WithTimeout(timeout time.Duration) *DeleteIamUsersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete iam users moid params
func (o *DeleteIamUsersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete iam users moid params
func (o *DeleteIamUsersMoidParams) WithContext(ctx context.Context) *DeleteIamUsersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete iam users moid params
func (o *DeleteIamUsersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete iam users moid params
func (o *DeleteIamUsersMoidParams) WithHTTPClient(client *http.Client) *DeleteIamUsersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete iam users moid params
func (o *DeleteIamUsersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete iam users moid params
func (o *DeleteIamUsersMoidParams) WithMoid(moid string) *DeleteIamUsersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete iam users moid params
func (o *DeleteIamUsersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIamUsersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param moid
	if err := r.SetPathParam("moid", o.Moid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}