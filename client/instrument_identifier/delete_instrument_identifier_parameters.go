// Code generated by go-swagger; DO NOT EDIT.

package instrument_identifier

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

// NewDeleteInstrumentIdentifierParams creates a new DeleteInstrumentIdentifierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteInstrumentIdentifierParams() *DeleteInstrumentIdentifierParams {
	return &DeleteInstrumentIdentifierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstrumentIdentifierParamsWithTimeout creates a new DeleteInstrumentIdentifierParams object
// with the ability to set a timeout on a request.
func NewDeleteInstrumentIdentifierParamsWithTimeout(timeout time.Duration) *DeleteInstrumentIdentifierParams {
	return &DeleteInstrumentIdentifierParams{
		timeout: timeout,
	}
}

// NewDeleteInstrumentIdentifierParamsWithContext creates a new DeleteInstrumentIdentifierParams object
// with the ability to set a context for a request.
func NewDeleteInstrumentIdentifierParamsWithContext(ctx context.Context) *DeleteInstrumentIdentifierParams {
	return &DeleteInstrumentIdentifierParams{
		Context: ctx,
	}
}

// NewDeleteInstrumentIdentifierParamsWithHTTPClient creates a new DeleteInstrumentIdentifierParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteInstrumentIdentifierParamsWithHTTPClient(client *http.Client) *DeleteInstrumentIdentifierParams {
	return &DeleteInstrumentIdentifierParams{
		HTTPClient: client,
	}
}

/* DeleteInstrumentIdentifierParams contains all the parameters to send to the API endpoint
   for the delete instrument identifier operation.

   Typically these are written to a http.Request.
*/
type DeleteInstrumentIdentifierParams struct {

	/* InstrumentIdentifierTokenID.

	   The TokenId of a Instrument Identifier.
	*/
	InstrumentIdentifierTokenID string

	/* ProfileID.

	   The id of a profile containing user specific TMS configuration.
	*/
	ProfileID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete instrument identifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstrumentIdentifierParams) WithDefaults() *DeleteInstrumentIdentifierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete instrument identifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstrumentIdentifierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) WithTimeout(timeout time.Duration) *DeleteInstrumentIdentifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) WithContext(ctx context.Context) *DeleteInstrumentIdentifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) WithHTTPClient(client *http.Client) *DeleteInstrumentIdentifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstrumentIdentifierTokenID adds the instrumentIdentifierTokenID to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) WithInstrumentIdentifierTokenID(instrumentIdentifierTokenID string) *DeleteInstrumentIdentifierParams {
	o.SetInstrumentIdentifierTokenID(instrumentIdentifierTokenID)
	return o
}

// SetInstrumentIdentifierTokenID adds the instrumentIdentifierTokenId to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) SetInstrumentIdentifierTokenID(instrumentIdentifierTokenID string) {
	o.InstrumentIdentifierTokenID = instrumentIdentifierTokenID
}

// WithProfileID adds the profileID to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) WithProfileID(profileID *string) *DeleteInstrumentIdentifierParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the delete instrument identifier params
func (o *DeleteInstrumentIdentifierParams) SetProfileID(profileID *string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstrumentIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instrumentIdentifierTokenId
	if err := r.SetPathParam("instrumentIdentifierTokenId", o.InstrumentIdentifierTokenID); err != nil {
		return err
	}

	if o.ProfileID != nil {

		// header param profile-id
		if err := r.SetHeaderParam("profile-id", *o.ProfileID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
