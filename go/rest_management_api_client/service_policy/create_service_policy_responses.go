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

package service_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/go/rest_model"
)

// CreateServicePolicyReader is a Reader for the CreateServicePolicy structure.
type CreateServicePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServicePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServicePolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServicePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateServicePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateServicePolicyCreated creates a CreateServicePolicyCreated with default headers values
func NewCreateServicePolicyCreated() *CreateServicePolicyCreated {
	return &CreateServicePolicyCreated{}
}

/* CreateServicePolicyCreated describes a response with status code 201, with default header values.

The create request was successful and the resource has been added at the following location
*/
type CreateServicePolicyCreated struct {
	Payload *rest_model.CreateEnvelope
}

func (o *CreateServicePolicyCreated) Error() string {
	return fmt.Sprintf("[POST /service-policies][%d] createServicePolicyCreated  %+v", 201, o.Payload)
}
func (o *CreateServicePolicyCreated) GetPayload() *rest_model.CreateEnvelope {
	return o.Payload
}

func (o *CreateServicePolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServicePolicyBadRequest creates a CreateServicePolicyBadRequest with default headers values
func NewCreateServicePolicyBadRequest() *CreateServicePolicyBadRequest {
	return &CreateServicePolicyBadRequest{}
}

/* CreateServicePolicyBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type CreateServicePolicyBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateServicePolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /service-policies][%d] createServicePolicyBadRequest  %+v", 400, o.Payload)
}
func (o *CreateServicePolicyBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateServicePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServicePolicyUnauthorized creates a CreateServicePolicyUnauthorized with default headers values
func NewCreateServicePolicyUnauthorized() *CreateServicePolicyUnauthorized {
	return &CreateServicePolicyUnauthorized{}
}

/* CreateServicePolicyUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type CreateServicePolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateServicePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /service-policies][%d] createServicePolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateServicePolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateServicePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}