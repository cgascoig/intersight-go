// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_ext_fc_storage_policy

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

// NewDeleteHyperflexExtFcStoragePoliciesMoidParams creates a new DeleteHyperflexExtFcStoragePoliciesMoidParams object
// with the default values initialized.
func NewDeleteHyperflexExtFcStoragePoliciesMoidParams() *DeleteHyperflexExtFcStoragePoliciesMoidParams {
	var ()
	return &DeleteHyperflexExtFcStoragePoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHyperflexExtFcStoragePoliciesMoidParamsWithTimeout creates a new DeleteHyperflexExtFcStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteHyperflexExtFcStoragePoliciesMoidParamsWithTimeout(timeout time.Duration) *DeleteHyperflexExtFcStoragePoliciesMoidParams {
	var ()
	return &DeleteHyperflexExtFcStoragePoliciesMoidParams{

		timeout: timeout,
	}
}

// NewDeleteHyperflexExtFcStoragePoliciesMoidParamsWithContext creates a new DeleteHyperflexExtFcStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteHyperflexExtFcStoragePoliciesMoidParamsWithContext(ctx context.Context) *DeleteHyperflexExtFcStoragePoliciesMoidParams {
	var ()
	return &DeleteHyperflexExtFcStoragePoliciesMoidParams{

		Context: ctx,
	}
}

// NewDeleteHyperflexExtFcStoragePoliciesMoidParamsWithHTTPClient creates a new DeleteHyperflexExtFcStoragePoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteHyperflexExtFcStoragePoliciesMoidParamsWithHTTPClient(client *http.Client) *DeleteHyperflexExtFcStoragePoliciesMoidParams {
	var ()
	return &DeleteHyperflexExtFcStoragePoliciesMoidParams{
		HTTPClient: client,
	}
}

/*DeleteHyperflexExtFcStoragePoliciesMoidParams contains all the parameters to send to the API endpoint
for the delete hyperflex ext fc storage policies moid operation typically these are written to a http.Request
*/
type DeleteHyperflexExtFcStoragePoliciesMoidParams struct {

	/*Moid
	  The moid of the hyperflexExtFcStoragePolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete hyperflex ext fc storage policies moid params
func (o *DeleteHyperflexExtFcStoragePoliciesMoidParams) WithTimeout(timeout time.Duration) *DeleteHyperflexExtFcStoragePoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete hyperflex ext fc storage policies moid params
func (o *DeleteHyperflexExtFcStoragePoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete hyperflex ext fc storage policies moid params
func (o *DeleteHyperflexExtFcStoragePoliciesMoidParams) WithContext(ctx context.Context) *DeleteHyperflexExtFcStoragePoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete hyperflex ext fc storage policies moid params
func (o *DeleteHyperflexExtFcStoragePoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete hyperflex ext fc storage policies moid params
func (o *DeleteHyperflexExtFcStoragePoliciesMoidParams) WithHTTPClient(client *http.Client) *DeleteHyperflexExtFcStoragePoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete hyperflex ext fc storage policies moid params
func (o *DeleteHyperflexExtFcStoragePoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the delete hyperflex ext fc storage policies moid params
func (o *DeleteHyperflexExtFcStoragePoliciesMoidParams) WithMoid(moid string) *DeleteHyperflexExtFcStoragePoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the delete hyperflex ext fc storage policies moid params
func (o *DeleteHyperflexExtFcStoragePoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHyperflexExtFcStoragePoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
