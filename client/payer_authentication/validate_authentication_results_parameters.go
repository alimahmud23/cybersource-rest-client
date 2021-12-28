// Code generated by go-swagger; DO NOT EDIT.

package payer_authentication

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

// NewValidateAuthenticationResultsParams creates a new ValidateAuthenticationResultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateAuthenticationResultsParams() *ValidateAuthenticationResultsParams {
	return &ValidateAuthenticationResultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateAuthenticationResultsParamsWithTimeout creates a new ValidateAuthenticationResultsParams object
// with the ability to set a timeout on a request.
func NewValidateAuthenticationResultsParamsWithTimeout(timeout time.Duration) *ValidateAuthenticationResultsParams {
	return &ValidateAuthenticationResultsParams{
		timeout: timeout,
	}
}

// NewValidateAuthenticationResultsParamsWithContext creates a new ValidateAuthenticationResultsParams object
// with the ability to set a context for a request.
func NewValidateAuthenticationResultsParamsWithContext(ctx context.Context) *ValidateAuthenticationResultsParams {
	return &ValidateAuthenticationResultsParams{
		Context: ctx,
	}
}

// NewValidateAuthenticationResultsParamsWithHTTPClient creates a new ValidateAuthenticationResultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateAuthenticationResultsParamsWithHTTPClient(client *http.Client) *ValidateAuthenticationResultsParams {
	return &ValidateAuthenticationResultsParams{
		HTTPClient: client,
	}
}

/* ValidateAuthenticationResultsParams contains all the parameters to send to the API endpoint
   for the validate authentication results operation.

   Typically these are written to a http.Request.
*/
type ValidateAuthenticationResultsParams struct {

	// ValidateRequest.
	ValidateRequest ValidateAuthenticationResultsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate authentication results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateAuthenticationResultsParams) WithDefaults() *ValidateAuthenticationResultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate authentication results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateAuthenticationResultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate authentication results params
func (o *ValidateAuthenticationResultsParams) WithTimeout(timeout time.Duration) *ValidateAuthenticationResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate authentication results params
func (o *ValidateAuthenticationResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate authentication results params
func (o *ValidateAuthenticationResultsParams) WithContext(ctx context.Context) *ValidateAuthenticationResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate authentication results params
func (o *ValidateAuthenticationResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate authentication results params
func (o *ValidateAuthenticationResultsParams) WithHTTPClient(client *http.Client) *ValidateAuthenticationResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate authentication results params
func (o *ValidateAuthenticationResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidateRequest adds the validateRequest to the validate authentication results params
func (o *ValidateAuthenticationResultsParams) WithValidateRequest(validateRequest ValidateAuthenticationResultsBody) *ValidateAuthenticationResultsParams {
	o.SetValidateRequest(validateRequest)
	return o
}

// SetValidateRequest adds the validateRequest to the validate authentication results params
func (o *ValidateAuthenticationResultsParams) SetValidateRequest(validateRequest ValidateAuthenticationResultsBody) {
	o.ValidateRequest = validateRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateAuthenticationResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.ValidateRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
