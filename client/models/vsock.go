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

// Vsock Defines a vsock device, backed by a set of Unix Domain Sockets, on the host side. For host-initiated connections, Firecracker will be listening on the Unix socket identified by the path `uds_path`. Firecracker will create this socket, bind and listen on it. Host-initiated connections will be performed by connection to this socket and issuing a connection forwarding request to the desired guest-side vsock port (i.e. `CONNECT 52\n`, to connect to port 52). For guest-initiated connections, Firecracker will expect host software to be bound and listening on Unix sockets at `uds_path_<PORT>`. E.g. "/path/to/host_vsock.sock_52" for port number 52.
//
// swagger:model Vsock
type Vsock struct {

	// Guest Vsock CID
	// Required: true
	// Minimum: 3
	GuestCid *int64 `json:"guest_cid"`

	// Path to UNIX domain socket, used to proxy vsock connections.
	// Required: true
	UdsPath *string `json:"uds_path"`

	// This parameter has been deprecated since v1.0.0.
	VsockID string `json:"vsock_id,omitempty"`
}

// Validate validates this vsock
func (m *Vsock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGuestCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUdsPath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vsock) validateGuestCid(formats strfmt.Registry) error {

	if err := validate.Required("guest_cid", "body", m.GuestCid); err != nil {
		return err
	}

	if err := validate.MinimumInt("guest_cid", "body", *m.GuestCid, 3, false); err != nil {
		return err
	}

	return nil
}

func (m *Vsock) validateUdsPath(formats strfmt.Registry) error {

	if err := validate.Required("uds_path", "body", m.UdsPath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vsock based on context it is used
func (m *Vsock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Vsock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vsock) UnmarshalBinary(b []byte) error {
	var res Vsock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
