// Code generated by go-swagger; DO NOT EDIT.

package storage_storage_policy

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

// NewPatchStorageStoragePoliciesMoidParams creates a new PatchStorageStoragePoliciesMoidParams object
// with the default values initialized.
func NewPatchStorageStoragePoliciesMoidParams() *PatchStorageStoragePoliciesMoidParams {
	var ()
	return &PatchStorageStoragePoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchStorageStoragePoliciesMoidParamsWithTimeout creates a new PatchStorageStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchStorageStoragePoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchStorageStoragePoliciesMoidParams {
	var ()
	return &PatchStorageStoragePoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchStorageStoragePoliciesMoidParamsWithContext creates a new PatchStorageStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchStorageStoragePoliciesMoidParamsWithContext(ctx context.Context) *PatchStorageStoragePoliciesMoidParams {
	var ()
	return &PatchStorageStoragePoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchStorageStoragePoliciesMoidParamsWithHTTPClient creates a new PatchStorageStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchStorageStoragePoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchStorageStoragePoliciesMoidParams {
	var ()
	return &PatchStorageStoragePoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchStorageStoragePoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch storage storage policies moid operation typically these are written to a http.Request
*/
type PatchStorageStoragePoliciesMoidParams struct {

	/*Body
	  storageStoragePolicy to update

	*/
	Body *models.StorageStoragePolicy
	/*Moid
	  The moid of the storageStoragePolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchStorageStoragePoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) WithContext(ctx context.Context) *PatchStorageStoragePoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchStorageStoragePoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) WithBody(body *models.StorageStoragePolicy) *PatchStorageStoragePoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) SetBody(body *models.StorageStoragePolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) WithMoid(moid string) *PatchStorageStoragePoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch storage storage policies moid params
func (o *PatchStorageStoragePoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchStorageStoragePoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
