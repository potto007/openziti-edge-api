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

// DialBind dial bind
//
// swagger:model dialBind
type DialBind string

func NewDialBind(value DialBind) *DialBind {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DialBind.
func (m DialBind) Pointer() *DialBind {
	return &m
}

const (

	// DialBindDial captures enum value "Dial"
	DialBindDial DialBind = "Dial"

	// DialBindBind captures enum value "Bind"
	DialBindBind DialBind = "Bind"
)

// for schema
var dialBindEnum []interface{}

func init() {
	var res []DialBind
	if err := json.Unmarshal([]byte(`["Dial","Bind"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dialBindEnum = append(dialBindEnum, v)
	}
}

func (m DialBind) validateDialBindEnum(path, location string, value DialBind) error {
	if err := validate.EnumCase(path, location, value, dialBindEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this dial bind
func (m DialBind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDialBindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this dial bind based on context it is used
func (m DialBind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
