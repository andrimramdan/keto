// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIsInstanceAliveParams creates a new IsInstanceAliveParams object
// with the default values initialized.
func NewIsInstanceAliveParams() *IsInstanceAliveParams {

	return &IsInstanceAliveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsInstanceAliveParamsWithTimeout creates a new IsInstanceAliveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsInstanceAliveParamsWithTimeout(timeout time.Duration) *IsInstanceAliveParams {

	return &IsInstanceAliveParams{

		timeout: timeout,
	}
}

// NewIsInstanceAliveParamsWithContext creates a new IsInstanceAliveParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsInstanceAliveParamsWithContext(ctx context.Context) *IsInstanceAliveParams {

	return &IsInstanceAliveParams{

		Context: ctx,
	}
}

// NewIsInstanceAliveParamsWithHTTPClient creates a new IsInstanceAliveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsInstanceAliveParamsWithHTTPClient(client *http.Client) *IsInstanceAliveParams {

	return &IsInstanceAliveParams{
		HTTPClient: client,
	}
}

/*IsInstanceAliveParams contains all the parameters to send to the API endpoint
for the is instance alive operation typically these are written to a http.Request
*/
type IsInstanceAliveParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is instance alive params
func (o *IsInstanceAliveParams) WithTimeout(timeout time.Duration) *IsInstanceAliveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is instance alive params
func (o *IsInstanceAliveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is instance alive params
func (o *IsInstanceAliveParams) WithContext(ctx context.Context) *IsInstanceAliveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is instance alive params
func (o *IsInstanceAliveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is instance alive params
func (o *IsInstanceAliveParams) WithHTTPClient(client *http.Client) *IsInstanceAliveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is instance alive params
func (o *IsInstanceAliveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IsInstanceAliveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
