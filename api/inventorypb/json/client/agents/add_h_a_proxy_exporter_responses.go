// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddHAProxyExporterReader is a Reader for the AddHAProxyExporter structure.
type AddHAProxyExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddHAProxyExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddHAProxyExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddHAProxyExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddHAProxyExporterOK creates a AddHAProxyExporterOK with default headers values
func NewAddHAProxyExporterOK() *AddHAProxyExporterOK {
	return &AddHAProxyExporterOK{}
}

/*AddHAProxyExporterOK handles this case with default header values.

A successful response.
*/
type AddHAProxyExporterOK struct {
	Payload *AddHAProxyExporterOKBody
}

func (o *AddHAProxyExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddHAProxyExporter][%d] addHAProxyExporterOk  %+v", 200, o.Payload)
}

func (o *AddHAProxyExporterOK) GetPayload() *AddHAProxyExporterOKBody {
	return o.Payload
}

func (o *AddHAProxyExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddHAProxyExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddHAProxyExporterDefault creates a AddHAProxyExporterDefault with default headers values
func NewAddHAProxyExporterDefault(code int) *AddHAProxyExporterDefault {
	return &AddHAProxyExporterDefault{
		_statusCode: code,
	}
}

/*AddHAProxyExporterDefault handles this case with default header values.

An unexpected error response.
*/
type AddHAProxyExporterDefault struct {
	_statusCode int

	Payload *AddHAProxyExporterDefaultBody
}

// Code gets the status code for the add h a proxy exporter default response
func (o *AddHAProxyExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddHAProxyExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddHAProxyExporter][%d] AddHAProxyExporter default  %+v", o._statusCode, o.Payload)
}

func (o *AddHAProxyExporterDefault) GetPayload() *AddHAProxyExporterDefaultBody {
	return o.Payload
}

func (o *AddHAProxyExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddHAProxyExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddHAProxyExporterBody add h a proxy exporter body
swagger:model AddHAProxyExporterBody
*/
type AddHAProxyExporterBody struct {

	// The node identifier where this instance is run.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// HTTP basic auth username for collecting metrics.
	Username string `json:"username,omitempty"`

	// HTTP basic auth password for collecting metrics.
	Password string `json:"password,omitempty"`

	// Scheme to generate URI to exporter metrics endpoints(default: http).
	Scheme string `json:"scheme,omitempty"`

	// Path under which metrics are exposed, used to generate URI(default: /metrics).
	MetricsPath string `json:"metrics_path,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`
}

// Validate validates this add h a proxy exporter body
func (o *AddHAProxyExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyExporterBody) UnmarshalBinary(b []byte) error {
	var res AddHAProxyExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyExporterDefaultBody add h a proxy exporter default body
swagger:model AddHAProxyExporterDefaultBody
*/
type AddHAProxyExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add h a proxy exporter default body
func (o *AddHAProxyExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddHAProxyExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddHAProxyExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddHAProxyExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyExporterOKBody add h a proxy exporter OK body
swagger:model AddHAProxyExporterOKBody
*/
type AddHAProxyExporterOKBody struct {

	// external exporter
	ExternalExporter *AddHAProxyExporterOKBodyExternalExporter `json:"external_exporter,omitempty"`
}

// Validate validates this add h a proxy exporter OK body
func (o *AddHAProxyExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExternalExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddHAProxyExporterOKBody) validateExternalExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalExporter) { // not required
		return nil
	}

	if o.ExternalExporter != nil {
		if err := o.ExternalExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addHAProxyExporterOk" + "." + "external_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddHAProxyExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyExporterOKBodyExternalExporter HAProxyExporter runs on any Node type, including Remote Node.
swagger:model AddHAProxyExporterOKBodyExternalExporter
*/
type AddHAProxyExporterOKBodyExternalExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// If disabled, metrics from this exporter will not be collected.
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// HTTP basic auth username for collecting metrics.
	Username string `json:"username,omitempty"`

	// Scheme to generate URI to exporter metrics endpoints.
	Scheme string `json:"scheme,omitempty"`

	// Path under which metrics are exposed, used to generate URI.
	MetricsPath string `json:"metrics_path,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`
}

// Validate validates this add h a proxy exporter OK body external exporter
func (o *AddHAProxyExporterOKBodyExternalExporter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyExporterOKBodyExternalExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyExporterOKBodyExternalExporter) UnmarshalBinary(b []byte) error {
	var res AddHAProxyExporterOKBodyExternalExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
