// Code generated by go-swagger; DO NOT EDIT.

package cluster_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewUpdateClusterParams creates a new UpdateClusterParams object
// with the default values initialized.
func NewUpdateClusterParams() *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClusterParamsWithTimeout creates a new UpdateClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateClusterParamsWithTimeout(timeout time.Duration) *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{

		timeout: timeout,
	}
}

// NewUpdateClusterParamsWithContext creates a new UpdateClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateClusterParamsWithContext(ctx context.Context) *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{

		Context: ctx,
	}
}

// NewUpdateClusterParamsWithHTTPClient creates a new UpdateClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateClusterParamsWithHTTPClient(client *http.Client) *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{
		HTTPClient: client,
	}
}

/*UpdateClusterParams contains all the parameters to send to the API endpoint
for the update cluster operation typically these are written to a http.Request
*/
type UpdateClusterParams struct {

	/*Body*/
	Body *models.OpenpitrixCluster
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) WithTimeout(timeout time.Duration) *UpdateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cluster params
func (o *UpdateClusterParams) WithContext(ctx context.Context) *UpdateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cluster params
func (o *UpdateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) WithHTTPClient(client *http.Client) *UpdateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update cluster params
func (o *UpdateClusterParams) WithBody(body *models.OpenpitrixCluster) *UpdateClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update cluster params
func (o *UpdateClusterParams) SetBody(body *models.OpenpitrixCluster) {
	o.Body = body
}

// WithID adds the id to the update cluster params
func (o *UpdateClusterParams) WithID(id string) *UpdateClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update cluster params
func (o *UpdateClusterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.OpenpitrixCluster)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
