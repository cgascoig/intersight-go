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

// NewPatchIamUserGroupsMoidParams creates a new PatchIamUserGroupsMoidParams object
// with the default values initialized.
func NewPatchIamUserGroupsMoidParams() *PatchIamUserGroupsMoidParams {
	var ()
	return &PatchIamUserGroupsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchIamUserGroupsMoidParamsWithTimeout creates a new PatchIamUserGroupsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchIamUserGroupsMoidParamsWithTimeout(timeout time.Duration) *PatchIamUserGroupsMoidParams {
	var ()
	return &PatchIamUserGroupsMoidParams{

		timeout: timeout,
	}
}

// NewPatchIamUserGroupsMoidParamsWithContext creates a new PatchIamUserGroupsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchIamUserGroupsMoidParamsWithContext(ctx context.Context) *PatchIamUserGroupsMoidParams {
	var ()
	return &PatchIamUserGroupsMoidParams{

		Context: ctx,
	}
}

// NewPatchIamUserGroupsMoidParamsWithHTTPClient creates a new PatchIamUserGroupsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchIamUserGroupsMoidParamsWithHTTPClient(client *http.Client) *PatchIamUserGroupsMoidParams {
	var ()
	return &PatchIamUserGroupsMoidParams{
		HTTPClient: client,
	}
}

/*PatchIamUserGroupsMoidParams contains all the parameters to send to the API endpoint
for the patch iam user groups moid operation typically these are written to a http.Request
*/
type PatchIamUserGroupsMoidParams struct {

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

// WithTimeout adds the timeout to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) WithTimeout(timeout time.Duration) *PatchIamUserGroupsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) WithContext(ctx context.Context) *PatchIamUserGroupsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) WithHTTPClient(client *http.Client) *PatchIamUserGroupsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) WithBody(body *models.IamUserGroup) *PatchIamUserGroupsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) SetBody(body *models.IamUserGroup) {
	o.Body = body
}

// WithMoid adds the moid to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) WithMoid(moid string) *PatchIamUserGroupsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch iam user groups moid params
func (o *PatchIamUserGroupsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchIamUserGroupsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
