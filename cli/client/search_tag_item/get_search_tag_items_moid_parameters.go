// Code generated by go-swagger; DO NOT EDIT.

package search_tag_item

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

// NewGetSearchTagItemsMoidParams creates a new GetSearchTagItemsMoidParams object
// with the default values initialized.
func NewGetSearchTagItemsMoidParams() *GetSearchTagItemsMoidParams {
	var ()
	return &GetSearchTagItemsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSearchTagItemsMoidParamsWithTimeout creates a new GetSearchTagItemsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSearchTagItemsMoidParamsWithTimeout(timeout time.Duration) *GetSearchTagItemsMoidParams {
	var ()
	return &GetSearchTagItemsMoidParams{

		timeout: timeout,
	}
}

// NewGetSearchTagItemsMoidParamsWithContext creates a new GetSearchTagItemsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSearchTagItemsMoidParamsWithContext(ctx context.Context) *GetSearchTagItemsMoidParams {
	var ()
	return &GetSearchTagItemsMoidParams{

		Context: ctx,
	}
}

// NewGetSearchTagItemsMoidParamsWithHTTPClient creates a new GetSearchTagItemsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSearchTagItemsMoidParamsWithHTTPClient(client *http.Client) *GetSearchTagItemsMoidParams {
	var ()
	return &GetSearchTagItemsMoidParams{
		HTTPClient: client,
	}
}

/*GetSearchTagItemsMoidParams contains all the parameters to send to the API endpoint
for the get search tag items moid operation typically these are written to a http.Request
*/
type GetSearchTagItemsMoidParams struct {

	/*Moid
	  The moid of the searchTagItem instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get search tag items moid params
func (o *GetSearchTagItemsMoidParams) WithTimeout(timeout time.Duration) *GetSearchTagItemsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get search tag items moid params
func (o *GetSearchTagItemsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get search tag items moid params
func (o *GetSearchTagItemsMoidParams) WithContext(ctx context.Context) *GetSearchTagItemsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get search tag items moid params
func (o *GetSearchTagItemsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get search tag items moid params
func (o *GetSearchTagItemsMoidParams) WithHTTPClient(client *http.Client) *GetSearchTagItemsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get search tag items moid params
func (o *GetSearchTagItemsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get search tag items moid params
func (o *GetSearchTagItemsMoidParams) WithMoid(moid string) *GetSearchTagItemsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get search tag items moid params
func (o *GetSearchTagItemsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetSearchTagItemsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
