// Code generated by go-swagger; DO NOT EDIT.

package iam_ldap_policy

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

// NewGetIamLdapPoliciesMoidParams creates a new GetIamLdapPoliciesMoidParams object
// with the default values initialized.
func NewGetIamLdapPoliciesMoidParams() *GetIamLdapPoliciesMoidParams {
	var ()
	return &GetIamLdapPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIamLdapPoliciesMoidParamsWithTimeout creates a new GetIamLdapPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIamLdapPoliciesMoidParamsWithTimeout(timeout time.Duration) *GetIamLdapPoliciesMoidParams {
	var ()
	return &GetIamLdapPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewGetIamLdapPoliciesMoidParamsWithContext creates a new GetIamLdapPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIamLdapPoliciesMoidParamsWithContext(ctx context.Context) *GetIamLdapPoliciesMoidParams {
	var ()
	return &GetIamLdapPoliciesMoidParams{

		Context: ctx,
	}
}

// NewGetIamLdapPoliciesMoidParamsWithHTTPClient creates a new GetIamLdapPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIamLdapPoliciesMoidParamsWithHTTPClient(client *http.Client) *GetIamLdapPoliciesMoidParams {
	var ()
	return &GetIamLdapPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*GetIamLdapPoliciesMoidParams contains all the parameters to send to the API endpoint
for the get iam ldap policies moid operation typically these are written to a http.Request
*/
type GetIamLdapPoliciesMoidParams struct {

	/*Moid
	  The moid of the iamLdapPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iam ldap policies moid params
func (o *GetIamLdapPoliciesMoidParams) WithTimeout(timeout time.Duration) *GetIamLdapPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iam ldap policies moid params
func (o *GetIamLdapPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iam ldap policies moid params
func (o *GetIamLdapPoliciesMoidParams) WithContext(ctx context.Context) *GetIamLdapPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iam ldap policies moid params
func (o *GetIamLdapPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iam ldap policies moid params
func (o *GetIamLdapPoliciesMoidParams) WithHTTPClient(client *http.Client) *GetIamLdapPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iam ldap policies moid params
func (o *GetIamLdapPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get iam ldap policies moid params
func (o *GetIamLdapPoliciesMoidParams) WithMoid(moid string) *GetIamLdapPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get iam ldap policies moid params
func (o *GetIamLdapPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetIamLdapPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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