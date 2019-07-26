// Code generated by go-swagger; DO NOT EDIT.

package iam_ldap_provider

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

// NewDeleteIamLdapProvidersMoidParams creates a new DeleteIamLdapProvidersMoidParams object
// with the default values initialized.
func NewDeleteIamLdapProvidersMoidParams() *DeleteIamLdapProvidersMoidParams {
	var ()
	return &DeleteIamLdapProvidersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIamLdapProvidersMoidParamsWithTimeout creates a new DeleteIamLdapProvidersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIamLdapProvidersMoidParamsWithTimeout(timeout time.Duration) *DeleteIamLdapProvidersMoidParams {
	var ()
	return &DeleteIamLdapProvidersMoidParams{

		timeout: timeout,
	}
}

// NewDeleteIamLdapProvidersMoidParamsWithContext creates a new DeleteIamLdapProvidersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIamLdapProvidersMoidParamsWithContext(ctx context.Context) *DeleteIamLdapProvidersMoidParams {
	var ()
	return &DeleteIamLdapProvidersMoidParams{

		Context: ctx,
	}
}

// NewDeleteIamLdapProvidersMoidParamsWithHTTPClient creates a new DeleteIamLdapProvidersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIamLdapProvidersMoidParamsWithHTTPClient(client *http.Client) *DeleteIamLdapProvidersMoidParams {
	var ()
	return &DeleteIamLdapProvidersMoidParams{
		HTTPClient: client,
	}
}

/*DeleteIamLdapProvidersMoidParams contains all the parameters to send to the API endpoint
for the delete iam ldap providers moid operation typically these are written to a http.Request
*/
type DeleteIamLdapProvidersMoidParams struct {

	/*Moid
	  The moid of the iamLdapProvider instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete iam ldap providers moid params
func (o *DeleteIamLdapProvidersMoidParams) WithTimeout(timeout time.Duration) *DeleteIamLdapProvidersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete iam ldap providers moid params
func (o *DeleteIamLdapProvidersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete iam ldap providers moid params
func (o *DeleteIamLdapProvidersMoidParams) WithContext(ctx context.Context) *DeleteIamLdapProvidersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete iam ldap providers moid params
func (o *DeleteIamLdapProvidersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete iam ldap providers moid params
func (o *DeleteIamLdapProvidersMoidParams) WithHTTPClient(client *http.Client) *DeleteIamLdapProvidersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete iam ldap providers moid params
func (o *DeleteIamLdapProvidersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete iam ldap providers moid params
func (o *DeleteIamLdapProvidersMoidParams) WithMoid(moid string) *DeleteIamLdapProvidersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete iam ldap providers moid params
func (o *DeleteIamLdapProvidersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIamLdapProvidersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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