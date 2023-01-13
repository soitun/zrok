// Code generated by go-swagger; DO NOT EDIT.

package admin

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

	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
)

// NewCreateFrontendParams creates a new CreateFrontendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFrontendParams() *CreateFrontendParams {
	return &CreateFrontendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFrontendParamsWithTimeout creates a new CreateFrontendParams object
// with the ability to set a timeout on a request.
func NewCreateFrontendParamsWithTimeout(timeout time.Duration) *CreateFrontendParams {
	return &CreateFrontendParams{
		timeout: timeout,
	}
}

// NewCreateFrontendParamsWithContext creates a new CreateFrontendParams object
// with the ability to set a context for a request.
func NewCreateFrontendParamsWithContext(ctx context.Context) *CreateFrontendParams {
	return &CreateFrontendParams{
		Context: ctx,
	}
}

// NewCreateFrontendParamsWithHTTPClient creates a new CreateFrontendParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFrontendParamsWithHTTPClient(client *http.Client) *CreateFrontendParams {
	return &CreateFrontendParams{
		HTTPClient: client,
	}
}

/*
CreateFrontendParams contains all the parameters to send to the API endpoint

	for the create frontend operation.

	Typically these are written to a http.Request.
*/
type CreateFrontendParams struct {

	// Body.
	Body *rest_model_zrok.CreateFrontendRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create frontend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFrontendParams) WithDefaults() *CreateFrontendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create frontend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFrontendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create frontend params
func (o *CreateFrontendParams) WithTimeout(timeout time.Duration) *CreateFrontendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create frontend params
func (o *CreateFrontendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create frontend params
func (o *CreateFrontendParams) WithContext(ctx context.Context) *CreateFrontendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create frontend params
func (o *CreateFrontendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create frontend params
func (o *CreateFrontendParams) WithHTTPClient(client *http.Client) *CreateFrontendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create frontend params
func (o *CreateFrontendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create frontend params
func (o *CreateFrontendParams) WithBody(body *rest_model_zrok.CreateFrontendRequest) *CreateFrontendParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create frontend params
func (o *CreateFrontendParams) SetBody(body *rest_model_zrok.CreateFrontendRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFrontendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
