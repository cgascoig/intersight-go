// Code generated by go-swagger; DO NOT EDIT.

package adapter_host_eth_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAdapterHostEthInterfacesParams creates a new GetAdapterHostEthInterfacesParams object
// with the default values initialized.
func NewGetAdapterHostEthInterfacesParams() *GetAdapterHostEthInterfacesParams {
	var ()
	return &GetAdapterHostEthInterfacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdapterHostEthInterfacesParamsWithTimeout creates a new GetAdapterHostEthInterfacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAdapterHostEthInterfacesParamsWithTimeout(timeout time.Duration) *GetAdapterHostEthInterfacesParams {
	var ()
	return &GetAdapterHostEthInterfacesParams{

		timeout: timeout,
	}
}

// NewGetAdapterHostEthInterfacesParamsWithContext creates a new GetAdapterHostEthInterfacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAdapterHostEthInterfacesParamsWithContext(ctx context.Context) *GetAdapterHostEthInterfacesParams {
	var ()
	return &GetAdapterHostEthInterfacesParams{

		Context: ctx,
	}
}

// NewGetAdapterHostEthInterfacesParamsWithHTTPClient creates a new GetAdapterHostEthInterfacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAdapterHostEthInterfacesParamsWithHTTPClient(client *http.Client) *GetAdapterHostEthInterfacesParams {
	var ()
	return &GetAdapterHostEthInterfacesParams{
		HTTPClient: client,
	}
}

/*GetAdapterHostEthInterfacesParams contains all the parameters to send to the API endpoint
for the get adapter host eth interfaces operation typically these are written to a http.Request
*/
type GetAdapterHostEthInterfacesParams struct {

	/*DollarApply
	  Specify one or more transformation operations to perform aggregation on documents. The transformations
	are processed in order with the output from a transformation being used as input for the subsequent
	transformation.
	Query examples:
	$apply=groupby((Model), aggregate($count as Total))
	$apply=groupby((Model), aggregate(AvailableMemory with average as AverageAvailableMemory))


	*/
	DollarApply *string
	/*DollarCount
	  The $count query option allows clients to request a count of the matching resources.

	*/
	DollarCount *bool
	/*DollarExpand
	  Specify additional attributes or related documents to return. Supports only 'DisplayNames' attribute now.
	Query examples:
	$expand=DisplayNames


	*/
	DollarExpand *string
	/*DollarFilter
	  Filter criteria for documents to return. A URI with a $filter System Query Option identifies a
	subset of the Entries from the Collection of Entries identified by the Resource Path section
	of the URI. The subset is determined by selecting only the Entries that satisfy the predicate
	expression specified by the query option.
	The expression language that is used in $filter operators supports references to properties and
	literals. The literal values can be strings enclosed in single quotes, numbers and boolean values
	(true or false) or any of the additional literal representations shown in the Abstract
	Type System section.
	Query examples:
	$filter=Name eq 'Bob'
	$filter=Tags/any(t: t/Key eq 'Site')
	$filter=Tags/any(t: t/Key eq 'Site' and t/Value eq 'London')


	*/
	DollarFilter *string
	/*DollarInlinecount
	  The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response

	*/
	DollarInlinecount *string
	/*DollarOrderby
	  Determines what values are used to order a collection of documents.

	*/
	DollarOrderby *string
	/*DollarSelect
	  Specifies a subset of properties to return.

	*/
	DollarSelect *string
	/*DollarSkip
	  The number of documents to skip.

	*/
	DollarSkip *int32
	/*DollarTop
	  The max number of documents to return.

	*/
	DollarTop *int32
	/*At
	  Similar to "$filter", but "at" is specifically used to filter versioning information properties for documents to return.
	A URI with an "at" Query Option identifies a subset of the Entries from the Collection of Entries identified by the
	Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate
	expression specified by the query option.
	The expression language that is used in at operators supports references to properties and
	literals. The literal values can be strings enclosed in single quotes, numbers and boolean values
	(true or false) or any of the additional literal representations shown in the Abstract
	Type System section.
	Query examples:
	at=VersionType eq 'Configured'
	at=InterestedMos.Moid eq '5b5877e56c6730367acf46cd'


	*/
	At *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithTimeout(timeout time.Duration) *GetAdapterHostEthInterfacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithContext(ctx context.Context) *GetAdapterHostEthInterfacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithHTTPClient(client *http.Client) *GetAdapterHostEthInterfacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarApply adds the dollarApply to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithDollarApply(dollarApply *string) *GetAdapterHostEthInterfacesParams {
	o.SetDollarApply(dollarApply)
	return o
}

// SetDollarApply adds the dollarApply to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetDollarApply(dollarApply *string) {
	o.DollarApply = dollarApply
}

// WithDollarCount adds the dollarCount to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithDollarCount(dollarCount *bool) *GetAdapterHostEthInterfacesParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarExpand adds the dollarExpand to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithDollarExpand(dollarExpand *string) *GetAdapterHostEthInterfacesParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithDollarFilter adds the dollarFilter to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithDollarFilter(dollarFilter *string) *GetAdapterHostEthInterfacesParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarInlinecount adds the dollarInlinecount to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithDollarInlinecount(dollarInlinecount *string) *GetAdapterHostEthInterfacesParams {
	o.SetDollarInlinecount(dollarInlinecount)
	return o
}

// SetDollarInlinecount adds the dollarInlinecount to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetDollarInlinecount(dollarInlinecount *string) {
	o.DollarInlinecount = dollarInlinecount
}

// WithDollarOrderby adds the dollarOrderby to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithDollarOrderby(dollarOrderby *string) *GetAdapterHostEthInterfacesParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithDollarSelect(dollarSelect *string) *GetAdapterHostEthInterfacesParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithDollarSkip(dollarSkip *int32) *GetAdapterHostEthInterfacesParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithDollarTop(dollarTop *int32) *GetAdapterHostEthInterfacesParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAt adds the at to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) WithAt(at *string) *GetAdapterHostEthInterfacesParams {
	o.SetAt(at)
	return o
}

// SetAt adds the at to the get adapter host eth interfaces params
func (o *GetAdapterHostEthInterfacesParams) SetAt(at *string) {
	o.At = at
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdapterHostEthInterfacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarApply != nil {

		// query param $apply
		var qrDollarApply string
		if o.DollarApply != nil {
			qrDollarApply = *o.DollarApply
		}
		qDollarApply := qrDollarApply
		if qDollarApply != "" {
			if err := r.SetQueryParam("$apply", qDollarApply); err != nil {
				return err
			}
		}

	}

	if o.DollarCount != nil {

		// query param $count
		var qrDollarCount bool
		if o.DollarCount != nil {
			qrDollarCount = *o.DollarCount
		}
		qDollarCount := swag.FormatBool(qrDollarCount)
		if qDollarCount != "" {
			if err := r.SetQueryParam("$count", qDollarCount); err != nil {
				return err
			}
		}

	}

	if o.DollarExpand != nil {

		// query param $expand
		var qrDollarExpand string
		if o.DollarExpand != nil {
			qrDollarExpand = *o.DollarExpand
		}
		qDollarExpand := qrDollarExpand
		if qDollarExpand != "" {
			if err := r.SetQueryParam("$expand", qDollarExpand); err != nil {
				return err
			}
		}

	}

	if o.DollarFilter != nil {

		// query param $filter
		var qrDollarFilter string
		if o.DollarFilter != nil {
			qrDollarFilter = *o.DollarFilter
		}
		qDollarFilter := qrDollarFilter
		if qDollarFilter != "" {
			if err := r.SetQueryParam("$filter", qDollarFilter); err != nil {
				return err
			}
		}

	}

	if o.DollarInlinecount != nil {

		// query param $inlinecount
		var qrDollarInlinecount string
		if o.DollarInlinecount != nil {
			qrDollarInlinecount = *o.DollarInlinecount
		}
		qDollarInlinecount := qrDollarInlinecount
		if qDollarInlinecount != "" {
			if err := r.SetQueryParam("$inlinecount", qDollarInlinecount); err != nil {
				return err
			}
		}

	}

	if o.DollarOrderby != nil {

		// query param $orderby
		var qrDollarOrderby string
		if o.DollarOrderby != nil {
			qrDollarOrderby = *o.DollarOrderby
		}
		qDollarOrderby := qrDollarOrderby
		if qDollarOrderby != "" {
			if err := r.SetQueryParam("$orderby", qDollarOrderby); err != nil {
				return err
			}
		}

	}

	if o.DollarSelect != nil {

		// query param $select
		var qrDollarSelect string
		if o.DollarSelect != nil {
			qrDollarSelect = *o.DollarSelect
		}
		qDollarSelect := qrDollarSelect
		if qDollarSelect != "" {
			if err := r.SetQueryParam("$select", qDollarSelect); err != nil {
				return err
			}
		}

	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip int32
		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := swag.FormatInt32(qrDollarSkip)
		if qDollarSkip != "" {
			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}

	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop int32
		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt32(qrDollarTop)
		if qDollarTop != "" {
			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
		}

	}

	if o.At != nil {

		// query param at
		var qrAt string
		if o.At != nil {
			qrAt = *o.At
		}
		qAt := qrAt
		if qAt != "" {
			if err := r.SetQueryParam("at", qAt); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
