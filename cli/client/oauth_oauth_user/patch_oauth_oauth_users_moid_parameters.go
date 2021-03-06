// Code generated by go-swagger; DO NOT EDIT.

package oauth_oauth_user

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

// NewPatchOauthOauthUsersMoidParams creates a new PatchOauthOauthUsersMoidParams object
// with the default values initialized.
func NewPatchOauthOauthUsersMoidParams() *PatchOauthOauthUsersMoidParams {
	var ()
	return &PatchOauthOauthUsersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchOauthOauthUsersMoidParamsWithTimeout creates a new PatchOauthOauthUsersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchOauthOauthUsersMoidParamsWithTimeout(timeout time.Duration) *PatchOauthOauthUsersMoidParams {
	var ()
	return &PatchOauthOauthUsersMoidParams{

		timeout: timeout,
	}
}

// NewPatchOauthOauthUsersMoidParamsWithContext creates a new PatchOauthOauthUsersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchOauthOauthUsersMoidParamsWithContext(ctx context.Context) *PatchOauthOauthUsersMoidParams {
	var ()
	return &PatchOauthOauthUsersMoidParams{

		Context: ctx,
	}
}

// NewPatchOauthOauthUsersMoidParamsWithHTTPClient creates a new PatchOauthOauthUsersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchOauthOauthUsersMoidParamsWithHTTPClient(client *http.Client) *PatchOauthOauthUsersMoidParams {
	var ()
	return &PatchOauthOauthUsersMoidParams{
		HTTPClient: client,
	}
}

/*PatchOauthOauthUsersMoidParams contains all the parameters to send to the API endpoint
for the patch oauth oauth users moid operation typically these are written to a http.Request
*/
type PatchOauthOauthUsersMoidParams struct {

	/*Body
	  oauthOauthUser to update

	*/
	Body *models.OauthOauthUser
	/*Moid
	  The moid of the oauthOauthUser instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) WithTimeout(timeout time.Duration) *PatchOauthOauthUsersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) WithContext(ctx context.Context) *PatchOauthOauthUsersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) WithHTTPClient(client *http.Client) *PatchOauthOauthUsersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) WithBody(body *models.OauthOauthUser) *PatchOauthOauthUsersMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) SetBody(body *models.OauthOauthUser) {
	o.Body = body
}

// WithMoid adds the moid to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) WithMoid(moid string) *PatchOauthOauthUsersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch oauth oauth users moid params
func (o *PatchOauthOauthUsersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchOauthOauthUsersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
