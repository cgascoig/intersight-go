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

// NewPostStorageStoragePoliciesMoidParams creates a new PostStorageStoragePoliciesMoidParams object
// with the default values initialized.
func NewPostStorageStoragePoliciesMoidParams() *PostStorageStoragePoliciesMoidParams {
	var ()
	return &PostStorageStoragePoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageStoragePoliciesMoidParamsWithTimeout creates a new PostStorageStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStorageStoragePoliciesMoidParamsWithTimeout(timeout time.Duration) *PostStorageStoragePoliciesMoidParams {
	var ()
	return &PostStorageStoragePoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPostStorageStoragePoliciesMoidParamsWithContext creates a new PostStorageStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostStorageStoragePoliciesMoidParamsWithContext(ctx context.Context) *PostStorageStoragePoliciesMoidParams {
	var ()
	return &PostStorageStoragePoliciesMoidParams{

		Context: ctx,
	}
}

// NewPostStorageStoragePoliciesMoidParamsWithHTTPClient creates a new PostStorageStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostStorageStoragePoliciesMoidParamsWithHTTPClient(client *http.Client) *PostStorageStoragePoliciesMoidParams {
	var ()
	return &PostStorageStoragePoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PostStorageStoragePoliciesMoidParams contains all the parameters to send to the API endpoint
for the post storage storage policies moid operation typically these are written to a http.Request
*/
type PostStorageStoragePoliciesMoidParams struct {

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

// WithTimeout adds the timeout to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) WithTimeout(timeout time.Duration) *PostStorageStoragePoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) WithContext(ctx context.Context) *PostStorageStoragePoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) WithHTTPClient(client *http.Client) *PostStorageStoragePoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) WithBody(body *models.StorageStoragePolicy) *PostStorageStoragePoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) SetBody(body *models.StorageStoragePolicy) {
	o.Body = body
}

// WithMoid adds the moid to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) WithMoid(moid string) *PostStorageStoragePoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post storage storage policies moid params
func (o *PostStorageStoragePoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageStoragePoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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