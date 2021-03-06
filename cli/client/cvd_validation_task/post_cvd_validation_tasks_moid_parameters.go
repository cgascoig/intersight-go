// Code generated by go-swagger; DO NOT EDIT.

package cvd_validation_task

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

// NewPostCvdValidationTasksMoidParams creates a new PostCvdValidationTasksMoidParams object
// with the default values initialized.
func NewPostCvdValidationTasksMoidParams() *PostCvdValidationTasksMoidParams {
	var ()
	return &PostCvdValidationTasksMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCvdValidationTasksMoidParamsWithTimeout creates a new PostCvdValidationTasksMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCvdValidationTasksMoidParamsWithTimeout(timeout time.Duration) *PostCvdValidationTasksMoidParams {
	var ()
	return &PostCvdValidationTasksMoidParams{

		timeout: timeout,
	}
}

// NewPostCvdValidationTasksMoidParamsWithContext creates a new PostCvdValidationTasksMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCvdValidationTasksMoidParamsWithContext(ctx context.Context) *PostCvdValidationTasksMoidParams {
	var ()
	return &PostCvdValidationTasksMoidParams{

		Context: ctx,
	}
}

// NewPostCvdValidationTasksMoidParamsWithHTTPClient creates a new PostCvdValidationTasksMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCvdValidationTasksMoidParamsWithHTTPClient(client *http.Client) *PostCvdValidationTasksMoidParams {
	var ()
	return &PostCvdValidationTasksMoidParams{
		HTTPClient: client,
	}
}

/*PostCvdValidationTasksMoidParams contains all the parameters to send to the API endpoint
for the post cvd validation tasks moid operation typically these are written to a http.Request
*/
type PostCvdValidationTasksMoidParams struct {

	/*Body
	  cvdValidationTask to update

	*/
	Body *models.CvdValidationTask
	/*Moid
	  The moid of the cvdValidationTask instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) WithTimeout(timeout time.Duration) *PostCvdValidationTasksMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) WithContext(ctx context.Context) *PostCvdValidationTasksMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) WithHTTPClient(client *http.Client) *PostCvdValidationTasksMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) WithBody(body *models.CvdValidationTask) *PostCvdValidationTasksMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) SetBody(body *models.CvdValidationTask) {
	o.Body = body
}

// WithMoid adds the moid to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) WithMoid(moid string) *PostCvdValidationTasksMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post cvd validation tasks moid params
func (o *PostCvdValidationTasksMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostCvdValidationTasksMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
