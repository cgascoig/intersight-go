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
)

// NewDeleteIamLdapGroupsMoidParams creates a new DeleteIamLdapGroupsMoidParams object
// with the default values initialized.
func NewDeleteIamLdapGroupsMoidParams() *DeleteIamLdapGroupsMoidParams {
	var ()
	return &DeleteIamLdapGroupsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIamLdapGroupsMoidParamsWithTimeout creates a new DeleteIamLdapGroupsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIamLdapGroupsMoidParamsWithTimeout(timeout time.Duration) *DeleteIamLdapGroupsMoidParams {
	var ()
	return &DeleteIamLdapGroupsMoidParams{

		timeout: timeout,
	}
}

// NewDeleteIamLdapGroupsMoidParamsWithContext creates a new DeleteIamLdapGroupsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIamLdapGroupsMoidParamsWithContext(ctx context.Context) *DeleteIamLdapGroupsMoidParams {
	var ()
	return &DeleteIamLdapGroupsMoidParams{

		Context: ctx,
	}
}

// NewDeleteIamLdapGroupsMoidParamsWithHTTPClient creates a new DeleteIamLdapGroupsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIamLdapGroupsMoidParamsWithHTTPClient(client *http.Client) *DeleteIamLdapGroupsMoidParams {
	var ()
	return &DeleteIamLdapGroupsMoidParams{
		HTTPClient: client,
	}
}

/*DeleteIamLdapGroupsMoidParams contains all the parameters to send to the API endpoint
for the delete iam ldap groups moid operation typically these are written to a http.Request
*/
type DeleteIamLdapGroupsMoidParams struct {

	/*Moid
	  The moid of the iamLdapGroup instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete iam ldap groups moid params
func (o *DeleteIamLdapGroupsMoidParams) WithTimeout(timeout time.Duration) *DeleteIamLdapGroupsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete iam ldap groups moid params
func (o *DeleteIamLdapGroupsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete iam ldap groups moid params
func (o *DeleteIamLdapGroupsMoidParams) WithContext(ctx context.Context) *DeleteIamLdapGroupsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete iam ldap groups moid params
func (o *DeleteIamLdapGroupsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete iam ldap groups moid params
func (o *DeleteIamLdapGroupsMoidParams) WithHTTPClient(client *http.Client) *DeleteIamLdapGroupsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete iam ldap groups moid params
func (o *DeleteIamLdapGroupsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete iam ldap groups moid params
func (o *DeleteIamLdapGroupsMoidParams) WithMoid(moid string) *DeleteIamLdapGroupsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete iam ldap groups moid params
func (o *DeleteIamLdapGroupsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIamLdapGroupsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
