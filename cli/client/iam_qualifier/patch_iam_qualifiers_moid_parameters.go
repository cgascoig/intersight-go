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

// NewPatchIamQualifiersMoidParams creates a new PatchIamQualifiersMoidParams object
// with the default values initialized.
func NewPatchIamQualifiersMoidParams() *PatchIamQualifiersMoidParams {
	var ()
	return &PatchIamQualifiersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchIamQualifiersMoidParamsWithTimeout creates a new PatchIamQualifiersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchIamQualifiersMoidParamsWithTimeout(timeout time.Duration) *PatchIamQualifiersMoidParams {
	var ()
	return &PatchIamQualifiersMoidParams{

		timeout: timeout,
	}
}

// NewPatchIamQualifiersMoidParamsWithContext creates a new PatchIamQualifiersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchIamQualifiersMoidParamsWithContext(ctx context.Context) *PatchIamQualifiersMoidParams {
	var ()
	return &PatchIamQualifiersMoidParams{

		Context: ctx,
	}
}

// NewPatchIamQualifiersMoidParamsWithHTTPClient creates a new PatchIamQualifiersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchIamQualifiersMoidParamsWithHTTPClient(client *http.Client) *PatchIamQualifiersMoidParams {
	var ()
	return &PatchIamQualifiersMoidParams{
		HTTPClient: client,
	}
}

/*PatchIamQualifiersMoidParams contains all the parameters to send to the API endpoint
for the patch iam qualifiers moid operation typically these are written to a http.Request
*/
type PatchIamQualifiersMoidParams struct {

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

// WithTimeout adds the timeout to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) WithTimeout(timeout time.Duration) *PatchIamQualifiersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) WithContext(ctx context.Context) *PatchIamQualifiersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) WithHTTPClient(client *http.Client) *PatchIamQualifiersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) WithBody(body *models.IamQualifier) *PatchIamQualifiersMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) SetBody(body *models.IamQualifier) {
	o.Body = body
}

// WithMoid adds the moid to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) WithMoid(moid string) *PatchIamQualifiersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch iam qualifiers moid params
func (o *PatchIamQualifiersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchIamQualifiersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
