// Code generated by go-swagger; DO NOT EDIT.

package invoice_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new invoice settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for invoice settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetInvoiceSettings(params *GetInvoiceSettingsParams, opts ...ClientOption) (*GetInvoiceSettingsOK, error)

	UpdateInvoiceSettings(params *UpdateInvoiceSettingsParams, opts ...ClientOption) (*UpdateInvoiceSettingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetInvoiceSettings gets invoice settings

  Get the invoice settings for the invoice payment page.
*/
func (a *Client) GetInvoiceSettings(params *GetInvoiceSettingsParams, opts ...ClientOption) (*GetInvoiceSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvoiceSettings",
		Method:             "GET",
		PathPattern:        "/invoicing/v2/invoiceSettings",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvoiceSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoiceSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetInvoiceSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateInvoiceSettings updates invoice settings

  Update invoice settings for the invoice payment page.
*/
func (a *Client) UpdateInvoiceSettings(params *UpdateInvoiceSettingsParams, opts ...ClientOption) (*UpdateInvoiceSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInvoiceSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateInvoiceSettings",
		Method:             "PUT",
		PathPattern:        "/invoicing/v2/invoiceSettings",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateInvoiceSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateInvoiceSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateInvoiceSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
