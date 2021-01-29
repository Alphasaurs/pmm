// Code generated by go-swagger; DO NOT EDIT.

package xtra_db_cluster

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

// NewGetXtraDBClusterCredentialsParams creates a new GetXtraDBClusterCredentialsParams object
// with the default values initialized.
func NewGetXtraDBClusterCredentialsParams() *GetXtraDBClusterCredentialsParams {
	var ()
	return &GetXtraDBClusterCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetXtraDBClusterCredentialsParamsWithTimeout creates a new GetXtraDBClusterCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetXtraDBClusterCredentialsParamsWithTimeout(timeout time.Duration) *GetXtraDBClusterCredentialsParams {
	var ()
	return &GetXtraDBClusterCredentialsParams{

		timeout: timeout,
	}
}

// NewGetXtraDBClusterCredentialsParamsWithContext creates a new GetXtraDBClusterCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetXtraDBClusterCredentialsParamsWithContext(ctx context.Context) *GetXtraDBClusterCredentialsParams {
	var ()
	return &GetXtraDBClusterCredentialsParams{

		Context: ctx,
	}
}

// NewGetXtraDBClusterCredentialsParamsWithHTTPClient creates a new GetXtraDBClusterCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetXtraDBClusterCredentialsParamsWithHTTPClient(client *http.Client) *GetXtraDBClusterCredentialsParams {
	var ()
	return &GetXtraDBClusterCredentialsParams{
		HTTPClient: client,
	}
}

/*GetXtraDBClusterCredentialsParams contains all the parameters to send to the API endpoint
for the get xtra DB cluster credentials operation typically these are written to a http.Request
*/
type GetXtraDBClusterCredentialsParams struct {

	/*Body*/
	Body GetXtraDBClusterCredentialsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get xtra DB cluster credentials params
func (o *GetXtraDBClusterCredentialsParams) WithTimeout(timeout time.Duration) *GetXtraDBClusterCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get xtra DB cluster credentials params
func (o *GetXtraDBClusterCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get xtra DB cluster credentials params
func (o *GetXtraDBClusterCredentialsParams) WithContext(ctx context.Context) *GetXtraDBClusterCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get xtra DB cluster credentials params
func (o *GetXtraDBClusterCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get xtra DB cluster credentials params
func (o *GetXtraDBClusterCredentialsParams) WithHTTPClient(client *http.Client) *GetXtraDBClusterCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get xtra DB cluster credentials params
func (o *GetXtraDBClusterCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get xtra DB cluster credentials params
func (o *GetXtraDBClusterCredentialsParams) WithBody(body GetXtraDBClusterCredentialsBody) *GetXtraDBClusterCredentialsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get xtra DB cluster credentials params
func (o *GetXtraDBClusterCredentialsParams) SetBody(body GetXtraDBClusterCredentialsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetXtraDBClusterCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
