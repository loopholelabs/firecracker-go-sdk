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

// PutGuestVsockReader is a Reader for the PutGuestVsock structure.
type PutGuestVsockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGuestVsockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutGuestVsockNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGuestVsockBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutGuestVsockDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutGuestVsockNoContent creates a PutGuestVsockNoContent with default headers values
func NewPutGuestVsockNoContent() *PutGuestVsockNoContent {
	return &PutGuestVsockNoContent{}
}

/*
PutGuestVsockNoContent describes a response with status code 204, with default header values.

Vsock created/updated
*/
type PutGuestVsockNoContent struct {
}

// IsSuccess returns true when this put guest vsock no content response has a 2xx status code
func (o *PutGuestVsockNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put guest vsock no content response has a 3xx status code
func (o *PutGuestVsockNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put guest vsock no content response has a 4xx status code
func (o *PutGuestVsockNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put guest vsock no content response has a 5xx status code
func (o *PutGuestVsockNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put guest vsock no content response a status code equal to that given
func (o *PutGuestVsockNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put guest vsock no content response
func (o *PutGuestVsockNoContent) Code() int {
	return 204
}

func (o *PutGuestVsockNoContent) Error() string {
	return fmt.Sprintf("[PUT /vsock][%d] putGuestVsockNoContent", 204)
}

func (o *PutGuestVsockNoContent) String() string {
	return fmt.Sprintf("[PUT /vsock][%d] putGuestVsockNoContent", 204)
}

func (o *PutGuestVsockNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutGuestVsockBadRequest creates a PutGuestVsockBadRequest with default headers values
func NewPutGuestVsockBadRequest() *PutGuestVsockBadRequest {
	return &PutGuestVsockBadRequest{}
}

/*
PutGuestVsockBadRequest describes a response with status code 400, with default header values.

Vsock cannot be created due to bad input
*/
type PutGuestVsockBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this put guest vsock bad request response has a 2xx status code
func (o *PutGuestVsockBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put guest vsock bad request response has a 3xx status code
func (o *PutGuestVsockBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put guest vsock bad request response has a 4xx status code
func (o *PutGuestVsockBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put guest vsock bad request response has a 5xx status code
func (o *PutGuestVsockBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put guest vsock bad request response a status code equal to that given
func (o *PutGuestVsockBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put guest vsock bad request response
func (o *PutGuestVsockBadRequest) Code() int {
	return 400
}

func (o *PutGuestVsockBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vsock][%d] putGuestVsockBadRequest %s", 400, payload)
}

func (o *PutGuestVsockBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vsock][%d] putGuestVsockBadRequest %s", 400, payload)
}

func (o *PutGuestVsockBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutGuestVsockBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGuestVsockDefault creates a PutGuestVsockDefault with default headers values
func NewPutGuestVsockDefault(code int) *PutGuestVsockDefault {
	return &PutGuestVsockDefault{
		_statusCode: code,
	}
}

/*
PutGuestVsockDefault describes a response with status code -1, with default header values.

Internal server error
*/
type PutGuestVsockDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this put guest vsock default response has a 2xx status code
func (o *PutGuestVsockDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put guest vsock default response has a 3xx status code
func (o *PutGuestVsockDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put guest vsock default response has a 4xx status code
func (o *PutGuestVsockDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put guest vsock default response has a 5xx status code
func (o *PutGuestVsockDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put guest vsock default response a status code equal to that given
func (o *PutGuestVsockDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put guest vsock default response
func (o *PutGuestVsockDefault) Code() int {
	return o._statusCode
}

func (o *PutGuestVsockDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vsock][%d] putGuestVsock default %s", o._statusCode, payload)
}

func (o *PutGuestVsockDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vsock][%d] putGuestVsock default %s", o._statusCode, payload)
}

func (o *PutGuestVsockDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutGuestVsockDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
