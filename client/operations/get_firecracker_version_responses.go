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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

// GetFirecrackerVersionReader is a Reader for the GetFirecrackerVersion structure.
type GetFirecrackerVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirecrackerVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFirecrackerVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFirecrackerVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFirecrackerVersionOK creates a GetFirecrackerVersionOK with default headers values
func NewGetFirecrackerVersionOK() *GetFirecrackerVersionOK {
	return &GetFirecrackerVersionOK{}
}

/*
GetFirecrackerVersionOK describes a response with status code 200, with default header values.

OK
*/
type GetFirecrackerVersionOK struct {
	Payload *models.FirecrackerVersion
}

// IsSuccess returns true when this get firecracker version o k response has a 2xx status code
func (o *GetFirecrackerVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get firecracker version o k response has a 3xx status code
func (o *GetFirecrackerVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get firecracker version o k response has a 4xx status code
func (o *GetFirecrackerVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get firecracker version o k response has a 5xx status code
func (o *GetFirecrackerVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get firecracker version o k response a status code equal to that given
func (o *GetFirecrackerVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get firecracker version o k response
func (o *GetFirecrackerVersionOK) Code() int {
	return 200
}

func (o *GetFirecrackerVersionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /version][%d] getFirecrackerVersionOK %s", 200, payload)
}

func (o *GetFirecrackerVersionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /version][%d] getFirecrackerVersionOK %s", 200, payload)
}

func (o *GetFirecrackerVersionOK) GetPayload() *models.FirecrackerVersion {
	return o.Payload
}

func (o *GetFirecrackerVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirecrackerVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirecrackerVersionDefault creates a GetFirecrackerVersionDefault with default headers values
func NewGetFirecrackerVersionDefault(code int) *GetFirecrackerVersionDefault {
	return &GetFirecrackerVersionDefault{
		_statusCode: code,
	}
}

/*
GetFirecrackerVersionDefault describes a response with status code -1, with default header values.

Internal server error
*/
type GetFirecrackerVersionDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get firecracker version default response has a 2xx status code
func (o *GetFirecrackerVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get firecracker version default response has a 3xx status code
func (o *GetFirecrackerVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get firecracker version default response has a 4xx status code
func (o *GetFirecrackerVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get firecracker version default response has a 5xx status code
func (o *GetFirecrackerVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get firecracker version default response a status code equal to that given
func (o *GetFirecrackerVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get firecracker version default response
func (o *GetFirecrackerVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetFirecrackerVersionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /version][%d] getFirecrackerVersion default %s", o._statusCode, payload)
}

func (o *GetFirecrackerVersionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /version][%d] getFirecrackerVersion default %s", o._statusCode, payload)
}

func (o *GetFirecrackerVersionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFirecrackerVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
