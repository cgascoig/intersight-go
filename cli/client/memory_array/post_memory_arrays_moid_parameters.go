// Code generated by go-swagger; DO NOT EDIT.

package memory_array

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

// NewPostMemoryArraysMoidParams creates a new PostMemoryArraysMoidParams object
// with the default values initialized.
func NewPostMemoryArraysMoidParams() *PostMemoryArraysMoidParams {
	var ()
	return &PostMemoryArraysMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMemoryArraysMoidParamsWithTimeout creates a new PostMemoryArraysMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMemoryArraysMoidParamsWithTimeout(timeout time.Duration) *PostMemoryArraysMoidParams {
	var ()
	return &PostMemoryArraysMoidParams{

		timeout: timeout,
	}
}

// NewPostMemoryArraysMoidParamsWithContext creates a new PostMemoryArraysMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMemoryArraysMoidParamsWithContext(ctx context.Context) *PostMemoryArraysMoidParams {
	var ()
	return &PostMemoryArraysMoidParams{

		Context: ctx,
	}
}

// NewPostMemoryArraysMoidParamsWithHTTPClient creates a new PostMemoryArraysMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMemoryArraysMoidParamsWithHTTPClient(client *http.Client) *PostMemoryArraysMoidParams {
	var ()
	return &PostMemoryArraysMoidParams{
		HTTPClient: client,
	}
}

/*PostMemoryArraysMoidParams contains all the parameters to send to the API endpoint
for the post memory arrays moid operation typically these are written to a http.Request
*/
type PostMemoryArraysMoidParams struct {

	/*Body
	  memoryArray to update

	*/
	Body *models.MemoryArray
	/*Moid
	  The moid of the memoryArray instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) WithTimeout(timeout time.Duration) *PostMemoryArraysMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) WithContext(ctx context.Context) *PostMemoryArraysMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) WithHTTPClient(client *http.Client) *PostMemoryArraysMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) WithBody(body *models.MemoryArray) *PostMemoryArraysMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) SetBody(body *models.MemoryArray) {
	o.Body = body
}

// WithMoid adds the moid to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) WithMoid(moid string) *PostMemoryArraysMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post memory arrays moid params
func (o *PostMemoryArraysMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostMemoryArraysMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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