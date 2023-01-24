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

// ListCurrentAPISessionCertificatesOKCode is the HTTP code returned for type ListCurrentAPISessionCertificatesOK
const ListCurrentAPISessionCertificatesOKCode int = 200

/*ListCurrentAPISessionCertificatesOK A list of the current API Session's certificate

swagger:response listCurrentApiSessionCertificatesOK
*/
type ListCurrentAPISessionCertificatesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListCurrentAPISessionCertificatesEnvelope `json:"body,omitempty"`
}

// NewListCurrentAPISessionCertificatesOK creates ListCurrentAPISessionCertificatesOK with default headers values
func NewListCurrentAPISessionCertificatesOK() *ListCurrentAPISessionCertificatesOK {

	return &ListCurrentAPISessionCertificatesOK{}
}

// WithPayload adds the payload to the list current Api session certificates o k response
func (o *ListCurrentAPISessionCertificatesOK) WithPayload(payload *rest_model.ListCurrentAPISessionCertificatesEnvelope) *ListCurrentAPISessionCertificatesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list current Api session certificates o k response
func (o *ListCurrentAPISessionCertificatesOK) SetPayload(payload *rest_model.ListCurrentAPISessionCertificatesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCurrentAPISessionCertificatesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCurrentAPISessionCertificatesBadRequestCode is the HTTP code returned for type ListCurrentAPISessionCertificatesBadRequest
const ListCurrentAPISessionCertificatesBadRequestCode int = 400

/*ListCurrentAPISessionCertificatesBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listCurrentApiSessionCertificatesBadRequest
*/
type ListCurrentAPISessionCertificatesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListCurrentAPISessionCertificatesBadRequest creates ListCurrentAPISessionCertificatesBadRequest with default headers values
func NewListCurrentAPISessionCertificatesBadRequest() *ListCurrentAPISessionCertificatesBadRequest {

	return &ListCurrentAPISessionCertificatesBadRequest{}
}

// WithPayload adds the payload to the list current Api session certificates bad request response
func (o *ListCurrentAPISessionCertificatesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListCurrentAPISessionCertificatesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list current Api session certificates bad request response
func (o *ListCurrentAPISessionCertificatesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCurrentAPISessionCertificatesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCurrentAPISessionCertificatesUnauthorizedCode is the HTTP code returned for type ListCurrentAPISessionCertificatesUnauthorized
const ListCurrentAPISessionCertificatesUnauthorizedCode int = 401

/*ListCurrentAPISessionCertificatesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listCurrentApiSessionCertificatesUnauthorized
*/
type ListCurrentAPISessionCertificatesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListCurrentAPISessionCertificatesUnauthorized creates ListCurrentAPISessionCertificatesUnauthorized with default headers values
func NewListCurrentAPISessionCertificatesUnauthorized() *ListCurrentAPISessionCertificatesUnauthorized {

	return &ListCurrentAPISessionCertificatesUnauthorized{}
}

// WithPayload adds the payload to the list current Api session certificates unauthorized response
func (o *ListCurrentAPISessionCertificatesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListCurrentAPISessionCertificatesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list current Api session certificates unauthorized response
func (o *ListCurrentAPISessionCertificatesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCurrentAPISessionCertificatesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}