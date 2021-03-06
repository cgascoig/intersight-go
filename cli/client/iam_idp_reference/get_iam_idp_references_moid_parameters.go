// Code generated by go-swagger; DO NOT EDIT.

package iam_idp_reference

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

// NewGetIamIdpReferencesMoidParams creates a new GetIamIdpReferencesMoidParams object
// with the default values initialized.
func NewGetIamIdpReferencesMoidParams() *GetIamIdpReferencesMoidParams {
	var ()
	return &GetIamIdpReferencesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIamIdpReferencesMoidParamsWithTimeout creates a new GetIamIdpReferencesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIamIdpReferencesMoidParamsWithTimeout(timeout time.Duration) *GetIamIdpReferencesMoidParams {
	var ()
	return &GetIamIdpReferencesMoidParams{

		timeout: timeout,
	}
}

// NewGetIamIdpReferencesMoidParamsWithContext creates a new GetIamIdpReferencesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIamIdpReferencesMoidParamsWithContext(ctx context.Context) *GetIamIdpReferencesMoidParams {
	var ()
	return &GetIamIdpReferencesMoidParams{

		Context: ctx,
	}
}

// NewGetIamIdpReferencesMoidParamsWithHTTPClient creates a new GetIamIdpReferencesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIamIdpReferencesMoidParamsWithHTTPClient(client *http.Client) *GetIamIdpReferencesMoidParams {
	var ()
	return &GetIamIdpReferencesMoidParams{
		HTTPClient: client,
	}
}

/*GetIamIdpReferencesMoidParams contains all the parameters to send to the API endpoint
for the get iam idp references moid operation typically these are written to a http.Request
*/
type GetIamIdpReferencesMoidParams struct {

	/*Moid
	  The moid of the iamIdpReference instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iam idp references moid params
func (o *GetIamIdpReferencesMoidParams) WithTimeout(timeout time.Duration) *GetIamIdpReferencesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iam idp references moid params
func (o *GetIamIdpReferencesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iam idp references moid params
func (o *GetIamIdpReferencesMoidParams) WithContext(ctx context.Context) *GetIamIdpReferencesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iam idp references moid params
func (o *GetIamIdpReferencesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iam idp references moid params
func (o *GetIamIdpReferencesMoidParams) WithHTTPClient(client *http.Client) *GetIamIdpReferencesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iam idp references moid params
func (o *GetIamIdpReferencesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get iam idp references moid params
func (o *GetIamIdpReferencesMoidParams) WithMoid(moid string) *GetIamIdpReferencesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get iam idp references moid params
func (o *GetIamIdpReferencesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetIamIdpReferencesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
