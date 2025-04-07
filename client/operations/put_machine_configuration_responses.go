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

	"github.com/loopholelabs/firecracker-go-sdk/client/models"
)

// PutMachineConfigurationReader is a Reader for the PutMachineConfiguration structure.
type PutMachineConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMachineConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutMachineConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutMachineConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutMachineConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMachineConfigurationNoContent creates a PutMachineConfigurationNoContent with default headers values
func NewPutMachineConfigurationNoContent() *PutMachineConfigurationNoContent {
	return &PutMachineConfigurationNoContent{}
}

/*
PutMachineConfigurationNoContent describes a response with status code 204, with default header values.

Machine Configuration created/updated
*/
type PutMachineConfigurationNoContent struct {
}

// IsSuccess returns true when this put machine configuration no content response has a 2xx status code
func (o *PutMachineConfigurationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put machine configuration no content response has a 3xx status code
func (o *PutMachineConfigurationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put machine configuration no content response has a 4xx status code
func (o *PutMachineConfigurationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put machine configuration no content response has a 5xx status code
func (o *PutMachineConfigurationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put machine configuration no content response a status code equal to that given
func (o *PutMachineConfigurationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put machine configuration no content response
func (o *PutMachineConfigurationNoContent) Code() int {
	return 204
}

func (o *PutMachineConfigurationNoContent) Error() string {
	return fmt.Sprintf("[PUT /machine-config][%d] putMachineConfigurationNoContent", 204)
}

func (o *PutMachineConfigurationNoContent) String() string {
	return fmt.Sprintf("[PUT /machine-config][%d] putMachineConfigurationNoContent", 204)
}

func (o *PutMachineConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMachineConfigurationBadRequest creates a PutMachineConfigurationBadRequest with default headers values
func NewPutMachineConfigurationBadRequest() *PutMachineConfigurationBadRequest {
	return &PutMachineConfigurationBadRequest{}
}

/*
PutMachineConfigurationBadRequest describes a response with status code 400, with default header values.

Machine Configuration cannot be updated due to bad input
*/
type PutMachineConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this put machine configuration bad request response has a 2xx status code
func (o *PutMachineConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put machine configuration bad request response has a 3xx status code
func (o *PutMachineConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put machine configuration bad request response has a 4xx status code
func (o *PutMachineConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put machine configuration bad request response has a 5xx status code
func (o *PutMachineConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put machine configuration bad request response a status code equal to that given
func (o *PutMachineConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put machine configuration bad request response
func (o *PutMachineConfigurationBadRequest) Code() int {
	return 400
}

func (o *PutMachineConfigurationBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /machine-config][%d] putMachineConfigurationBadRequest %s", 400, payload)
}

func (o *PutMachineConfigurationBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /machine-config][%d] putMachineConfigurationBadRequest %s", 400, payload)
}

func (o *PutMachineConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutMachineConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutMachineConfigurationDefault creates a PutMachineConfigurationDefault with default headers values
func NewPutMachineConfigurationDefault(code int) *PutMachineConfigurationDefault {
	return &PutMachineConfigurationDefault{
		_statusCode: code,
	}
}

/*
PutMachineConfigurationDefault describes a response with status code -1, with default header values.

Internal server error
*/
type PutMachineConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this put machine configuration default response has a 2xx status code
func (o *PutMachineConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put machine configuration default response has a 3xx status code
func (o *PutMachineConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put machine configuration default response has a 4xx status code
func (o *PutMachineConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put machine configuration default response has a 5xx status code
func (o *PutMachineConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put machine configuration default response a status code equal to that given
func (o *PutMachineConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put machine configuration default response
func (o *PutMachineConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *PutMachineConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /machine-config][%d] putMachineConfiguration default %s", o._statusCode, payload)
}

func (o *PutMachineConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /machine-config][%d] putMachineConfiguration default %s", o._statusCode, payload)
}

func (o *PutMachineConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutMachineConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
