// Code generated by go-swagger; DO NOT EDIT.

package app_runtime_service

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

// NewUpdateAppRuntimeParams creates a new UpdateAppRuntimeParams object
// with the default values initialized.
func NewUpdateAppRuntimeParams() *UpdateAppRuntimeParams {
	var ()
	return &UpdateAppRuntimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppRuntimeParamsWithTimeout creates a new UpdateAppRuntimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAppRuntimeParamsWithTimeout(timeout time.Duration) *UpdateAppRuntimeParams {
	var ()
	return &UpdateAppRuntimeParams{

		timeout: timeout,
	}
}

// NewUpdateAppRuntimeParamsWithContext creates a new UpdateAppRuntimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAppRuntimeParamsWithContext(ctx context.Context) *UpdateAppRuntimeParams {
	var ()
	return &UpdateAppRuntimeParams{

		Context: ctx,
	}
}

// NewUpdateAppRuntimeParamsWithHTTPClient creates a new UpdateAppRuntimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAppRuntimeParamsWithHTTPClient(client *http.Client) *UpdateAppRuntimeParams {
	var ()
	return &UpdateAppRuntimeParams{
		HTTPClient: client,
	}
}

/*UpdateAppRuntimeParams contains all the parameters to send to the API endpoint
for the update app runtime operation typically these are written to a http.Request
*/
type UpdateAppRuntimeParams struct {

	/*Body*/
	Body *models.OpenpitrixAppRuntime
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update app runtime params
func (o *UpdateAppRuntimeParams) WithTimeout(timeout time.Duration) *UpdateAppRuntimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update app runtime params
func (o *UpdateAppRuntimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update app runtime params
func (o *UpdateAppRuntimeParams) WithContext(ctx context.Context) *UpdateAppRuntimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update app runtime params
func (o *UpdateAppRuntimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update app runtime params
func (o *UpdateAppRuntimeParams) WithHTTPClient(client *http.Client) *UpdateAppRuntimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update app runtime params
func (o *UpdateAppRuntimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update app runtime params
func (o *UpdateAppRuntimeParams) WithBody(body *models.OpenpitrixAppRuntime) *UpdateAppRuntimeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update app runtime params
func (o *UpdateAppRuntimeParams) SetBody(body *models.OpenpitrixAppRuntime) {
	o.Body = body
}

// WithID adds the id to the update app runtime params
func (o *UpdateAppRuntimeParams) WithID(id string) *UpdateAppRuntimeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update app runtime params
func (o *UpdateAppRuntimeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppRuntimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.OpenpitrixAppRuntime)
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
