// Code generated by go-swagger; DO NOT EDIT.

package void

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new void API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for void API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	MitVoid(params *MitVoidParams, opts ...ClientOption) (*MitVoidCreated, error)

	VoidCapture(params *VoidCaptureParams, opts ...ClientOption) (*VoidCaptureCreated, error)

	VoidCredit(params *VoidCreditParams, opts ...ClientOption) (*VoidCreditCreated, error)

	VoidPayment(params *VoidPaymentParams, opts ...ClientOption) (*VoidPaymentCreated, error)

	VoidRefund(params *VoidRefundParams, opts ...ClientOption) (*VoidRefundCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  MitVoid timeouts void

  This is to void a previous payment, capture, refund, or credit that merchant does not receive a reply(Mostly due to timeout). This is to void a previous payment, capture, refund, or credit that merchant does not receive a reply(Mostly due to Timeout). To use this feature/API, make sure to pass unique value to field - clientReferenceInformation -> transactionId in your payment, capture, refund, or credit API call and use same transactionId in this API request payload to reverse the payment.
*/
func (a *Client) MitVoid(params *MitVoidParams, opts ...ClientOption) (*MitVoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMitVoidParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "mitVoid",
		Method:             "POST",
		PathPattern:        "/pts/v2/voids/",
		ProducesMediaTypes: []string{"application/hal+json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MitVoidReader{formats: a.formats},
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
	success, ok := result.(*MitVoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for mitVoid: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VoidCapture voids a capture

  Refund a capture API is only used, if you have requested Capture independenlty using [/pts/v2/payments/{id}/captures](https://developer.cybersource.com/api-reference-assets/index.html#payments_capture) API call. Include the capture ID in the POST request to cancel the capture.

*/
func (a *Client) VoidCapture(params *VoidCaptureParams, opts ...ClientOption) (*VoidCaptureCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoidCaptureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "voidCapture",
		Method:             "POST",
		PathPattern:        "/pts/v2/captures/{id}/voids",
		ProducesMediaTypes: []string{"application/hal+json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoidCaptureReader{formats: a.formats},
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
	success, ok := result.(*VoidCaptureCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for voidCapture: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VoidCredit voids a credit

  Include the credit ID in the POST request to cancel the credit.
*/
func (a *Client) VoidCredit(params *VoidCreditParams, opts ...ClientOption) (*VoidCreditCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoidCreditParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "voidCredit",
		Method:             "POST",
		PathPattern:        "/pts/v2/credits/{id}/voids",
		ProducesMediaTypes: []string{"application/hal+json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoidCreditReader{formats: a.formats},
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
	success, ok := result.(*VoidCreditCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for voidCredit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VoidPayment voids a payment

  Void a Payment API is only used, if you have requested Authorization and Capture together in [/pts/v2/payments](https://developer.cybersource.com/api-reference-assets/index.html#payments_payments) API call. Include the payment ID in the POST request to cancel the payment.

*/
func (a *Client) VoidPayment(params *VoidPaymentParams, opts ...ClientOption) (*VoidPaymentCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoidPaymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "voidPayment",
		Method:             "POST",
		PathPattern:        "/pts/v2/payments/{id}/voids",
		ProducesMediaTypes: []string{"application/hal+json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoidPaymentReader{formats: a.formats},
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
	success, ok := result.(*VoidPaymentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for voidPayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VoidRefund voids a refund

  Include the refund ID in the POST request to cancel the refund.
*/
func (a *Client) VoidRefund(params *VoidRefundParams, opts ...ClientOption) (*VoidRefundCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoidRefundParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "voidRefund",
		Method:             "POST",
		PathPattern:        "/pts/v2/refunds/{id}/voids",
		ProducesMediaTypes: []string{"application/hal+json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoidRefundReader{formats: a.formats},
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
	success, ok := result.(*VoidRefundCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for voidRefund: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
