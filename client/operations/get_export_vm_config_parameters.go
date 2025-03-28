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

// NewGetExportVMConfigParams creates a new GetExportVMConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExportVMConfigParams() *GetExportVMConfigParams {
	return &GetExportVMConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExportVMConfigParamsWithTimeout creates a new GetExportVMConfigParams object
// with the ability to set a timeout on a request.
func NewGetExportVMConfigParamsWithTimeout(timeout time.Duration) *GetExportVMConfigParams {
	return &GetExportVMConfigParams{
		timeout: timeout,
	}
}

// NewGetExportVMConfigParamsWithContext creates a new GetExportVMConfigParams object
// with the ability to set a context for a request.
func NewGetExportVMConfigParamsWithContext(ctx context.Context) *GetExportVMConfigParams {
	return &GetExportVMConfigParams{
		Context: ctx,
	}
}

// NewGetExportVMConfigParamsWithHTTPClient creates a new GetExportVMConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExportVMConfigParamsWithHTTPClient(client *http.Client) *GetExportVMConfigParams {
	return &GetExportVMConfigParams{
		HTTPClient: client,
	}
}

/*
GetExportVMConfigParams contains all the parameters to send to the API endpoint

	for the get export Vm config operation.

	Typically these are written to a http.Request.
*/
type GetExportVMConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get export Vm config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExportVMConfigParams) WithDefaults() *GetExportVMConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get export Vm config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExportVMConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get export Vm config params
func (o *GetExportVMConfigParams) WithTimeout(timeout time.Duration) *GetExportVMConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get export Vm config params
func (o *GetExportVMConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get export Vm config params
func (o *GetExportVMConfigParams) WithContext(ctx context.Context) *GetExportVMConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get export Vm config params
func (o *GetExportVMConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get export Vm config params
func (o *GetExportVMConfigParams) WithHTTPClient(client *http.Client) *GetExportVMConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get export Vm config params
func (o *GetExportVMConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetExportVMConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
