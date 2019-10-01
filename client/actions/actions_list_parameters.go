//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewActionsListParams creates a new ActionsListParams object
// with the default values initialized.
func NewActionsListParams() *ActionsListParams {
	var ()
	return &ActionsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActionsListParamsWithTimeout creates a new ActionsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActionsListParamsWithTimeout(timeout time.Duration) *ActionsListParams {
	var ()
	return &ActionsListParams{

		timeout: timeout,
	}
}

// NewActionsListParamsWithContext creates a new ActionsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewActionsListParamsWithContext(ctx context.Context) *ActionsListParams {
	var ()
	return &ActionsListParams{

		Context: ctx,
	}
}

// NewActionsListParamsWithHTTPClient creates a new ActionsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActionsListParamsWithHTTPClient(client *http.Client) *ActionsListParams {
	var ()
	return &ActionsListParams{
		HTTPClient: client,
	}
}

/*ActionsListParams contains all the parameters to send to the API endpoint
for the actions list operation typically these are written to a http.Request
*/
type ActionsListParams struct {

	/*Limit
	  The maximum number of items to be returned per page. Default value is set in Weaviate config.

	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the actions list params
func (o *ActionsListParams) WithTimeout(timeout time.Duration) *ActionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the actions list params
func (o *ActionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the actions list params
func (o *ActionsListParams) WithContext(ctx context.Context) *ActionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the actions list params
func (o *ActionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the actions list params
func (o *ActionsListParams) WithHTTPClient(client *http.Client) *ActionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the actions list params
func (o *ActionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the actions list params
func (o *ActionsListParams) WithLimit(limit *int64) *ActionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the actions list params
func (o *ActionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *ActionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}