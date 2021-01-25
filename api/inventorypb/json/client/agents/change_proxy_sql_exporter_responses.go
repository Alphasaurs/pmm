// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeProxySQLExporterReader is a Reader for the ChangeProxySQLExporter structure.
type ChangeProxySQLExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeProxySQLExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeProxySQLExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeProxySQLExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeProxySQLExporterOK creates a ChangeProxySQLExporterOK with default headers values
func NewChangeProxySQLExporterOK() *ChangeProxySQLExporterOK {
	return &ChangeProxySQLExporterOK{}
}

/*ChangeProxySQLExporterOK handles this case with default header values.

A successful response.
*/
type ChangeProxySQLExporterOK struct {
	Payload *ChangeProxySQLExporterOKBody
}

func (o *ChangeProxySQLExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeProxySQLExporter][%d] changeProxySqlExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeProxySQLExporterOK) GetPayload() *ChangeProxySQLExporterOKBody {
	return o.Payload
}

func (o *ChangeProxySQLExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeProxySQLExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeProxySQLExporterDefault creates a ChangeProxySQLExporterDefault with default headers values
func NewChangeProxySQLExporterDefault(code int) *ChangeProxySQLExporterDefault {
	return &ChangeProxySQLExporterDefault{
		_statusCode: code,
	}
}

/*ChangeProxySQLExporterDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeProxySQLExporterDefault struct {
	_statusCode int

	Payload *ChangeProxySQLExporterDefaultBody
}

// Code gets the status code for the change proxy SQL exporter default response
func (o *ChangeProxySQLExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeProxySQLExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeProxySQLExporter][%d] ChangeProxySQLExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeProxySQLExporterDefault) GetPayload() *ChangeProxySQLExporterDefaultBody {
	return o.Payload
}

func (o *ChangeProxySQLExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeProxySQLExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeProxySQLExporterBody change proxy SQL exporter body
swagger:model ChangeProxySQLExporterBody
*/
type ChangeProxySQLExporterBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeProxySQLExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change proxy SQL exporter body
func (o *ChangeProxySQLExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeProxySQLExporterBody) validateCommon(formats strfmt.Registry) error {

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
func (o *ChangeProxySQLExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeProxySQLExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeProxySQLExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeProxySQLExporterDefaultBody change proxy SQL exporter default body
swagger:model ChangeProxySQLExporterDefaultBody
*/
type ChangeProxySQLExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change proxy SQL exporter default body
func (o *ChangeProxySQLExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeProxySQLExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ChangeProxySQLExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeProxySQLExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeProxySQLExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeProxySQLExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeProxySQLExporterOKBody change proxy SQL exporter OK body
swagger:model ChangeProxySQLExporterOKBody
*/
type ChangeProxySQLExporterOKBody struct {

	// proxysql exporter
	ProxysqlExporter *ChangeProxySQLExporterOKBodyProxysqlExporter `json:"proxysql_exporter,omitempty"`
}

// Validate validates this change proxy SQL exporter OK body
func (o *ChangeProxySQLExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProxysqlExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeProxySQLExporterOKBody) validateProxysqlExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.ProxysqlExporter) { // not required
		return nil
	}

	if o.ProxysqlExporter != nil {
		if err := o.ProxysqlExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeProxySqlExporterOk" + "." + "proxysql_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeProxySQLExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeProxySQLExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeProxySQLExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeProxySQLExporterOKBodyProxysqlExporter ProxySQLExporter runs on Generic or Container Node and exposes ProxySQL Service metrics.
swagger:model ChangeProxySQLExporterOKBodyProxysqlExporter
*/
type ChangeProxySQLExporterOKBodyProxysqlExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// ProxySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of collector names to exclude from exporter.
	DisableCollectors []string `json:"disable_collectors"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`
}

// Validate validates this change proxy SQL exporter OK body proxysql exporter
func (o *ChangeProxySQLExporterOKBodyProxysqlExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum = append(changeProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangeProxySQLExporterOKBodyProxysqlExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeProxySQLExporterOKBodyProxysqlExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeProxySQLExporterOKBodyProxysqlExporterStatusSTARTING captures enum value "STARTING"
	ChangeProxySQLExporterOKBodyProxysqlExporterStatusSTARTING string = "STARTING"

	// ChangeProxySQLExporterOKBodyProxysqlExporterStatusRUNNING captures enum value "RUNNING"
	ChangeProxySQLExporterOKBodyProxysqlExporterStatusRUNNING string = "RUNNING"

	// ChangeProxySQLExporterOKBodyProxysqlExporterStatusWAITING captures enum value "WAITING"
	ChangeProxySQLExporterOKBodyProxysqlExporterStatusWAITING string = "WAITING"

	// ChangeProxySQLExporterOKBodyProxysqlExporterStatusSTOPPING captures enum value "STOPPING"
	ChangeProxySQLExporterOKBodyProxysqlExporterStatusSTOPPING string = "STOPPING"

	// ChangeProxySQLExporterOKBodyProxysqlExporterStatusDONE captures enum value "DONE"
	ChangeProxySQLExporterOKBodyProxysqlExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *ChangeProxySQLExporterOKBodyProxysqlExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeProxySQLExporterOKBodyProxysqlExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeProxySqlExporterOk"+"."+"proxysql_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeProxySQLExporterOKBodyProxysqlExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeProxySQLExporterOKBodyProxysqlExporter) UnmarshalBinary(b []byte) error {
	var res ChangeProxySQLExporterOKBodyProxysqlExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeProxySQLExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeProxySQLExporterParamsBodyCommon
*/
type ChangeProxySQLExporterParamsBodyCommon struct {

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

// Validate validates this change proxy SQL exporter params body common
func (o *ChangeProxySQLExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeProxySQLExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeProxySQLExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeProxySQLExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
