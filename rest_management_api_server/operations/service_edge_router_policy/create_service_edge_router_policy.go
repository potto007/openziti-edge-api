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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateServiceEdgeRouterPolicyHandlerFunc turns a function with the right signature into a create service edge router policy handler
type CreateServiceEdgeRouterPolicyHandlerFunc func(CreateServiceEdgeRouterPolicyParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateServiceEdgeRouterPolicyHandlerFunc) Handle(params CreateServiceEdgeRouterPolicyParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateServiceEdgeRouterPolicyHandler interface for that can handle valid create service edge router policy params
type CreateServiceEdgeRouterPolicyHandler interface {
	Handle(CreateServiceEdgeRouterPolicyParams, interface{}) middleware.Responder
}

// NewCreateServiceEdgeRouterPolicy creates a new http.Handler for the create service edge router policy operation
func NewCreateServiceEdgeRouterPolicy(ctx *middleware.Context, handler CreateServiceEdgeRouterPolicyHandler) *CreateServiceEdgeRouterPolicy {
	return &CreateServiceEdgeRouterPolicy{Context: ctx, Handler: handler}
}

/* CreateServiceEdgeRouterPolicy swagger:route POST /service-edge-router-policies Service Edge Router Policy createServiceEdgeRouterPolicy

Create a service edge router policy resource

Create a service edge router policy resource. Requires admin access.

*/
type CreateServiceEdgeRouterPolicy struct {
	Context *middleware.Context
	Handler CreateServiceEdgeRouterPolicyHandler
}

func (o *CreateServiceEdgeRouterPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateServiceEdgeRouterPolicyParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}