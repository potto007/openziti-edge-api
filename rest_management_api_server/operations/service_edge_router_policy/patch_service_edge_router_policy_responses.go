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
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// PatchServiceEdgeRouterPolicyOKCode is the HTTP code returned for type PatchServiceEdgeRouterPolicyOK
const PatchServiceEdgeRouterPolicyOKCode int = 200

/*PatchServiceEdgeRouterPolicyOK The patch request was successful and the resource has been altered

swagger:response patchServiceEdgeRouterPolicyOK
*/
type PatchServiceEdgeRouterPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchServiceEdgeRouterPolicyOK creates PatchServiceEdgeRouterPolicyOK with default headers values
func NewPatchServiceEdgeRouterPolicyOK() *PatchServiceEdgeRouterPolicyOK {

	return &PatchServiceEdgeRouterPolicyOK{}
}

// WithPayload adds the payload to the patch service edge router policy o k response
func (o *PatchServiceEdgeRouterPolicyOK) WithPayload(payload *rest_model.Empty) *PatchServiceEdgeRouterPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service edge router policy o k response
func (o *PatchServiceEdgeRouterPolicyOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServiceEdgeRouterPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServiceEdgeRouterPolicyBadRequestCode is the HTTP code returned for type PatchServiceEdgeRouterPolicyBadRequest
const PatchServiceEdgeRouterPolicyBadRequestCode int = 400

/*PatchServiceEdgeRouterPolicyBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchServiceEdgeRouterPolicyBadRequest
*/
type PatchServiceEdgeRouterPolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServiceEdgeRouterPolicyBadRequest creates PatchServiceEdgeRouterPolicyBadRequest with default headers values
func NewPatchServiceEdgeRouterPolicyBadRequest() *PatchServiceEdgeRouterPolicyBadRequest {

	return &PatchServiceEdgeRouterPolicyBadRequest{}
}

// WithPayload adds the payload to the patch service edge router policy bad request response
func (o *PatchServiceEdgeRouterPolicyBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServiceEdgeRouterPolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service edge router policy bad request response
func (o *PatchServiceEdgeRouterPolicyBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServiceEdgeRouterPolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServiceEdgeRouterPolicyUnauthorizedCode is the HTTP code returned for type PatchServiceEdgeRouterPolicyUnauthorized
const PatchServiceEdgeRouterPolicyUnauthorizedCode int = 401

/*PatchServiceEdgeRouterPolicyUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response patchServiceEdgeRouterPolicyUnauthorized
*/
type PatchServiceEdgeRouterPolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServiceEdgeRouterPolicyUnauthorized creates PatchServiceEdgeRouterPolicyUnauthorized with default headers values
func NewPatchServiceEdgeRouterPolicyUnauthorized() *PatchServiceEdgeRouterPolicyUnauthorized {

	return &PatchServiceEdgeRouterPolicyUnauthorized{}
}

// WithPayload adds the payload to the patch service edge router policy unauthorized response
func (o *PatchServiceEdgeRouterPolicyUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServiceEdgeRouterPolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service edge router policy unauthorized response
func (o *PatchServiceEdgeRouterPolicyUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServiceEdgeRouterPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServiceEdgeRouterPolicyNotFoundCode is the HTTP code returned for type PatchServiceEdgeRouterPolicyNotFound
const PatchServiceEdgeRouterPolicyNotFoundCode int = 404

/*PatchServiceEdgeRouterPolicyNotFound The requested resource does not exist

swagger:response patchServiceEdgeRouterPolicyNotFound
*/
type PatchServiceEdgeRouterPolicyNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServiceEdgeRouterPolicyNotFound creates PatchServiceEdgeRouterPolicyNotFound with default headers values
func NewPatchServiceEdgeRouterPolicyNotFound() *PatchServiceEdgeRouterPolicyNotFound {

	return &PatchServiceEdgeRouterPolicyNotFound{}
}

// WithPayload adds the payload to the patch service edge router policy not found response
func (o *PatchServiceEdgeRouterPolicyNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServiceEdgeRouterPolicyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service edge router policy not found response
func (o *PatchServiceEdgeRouterPolicyNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServiceEdgeRouterPolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}