// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

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

// NewGetFirecrackerVersionParams creates a new GetFirecrackerVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFirecrackerVersionParams() *GetFirecrackerVersionParams {
	return &GetFirecrackerVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFirecrackerVersionParamsWithTimeout creates a new GetFirecrackerVersionParams object
// with the ability to set a timeout on a request.
func NewGetFirecrackerVersionParamsWithTimeout(timeout time.Duration) *GetFirecrackerVersionParams {
	return &GetFirecrackerVersionParams{
		timeout: timeout,
	}
}

// NewGetFirecrackerVersionParamsWithContext creates a new GetFirecrackerVersionParams object
// with the ability to set a context for a request.
func NewGetFirecrackerVersionParamsWithContext(ctx context.Context) *GetFirecrackerVersionParams {
	return &GetFirecrackerVersionParams{
		Context: ctx,
	}
}

// NewGetFirecrackerVersionParamsWithHTTPClient creates a new GetFirecrackerVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFirecrackerVersionParamsWithHTTPClient(client *http.Client) *GetFirecrackerVersionParams {
	return &GetFirecrackerVersionParams{
		HTTPClient: client,
	}
}

/*
GetFirecrackerVersionParams contains all the parameters to send to the API endpoint

	for the get firecracker version operation.

	Typically these are written to a http.Request.
*/
type GetFirecrackerVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get firecracker version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFirecrackerVersionParams) WithDefaults() *GetFirecrackerVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get firecracker version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFirecrackerVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get firecracker version params
func (o *GetFirecrackerVersionParams) WithTimeout(timeout time.Duration) *GetFirecrackerVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get firecracker version params
func (o *GetFirecrackerVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get firecracker version params
func (o *GetFirecrackerVersionParams) WithContext(ctx context.Context) *GetFirecrackerVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get firecracker version params
func (o *GetFirecrackerVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get firecracker version params
func (o *GetFirecrackerVersionParams) WithHTTPClient(client *http.Client) *GetFirecrackerVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get firecracker version params
func (o *GetFirecrackerVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFirecrackerVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
