// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

	models "github.com/percona/pmm/api/json/models"
)

// NewAddNodeParams creates a new AddNodeParams object
// with the default values initialized.
func NewAddNodeParams() *AddNodeParams {
	var ()
	return &AddNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddNodeParamsWithTimeout creates a new AddNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddNodeParamsWithTimeout(timeout time.Duration) *AddNodeParams {
	var ()
	return &AddNodeParams{

		timeout: timeout,
	}
}

// NewAddNodeParamsWithContext creates a new AddNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddNodeParamsWithContext(ctx context.Context) *AddNodeParams {
	var ()
	return &AddNodeParams{

		Context: ctx,
	}
}

// NewAddNodeParamsWithHTTPClient creates a new AddNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddNodeParamsWithHTTPClient(client *http.Client) *AddNodeParams {
	var ()
	return &AddNodeParams{
		HTTPClient: client,
	}
}

/*AddNodeParams contains all the parameters to send to the API endpoint
for the add node operation typically these are written to a http.Request
*/
type AddNodeParams struct {

	/*Body*/
	Body *models.InventoryAddNodeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add node params
func (o *AddNodeParams) WithTimeout(timeout time.Duration) *AddNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add node params
func (o *AddNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add node params
func (o *AddNodeParams) WithContext(ctx context.Context) *AddNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add node params
func (o *AddNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add node params
func (o *AddNodeParams) WithHTTPClient(client *http.Client) *AddNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add node params
func (o *AddNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add node params
func (o *AddNodeParams) WithBody(body *models.InventoryAddNodeRequest) *AddNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add node params
func (o *AddNodeParams) SetBody(body *models.InventoryAddNodeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}