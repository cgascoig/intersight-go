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

// NewPatchManagementEntitiesMoidParams creates a new PatchManagementEntitiesMoidParams object
// with the default values initialized.
func NewPatchManagementEntitiesMoidParams() *PatchManagementEntitiesMoidParams {
	var ()
	return &PatchManagementEntitiesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchManagementEntitiesMoidParamsWithTimeout creates a new PatchManagementEntitiesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchManagementEntitiesMoidParamsWithTimeout(timeout time.Duration) *PatchManagementEntitiesMoidParams {
	var ()
	return &PatchManagementEntitiesMoidParams{

		timeout: timeout,
	}
}

// NewPatchManagementEntitiesMoidParamsWithContext creates a new PatchManagementEntitiesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchManagementEntitiesMoidParamsWithContext(ctx context.Context) *PatchManagementEntitiesMoidParams {
	var ()
	return &PatchManagementEntitiesMoidParams{

		Context: ctx,
	}
}

// NewPatchManagementEntitiesMoidParamsWithHTTPClient creates a new PatchManagementEntitiesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchManagementEntitiesMoidParamsWithHTTPClient(client *http.Client) *PatchManagementEntitiesMoidParams {
	var ()
	return &PatchManagementEntitiesMoidParams{
		HTTPClient: client,
	}
}

/*PatchManagementEntitiesMoidParams contains all the parameters to send to the API endpoint
for the patch management entities moid operation typically these are written to a http.Request
*/
type PatchManagementEntitiesMoidParams struct {

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

// WithTimeout adds the timeout to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) WithTimeout(timeout time.Duration) *PatchManagementEntitiesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) WithContext(ctx context.Context) *PatchManagementEntitiesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) WithHTTPClient(client *http.Client) *PatchManagementEntitiesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) WithBody(body *models.ManagementEntity) *PatchManagementEntitiesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) SetBody(body *models.ManagementEntity) {
	o.Body = body
}

// WithMoid adds the moid to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) WithMoid(moid string) *PatchManagementEntitiesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch management entities moid params
func (o *PatchManagementEntitiesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchManagementEntitiesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
