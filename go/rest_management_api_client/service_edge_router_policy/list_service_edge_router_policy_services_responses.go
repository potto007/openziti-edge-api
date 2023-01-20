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

package service_edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/go/rest_model"
)

// ListServiceEdgeRouterPolicyServicesReader is a Reader for the ListServiceEdgeRouterPolicyServices structure.
type ListServiceEdgeRouterPolicyServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceEdgeRouterPolicyServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceEdgeRouterPolicyServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListServiceEdgeRouterPolicyServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListServiceEdgeRouterPolicyServicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServiceEdgeRouterPolicyServicesOK creates a ListServiceEdgeRouterPolicyServicesOK with default headers values
func NewListServiceEdgeRouterPolicyServicesOK() *ListServiceEdgeRouterPolicyServicesOK {
	return &ListServiceEdgeRouterPolicyServicesOK{}
}

/* ListServiceEdgeRouterPolicyServicesOK describes a response with status code 200, with default header values.

A list of services
*/
type ListServiceEdgeRouterPolicyServicesOK struct {
	Payload *rest_model.ListServicesEnvelope
}

func (o *ListServiceEdgeRouterPolicyServicesOK) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}/services][%d] listServiceEdgeRouterPolicyServicesOK  %+v", 200, o.Payload)
}
func (o *ListServiceEdgeRouterPolicyServicesOK) GetPayload() *rest_model.ListServicesEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPolicyServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListServicesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceEdgeRouterPolicyServicesUnauthorized creates a ListServiceEdgeRouterPolicyServicesUnauthorized with default headers values
func NewListServiceEdgeRouterPolicyServicesUnauthorized() *ListServiceEdgeRouterPolicyServicesUnauthorized {
	return &ListServiceEdgeRouterPolicyServicesUnauthorized{}
}

/* ListServiceEdgeRouterPolicyServicesUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListServiceEdgeRouterPolicyServicesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceEdgeRouterPolicyServicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}/services][%d] listServiceEdgeRouterPolicyServicesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServiceEdgeRouterPolicyServicesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPolicyServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceEdgeRouterPolicyServicesNotFound creates a ListServiceEdgeRouterPolicyServicesNotFound with default headers values
func NewListServiceEdgeRouterPolicyServicesNotFound() *ListServiceEdgeRouterPolicyServicesNotFound {
	return &ListServiceEdgeRouterPolicyServicesNotFound{}
}

/* ListServiceEdgeRouterPolicyServicesNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type ListServiceEdgeRouterPolicyServicesNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceEdgeRouterPolicyServicesNotFound) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}/services][%d] listServiceEdgeRouterPolicyServicesNotFound  %+v", 404, o.Payload)
}
func (o *ListServiceEdgeRouterPolicyServicesNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRouterPolicyServicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}