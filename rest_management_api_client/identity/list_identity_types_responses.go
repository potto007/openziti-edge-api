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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListIdentityTypesReader is a Reader for the ListIdentityTypes structure.
type ListIdentityTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIdentityTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIdentityTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListIdentityTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListIdentityTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListIdentityTypesOK creates a ListIdentityTypesOK with default headers values
func NewListIdentityTypesOK() *ListIdentityTypesOK {
	return &ListIdentityTypesOK{}
}

/* ListIdentityTypesOK describes a response with status code 200, with default header values.

A list of identity types
*/
type ListIdentityTypesOK struct {
	Payload *rest_model.ListIdentityTypesEnvelope
}

func (o *ListIdentityTypesOK) Error() string {
	return fmt.Sprintf("[GET /identity-types][%d] listIdentityTypesOK  %+v", 200, o.Payload)
}
func (o *ListIdentityTypesOK) GetPayload() *rest_model.ListIdentityTypesEnvelope {
	return o.Payload
}

func (o *ListIdentityTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListIdentityTypesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentityTypesBadRequest creates a ListIdentityTypesBadRequest with default headers values
func NewListIdentityTypesBadRequest() *ListIdentityTypesBadRequest {
	return &ListIdentityTypesBadRequest{}
}

/* ListIdentityTypesBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListIdentityTypesBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentityTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /identity-types][%d] listIdentityTypesBadRequest  %+v", 400, o.Payload)
}
func (o *ListIdentityTypesBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentityTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentityTypesUnauthorized creates a ListIdentityTypesUnauthorized with default headers values
func NewListIdentityTypesUnauthorized() *ListIdentityTypesUnauthorized {
	return &ListIdentityTypesUnauthorized{}
}

/* ListIdentityTypesUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListIdentityTypesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentityTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /identity-types][%d] listIdentityTypesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListIdentityTypesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentityTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}