// Code generated by go-swagger; DO NOT EDIT.

package invoices

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

// NewPerformSendActionParams creates a new PerformSendActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformSendActionParams() *PerformSendActionParams {
	return &PerformSendActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformSendActionParamsWithTimeout creates a new PerformSendActionParams object
// with the ability to set a timeout on a request.
func NewPerformSendActionParamsWithTimeout(timeout time.Duration) *PerformSendActionParams {
	return &PerformSendActionParams{
		timeout: timeout,
	}
}

// NewPerformSendActionParamsWithContext creates a new PerformSendActionParams object
// with the ability to set a context for a request.
func NewPerformSendActionParamsWithContext(ctx context.Context) *PerformSendActionParams {
	return &PerformSendActionParams{
		Context: ctx,
	}
}

// NewPerformSendActionParamsWithHTTPClient creates a new PerformSendActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformSendActionParamsWithHTTPClient(client *http.Client) *PerformSendActionParams {
	return &PerformSendActionParams{
		HTTPClient: client,
	}
}

/* PerformSendActionParams contains all the parameters to send to the API endpoint
   for the perform send action operation.

   Typically these are written to a http.Request.
*/
type PerformSendActionParams struct {

	/* ID.

	   The invoice number.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform send action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformSendActionParams) WithDefaults() *PerformSendActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform send action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformSendActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the perform send action params
func (o *PerformSendActionParams) WithTimeout(timeout time.Duration) *PerformSendActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform send action params
func (o *PerformSendActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform send action params
func (o *PerformSendActionParams) WithContext(ctx context.Context) *PerformSendActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform send action params
func (o *PerformSendActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform send action params
func (o *PerformSendActionParams) WithHTTPClient(client *http.Client) *PerformSendActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform send action params
func (o *PerformSendActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the perform send action params
func (o *PerformSendActionParams) WithID(id string) *PerformSendActionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the perform send action params
func (o *PerformSendActionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PerformSendActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
