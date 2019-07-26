// Code generated by go-swagger; DO NOT EDIT.

package externalsite_authorization

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

// NewPostExternalsiteAuthorizationsMoidParams creates a new PostExternalsiteAuthorizationsMoidParams object
// with the default values initialized.
func NewPostExternalsiteAuthorizationsMoidParams() *PostExternalsiteAuthorizationsMoidParams {
	var ()
	return &PostExternalsiteAuthorizationsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostExternalsiteAuthorizationsMoidParamsWithTimeout creates a new PostExternalsiteAuthorizationsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostExternalsiteAuthorizationsMoidParamsWithTimeout(timeout time.Duration) *PostExternalsiteAuthorizationsMoidParams {
	var ()
	return &PostExternalsiteAuthorizationsMoidParams{

		timeout: timeout,
	}
}

// NewPostExternalsiteAuthorizationsMoidParamsWithContext creates a new PostExternalsiteAuthorizationsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostExternalsiteAuthorizationsMoidParamsWithContext(ctx context.Context) *PostExternalsiteAuthorizationsMoidParams {
	var ()
	return &PostExternalsiteAuthorizationsMoidParams{

		Context: ctx,
	}
}

// NewPostExternalsiteAuthorizationsMoidParamsWithHTTPClient creates a new PostExternalsiteAuthorizationsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostExternalsiteAuthorizationsMoidParamsWithHTTPClient(client *http.Client) *PostExternalsiteAuthorizationsMoidParams {
	var ()
	return &PostExternalsiteAuthorizationsMoidParams{
		HTTPClient: client,
	}
}

/*PostExternalsiteAuthorizationsMoidParams contains all the parameters to send to the API endpoint
for the post externalsite authorizations moid operation typically these are written to a http.Request
*/
type PostExternalsiteAuthorizationsMoidParams struct {

	/*Body
	  externalsiteAuthorization to update

	*/
	Body *models.ExternalsiteAuthorization
	/*Moid
	  The moid of the externalsiteAuthorization instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) WithTimeout(timeout time.Duration) *PostExternalsiteAuthorizationsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) WithContext(ctx context.Context) *PostExternalsiteAuthorizationsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) WithHTTPClient(client *http.Client) *PostExternalsiteAuthorizationsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) WithBody(body *models.ExternalsiteAuthorization) *PostExternalsiteAuthorizationsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) SetBody(body *models.ExternalsiteAuthorization) {
	o.Body = body
}

// WithMoid adds the moid to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) WithMoid(moid string) *PostExternalsiteAuthorizationsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post externalsite authorizations moid params
func (o *PostExternalsiteAuthorizationsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostExternalsiteAuthorizationsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
