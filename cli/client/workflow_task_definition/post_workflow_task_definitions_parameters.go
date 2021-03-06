// Code generated by go-swagger; DO NOT EDIT.

package workflow_task_definition

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

// NewPostWorkflowTaskDefinitionsParams creates a new PostWorkflowTaskDefinitionsParams object
// with the default values initialized.
func NewPostWorkflowTaskDefinitionsParams() *PostWorkflowTaskDefinitionsParams {
	var ()
	return &PostWorkflowTaskDefinitionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkflowTaskDefinitionsParamsWithTimeout creates a new PostWorkflowTaskDefinitionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWorkflowTaskDefinitionsParamsWithTimeout(timeout time.Duration) *PostWorkflowTaskDefinitionsParams {
	var ()
	return &PostWorkflowTaskDefinitionsParams{

		timeout: timeout,
	}
}

// NewPostWorkflowTaskDefinitionsParamsWithContext creates a new PostWorkflowTaskDefinitionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWorkflowTaskDefinitionsParamsWithContext(ctx context.Context) *PostWorkflowTaskDefinitionsParams {
	var ()
	return &PostWorkflowTaskDefinitionsParams{

		Context: ctx,
	}
}

// NewPostWorkflowTaskDefinitionsParamsWithHTTPClient creates a new PostWorkflowTaskDefinitionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWorkflowTaskDefinitionsParamsWithHTTPClient(client *http.Client) *PostWorkflowTaskDefinitionsParams {
	var ()
	return &PostWorkflowTaskDefinitionsParams{
		HTTPClient: client,
	}
}

/*PostWorkflowTaskDefinitionsParams contains all the parameters to send to the API endpoint
for the post workflow task definitions operation typically these are written to a http.Request
*/
type PostWorkflowTaskDefinitionsParams struct {

	/*Body
	  workflowTaskDefinition to add

	*/
	Body *models.WorkflowTaskDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post workflow task definitions params
func (o *PostWorkflowTaskDefinitionsParams) WithTimeout(timeout time.Duration) *PostWorkflowTaskDefinitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workflow task definitions params
func (o *PostWorkflowTaskDefinitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workflow task definitions params
func (o *PostWorkflowTaskDefinitionsParams) WithContext(ctx context.Context) *PostWorkflowTaskDefinitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workflow task definitions params
func (o *PostWorkflowTaskDefinitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workflow task definitions params
func (o *PostWorkflowTaskDefinitionsParams) WithHTTPClient(client *http.Client) *PostWorkflowTaskDefinitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workflow task definitions params
func (o *PostWorkflowTaskDefinitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workflow task definitions params
func (o *PostWorkflowTaskDefinitionsParams) WithBody(body *models.WorkflowTaskDefinition) *PostWorkflowTaskDefinitionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workflow task definitions params
func (o *PostWorkflowTaskDefinitionsParams) SetBody(body *models.WorkflowTaskDefinition) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkflowTaskDefinitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
