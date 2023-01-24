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

package edge_router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListEdgeRoutersReader is a Reader for the ListEdgeRouters structure.
type ListEdgeRoutersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEdgeRoutersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEdgeRoutersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListEdgeRoutersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListEdgeRoutersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEdgeRoutersOK creates a ListEdgeRoutersOK with default headers values
func NewListEdgeRoutersOK() *ListEdgeRoutersOK {
	return &ListEdgeRoutersOK{}
}

/* ListEdgeRoutersOK describes a response with status code 200, with default header values.

A list of edge routers
*/
type ListEdgeRoutersOK struct {
	Payload *rest_model.ListEdgeRoutersEnvelope
}

func (o *ListEdgeRoutersOK) Error() string {
	return fmt.Sprintf("[GET /edge-routers][%d] listEdgeRoutersOK  %+v", 200, o.Payload)
}
func (o *ListEdgeRoutersOK) GetPayload() *rest_model.ListEdgeRoutersEnvelope {
	return o.Payload
}

func (o *ListEdgeRoutersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListEdgeRoutersEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRoutersBadRequest creates a ListEdgeRoutersBadRequest with default headers values
func NewListEdgeRoutersBadRequest() *ListEdgeRoutersBadRequest {
	return &ListEdgeRoutersBadRequest{}
}

/* ListEdgeRoutersBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListEdgeRoutersBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRoutersBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge-routers][%d] listEdgeRoutersBadRequest  %+v", 400, o.Payload)
}
func (o *ListEdgeRoutersBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRoutersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRoutersUnauthorized creates a ListEdgeRoutersUnauthorized with default headers values
func NewListEdgeRoutersUnauthorized() *ListEdgeRoutersUnauthorized {
	return &ListEdgeRoutersUnauthorized{}
}

/* ListEdgeRoutersUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListEdgeRoutersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRoutersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /edge-routers][%d] listEdgeRoutersUnauthorized  %+v", 401, o.Payload)
}
func (o *ListEdgeRoutersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRoutersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}