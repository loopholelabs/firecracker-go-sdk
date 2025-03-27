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
)

// RateLimiter Defines an IO rate limiter with independent bytes/s and ops/s limits. Limits are defined by configuring each of the _bandwidth_ and _ops_ token buckets.
//
// swagger:model RateLimiter
type RateLimiter struct {

	// Token bucket with bytes as tokens
	Bandwidth *TokenBucket `json:"bandwidth,omitempty"`

	// Token bucket with operations as tokens
	Ops *TokenBucket `json:"ops,omitempty"`
}

// Validate validates this rate limiter
func (m *RateLimiter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBandwidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RateLimiter) validateBandwidth(formats strfmt.Registry) error {
	if swag.IsZero(m.Bandwidth) { // not required
		return nil
	}

	if m.Bandwidth != nil {
		if err := m.Bandwidth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bandwidth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bandwidth")
			}
			return err
		}
	}

	return nil
}

func (m *RateLimiter) validateOps(formats strfmt.Registry) error {
	if swag.IsZero(m.Ops) { // not required
		return nil
	}

	if m.Ops != nil {
		if err := m.Ops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ops")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ops")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rate limiter based on the context it is used
func (m *RateLimiter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBandwidth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RateLimiter) contextValidateBandwidth(ctx context.Context, formats strfmt.Registry) error {

	if m.Bandwidth != nil {

		if swag.IsZero(m.Bandwidth) { // not required
			return nil
		}

		if err := m.Bandwidth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bandwidth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bandwidth")
			}
			return err
		}
	}

	return nil
}

func (m *RateLimiter) contextValidateOps(ctx context.Context, formats strfmt.Registry) error {

	if m.Ops != nil {

		if swag.IsZero(m.Ops) { // not required
			return nil
		}

		if err := m.Ops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ops")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ops")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RateLimiter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateLimiter) UnmarshalBinary(b []byte) error {
	var res RateLimiter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
