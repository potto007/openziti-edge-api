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

package well_known

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListWellKnownCasParams creates a new ListWellKnownCasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListWellKnownCasParams() *ListWellKnownCasParams {
	return &ListWellKnownCasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListWellKnownCasParamsWithTimeout creates a new ListWellKnownCasParams object
// with the ability to set a timeout on a request.
func NewListWellKnownCasParamsWithTimeout(timeout time.Duration) *ListWellKnownCasParams {
	return &ListWellKnownCasParams{
		timeout: timeout,
	}
}

// NewListWellKnownCasParamsWithContext creates a new ListWellKnownCasParams object
// with the ability to set a context for a request.
func NewListWellKnownCasParamsWithContext(ctx context.Context) *ListWellKnownCasParams {
	return &ListWellKnownCasParams{
		Context: ctx,
	}
}

// NewListWellKnownCasParamsWithHTTPClient creates a new ListWellKnownCasParams object
// with the ability to set a custom HTTPClient for a request.
func NewListWellKnownCasParamsWithHTTPClient(client *http.Client) *ListWellKnownCasParams {
	return &ListWellKnownCasParams{
		HTTPClient: client,
	}
}

/* ListWellKnownCasParams contains all the parameters to send to the API endpoint
   for the list well known cas operation.

   Typically these are written to a http.Request.
*/
type ListWellKnownCasParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list well known cas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWellKnownCasParams) WithDefaults() *ListWellKnownCasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list well known cas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWellKnownCasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list well known cas params
func (o *ListWellKnownCasParams) WithTimeout(timeout time.Duration) *ListWellKnownCasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list well known cas params
func (o *ListWellKnownCasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list well known cas params
func (o *ListWellKnownCasParams) WithContext(ctx context.Context) *ListWellKnownCasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list well known cas params
func (o *ListWellKnownCasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list well known cas params
func (o *ListWellKnownCasParams) WithHTTPClient(client *http.Client) *ListWellKnownCasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list well known cas params
func (o *ListWellKnownCasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListWellKnownCasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}