// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewServeBuildParams creates a new ServeBuildParams object
// with the default values initialized.
func NewServeBuildParams() *ServeBuildParams {
	var ()
	return &ServeBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildParamsWithTimeout creates a new ServeBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildParamsWithTimeout(timeout time.Duration) *ServeBuildParams {
	var ()
	return &ServeBuildParams{

		timeout: timeout,
	}
}

// NewServeBuildParamsWithContext creates a new ServeBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildParamsWithContext(ctx context.Context) *ServeBuildParams {
	var ()
	return &ServeBuildParams{

		Context: ctx,
	}
}

// NewServeBuildParamsWithHTTPClient creates a new ServeBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildParamsWithHTTPClient(client *http.Client) *ServeBuildParams {
	var ()
	return &ServeBuildParams{
		HTTPClient: client,
	}
}

/*ServeBuildParams contains all the parameters to send to the API endpoint
for the serve build operation typically these are written to a http.Request
*/
type ServeBuildParams struct {

	/*BuildLocator*/
	BuildLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build params
func (o *ServeBuildParams) WithTimeout(timeout time.Duration) *ServeBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build params
func (o *ServeBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build params
func (o *ServeBuildParams) WithContext(ctx context.Context) *ServeBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build params
func (o *ServeBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build params
func (o *ServeBuildParams) WithHTTPClient(client *http.Client) *ServeBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build params
func (o *ServeBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the serve build params
func (o *ServeBuildParams) WithBuildLocator(buildLocator string) *ServeBuildParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve build params
func (o *ServeBuildParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the serve build params
func (o *ServeBuildParams) WithFields(fields *string) *ServeBuildParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve build params
func (o *ServeBuildParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}