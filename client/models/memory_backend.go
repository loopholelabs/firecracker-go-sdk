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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MemoryBackend memory backend
//
// swagger:model MemoryBackend
type MemoryBackend struct {

	// Based on 'backend_type' it is either 1) Path to the file that contains the guest memory to be loaded 2) Path to the UDS where a process is listening for a UFFD initialization control payload and open file descriptor that it can use to serve this process's guest memory page faults
	// Required: true
	BackendPath *string `json:"backend_path"`

	// backend type
	// Required: true
	// Enum: ["File","Uffd"]
	BackendType *string `json:"backend_type"`
}

// Validate validates this memory backend
func (m *MemoryBackend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackendPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackendType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemoryBackend) validateBackendPath(formats strfmt.Registry) error {

	if err := validate.Required("backend_path", "body", m.BackendPath); err != nil {
		return err
	}

	return nil
}

var memoryBackendTypeBackendTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["File","Uffd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memoryBackendTypeBackendTypePropEnum = append(memoryBackendTypeBackendTypePropEnum, v)
	}
}

const (

	// MemoryBackendBackendTypeFile captures enum value "File"
	MemoryBackendBackendTypeFile string = "File"

	// MemoryBackendBackendTypeUffd captures enum value "Uffd"
	MemoryBackendBackendTypeUffd string = "Uffd"
)

// prop value enum
func (m *MemoryBackend) validateBackendTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, memoryBackendTypeBackendTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MemoryBackend) validateBackendType(formats strfmt.Registry) error {

	if err := validate.Required("backend_type", "body", m.BackendType); err != nil {
		return err
	}

	// value enum
	if err := m.validateBackendTypeEnum("backend_type", "body", *m.BackendType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this memory backend based on context it is used
func (m *MemoryBackend) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MemoryBackend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemoryBackend) UnmarshalBinary(b []byte) error {
	var res MemoryBackend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
