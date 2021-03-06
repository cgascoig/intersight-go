// Code generated by go-swagger; DO NOT EDIT.

package workflow_workflow_info

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

// NewPatchWorkflowWorkflowInfosMoidParams creates a new PatchWorkflowWorkflowInfosMoidParams object
// with the default values initialized.
func NewPatchWorkflowWorkflowInfosMoidParams() *PatchWorkflowWorkflowInfosMoidParams {
	var ()
	return &PatchWorkflowWorkflowInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWorkflowWorkflowInfosMoidParamsWithTimeout creates a new PatchWorkflowWorkflowInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWorkflowWorkflowInfosMoidParamsWithTimeout(timeout time.Duration) *PatchWorkflowWorkflowInfosMoidParams {
	var ()
	return &PatchWorkflowWorkflowInfosMoidParams{

		timeout: timeout,
	}
}

// NewPatchWorkflowWorkflowInfosMoidParamsWithContext creates a new PatchWorkflowWorkflowInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWorkflowWorkflowInfosMoidParamsWithContext(ctx context.Context) *PatchWorkflowWorkflowInfosMoidParams {
	var ()
	return &PatchWorkflowWorkflowInfosMoidParams{

		Context: ctx,
	}
}

// NewPatchWorkflowWorkflowInfosMoidParamsWithHTTPClient creates a new PatchWorkflowWorkflowInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWorkflowWorkflowInfosMoidParamsWithHTTPClient(client *http.Client) *PatchWorkflowWorkflowInfosMoidParams {
	var ()
	return &PatchWorkflowWorkflowInfosMoidParams{
		HTTPClient: client,
	}
}

/*PatchWorkflowWorkflowInfosMoidParams contains all the parameters to send to the API endpoint
for the patch workflow workflow infos moid operation typically these are written to a http.Request
*/
type PatchWorkflowWorkflowInfosMoidParams struct {

	/*Body
	  workflowWorkflowInfo to update

	*/
	Body *models.WorkflowWorkflowInfo
	/*Moid
	  The moid of the workflowWorkflowInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) WithTimeout(timeout time.Duration) *PatchWorkflowWorkflowInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) WithContext(ctx context.Context) *PatchWorkflowWorkflowInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) WithHTTPClient(client *http.Client) *PatchWorkflowWorkflowInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) WithBody(body *models.WorkflowWorkflowInfo) *PatchWorkflowWorkflowInfosMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) SetBody(body *models.WorkflowWorkflowInfo) {
	o.Body = body
}

// WithMoid adds the moid to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) WithMoid(moid string) *PatchWorkflowWorkflowInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch workflow workflow infos moid params
func (o *PatchWorkflowWorkflowInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWorkflowWorkflowInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
