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

// ChangeHAProxyExporterReader is a Reader for the ChangeHAProxyExporter structure.
type ChangeHAProxyExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeHAProxyExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeHAProxyExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeHAProxyExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeHAProxyExporterOK creates a ChangeHAProxyExporterOK with default headers values
func NewChangeHAProxyExporterOK() *ChangeHAProxyExporterOK {
	return &ChangeHAProxyExporterOK{}
}

/*ChangeHAProxyExporterOK handles this case with default header values.

A successful response.
*/
type ChangeHAProxyExporterOK struct {
	Payload *ChangeHAProxyExporterOKBody
}

func (o *ChangeHAProxyExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeHAProxyExporter][%d] changeHAProxyExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeHAProxyExporterOK) GetPayload() *ChangeHAProxyExporterOKBody {
	return o.Payload
}

func (o *ChangeHAProxyExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeHAProxyExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeHAProxyExporterDefault creates a ChangeHAProxyExporterDefault with default headers values
func NewChangeHAProxyExporterDefault(code int) *ChangeHAProxyExporterDefault {
	return &ChangeHAProxyExporterDefault{
		_statusCode: code,
	}
}

/*ChangeHAProxyExporterDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeHAProxyExporterDefault struct {
	_statusCode int

	Payload *ChangeHAProxyExporterDefaultBody
}

// Code gets the status code for the change h a proxy exporter default response
func (o *ChangeHAProxyExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeHAProxyExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeHAProxyExporter][%d] ChangeHAProxyExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeHAProxyExporterDefault) GetPayload() *ChangeHAProxyExporterDefaultBody {
	return o.Payload
}

func (o *ChangeHAProxyExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeHAProxyExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeHAProxyExporterBody change h a proxy exporter body
swagger:model ChangeHAProxyExporterBody
*/
type ChangeHAProxyExporterBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeHAProxyExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change h a proxy exporter body
func (o *ChangeHAProxyExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeHAProxyExporterBody) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeHAProxyExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeHAProxyExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeHAProxyExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeHAProxyExporterDefaultBody change h a proxy exporter default body
swagger:model ChangeHAProxyExporterDefaultBody
*/
type ChangeHAProxyExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change h a proxy exporter default body
func (o *ChangeHAProxyExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeHAProxyExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ChangeHAProxyExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeHAProxyExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeHAProxyExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeHAProxyExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeHAProxyExporterOKBody change h a proxy exporter OK body
swagger:model ChangeHAProxyExporterOKBody
*/
type ChangeHAProxyExporterOKBody struct {

	// external exporter
	ExternalExporter *ChangeHAProxyExporterOKBodyExternalExporter `json:"external_exporter,omitempty"`
}

// Validate validates this change h a proxy exporter OK body
func (o *ChangeHAProxyExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExternalExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeHAProxyExporterOKBody) validateExternalExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalExporter) { // not required
		return nil
	}

	if o.ExternalExporter != nil {
		if err := o.ExternalExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeHAProxyExporterOk" + "." + "external_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeHAProxyExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeHAProxyExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeHAProxyExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeHAProxyExporterOKBodyExternalExporter HAProxyExporter runs on any Node type, including Remote Node.
swagger:model ChangeHAProxyExporterOKBodyExternalExporter
*/
type ChangeHAProxyExporterOKBodyExternalExporter struct {

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

// Validate validates this change h a proxy exporter OK body external exporter
func (o *ChangeHAProxyExporterOKBodyExternalExporter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeHAProxyExporterOKBodyExternalExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeHAProxyExporterOKBodyExternalExporter) UnmarshalBinary(b []byte) error {
	var res ChangeHAProxyExporterOKBodyExternalExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeHAProxyExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeHAProxyExporterParamsBodyCommon
*/
type ChangeHAProxyExporterParamsBodyCommon struct {

	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`

	// Enables push metrics with vmagent, can't be used with disable_push_metrics.
	// Can't be used with agent version lower then 2.12 and unsupported agents.
	EnablePushMetrics bool `json:"enable_push_metrics,omitempty"`

	// Disables push metrics, pmm-server starts to pull it, can't be used with enable_push_metrics.
	DisablePushMetrics bool `json:"disable_push_metrics,omitempty"`
}

// Validate validates this change h a proxy exporter params body common
func (o *ChangeHAProxyExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeHAProxyExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeHAProxyExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeHAProxyExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
