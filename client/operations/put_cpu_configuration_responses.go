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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

// PutCPUConfigurationReader is a Reader for the PutCPUConfiguration structure.
type PutCPUConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCPUConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCPUConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutCPUConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutCPUConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCPUConfigurationNoContent creates a PutCPUConfigurationNoContent with default headers values
func NewPutCPUConfigurationNoContent() *PutCPUConfigurationNoContent {
	return &PutCPUConfigurationNoContent{}
}

/*PutCPUConfigurationNoContent handles this case with default header values.

CPU configuration set successfully
*/
type PutCPUConfigurationNoContent struct {
}

func (o *PutCPUConfigurationNoContent) Error() string {
	return fmt.Sprintf("[PUT /cpu-config][%d] putCpuConfigurationNoContent ", 204)
}

func (o *PutCPUConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCPUConfigurationBadRequest creates a PutCPUConfigurationBadRequest with default headers values
func NewPutCPUConfigurationBadRequest() *PutCPUConfigurationBadRequest {
	return &PutCPUConfigurationBadRequest{}
}

/*PutCPUConfigurationBadRequest handles this case with default header values.

CPU configuration cannot be updated due to invalid input format
*/
type PutCPUConfigurationBadRequest struct {
	Payload *models.Error
}

func (o *PutCPUConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /cpu-config][%d] putCpuConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *PutCPUConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCPUConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCPUConfigurationDefault creates a PutCPUConfigurationDefault with default headers values
func NewPutCPUConfigurationDefault(code int) *PutCPUConfigurationDefault {
	return &PutCPUConfigurationDefault{
		_statusCode: code,
	}
}

/*PutCPUConfigurationDefault handles this case with default header values.

Internal server error
*/
type PutCPUConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put Cpu configuration default response
func (o *PutCPUConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *PutCPUConfigurationDefault) Error() string {
	return fmt.Sprintf("[PUT /cpu-config][%d] putCpuConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *PutCPUConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCPUConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
