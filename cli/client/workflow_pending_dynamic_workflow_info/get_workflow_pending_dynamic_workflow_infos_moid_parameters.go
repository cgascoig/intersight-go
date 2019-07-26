// Code generated by go-swagger; DO NOT EDIT.

package workflow_pending_dynamic_workflow_info

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

// NewGetWorkflowPendingDynamicWorkflowInfosMoidParams creates a new GetWorkflowPendingDynamicWorkflowInfosMoidParams object
// with the default values initialized.
func NewGetWorkflowPendingDynamicWorkflowInfosMoidParams() *GetWorkflowPendingDynamicWorkflowInfosMoidParams {
	var ()
	return &GetWorkflowPendingDynamicWorkflowInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowPendingDynamicWorkflowInfosMoidParamsWithTimeout creates a new GetWorkflowPendingDynamicWorkflowInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkflowPendingDynamicWorkflowInfosMoidParamsWithTimeout(timeout time.Duration) *GetWorkflowPendingDynamicWorkflowInfosMoidParams {
	var ()
	return &GetWorkflowPendingDynamicWorkflowInfosMoidParams{

		timeout: timeout,
	}
}

// NewGetWorkflowPendingDynamicWorkflowInfosMoidParamsWithContext creates a new GetWorkflowPendingDynamicWorkflowInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkflowPendingDynamicWorkflowInfosMoidParamsWithContext(ctx context.Context) *GetWorkflowPendingDynamicWorkflowInfosMoidParams {
	var ()
	return &GetWorkflowPendingDynamicWorkflowInfosMoidParams{

		Context: ctx,
	}
}

// NewGetWorkflowPendingDynamicWorkflowInfosMoidParamsWithHTTPClient creates a new GetWorkflowPendingDynamicWorkflowInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkflowPendingDynamicWorkflowInfosMoidParamsWithHTTPClient(client *http.Client) *GetWorkflowPendingDynamicWorkflowInfosMoidParams {
	var ()
	return &GetWorkflowPendingDynamicWorkflowInfosMoidParams{
		HTTPClient: client,
	}
}

/*GetWorkflowPendingDynamicWorkflowInfosMoidParams contains all the parameters to send to the API endpoint
for the get workflow pending dynamic workflow infos moid operation typically these are written to a http.Request
*/
type GetWorkflowPendingDynamicWorkflowInfosMoidParams struct {

	/*Moid
	  The moid of the workflowPendingDynamicWorkflowInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workflow pending dynamic workflow infos moid params
func (o *GetWorkflowPendingDynamicWorkflowInfosMoidParams) WithTimeout(timeout time.Duration) *GetWorkflowPendingDynamicWorkflowInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow pending dynamic workflow infos moid params
func (o *GetWorkflowPendingDynamicWorkflowInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow pending dynamic workflow infos moid params
func (o *GetWorkflowPendingDynamicWorkflowInfosMoidParams) WithContext(ctx context.Context) *GetWorkflowPendingDynamicWorkflowInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow pending dynamic workflow infos moid params
func (o *GetWorkflowPendingDynamicWorkflowInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow pending dynamic workflow infos moid params
func (o *GetWorkflowPendingDynamicWorkflowInfosMoidParams) WithHTTPClient(client *http.Client) *GetWorkflowPendingDynamicWorkflowInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow pending dynamic workflow infos moid params
func (o *GetWorkflowPendingDynamicWorkflowInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get workflow pending dynamic workflow infos moid params
func (o *GetWorkflowPendingDynamicWorkflowInfosMoidParams) WithMoid(moid string) *GetWorkflowPendingDynamicWorkflowInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get workflow pending dynamic workflow infos moid params
func (o *GetWorkflowPendingDynamicWorkflowInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowPendingDynamicWorkflowInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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