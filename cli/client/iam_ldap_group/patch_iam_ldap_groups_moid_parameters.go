// Code generated by go-swagger; DO NOT EDIT.

package iam_ldap_group

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

// NewPatchIamLdapGroupsMoidParams creates a new PatchIamLdapGroupsMoidParams object
// with the default values initialized.
func NewPatchIamLdapGroupsMoidParams() *PatchIamLdapGroupsMoidParams {
	var ()
	return &PatchIamLdapGroupsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchIamLdapGroupsMoidParamsWithTimeout creates a new PatchIamLdapGroupsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchIamLdapGroupsMoidParamsWithTimeout(timeout time.Duration) *PatchIamLdapGroupsMoidParams {
	var ()
	return &PatchIamLdapGroupsMoidParams{

		timeout: timeout,
	}
}

// NewPatchIamLdapGroupsMoidParamsWithContext creates a new PatchIamLdapGroupsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchIamLdapGroupsMoidParamsWithContext(ctx context.Context) *PatchIamLdapGroupsMoidParams {
	var ()
	return &PatchIamLdapGroupsMoidParams{

		Context: ctx,
	}
}

// NewPatchIamLdapGroupsMoidParamsWithHTTPClient creates a new PatchIamLdapGroupsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchIamLdapGroupsMoidParamsWithHTTPClient(client *http.Client) *PatchIamLdapGroupsMoidParams {
	var ()
	return &PatchIamLdapGroupsMoidParams{
		HTTPClient: client,
	}
}

/*PatchIamLdapGroupsMoidParams contains all the parameters to send to the API endpoint
for the patch iam ldap groups moid operation typically these are written to a http.Request
*/
type PatchIamLdapGroupsMoidParams struct {

	/*Body
	  iamLdapGroup to update

	*/
	Body *models.IamLdapGroup
	/*Moid
	  The moid of the iamLdapGroup instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) WithTimeout(timeout time.Duration) *PatchIamLdapGroupsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) WithContext(ctx context.Context) *PatchIamLdapGroupsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) WithHTTPClient(client *http.Client) *PatchIamLdapGroupsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) WithBody(body *models.IamLdapGroup) *PatchIamLdapGroupsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) SetBody(body *models.IamLdapGroup) {
	o.Body = body
}

// WithMoid adds the moid to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) WithMoid(moid string) *PatchIamLdapGroupsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch iam ldap groups moid params
func (o *PatchIamLdapGroupsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchIamLdapGroupsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
