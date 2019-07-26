// Code generated by go-swagger; DO NOT EDIT.

package management_entity

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

// NewPostManagementEntitiesMoidParams creates a new PostManagementEntitiesMoidParams object
// with the default values initialized.
func NewPostManagementEntitiesMoidParams() *PostManagementEntitiesMoidParams {
	var ()
	return &PostManagementEntitiesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostManagementEntitiesMoidParamsWithTimeout creates a new PostManagementEntitiesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostManagementEntitiesMoidParamsWithTimeout(timeout time.Duration) *PostManagementEntitiesMoidParams {
	var ()
	return &PostManagementEntitiesMoidParams{

		timeout: timeout,
	}
}

// NewPostManagementEntitiesMoidParamsWithContext creates a new PostManagementEntitiesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostManagementEntitiesMoidParamsWithContext(ctx context.Context) *PostManagementEntitiesMoidParams {
	var ()
	return &PostManagementEntitiesMoidParams{

		Context: ctx,
	}
}

// NewPostManagementEntitiesMoidParamsWithHTTPClient creates a new PostManagementEntitiesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostManagementEntitiesMoidParamsWithHTTPClient(client *http.Client) *PostManagementEntitiesMoidParams {
	var ()
	return &PostManagementEntitiesMoidParams{
		HTTPClient: client,
	}
}

/*PostManagementEntitiesMoidParams contains all the parameters to send to the API endpoint
for the post management entities moid operation typically these are written to a http.Request
*/
type PostManagementEntitiesMoidParams struct {

	/*Body
	  managementEntity to update

	*/
	Body *models.ManagementEntity
	/*Moid
	  The moid of the managementEntity instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) WithTimeout(timeout time.Duration) *PostManagementEntitiesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) WithContext(ctx context.Context) *PostManagementEntitiesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) WithHTTPClient(client *http.Client) *PostManagementEntitiesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) WithBody(body *models.ManagementEntity) *PostManagementEntitiesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) SetBody(body *models.ManagementEntity) {
	o.Body = body
}

// WithMoid adds the moid to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) WithMoid(moid string) *PostManagementEntitiesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post management entities moid params
func (o *PostManagementEntitiesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostManagementEntitiesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
