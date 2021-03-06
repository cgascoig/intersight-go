// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_ext_iscsi_storage_policy

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

// NewDeleteHyperflexExtIscsiStoragePoliciesMoidParams creates a new DeleteHyperflexExtIscsiStoragePoliciesMoidParams object
// with the default values initialized.
func NewDeleteHyperflexExtIscsiStoragePoliciesMoidParams() *DeleteHyperflexExtIscsiStoragePoliciesMoidParams {
	var ()
	return &DeleteHyperflexExtIscsiStoragePoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHyperflexExtIscsiStoragePoliciesMoidParamsWithTimeout creates a new DeleteHyperflexExtIscsiStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteHyperflexExtIscsiStoragePoliciesMoidParamsWithTimeout(timeout time.Duration) *DeleteHyperflexExtIscsiStoragePoliciesMoidParams {
	var ()
	return &DeleteHyperflexExtIscsiStoragePoliciesMoidParams{

		timeout: timeout,
	}
}

// NewDeleteHyperflexExtIscsiStoragePoliciesMoidParamsWithContext creates a new DeleteHyperflexExtIscsiStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteHyperflexExtIscsiStoragePoliciesMoidParamsWithContext(ctx context.Context) *DeleteHyperflexExtIscsiStoragePoliciesMoidParams {
	var ()
	return &DeleteHyperflexExtIscsiStoragePoliciesMoidParams{

		Context: ctx,
	}
}

// NewDeleteHyperflexExtIscsiStoragePoliciesMoidParamsWithHTTPClient creates a new DeleteHyperflexExtIscsiStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteHyperflexExtIscsiStoragePoliciesMoidParamsWithHTTPClient(client *http.Client) *DeleteHyperflexExtIscsiStoragePoliciesMoidParams {
	var ()
	return &DeleteHyperflexExtIscsiStoragePoliciesMoidParams{
		HTTPClient: client,
	}
}

/*DeleteHyperflexExtIscsiStoragePoliciesMoidParams contains all the parameters to send to the API endpoint
for the delete hyperflex ext iscsi storage policies moid operation typically these are written to a http.Request
*/
type DeleteHyperflexExtIscsiStoragePoliciesMoidParams struct {

	/*Moid
	  The moid of the hyperflexExtIscsiStoragePolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete hyperflex ext iscsi storage policies moid params
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidParams) WithTimeout(timeout time.Duration) *DeleteHyperflexExtIscsiStoragePoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete hyperflex ext iscsi storage policies moid params
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete hyperflex ext iscsi storage policies moid params
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidParams) WithContext(ctx context.Context) *DeleteHyperflexExtIscsiStoragePoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete hyperflex ext iscsi storage policies moid params
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete hyperflex ext iscsi storage policies moid params
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidParams) WithHTTPClient(client *http.Client) *DeleteHyperflexExtIscsiStoragePoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete hyperflex ext iscsi storage policies moid params
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete hyperflex ext iscsi storage policies moid params
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidParams) WithMoid(moid string) *DeleteHyperflexExtIscsiStoragePoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete hyperflex ext iscsi storage policies moid params
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
