// Code generated by go-swagger; DO NOT EDIT.

package compute_rack_unit

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

// NewPostComputeRackUnitsMoidParams creates a new PostComputeRackUnitsMoidParams object
// with the default values initialized.
func NewPostComputeRackUnitsMoidParams() *PostComputeRackUnitsMoidParams {
	var ()
	return &PostComputeRackUnitsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostComputeRackUnitsMoidParamsWithTimeout creates a new PostComputeRackUnitsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostComputeRackUnitsMoidParamsWithTimeout(timeout time.Duration) *PostComputeRackUnitsMoidParams {
	var ()
	return &PostComputeRackUnitsMoidParams{

		timeout: timeout,
	}
}

// NewPostComputeRackUnitsMoidParamsWithContext creates a new PostComputeRackUnitsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostComputeRackUnitsMoidParamsWithContext(ctx context.Context) *PostComputeRackUnitsMoidParams {
	var ()
	return &PostComputeRackUnitsMoidParams{

		Context: ctx,
	}
}

// NewPostComputeRackUnitsMoidParamsWithHTTPClient creates a new PostComputeRackUnitsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostComputeRackUnitsMoidParamsWithHTTPClient(client *http.Client) *PostComputeRackUnitsMoidParams {
	var ()
	return &PostComputeRackUnitsMoidParams{
		HTTPClient: client,
	}
}

/*PostComputeRackUnitsMoidParams contains all the parameters to send to the API endpoint
for the post compute rack units moid operation typically these are written to a http.Request
*/
type PostComputeRackUnitsMoidParams struct {

	/*Body
	  computeRackUnit to update

	*/
	Body *models.ComputeRackUnit
	/*Moid
	  The moid of the computeRackUnit instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) WithTimeout(timeout time.Duration) *PostComputeRackUnitsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) WithContext(ctx context.Context) *PostComputeRackUnitsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) WithHTTPClient(client *http.Client) *PostComputeRackUnitsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) WithBody(body *models.ComputeRackUnit) *PostComputeRackUnitsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) SetBody(body *models.ComputeRackUnit) {
	o.Body = body
}

// WithMoid adds the moid to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) WithMoid(moid string) *PostComputeRackUnitsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post compute rack units moid params
func (o *PostComputeRackUnitsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostComputeRackUnitsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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