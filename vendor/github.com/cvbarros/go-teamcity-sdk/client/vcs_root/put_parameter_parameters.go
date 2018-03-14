// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

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

// NewPutParameterParams creates a new PutParameterParams object
// with the default values initialized.
func NewPutParameterParams() *PutParameterParams {
	var ()
	return &PutParameterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutParameterParamsWithTimeout creates a new PutParameterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutParameterParamsWithTimeout(timeout time.Duration) *PutParameterParams {
	var ()
	return &PutParameterParams{

		timeout: timeout,
	}
}

// NewPutParameterParamsWithContext creates a new PutParameterParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutParameterParamsWithContext(ctx context.Context) *PutParameterParams {
	var ()
	return &PutParameterParams{

		Context: ctx,
	}
}

// NewPutParameterParamsWithHTTPClient creates a new PutParameterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutParameterParamsWithHTTPClient(client *http.Client) *PutParameterParams {
	var ()
	return &PutParameterParams{
		HTTPClient: client,
	}
}

/*PutParameterParams contains all the parameters to send to the API endpoint
for the put parameter operation typically these are written to a http.Request
*/
type PutParameterParams struct {

	/*Body*/
	Body string
	/*Name*/
	Name string
	/*VcsRootLocator*/
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put parameter params
func (o *PutParameterParams) WithTimeout(timeout time.Duration) *PutParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put parameter params
func (o *PutParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put parameter params
func (o *PutParameterParams) WithContext(ctx context.Context) *PutParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put parameter params
func (o *PutParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put parameter params
func (o *PutParameterParams) WithHTTPClient(client *http.Client) *PutParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put parameter params
func (o *PutParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put parameter params
func (o *PutParameterParams) WithBody(body string) *PutParameterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put parameter params
func (o *PutParameterParams) SetBody(body string) {
	o.Body = body
}

// WithName adds the name to the put parameter params
func (o *PutParameterParams) WithName(name string) *PutParameterParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put parameter params
func (o *PutParameterParams) SetName(name string) {
	o.Name = name
}

// WithVcsRootLocator adds the vcsRootLocator to the put parameter params
func (o *PutParameterParams) WithVcsRootLocator(vcsRootLocator string) *PutParameterParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the put parameter params
func (o *PutParameterParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *PutParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param vcsRootLocator
	if err := r.SetPathParam("vcsRootLocator", o.VcsRootLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}