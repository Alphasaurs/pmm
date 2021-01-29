// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewAddHAProxyExporterParams creates a new AddHAProxyExporterParams object
// with the default values initialized.
func NewAddHAProxyExporterParams() *AddHAProxyExporterParams {
	var ()
	return &AddHAProxyExporterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddHAProxyExporterParamsWithTimeout creates a new AddHAProxyExporterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddHAProxyExporterParamsWithTimeout(timeout time.Duration) *AddHAProxyExporterParams {
	var ()
	return &AddHAProxyExporterParams{

		timeout: timeout,
	}
}

// NewAddHAProxyExporterParamsWithContext creates a new AddHAProxyExporterParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddHAProxyExporterParamsWithContext(ctx context.Context) *AddHAProxyExporterParams {
	var ()
	return &AddHAProxyExporterParams{

		Context: ctx,
	}
}

// NewAddHAProxyExporterParamsWithHTTPClient creates a new AddHAProxyExporterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddHAProxyExporterParamsWithHTTPClient(client *http.Client) *AddHAProxyExporterParams {
	var ()
	return &AddHAProxyExporterParams{
		HTTPClient: client,
	}
}

/*AddHAProxyExporterParams contains all the parameters to send to the API endpoint
for the add h a proxy exporter operation typically these are written to a http.Request
*/
type AddHAProxyExporterParams struct {

	/*Body*/
	Body AddHAProxyExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add h a proxy exporter params
func (o *AddHAProxyExporterParams) WithTimeout(timeout time.Duration) *AddHAProxyExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add h a proxy exporter params
func (o *AddHAProxyExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add h a proxy exporter params
func (o *AddHAProxyExporterParams) WithContext(ctx context.Context) *AddHAProxyExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add h a proxy exporter params
func (o *AddHAProxyExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add h a proxy exporter params
func (o *AddHAProxyExporterParams) WithHTTPClient(client *http.Client) *AddHAProxyExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add h a proxy exporter params
func (o *AddHAProxyExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add h a proxy exporter params
func (o *AddHAProxyExporterParams) WithBody(body AddHAProxyExporterBody) *AddHAProxyExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add h a proxy exporter params
func (o *AddHAProxyExporterParams) SetBody(body AddHAProxyExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddHAProxyExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}