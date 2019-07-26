// Code generated by go-swagger; DO NOT EDIT.

package tam_advisory_instance

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

// NewPostTamAdvisoryInstancesMoidParams creates a new PostTamAdvisoryInstancesMoidParams object
// with the default values initialized.
func NewPostTamAdvisoryInstancesMoidParams() *PostTamAdvisoryInstancesMoidParams {
	var ()
	return &PostTamAdvisoryInstancesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTamAdvisoryInstancesMoidParamsWithTimeout creates a new PostTamAdvisoryInstancesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTamAdvisoryInstancesMoidParamsWithTimeout(timeout time.Duration) *PostTamAdvisoryInstancesMoidParams {
	var ()
	return &PostTamAdvisoryInstancesMoidParams{

		timeout: timeout,
	}
}

// NewPostTamAdvisoryInstancesMoidParamsWithContext creates a new PostTamAdvisoryInstancesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTamAdvisoryInstancesMoidParamsWithContext(ctx context.Context) *PostTamAdvisoryInstancesMoidParams {
	var ()
	return &PostTamAdvisoryInstancesMoidParams{

		Context: ctx,
	}
}

// NewPostTamAdvisoryInstancesMoidParamsWithHTTPClient creates a new PostTamAdvisoryInstancesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTamAdvisoryInstancesMoidParamsWithHTTPClient(client *http.Client) *PostTamAdvisoryInstancesMoidParams {
	var ()
	return &PostTamAdvisoryInstancesMoidParams{
		HTTPClient: client,
	}
}

/*PostTamAdvisoryInstancesMoidParams contains all the parameters to send to the API endpoint
for the post tam advisory instances moid operation typically these are written to a http.Request
*/
type PostTamAdvisoryInstancesMoidParams struct {

	/*Body
	  tamAdvisoryInstance to update

	*/
	Body *models.TamAdvisoryInstance
	/*Moid
	  The moid of the tamAdvisoryInstance instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) WithTimeout(timeout time.Duration) *PostTamAdvisoryInstancesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) WithContext(ctx context.Context) *PostTamAdvisoryInstancesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) WithHTTPClient(client *http.Client) *PostTamAdvisoryInstancesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) WithBody(body *models.TamAdvisoryInstance) *PostTamAdvisoryInstancesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) SetBody(body *models.TamAdvisoryInstance) {
	o.Body = body
}

// WithMoid adds the moid to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) WithMoid(moid string) *PostTamAdvisoryInstancesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post tam advisory instances moid params
func (o *PostTamAdvisoryInstancesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostTamAdvisoryInstancesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
