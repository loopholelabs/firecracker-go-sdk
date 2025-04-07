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

// PatchBalloonStatsIntervalReader is a Reader for the PatchBalloonStatsInterval structure.
type PatchBalloonStatsIntervalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchBalloonStatsIntervalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchBalloonStatsIntervalNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchBalloonStatsIntervalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchBalloonStatsIntervalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchBalloonStatsIntervalNoContent creates a PatchBalloonStatsIntervalNoContent with default headers values
func NewPatchBalloonStatsIntervalNoContent() *PatchBalloonStatsIntervalNoContent {
	return &PatchBalloonStatsIntervalNoContent{}
}

/*
PatchBalloonStatsIntervalNoContent describes a response with status code 204, with default header values.

Balloon statistics interval updated
*/
type PatchBalloonStatsIntervalNoContent struct {
}

// IsSuccess returns true when this patch balloon stats interval no content response has a 2xx status code
func (o *PatchBalloonStatsIntervalNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch balloon stats interval no content response has a 3xx status code
func (o *PatchBalloonStatsIntervalNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch balloon stats interval no content response has a 4xx status code
func (o *PatchBalloonStatsIntervalNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch balloon stats interval no content response has a 5xx status code
func (o *PatchBalloonStatsIntervalNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch balloon stats interval no content response a status code equal to that given
func (o *PatchBalloonStatsIntervalNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the patch balloon stats interval no content response
func (o *PatchBalloonStatsIntervalNoContent) Code() int {
	return 204
}

func (o *PatchBalloonStatsIntervalNoContent) Error() string {
	return fmt.Sprintf("[PATCH /balloon/statistics][%d] patchBalloonStatsIntervalNoContent", 204)
}

func (o *PatchBalloonStatsIntervalNoContent) String() string {
	return fmt.Sprintf("[PATCH /balloon/statistics][%d] patchBalloonStatsIntervalNoContent", 204)
}

func (o *PatchBalloonStatsIntervalNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchBalloonStatsIntervalBadRequest creates a PatchBalloonStatsIntervalBadRequest with default headers values
func NewPatchBalloonStatsIntervalBadRequest() *PatchBalloonStatsIntervalBadRequest {
	return &PatchBalloonStatsIntervalBadRequest{}
}

/*
PatchBalloonStatsIntervalBadRequest describes a response with status code 400, with default header values.

Balloon statistics interval cannot be updated due to bad input
*/
type PatchBalloonStatsIntervalBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch balloon stats interval bad request response has a 2xx status code
func (o *PatchBalloonStatsIntervalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch balloon stats interval bad request response has a 3xx status code
func (o *PatchBalloonStatsIntervalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch balloon stats interval bad request response has a 4xx status code
func (o *PatchBalloonStatsIntervalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch balloon stats interval bad request response has a 5xx status code
func (o *PatchBalloonStatsIntervalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch balloon stats interval bad request response a status code equal to that given
func (o *PatchBalloonStatsIntervalBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch balloon stats interval bad request response
func (o *PatchBalloonStatsIntervalBadRequest) Code() int {
	return 400
}

func (o *PatchBalloonStatsIntervalBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /balloon/statistics][%d] patchBalloonStatsIntervalBadRequest %s", 400, payload)
}

func (o *PatchBalloonStatsIntervalBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /balloon/statistics][%d] patchBalloonStatsIntervalBadRequest %s", 400, payload)
}

func (o *PatchBalloonStatsIntervalBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchBalloonStatsIntervalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBalloonStatsIntervalDefault creates a PatchBalloonStatsIntervalDefault with default headers values
func NewPatchBalloonStatsIntervalDefault(code int) *PatchBalloonStatsIntervalDefault {
	return &PatchBalloonStatsIntervalDefault{
		_statusCode: code,
	}
}

/*
PatchBalloonStatsIntervalDefault describes a response with status code -1, with default header values.

Internal server error
*/
type PatchBalloonStatsIntervalDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this patch balloon stats interval default response has a 2xx status code
func (o *PatchBalloonStatsIntervalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch balloon stats interval default response has a 3xx status code
func (o *PatchBalloonStatsIntervalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch balloon stats interval default response has a 4xx status code
func (o *PatchBalloonStatsIntervalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch balloon stats interval default response has a 5xx status code
func (o *PatchBalloonStatsIntervalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch balloon stats interval default response a status code equal to that given
func (o *PatchBalloonStatsIntervalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch balloon stats interval default response
func (o *PatchBalloonStatsIntervalDefault) Code() int {
	return o._statusCode
}

func (o *PatchBalloonStatsIntervalDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /balloon/statistics][%d] patchBalloonStatsInterval default %s", o._statusCode, payload)
}

func (o *PatchBalloonStatsIntervalDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /balloon/statistics][%d] patchBalloonStatsInterval default %s", o._statusCode, payload)
}

func (o *PatchBalloonStatsIntervalDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchBalloonStatsIntervalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
