// Code generated by go-swagger; DO NOT EDIT.

package cvd_deployment_task

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

// NewPatchCvdDeploymentTasksMoidParams creates a new PatchCvdDeploymentTasksMoidParams object
// with the default values initialized.
func NewPatchCvdDeploymentTasksMoidParams() *PatchCvdDeploymentTasksMoidParams {
	var ()
	return &PatchCvdDeploymentTasksMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCvdDeploymentTasksMoidParamsWithTimeout creates a new PatchCvdDeploymentTasksMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCvdDeploymentTasksMoidParamsWithTimeout(timeout time.Duration) *PatchCvdDeploymentTasksMoidParams {
	var ()
	return &PatchCvdDeploymentTasksMoidParams{

		timeout: timeout,
	}
}

// NewPatchCvdDeploymentTasksMoidParamsWithContext creates a new PatchCvdDeploymentTasksMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCvdDeploymentTasksMoidParamsWithContext(ctx context.Context) *PatchCvdDeploymentTasksMoidParams {
	var ()
	return &PatchCvdDeploymentTasksMoidParams{

		Context: ctx,
	}
}

// NewPatchCvdDeploymentTasksMoidParamsWithHTTPClient creates a new PatchCvdDeploymentTasksMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCvdDeploymentTasksMoidParamsWithHTTPClient(client *http.Client) *PatchCvdDeploymentTasksMoidParams {
	var ()
	return &PatchCvdDeploymentTasksMoidParams{
		HTTPClient: client,
	}
}

/*PatchCvdDeploymentTasksMoidParams contains all the parameters to send to the API endpoint
for the patch cvd deployment tasks moid operation typically these are written to a http.Request
*/
type PatchCvdDeploymentTasksMoidParams struct {

	/*Body
	  cvdDeploymentTask to update

	*/
	Body *models.CvdDeploymentTask
	/*Moid
	  The moid of the cvdDeploymentTask instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) WithTimeout(timeout time.Duration) *PatchCvdDeploymentTasksMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) WithContext(ctx context.Context) *PatchCvdDeploymentTasksMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) WithHTTPClient(client *http.Client) *PatchCvdDeploymentTasksMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) WithBody(body *models.CvdDeploymentTask) *PatchCvdDeploymentTasksMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) SetBody(body *models.CvdDeploymentTask) {
	o.Body = body
}

// WithMoid adds the moid to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) WithMoid(moid string) *PatchCvdDeploymentTasksMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch cvd deployment tasks moid params
func (o *PatchCvdDeploymentTasksMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCvdDeploymentTasksMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
