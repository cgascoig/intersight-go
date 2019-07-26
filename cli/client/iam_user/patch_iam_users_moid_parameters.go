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

// NewPatchIamUsersMoidParams creates a new PatchIamUsersMoidParams object
// with the default values initialized.
func NewPatchIamUsersMoidParams() *PatchIamUsersMoidParams {
	var ()
	return &PatchIamUsersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchIamUsersMoidParamsWithTimeout creates a new PatchIamUsersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchIamUsersMoidParamsWithTimeout(timeout time.Duration) *PatchIamUsersMoidParams {
	var ()
	return &PatchIamUsersMoidParams{

		timeout: timeout,
	}
}

// NewPatchIamUsersMoidParamsWithContext creates a new PatchIamUsersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchIamUsersMoidParamsWithContext(ctx context.Context) *PatchIamUsersMoidParams {
	var ()
	return &PatchIamUsersMoidParams{

		Context: ctx,
	}
}

// NewPatchIamUsersMoidParamsWithHTTPClient creates a new PatchIamUsersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchIamUsersMoidParamsWithHTTPClient(client *http.Client) *PatchIamUsersMoidParams {
	var ()
	return &PatchIamUsersMoidParams{
		HTTPClient: client,
	}
}

/*PatchIamUsersMoidParams contains all the parameters to send to the API endpoint
for the patch iam users moid operation typically these are written to a http.Request
*/
type PatchIamUsersMoidParams struct {

	/*Body
	  iamUser to update

	*/
	Body *models.IamUser
	/*Moid
	  The moid of the iamUser instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch iam users moid params
func (o *PatchIamUsersMoidParams) WithTimeout(timeout time.Duration) *PatchIamUsersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch iam users moid params
func (o *PatchIamUsersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch iam users moid params
func (o *PatchIamUsersMoidParams) WithContext(ctx context.Context) *PatchIamUsersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch iam users moid params
func (o *PatchIamUsersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch iam users moid params
func (o *PatchIamUsersMoidParams) WithHTTPClient(client *http.Client) *PatchIamUsersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch iam users moid params
func (o *PatchIamUsersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch iam users moid params
func (o *PatchIamUsersMoidParams) WithBody(body *models.IamUser) *PatchIamUsersMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch iam users moid params
func (o *PatchIamUsersMoidParams) SetBody(body *models.IamUser) {
	o.Body = body
}

// WithMoid adds the moid to the patch iam users moid params
func (o *PatchIamUsersMoidParams) WithMoid(moid string) *PatchIamUsersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch iam users moid params
func (o *PatchIamUsersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchIamUsersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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