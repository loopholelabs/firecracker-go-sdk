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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BalloonStats Describes the balloon device statistics.
//
// swagger:model BalloonStats
type BalloonStats struct {

	// Actual amount of memory (in MiB) the device is holding.
	// Required: true
	ActualMib *int64 `json:"actual_mib"`

	// Actual number of pages the device is holding.
	// Required: true
	ActualPages *int64 `json:"actual_pages"`

	// An estimate of how much memory is available (in bytes) for starting new applications, without pushing the system to swap.
	AvailableMemory int64 `json:"available_memory,omitempty"`

	// The amount of memory, in bytes, that can be quickly reclaimed without additional I/O. Typically these pages are used for caching files from disk.
	DiskCaches int64 `json:"disk_caches,omitempty"`

	// The amount of memory not being used for any purpose (in bytes).
	FreeMemory int64 `json:"free_memory,omitempty"`

	// The number of successful hugetlb page allocations in the guest.
	HugetlbAllocations int64 `json:"hugetlb_allocations,omitempty"`

	// The number of failed hugetlb page allocations in the guest.
	HugetlbFailures int64 `json:"hugetlb_failures,omitempty"`

	// The number of major page faults that have occurred.
	MajorFaults int64 `json:"major_faults,omitempty"`

	// The number of minor page faults that have occurred.
	MinorFaults int64 `json:"minor_faults,omitempty"`

	// The amount of memory that has been swapped in (in bytes).
	SwapIn int64 `json:"swap_in,omitempty"`

	// The amount of memory that has been swapped out to disk (in bytes).
	SwapOut int64 `json:"swap_out,omitempty"`

	// Target amount of memory (in MiB) the device aims to hold.
	// Required: true
	TargetMib *int64 `json:"target_mib"`

	// Target number of pages the device aims to hold.
	// Required: true
	TargetPages *int64 `json:"target_pages"`

	// The total amount of memory available (in bytes).
	TotalMemory int64 `json:"total_memory,omitempty"`
}

// Validate validates this balloon stats
func (m *BalloonStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualMib(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActualPages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetMib(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetPages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BalloonStats) validateActualMib(formats strfmt.Registry) error {

	if err := validate.Required("actual_mib", "body", m.ActualMib); err != nil {
		return err
	}

	return nil
}

func (m *BalloonStats) validateActualPages(formats strfmt.Registry) error {

	if err := validate.Required("actual_pages", "body", m.ActualPages); err != nil {
		return err
	}

	return nil
}

func (m *BalloonStats) validateTargetMib(formats strfmt.Registry) error {

	if err := validate.Required("target_mib", "body", m.TargetMib); err != nil {
		return err
	}

	return nil
}

func (m *BalloonStats) validateTargetPages(formats strfmt.Registry) error {

	if err := validate.Required("target_pages", "body", m.TargetPages); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this balloon stats based on context it is used
func (m *BalloonStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BalloonStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BalloonStats) UnmarshalBinary(b []byte) error {
	var res BalloonStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
