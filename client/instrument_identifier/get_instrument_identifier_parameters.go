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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInstrumentIdentifierParams creates a new GetInstrumentIdentifierParams object
// with the default values initialized.
func NewGetInstrumentIdentifierParams() *GetInstrumentIdentifierParams {
	var ()
	return &GetInstrumentIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstrumentIdentifierParamsWithTimeout creates a new GetInstrumentIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstrumentIdentifierParamsWithTimeout(timeout time.Duration) *GetInstrumentIdentifierParams {
	var ()
	return &GetInstrumentIdentifierParams{

		timeout: timeout,
	}
}

// NewGetInstrumentIdentifierParamsWithContext creates a new GetInstrumentIdentifierParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstrumentIdentifierParamsWithContext(ctx context.Context) *GetInstrumentIdentifierParams {
	var ()
	return &GetInstrumentIdentifierParams{

		Context: ctx,
	}
}

// NewGetInstrumentIdentifierParamsWithHTTPClient creates a new GetInstrumentIdentifierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstrumentIdentifierParamsWithHTTPClient(client *http.Client) *GetInstrumentIdentifierParams {
	var ()
	return &GetInstrumentIdentifierParams{
		HTTPClient: client,
	}
}

/*GetInstrumentIdentifierParams contains all the parameters to send to the API endpoint
for the get instrument identifier operation typically these are written to a http.Request
*/
type GetInstrumentIdentifierParams struct {

	/*ProfileID
	  The id of a profile containing user specific TMS configuration.

	*/
	ProfileID string
	/*TokenID
	  The TokenId of an Instrument Identifier.

	*/
	TokenID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) WithTimeout(timeout time.Duration) *GetInstrumentIdentifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) WithContext(ctx context.Context) *GetInstrumentIdentifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) WithHTTPClient(client *http.Client) *GetInstrumentIdentifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProfileID adds the profileID to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) WithProfileID(profileID string) *GetInstrumentIdentifierParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WithTokenID adds the tokenID to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) WithTokenID(tokenID string) *GetInstrumentIdentifierParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the get instrument identifier params
func (o *GetInstrumentIdentifierParams) SetTokenID(tokenID string) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstrumentIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param profile-id
	if err := r.SetHeaderParam("profile-id", o.ProfileID); err != nil {
		return err
	}

	// path param tokenId
	if err := r.SetPathParam("tokenId", o.TokenID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}