// Code generated by go-swagger; DO NOT EDIT.

package appliance_restore

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

// NewPostApplianceRestoresParams creates a new PostApplianceRestoresParams object
// with the default values initialized.
func NewPostApplianceRestoresParams() *PostApplianceRestoresParams {
	var ()
	return &PostApplianceRestoresParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostApplianceRestoresParamsWithTimeout creates a new PostApplianceRestoresParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostApplianceRestoresParamsWithTimeout(timeout time.Duration) *PostApplianceRestoresParams {
	var ()
	return &PostApplianceRestoresParams{

		timeout: timeout,
	}
}

// NewPostApplianceRestoresParamsWithContext creates a new PostApplianceRestoresParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostApplianceRestoresParamsWithContext(ctx context.Context) *PostApplianceRestoresParams {
	var ()
	return &PostApplianceRestoresParams{

		Context: ctx,
	}
}

// NewPostApplianceRestoresParamsWithHTTPClient creates a new PostApplianceRestoresParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostApplianceRestoresParamsWithHTTPClient(client *http.Client) *PostApplianceRestoresParams {
	var ()
	return &PostApplianceRestoresParams{
		HTTPClient: client,
	}
}

/*PostApplianceRestoresParams contains all the parameters to send to the API endpoint
for the post appliance restores operation typically these are written to a http.Request
*/
type PostApplianceRestoresParams struct {

	/*Body
	  applianceRestore to add

	*/
	Body *models.ApplianceRestore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post appliance restores params
func (o *PostApplianceRestoresParams) WithTimeout(timeout time.Duration) *PostApplianceRestoresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post appliance restores params
func (o *PostApplianceRestoresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post appliance restores params
func (o *PostApplianceRestoresParams) WithContext(ctx context.Context) *PostApplianceRestoresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post appliance restores params
func (o *PostApplianceRestoresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post appliance restores params
func (o *PostApplianceRestoresParams) WithHTTPClient(client *http.Client) *PostApplianceRestoresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post appliance restores params
func (o *PostApplianceRestoresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post appliance restores params
func (o *PostApplianceRestoresParams) WithBody(body *models.ApplianceRestore) *PostApplianceRestoresParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post appliance restores params
func (o *PostApplianceRestoresParams) SetBody(body *models.ApplianceRestore) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostApplianceRestoresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
