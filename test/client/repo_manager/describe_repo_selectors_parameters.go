// Code generated by go-swagger; DO NOT EDIT.

package repo_manager

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

// NewDescribeRepoSelectorsParams creates a new DescribeRepoSelectorsParams object
// with the default values initialized.
func NewDescribeRepoSelectorsParams() *DescribeRepoSelectorsParams {
	var ()
	return &DescribeRepoSelectorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeRepoSelectorsParamsWithTimeout creates a new DescribeRepoSelectorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeRepoSelectorsParamsWithTimeout(timeout time.Duration) *DescribeRepoSelectorsParams {
	var ()
	return &DescribeRepoSelectorsParams{

		timeout: timeout,
	}
}

// NewDescribeRepoSelectorsParamsWithContext creates a new DescribeRepoSelectorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeRepoSelectorsParamsWithContext(ctx context.Context) *DescribeRepoSelectorsParams {
	var ()
	return &DescribeRepoSelectorsParams{

		Context: ctx,
	}
}

// NewDescribeRepoSelectorsParamsWithHTTPClient creates a new DescribeRepoSelectorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeRepoSelectorsParamsWithHTTPClient(client *http.Client) *DescribeRepoSelectorsParams {
	var ()
	return &DescribeRepoSelectorsParams{
		HTTPClient: client,
	}
}

/*DescribeRepoSelectorsParams contains all the parameters to send to the API endpoint
for the describe repo selectors operation typically these are written to a http.Request
*/
type DescribeRepoSelectorsParams struct {

	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*RepoID*/
	RepoID []string
	/*RepoSelectorID*/
	RepoSelectorID []string
	/*Status*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) WithTimeout(timeout time.Duration) *DescribeRepoSelectorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) WithContext(ctx context.Context) *DescribeRepoSelectorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) WithHTTPClient(client *http.Client) *DescribeRepoSelectorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) WithLimit(limit *int64) *DescribeRepoSelectorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) WithOffset(offset *int64) *DescribeRepoSelectorsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithRepoID adds the repoID to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) WithRepoID(repoID []string) *DescribeRepoSelectorsParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) SetRepoID(repoID []string) {
	o.RepoID = repoID
}

// WithRepoSelectorID adds the repoSelectorID to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) WithRepoSelectorID(repoSelectorID []string) *DescribeRepoSelectorsParams {
	o.SetRepoSelectorID(repoSelectorID)
	return o
}

// SetRepoSelectorID adds the repoSelectorId to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) SetRepoSelectorID(repoSelectorID []string) {
	o.RepoSelectorID = repoSelectorID
}

// WithStatus adds the status to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) WithStatus(status []string) *DescribeRepoSelectorsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe repo selectors params
func (o *DescribeRepoSelectorsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeRepoSelectorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesRepoID := o.RepoID

	joinedRepoID := swag.JoinByFormat(valuesRepoID, "")
	// query array param repo_id
	if err := r.SetQueryParam("repo_id", joinedRepoID...); err != nil {
		return err
	}

	valuesRepoSelectorID := o.RepoSelectorID

	joinedRepoSelectorID := swag.JoinByFormat(valuesRepoSelectorID, "")
	// query array param repo_selector_id
	if err := r.SetQueryParam("repo_selector_id", joinedRepoSelectorID...); err != nil {
		return err
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
