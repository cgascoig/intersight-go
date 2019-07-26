// Code generated by go-swagger; DO NOT EDIT.

package bios_unit

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

// NewPostBiosUnitsMoidParams creates a new PostBiosUnitsMoidParams object
// with the default values initialized.
func NewPostBiosUnitsMoidParams() *PostBiosUnitsMoidParams {
	var ()
	return &PostBiosUnitsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostBiosUnitsMoidParamsWithTimeout creates a new PostBiosUnitsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostBiosUnitsMoidParamsWithTimeout(timeout time.Duration) *PostBiosUnitsMoidParams {
	var ()
	return &PostBiosUnitsMoidParams{

		timeout: timeout,
	}
}

// NewPostBiosUnitsMoidParamsWithContext creates a new PostBiosUnitsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostBiosUnitsMoidParamsWithContext(ctx context.Context) *PostBiosUnitsMoidParams {
	var ()
	return &PostBiosUnitsMoidParams{

		Context: ctx,
	}
}

// NewPostBiosUnitsMoidParamsWithHTTPClient creates a new PostBiosUnitsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostBiosUnitsMoidParamsWithHTTPClient(client *http.Client) *PostBiosUnitsMoidParams {
	var ()
	return &PostBiosUnitsMoidParams{
		HTTPClient: client,
	}
}

/*PostBiosUnitsMoidParams contains all the parameters to send to the API endpoint
for the post bios units moid operation typically these are written to a http.Request
*/
type PostBiosUnitsMoidParams struct {

	/*Body
	  biosUnit to update

	*/
	Body *models.BiosUnit
	/*Moid
	  The moid of the biosUnit instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post bios units moid params
func (o *PostBiosUnitsMoidParams) WithTimeout(timeout time.Duration) *PostBiosUnitsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post bios units moid params
func (o *PostBiosUnitsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post bios units moid params
func (o *PostBiosUnitsMoidParams) WithContext(ctx context.Context) *PostBiosUnitsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post bios units moid params
func (o *PostBiosUnitsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post bios units moid params
func (o *PostBiosUnitsMoidParams) WithHTTPClient(client *http.Client) *PostBiosUnitsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post bios units moid params
func (o *PostBiosUnitsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post bios units moid params
func (o *PostBiosUnitsMoidParams) WithBody(body *models.BiosUnit) *PostBiosUnitsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post bios units moid params
func (o *PostBiosUnitsMoidParams) SetBody(body *models.BiosUnit) {
	o.Body = body
}

// WithMoid adds the moid to the post bios units moid params
func (o *PostBiosUnitsMoidParams) WithMoid(moid string) *PostBiosUnitsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post bios units moid params
func (o *PostBiosUnitsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostBiosUnitsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
