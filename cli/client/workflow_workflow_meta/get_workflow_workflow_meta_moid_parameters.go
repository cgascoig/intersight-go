// Code generated by go-swagger; DO NOT EDIT.

package workflow_workflow_meta

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

// NewGetWorkflowWorkflowMetaMoidParams creates a new GetWorkflowWorkflowMetaMoidParams object
// with the default values initialized.
func NewGetWorkflowWorkflowMetaMoidParams() *GetWorkflowWorkflowMetaMoidParams {
	var ()
	return &GetWorkflowWorkflowMetaMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowWorkflowMetaMoidParamsWithTimeout creates a new GetWorkflowWorkflowMetaMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkflowWorkflowMetaMoidParamsWithTimeout(timeout time.Duration) *GetWorkflowWorkflowMetaMoidParams {
	var ()
	return &GetWorkflowWorkflowMetaMoidParams{

		timeout: timeout,
	}
}

// NewGetWorkflowWorkflowMetaMoidParamsWithContext creates a new GetWorkflowWorkflowMetaMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkflowWorkflowMetaMoidParamsWithContext(ctx context.Context) *GetWorkflowWorkflowMetaMoidParams {
	var ()
	return &GetWorkflowWorkflowMetaMoidParams{

		Context: ctx,
	}
}

// NewGetWorkflowWorkflowMetaMoidParamsWithHTTPClient creates a new GetWorkflowWorkflowMetaMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkflowWorkflowMetaMoidParamsWithHTTPClient(client *http.Client) *GetWorkflowWorkflowMetaMoidParams {
	var ()
	return &GetWorkflowWorkflowMetaMoidParams{
		HTTPClient: client,
	}
}

/*GetWorkflowWorkflowMetaMoidParams contains all the parameters to send to the API endpoint
for the get workflow workflow meta moid operation typically these are written to a http.Request
*/
type GetWorkflowWorkflowMetaMoidParams struct {

	/*Moid
	  The moid of the workflowWorkflowMeta instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workflow workflow meta moid params
func (o *GetWorkflowWorkflowMetaMoidParams) WithTimeout(timeout time.Duration) *GetWorkflowWorkflowMetaMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow workflow meta moid params
func (o *GetWorkflowWorkflowMetaMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow workflow meta moid params
func (o *GetWorkflowWorkflowMetaMoidParams) WithContext(ctx context.Context) *GetWorkflowWorkflowMetaMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow workflow meta moid params
func (o *GetWorkflowWorkflowMetaMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow workflow meta moid params
func (o *GetWorkflowWorkflowMetaMoidParams) WithHTTPClient(client *http.Client) *GetWorkflowWorkflowMetaMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow workflow meta moid params
func (o *GetWorkflowWorkflowMetaMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get workflow workflow meta moid params
func (o *GetWorkflowWorkflowMetaMoidParams) WithMoid(moid string) *GetWorkflowWorkflowMetaMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get workflow workflow meta moid params
func (o *GetWorkflowWorkflowMetaMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowWorkflowMetaMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
