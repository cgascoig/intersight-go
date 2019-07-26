// Code generated by go-swagger; DO NOT EDIT.

package storage_sas_expander

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

// NewPostStorageSasExpandersMoidParams creates a new PostStorageSasExpandersMoidParams object
// with the default values initialized.
func NewPostStorageSasExpandersMoidParams() *PostStorageSasExpandersMoidParams {
	var ()
	return &PostStorageSasExpandersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageSasExpandersMoidParamsWithTimeout creates a new PostStorageSasExpandersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStorageSasExpandersMoidParamsWithTimeout(timeout time.Duration) *PostStorageSasExpandersMoidParams {
	var ()
	return &PostStorageSasExpandersMoidParams{

		timeout: timeout,
	}
}

// NewPostStorageSasExpandersMoidParamsWithContext creates a new PostStorageSasExpandersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostStorageSasExpandersMoidParamsWithContext(ctx context.Context) *PostStorageSasExpandersMoidParams {
	var ()
	return &PostStorageSasExpandersMoidParams{

		Context: ctx,
	}
}

// NewPostStorageSasExpandersMoidParamsWithHTTPClient creates a new PostStorageSasExpandersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostStorageSasExpandersMoidParamsWithHTTPClient(client *http.Client) *PostStorageSasExpandersMoidParams {
	var ()
	return &PostStorageSasExpandersMoidParams{
		HTTPClient: client,
	}
}

/*PostStorageSasExpandersMoidParams contains all the parameters to send to the API endpoint
for the post storage sas expanders moid operation typically these are written to a http.Request
*/
type PostStorageSasExpandersMoidParams struct {

	/*Body
	  storageSasExpander to update

	*/
	Body *models.StorageSasExpander
	/*Moid
	  The moid of the storageSasExpander instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) WithTimeout(timeout time.Duration) *PostStorageSasExpandersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) WithContext(ctx context.Context) *PostStorageSasExpandersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) WithHTTPClient(client *http.Client) *PostStorageSasExpandersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) WithBody(body *models.StorageSasExpander) *PostStorageSasExpandersMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) SetBody(body *models.StorageSasExpander) {
	o.Body = body
}

// WithMoid adds the moid to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) WithMoid(moid string) *PostStorageSasExpandersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post storage sas expanders moid params
func (o *PostStorageSasExpandersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageSasExpandersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
