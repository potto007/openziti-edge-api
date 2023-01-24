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

package current_api_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListServiceUpdatesOKCode is the HTTP code returned for type ListServiceUpdatesOK
const ListServiceUpdatesOKCode int = 200

/*ListServiceUpdatesOK Data indicating necessary service updates

swagger:response listServiceUpdatesOK
*/
type ListServiceUpdatesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListCurrentAPISessionServiceUpdatesEnvelope `json:"body,omitempty"`
}

// NewListServiceUpdatesOK creates ListServiceUpdatesOK with default headers values
func NewListServiceUpdatesOK() *ListServiceUpdatesOK {

	return &ListServiceUpdatesOK{}
}

// WithPayload adds the payload to the list service updates o k response
func (o *ListServiceUpdatesOK) WithPayload(payload *rest_model.ListCurrentAPISessionServiceUpdatesEnvelope) *ListServiceUpdatesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service updates o k response
func (o *ListServiceUpdatesOK) SetPayload(payload *rest_model.ListCurrentAPISessionServiceUpdatesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceUpdatesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceUpdatesUnauthorizedCode is the HTTP code returned for type ListServiceUpdatesUnauthorized
const ListServiceUpdatesUnauthorizedCode int = 401

/*ListServiceUpdatesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listServiceUpdatesUnauthorized
*/
type ListServiceUpdatesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceUpdatesUnauthorized creates ListServiceUpdatesUnauthorized with default headers values
func NewListServiceUpdatesUnauthorized() *ListServiceUpdatesUnauthorized {

	return &ListServiceUpdatesUnauthorized{}
}

// WithPayload adds the payload to the list service updates unauthorized response
func (o *ListServiceUpdatesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceUpdatesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service updates unauthorized response
func (o *ListServiceUpdatesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceUpdatesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}