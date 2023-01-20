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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/go/rest_model"
)

// ListServiceServiceEdgeRouterPoliciesOKCode is the HTTP code returned for type ListServiceServiceEdgeRouterPoliciesOK
const ListServiceServiceEdgeRouterPoliciesOKCode int = 200

/*ListServiceServiceEdgeRouterPoliciesOK A list of service edge router policies

swagger:response listServiceServiceEdgeRouterPoliciesOK
*/
type ListServiceServiceEdgeRouterPoliciesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListServiceEdgeRouterPoliciesEnvelope `json:"body,omitempty"`
}

// NewListServiceServiceEdgeRouterPoliciesOK creates ListServiceServiceEdgeRouterPoliciesOK with default headers values
func NewListServiceServiceEdgeRouterPoliciesOK() *ListServiceServiceEdgeRouterPoliciesOK {

	return &ListServiceServiceEdgeRouterPoliciesOK{}
}

// WithPayload adds the payload to the list service service edge router policies o k response
func (o *ListServiceServiceEdgeRouterPoliciesOK) WithPayload(payload *rest_model.ListServiceEdgeRouterPoliciesEnvelope) *ListServiceServiceEdgeRouterPoliciesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service service edge router policies o k response
func (o *ListServiceServiceEdgeRouterPoliciesOK) SetPayload(payload *rest_model.ListServiceEdgeRouterPoliciesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceServiceEdgeRouterPoliciesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceServiceEdgeRouterPoliciesBadRequestCode is the HTTP code returned for type ListServiceServiceEdgeRouterPoliciesBadRequest
const ListServiceServiceEdgeRouterPoliciesBadRequestCode int = 400

/*ListServiceServiceEdgeRouterPoliciesBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listServiceServiceEdgeRouterPoliciesBadRequest
*/
type ListServiceServiceEdgeRouterPoliciesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceServiceEdgeRouterPoliciesBadRequest creates ListServiceServiceEdgeRouterPoliciesBadRequest with default headers values
func NewListServiceServiceEdgeRouterPoliciesBadRequest() *ListServiceServiceEdgeRouterPoliciesBadRequest {

	return &ListServiceServiceEdgeRouterPoliciesBadRequest{}
}

// WithPayload adds the payload to the list service service edge router policies bad request response
func (o *ListServiceServiceEdgeRouterPoliciesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceServiceEdgeRouterPoliciesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service service edge router policies bad request response
func (o *ListServiceServiceEdgeRouterPoliciesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceServiceEdgeRouterPoliciesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceServiceEdgeRouterPoliciesUnauthorizedCode is the HTTP code returned for type ListServiceServiceEdgeRouterPoliciesUnauthorized
const ListServiceServiceEdgeRouterPoliciesUnauthorizedCode int = 401

/*ListServiceServiceEdgeRouterPoliciesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listServiceServiceEdgeRouterPoliciesUnauthorized
*/
type ListServiceServiceEdgeRouterPoliciesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceServiceEdgeRouterPoliciesUnauthorized creates ListServiceServiceEdgeRouterPoliciesUnauthorized with default headers values
func NewListServiceServiceEdgeRouterPoliciesUnauthorized() *ListServiceServiceEdgeRouterPoliciesUnauthorized {

	return &ListServiceServiceEdgeRouterPoliciesUnauthorized{}
}

// WithPayload adds the payload to the list service service edge router policies unauthorized response
func (o *ListServiceServiceEdgeRouterPoliciesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceServiceEdgeRouterPoliciesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service service edge router policies unauthorized response
func (o *ListServiceServiceEdgeRouterPoliciesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceServiceEdgeRouterPoliciesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}