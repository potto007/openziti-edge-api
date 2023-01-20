// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Version version
//
// swagger:model version
type Version struct {

	// api versions
	APIVersions map[string]map[string]APIVersion `json:"apiVersions,omitempty"`

	// build date
	// Example: 2020-02-11 16:09:08
	BuildDate string `json:"buildDate,omitempty"`

	// revision
	// Example: ea556fc18740
	Revision string `json:"revision,omitempty"`

	// runtime version
	// Example: go1.13.5
	RuntimeVersion string `json:"runtimeVersion,omitempty"`

	// version
	// Example: v0.9.0
	Version string `json:"version,omitempty"`
}

// Validate validates this version
func (m *Version) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Version) validateAPIVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.APIVersions) { // not required
		return nil
	}

	for k := range m.APIVersions {

		for kk := range m.APIVersions[k] {

			if err := validate.Required("apiVersions"+"."+k+"."+kk, "body", m.APIVersions[k][kk]); err != nil {
				return err
			}
			if val, ok := m.APIVersions[k][kk]; ok {
				if err := val.Validate(formats); err != nil {
					return err
				}
			}

		}

	}

	return nil
}

// ContextValidate validate this version based on the context it is used
func (m *Version) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Version) contextValidateAPIVersions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.APIVersions {

		for kk := range m.APIVersions[k] {

			if val, ok := m.APIVersions[k][kk]; ok {
				if err := val.ContextValidate(ctx, formats); err != nil {
					return err
				}
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Version) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Version) UnmarshalBinary(b []byte) error {
	var res Version
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}