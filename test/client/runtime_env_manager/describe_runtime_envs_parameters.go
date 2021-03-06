// Code generated by go-swagger; DO NOT EDIT.

package runtime_env_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeRuntimeEnvsParams creates a new DescribeRuntimeEnvsParams object
// with the default values initialized.
func NewDescribeRuntimeEnvsParams() *DescribeRuntimeEnvsParams {
	var ()
	return &DescribeRuntimeEnvsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeRuntimeEnvsParamsWithTimeout creates a new DescribeRuntimeEnvsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeRuntimeEnvsParamsWithTimeout(timeout time.Duration) *DescribeRuntimeEnvsParams {
	var ()
	return &DescribeRuntimeEnvsParams{

		timeout: timeout,
	}
}

// NewDescribeRuntimeEnvsParamsWithContext creates a new DescribeRuntimeEnvsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeRuntimeEnvsParamsWithContext(ctx context.Context) *DescribeRuntimeEnvsParams {
	var ()
	return &DescribeRuntimeEnvsParams{

		Context: ctx,
	}
}

// NewDescribeRuntimeEnvsParamsWithHTTPClient creates a new DescribeRuntimeEnvsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeRuntimeEnvsParamsWithHTTPClient(client *http.Client) *DescribeRuntimeEnvsParams {
	var ()
	return &DescribeRuntimeEnvsParams{
		HTTPClient: client,
	}
}

/*DescribeRuntimeEnvsParams contains all the parameters to send to the API endpoint
for the describe runtime envs operation typically these are written to a http.Request
*/
type DescribeRuntimeEnvsParams struct {

	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*Owner*/
	Owner []string
	/*RuntimeEnvID*/
	RuntimeEnvID []string
	/*SearchWord*/
	SearchWord *string
	/*Selector*/
	Selector *string
	/*Status*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithTimeout(timeout time.Duration) *DescribeRuntimeEnvsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithContext(ctx context.Context) *DescribeRuntimeEnvsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithHTTPClient(client *http.Client) *DescribeRuntimeEnvsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithLimit(limit *int64) *DescribeRuntimeEnvsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithOffset(offset *int64) *DescribeRuntimeEnvsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithOwner(owner []string) *DescribeRuntimeEnvsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithRuntimeEnvID adds the runtimeEnvID to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithRuntimeEnvID(runtimeEnvID []string) *DescribeRuntimeEnvsParams {
	o.SetRuntimeEnvID(runtimeEnvID)
	return o
}

// SetRuntimeEnvID adds the runtimeEnvId to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetRuntimeEnvID(runtimeEnvID []string) {
	o.RuntimeEnvID = runtimeEnvID
}

// WithSearchWord adds the searchWord to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithSearchWord(searchWord *string) *DescribeRuntimeEnvsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSelector adds the selector to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithSelector(selector *string) *DescribeRuntimeEnvsParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetSelector(selector *string) {
	o.Selector = selector
}

// WithStatus adds the status to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) WithStatus(status []string) *DescribeRuntimeEnvsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe runtime envs params
func (o *DescribeRuntimeEnvsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeRuntimeEnvsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	valuesRuntimeEnvID := o.RuntimeEnvID

	joinedRuntimeEnvID := swag.JoinByFormat(valuesRuntimeEnvID, "")
	// query array param runtime_env_id
	if err := r.SetQueryParam("runtime_env_id", joinedRuntimeEnvID...); err != nil {
		return err
	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.Selector != nil {

		// query param selector
		var qrSelector string
		if o.Selector != nil {
			qrSelector = *o.Selector
		}
		qSelector := qrSelector
		if qSelector != "" {
			if err := r.SetQueryParam("selector", qSelector); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
