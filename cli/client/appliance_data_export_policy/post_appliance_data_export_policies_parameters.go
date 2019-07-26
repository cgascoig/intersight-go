// Code generated by go-swagger; DO NOT EDIT.

package appliance_data_export_policy

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

// NewPostApplianceDataExportPoliciesParams creates a new PostApplianceDataExportPoliciesParams object
// with the default values initialized.
func NewPostApplianceDataExportPoliciesParams() *PostApplianceDataExportPoliciesParams {
	var ()
	return &PostApplianceDataExportPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostApplianceDataExportPoliciesParamsWithTimeout creates a new PostApplianceDataExportPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostApplianceDataExportPoliciesParamsWithTimeout(timeout time.Duration) *PostApplianceDataExportPoliciesParams {
	var ()
	return &PostApplianceDataExportPoliciesParams{

		timeout: timeout,
	}
}

// NewPostApplianceDataExportPoliciesParamsWithContext creates a new PostApplianceDataExportPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostApplianceDataExportPoliciesParamsWithContext(ctx context.Context) *PostApplianceDataExportPoliciesParams {
	var ()
	return &PostApplianceDataExportPoliciesParams{

		Context: ctx,
	}
}

// NewPostApplianceDataExportPoliciesParamsWithHTTPClient creates a new PostApplianceDataExportPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostApplianceDataExportPoliciesParamsWithHTTPClient(client *http.Client) *PostApplianceDataExportPoliciesParams {
	var ()
	return &PostApplianceDataExportPoliciesParams{
		HTTPClient: client,
	}
}

/*PostApplianceDataExportPoliciesParams contains all the parameters to send to the API endpoint
for the post appliance data export policies operation typically these are written to a http.Request
*/
type PostApplianceDataExportPoliciesParams struct {

	/*Body
	  applianceDataExportPolicy to add

	*/
	Body *models.ApplianceDataExportPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post appliance data export policies params
func (o *PostApplianceDataExportPoliciesParams) WithTimeout(timeout time.Duration) *PostApplianceDataExportPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post appliance data export policies params
func (o *PostApplianceDataExportPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post appliance data export policies params
func (o *PostApplianceDataExportPoliciesParams) WithContext(ctx context.Context) *PostApplianceDataExportPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post appliance data export policies params
func (o *PostApplianceDataExportPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post appliance data export policies params
func (o *PostApplianceDataExportPoliciesParams) WithHTTPClient(client *http.Client) *PostApplianceDataExportPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post appliance data export policies params
func (o *PostApplianceDataExportPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post appliance data export policies params
func (o *PostApplianceDataExportPoliciesParams) WithBody(body *models.ApplianceDataExportPolicy) *PostApplianceDataExportPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post appliance data export policies params
func (o *PostApplianceDataExportPoliciesParams) SetBody(body *models.ApplianceDataExportPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostApplianceDataExportPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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