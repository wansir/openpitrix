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

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewCreateRepoSelectorParams creates a new CreateRepoSelectorParams object
// with the default values initialized.
func NewCreateRepoSelectorParams() *CreateRepoSelectorParams {
	var ()
	return &CreateRepoSelectorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepoSelectorParamsWithTimeout creates a new CreateRepoSelectorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRepoSelectorParamsWithTimeout(timeout time.Duration) *CreateRepoSelectorParams {
	var ()
	return &CreateRepoSelectorParams{

		timeout: timeout,
	}
}

// NewCreateRepoSelectorParamsWithContext creates a new CreateRepoSelectorParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRepoSelectorParamsWithContext(ctx context.Context) *CreateRepoSelectorParams {
	var ()
	return &CreateRepoSelectorParams{

		Context: ctx,
	}
}

// NewCreateRepoSelectorParamsWithHTTPClient creates a new CreateRepoSelectorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRepoSelectorParamsWithHTTPClient(client *http.Client) *CreateRepoSelectorParams {
	var ()
	return &CreateRepoSelectorParams{
		HTTPClient: client,
	}
}

/*CreateRepoSelectorParams contains all the parameters to send to the API endpoint
for the create repo selector operation typically these are written to a http.Request
*/
type CreateRepoSelectorParams struct {

	/*Body*/
	Body *models.OpenpitrixCreateRepoSelectorRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create repo selector params
func (o *CreateRepoSelectorParams) WithTimeout(timeout time.Duration) *CreateRepoSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repo selector params
func (o *CreateRepoSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repo selector params
func (o *CreateRepoSelectorParams) WithContext(ctx context.Context) *CreateRepoSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repo selector params
func (o *CreateRepoSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repo selector params
func (o *CreateRepoSelectorParams) WithHTTPClient(client *http.Client) *CreateRepoSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repo selector params
func (o *CreateRepoSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repo selector params
func (o *CreateRepoSelectorParams) WithBody(body *models.OpenpitrixCreateRepoSelectorRequest) *CreateRepoSelectorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repo selector params
func (o *CreateRepoSelectorParams) SetBody(body *models.OpenpitrixCreateRepoSelectorRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepoSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
