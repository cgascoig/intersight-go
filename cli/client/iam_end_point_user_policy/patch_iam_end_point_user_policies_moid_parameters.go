// Code generated by go-swagger; DO NOT EDIT.

package iam_end_point_user_policy

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

// NewPatchIamEndPointUserPoliciesMoidParams creates a new PatchIamEndPointUserPoliciesMoidParams object
// with the default values initialized.
func NewPatchIamEndPointUserPoliciesMoidParams() *PatchIamEndPointUserPoliciesMoidParams {
	var ()
	return &PatchIamEndPointUserPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchIamEndPointUserPoliciesMoidParamsWithTimeout creates a new PatchIamEndPointUserPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchIamEndPointUserPoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchIamEndPointUserPoliciesMoidParams {
	var ()
	return &PatchIamEndPointUserPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchIamEndPointUserPoliciesMoidParamsWithContext creates a new PatchIamEndPointUserPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchIamEndPointUserPoliciesMoidParamsWithContext(ctx context.Context) *PatchIamEndPointUserPoliciesMoidParams {
	var ()
	return &PatchIamEndPointUserPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchIamEndPointUserPoliciesMoidParamsWithHTTPClient creates a new PatchIamEndPointUserPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchIamEndPointUserPoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchIamEndPointUserPoliciesMoidParams {
	var ()
	return &PatchIamEndPointUserPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchIamEndPointUserPoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch iam end point user policies moid operation typically these are written to a http.Request
*/
type PatchIamEndPointUserPoliciesMoidParams struct {

	/*Body
	  iamEndPointUserPolicy to update

	*/
	Body *models.IamEndPointUserPolicy
	/*Moid
	  The moid of the iamEndPointUserPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchIamEndPointUserPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) WithContext(ctx context.Context) *PatchIamEndPointUserPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchIamEndPointUserPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) WithBody(body *models.IamEndPointUserPolicy) *PatchIamEndPointUserPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) SetBody(body *models.IamEndPointUserPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) WithMoid(moid string) *PatchIamEndPointUserPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch iam end point user policies moid params
func (o *PatchIamEndPointUserPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchIamEndPointUserPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
