// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_initiate_hx_cluster_upgrade

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

// NewPostHyperflexInitiateHxClusterUpgradesParams creates a new PostHyperflexInitiateHxClusterUpgradesParams object
// with the default values initialized.
func NewPostHyperflexInitiateHxClusterUpgradesParams() *PostHyperflexInitiateHxClusterUpgradesParams {
	var ()
	return &PostHyperflexInitiateHxClusterUpgradesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHyperflexInitiateHxClusterUpgradesParamsWithTimeout creates a new PostHyperflexInitiateHxClusterUpgradesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHyperflexInitiateHxClusterUpgradesParamsWithTimeout(timeout time.Duration) *PostHyperflexInitiateHxClusterUpgradesParams {
	var ()
	return &PostHyperflexInitiateHxClusterUpgradesParams{

		timeout: timeout,
	}
}

// NewPostHyperflexInitiateHxClusterUpgradesParamsWithContext creates a new PostHyperflexInitiateHxClusterUpgradesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHyperflexInitiateHxClusterUpgradesParamsWithContext(ctx context.Context) *PostHyperflexInitiateHxClusterUpgradesParams {
	var ()
	return &PostHyperflexInitiateHxClusterUpgradesParams{

		Context: ctx,
	}
}

// NewPostHyperflexInitiateHxClusterUpgradesParamsWithHTTPClient creates a new PostHyperflexInitiateHxClusterUpgradesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHyperflexInitiateHxClusterUpgradesParamsWithHTTPClient(client *http.Client) *PostHyperflexInitiateHxClusterUpgradesParams {
	var ()
	return &PostHyperflexInitiateHxClusterUpgradesParams{
		HTTPClient: client,
	}
}

/*PostHyperflexInitiateHxClusterUpgradesParams contains all the parameters to send to the API endpoint
for the post hyperflex initiate hx cluster upgrades operation typically these are written to a http.Request
*/
type PostHyperflexInitiateHxClusterUpgradesParams struct {

	/*Body
	  hyperflexInitiateHxClusterUpgrade to add

	*/
	Body *models.HyperflexInitiateHxClusterUpgrade

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post hyperflex initiate hx cluster upgrades params
func (o *PostHyperflexInitiateHxClusterUpgradesParams) WithTimeout(timeout time.Duration) *PostHyperflexInitiateHxClusterUpgradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hyperflex initiate hx cluster upgrades params
func (o *PostHyperflexInitiateHxClusterUpgradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hyperflex initiate hx cluster upgrades params
func (o *PostHyperflexInitiateHxClusterUpgradesParams) WithContext(ctx context.Context) *PostHyperflexInitiateHxClusterUpgradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hyperflex initiate hx cluster upgrades params
func (o *PostHyperflexInitiateHxClusterUpgradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post hyperflex initiate hx cluster upgrades params
func (o *PostHyperflexInitiateHxClusterUpgradesParams) WithHTTPClient(client *http.Client) *PostHyperflexInitiateHxClusterUpgradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post hyperflex initiate hx cluster upgrades params
func (o *PostHyperflexInitiateHxClusterUpgradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post hyperflex initiate hx cluster upgrades params
func (o *PostHyperflexInitiateHxClusterUpgradesParams) WithBody(body *models.HyperflexInitiateHxClusterUpgrade) *PostHyperflexInitiateHxClusterUpgradesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post hyperflex initiate hx cluster upgrades params
func (o *PostHyperflexInitiateHxClusterUpgradesParams) SetBody(body *models.HyperflexInitiateHxClusterUpgrade) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostHyperflexInitiateHxClusterUpgradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
