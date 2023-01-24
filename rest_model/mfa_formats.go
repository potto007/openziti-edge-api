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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MfaFormats mfa formats
//
// swagger:model mfaFormats
type MfaFormats string

func NewMfaFormats(value MfaFormats) *MfaFormats {
	v := value
	return &v
}

const (

	// MfaFormatsNumeric captures enum value "numeric"
	MfaFormatsNumeric MfaFormats = "numeric"

	// MfaFormatsAlpha captures enum value "alpha"
	MfaFormatsAlpha MfaFormats = "alpha"

	// MfaFormatsAlphaNumeric captures enum value "alphaNumeric"
	MfaFormatsAlphaNumeric MfaFormats = "alphaNumeric"
)

// for schema
var mfaFormatsEnum []interface{}

func init() {
	var res []MfaFormats
	if err := json.Unmarshal([]byte(`["numeric","alpha","alphaNumeric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mfaFormatsEnum = append(mfaFormatsEnum, v)
	}
}

func (m MfaFormats) validateMfaFormatsEnum(path, location string, value MfaFormats) error {
	if err := validate.EnumCase(path, location, value, mfaFormatsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this mfa formats
func (m MfaFormats) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMfaFormatsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this mfa formats based on context it is used
func (m MfaFormats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}