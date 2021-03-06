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
)

// NewDeleteIamAccountsMoidParams creates a new DeleteIamAccountsMoidParams object
// with the default values initialized.
func NewDeleteIamAccountsMoidParams() *DeleteIamAccountsMoidParams {
	var ()
	return &DeleteIamAccountsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIamAccountsMoidParamsWithTimeout creates a new DeleteIamAccountsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIamAccountsMoidParamsWithTimeout(timeout time.Duration) *DeleteIamAccountsMoidParams {
	var ()
	return &DeleteIamAccountsMoidParams{

		timeout: timeout,
	}
}

// NewDeleteIamAccountsMoidParamsWithContext creates a new DeleteIamAccountsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIamAccountsMoidParamsWithContext(ctx context.Context) *DeleteIamAccountsMoidParams {
	var ()
	return &DeleteIamAccountsMoidParams{

		Context: ctx,
	}
}

// NewDeleteIamAccountsMoidParamsWithHTTPClient creates a new DeleteIamAccountsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIamAccountsMoidParamsWithHTTPClient(client *http.Client) *DeleteIamAccountsMoidParams {
	var ()
	return &DeleteIamAccountsMoidParams{
		HTTPClient: client,
	}
}

/*DeleteIamAccountsMoidParams contains all the parameters to send to the API endpoint
for the delete iam accounts moid operation typically these are written to a http.Request
*/
type DeleteIamAccountsMoidParams struct {

	/*Moid
	  The moid of the iamAccount instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete iam accounts moid params
func (o *DeleteIamAccountsMoidParams) WithTimeout(timeout time.Duration) *DeleteIamAccountsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete iam accounts moid params
func (o *DeleteIamAccountsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete iam accounts moid params
func (o *DeleteIamAccountsMoidParams) WithContext(ctx context.Context) *DeleteIamAccountsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete iam accounts moid params
func (o *DeleteIamAccountsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete iam accounts moid params
func (o *DeleteIamAccountsMoidParams) WithHTTPClient(client *http.Client) *DeleteIamAccountsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete iam accounts moid params
func (o *DeleteIamAccountsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete iam accounts moid params
func (o *DeleteIamAccountsMoidParams) WithMoid(moid string) *DeleteIamAccountsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete iam accounts moid params
func (o *DeleteIamAccountsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIamAccountsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
