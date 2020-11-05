// Code generated by go-swagger; DO NOT EDIT.

package xtra_db_cluster

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

// GetXtraDBClusterReader is a Reader for the GetXtraDBCluster structure.
type GetXtraDBClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetXtraDBClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetXtraDBClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetXtraDBClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetXtraDBClusterOK creates a GetXtraDBClusterOK with default headers values
func NewGetXtraDBClusterOK() *GetXtraDBClusterOK {
	return &GetXtraDBClusterOK{}
}

/*GetXtraDBClusterOK handles this case with default header values.

A successful response.
*/
type GetXtraDBClusterOK struct {
	Payload *GetXtraDBClusterOKBody
}

func (o *GetXtraDBClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBClusters/Get][%d] getXtraDbClusterOk  %+v", 200, o.Payload)
}

func (o *GetXtraDBClusterOK) GetPayload() *GetXtraDBClusterOKBody {
	return o.Payload
}

func (o *GetXtraDBClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetXtraDBClusterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetXtraDBClusterDefault creates a GetXtraDBClusterDefault with default headers values
func NewGetXtraDBClusterDefault(code int) *GetXtraDBClusterDefault {
	return &GetXtraDBClusterDefault{
		_statusCode: code,
	}
}

/*GetXtraDBClusterDefault handles this case with default header values.

An unexpected error response.
*/
type GetXtraDBClusterDefault struct {
	_statusCode int

	Payload *GetXtraDBClusterDefaultBody
}

// Code gets the status code for the get xtra DB cluster default response
func (o *GetXtraDBClusterDefault) Code() int {
	return o._statusCode
}

func (o *GetXtraDBClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBClusters/Get][%d] GetXtraDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *GetXtraDBClusterDefault) GetPayload() *GetXtraDBClusterDefaultBody {
	return o.Payload
}

func (o *GetXtraDBClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetXtraDBClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetXtraDBClusterBody get xtra DB cluster body
swagger:model GetXtraDBClusterBody
*/
type GetXtraDBClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// XtraDB cluster name.
	Name string `json:"name,omitempty"`
}

// Validate validates this get xtra DB cluster body
func (o *GetXtraDBClusterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterBody) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetXtraDBClusterDefaultBody get xtra DB cluster default body
swagger:model GetXtraDBClusterDefaultBody
*/
type GetXtraDBClusterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get xtra DB cluster default body
func (o *GetXtraDBClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetXtraDBClusterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("GetXtraDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetXtraDBClusterOKBody get xtra DB cluster OK body
swagger:model GetXtraDBClusterOKBody
*/
type GetXtraDBClusterOKBody struct {

	// XtraDBClusterState represents XtraDB cluster CR state.
	//
	//  - XTRA_DB_CLUSTER_STATE_INVALID: XTRA_DB_CLUSTER_STATE_INVALID represents unknown state.
	//  - XTRA_DB_CLUSTER_STATE_CHANGING: XTRA_DB_CLUSTER_STATE_CHANGING represents a cluster being changed.
	//  - XTRA_DB_CLUSTER_STATE_READY: XTRA_DB_CLUSTER_STATE_READY represents a cluster without pending changes.
	//  - XTRA_DB_CLUSTER_STATE_FAILED: XTRA_DB_CLUSTER_STATE_FAILED represents a failed cluster.
	//  - XTRA_DB_CLUSTER_STATE_DELETING: XTRA_DB_CLUSTER_STATE_DELETING represents a cluster being deleting.
	// Enum: [XTRA_DB_CLUSTER_STATE_INVALID XTRA_DB_CLUSTER_STATE_CHANGING XTRA_DB_CLUSTER_STATE_READY XTRA_DB_CLUSTER_STATE_FAILED XTRA_DB_CLUSTER_STATE_DELETING]
	State *string `json:"state,omitempty"`

	// connection credentials
	ConnectionCredentials *GetXtraDBClusterOKBodyConnectionCredentials `json:"connection_credentials,omitempty"`

	// operation
	Operation *GetXtraDBClusterOKBodyOperation `json:"operation,omitempty"`

	// params
	Params *GetXtraDBClusterOKBodyParams `json:"params,omitempty"`
}

// Validate validates this get xtra DB cluster OK body
func (o *GetXtraDBClusterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConnectionCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getXtraDbClusterOkBodyTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["XTRA_DB_CLUSTER_STATE_INVALID","XTRA_DB_CLUSTER_STATE_CHANGING","XTRA_DB_CLUSTER_STATE_READY","XTRA_DB_CLUSTER_STATE_FAILED","XTRA_DB_CLUSTER_STATE_DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getXtraDbClusterOkBodyTypeStatePropEnum = append(getXtraDbClusterOkBodyTypeStatePropEnum, v)
	}
}

const (

	// GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATEINVALID captures enum value "XTRA_DB_CLUSTER_STATE_INVALID"
	GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATEINVALID string = "XTRA_DB_CLUSTER_STATE_INVALID"

	// GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATECHANGING captures enum value "XTRA_DB_CLUSTER_STATE_CHANGING"
	GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATECHANGING string = "XTRA_DB_CLUSTER_STATE_CHANGING"

	// GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATEREADY captures enum value "XTRA_DB_CLUSTER_STATE_READY"
	GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATEREADY string = "XTRA_DB_CLUSTER_STATE_READY"

	// GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATEFAILED captures enum value "XTRA_DB_CLUSTER_STATE_FAILED"
	GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATEFAILED string = "XTRA_DB_CLUSTER_STATE_FAILED"

	// GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATEDELETING captures enum value "XTRA_DB_CLUSTER_STATE_DELETING"
	GetXtraDBClusterOKBodyStateXTRADBCLUSTERSTATEDELETING string = "XTRA_DB_CLUSTER_STATE_DELETING"
)

// prop value enum
func (o *GetXtraDBClusterOKBody) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getXtraDbClusterOkBodyTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetXtraDBClusterOKBody) validateState(formats strfmt.Registry) error {

	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("getXtraDbClusterOk"+"."+"state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *GetXtraDBClusterOKBody) validateConnectionCredentials(formats strfmt.Registry) error {

	if swag.IsZero(o.ConnectionCredentials) { // not required
		return nil
	}

	if o.ConnectionCredentials != nil {
		if err := o.ConnectionCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getXtraDbClusterOk" + "." + "connection_credentials")
			}
			return err
		}
	}

	return nil
}

func (o *GetXtraDBClusterOKBody) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(o.Operation) { // not required
		return nil
	}

	if o.Operation != nil {
		if err := o.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getXtraDbClusterOk" + "." + "operation")
			}
			return err
		}
	}

	return nil
}

func (o *GetXtraDBClusterOKBody) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getXtraDbClusterOk" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterOKBody) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetXtraDBClusterOKBodyConnectionCredentials XtraDBClusterConnectionCredentials is cluster connection credentials.
swagger:model GetXtraDBClusterOKBodyConnectionCredentials
*/
type GetXtraDBClusterOKBodyConnectionCredentials struct {

	// ProxySQL username.
	Username string `json:"username,omitempty"`

	// ProxySQL password.
	Password string `json:"password,omitempty"`

	// ProxySQL host.
	Host string `json:"host,omitempty"`

	// ProxySQL port.
	Port int32 `json:"port,omitempty"`
}

// Validate validates this get xtra DB cluster OK body connection credentials
func (o *GetXtraDBClusterOKBodyConnectionCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyConnectionCredentials) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyConnectionCredentials) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterOKBodyConnectionCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetXtraDBClusterOKBodyOperation RunningOperation respresents a long-running operation.
swagger:model GetXtraDBClusterOKBodyOperation
*/
type GetXtraDBClusterOKBodyOperation struct {

	// Progress from 0.0 to 1.0; can decrease compared to the previous value.
	Progress float32 `json:"progress,omitempty"`

	// Text describing the current operation progress step.
	Message string `json:"message,omitempty"`
}

// Validate validates this get xtra DB cluster OK body operation
func (o *GetXtraDBClusterOKBodyOperation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyOperation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyOperation) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterOKBodyOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetXtraDBClusterOKBodyParams XtraDBClusterParams represents XtraDB cluster parameters that can be updated.
swagger:model GetXtraDBClusterOKBodyParams
*/
type GetXtraDBClusterOKBodyParams struct {

	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// proxysql
	Proxysql *GetXtraDBClusterOKBodyParamsProxysql `json:"proxysql,omitempty"`

	// pxc
	Pxc *GetXtraDBClusterOKBodyParamsPxc `json:"pxc,omitempty"`
}

// Validate validates this get xtra DB cluster OK body params
func (o *GetXtraDBClusterOKBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePxc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetXtraDBClusterOKBodyParams) validateProxysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	if o.Proxysql != nil {
		if err := o.Proxysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getXtraDbClusterOk" + "." + "params" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *GetXtraDBClusterOKBodyParams) validatePxc(formats strfmt.Registry) error {

	if swag.IsZero(o.Pxc) { // not required
		return nil
	}

	if o.Pxc != nil {
		if err := o.Pxc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getXtraDbClusterOk" + "." + "params" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParams) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterOKBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetXtraDBClusterOKBodyParamsProxysql ProxySQL container parameters.
swagger:model GetXtraDBClusterOKBodyParamsProxysql
*/
type GetXtraDBClusterOKBodyParamsProxysql struct {

	// compute resources
	ComputeResources *GetXtraDBClusterOKBodyParamsProxysqlComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this get xtra DB cluster OK body params proxysql
func (o *GetXtraDBClusterOKBodyParamsProxysql) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetXtraDBClusterOKBodyParamsProxysql) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getXtraDbClusterOk" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParamsProxysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParamsProxysql) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterOKBodyParamsProxysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetXtraDBClusterOKBodyParamsProxysqlComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model GetXtraDBClusterOKBodyParamsProxysqlComputeResources
*/
type GetXtraDBClusterOKBodyParamsProxysqlComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this get xtra DB cluster OK body params proxysql compute resources
func (o *GetXtraDBClusterOKBodyParamsProxysqlComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParamsProxysqlComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParamsProxysqlComputeResources) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterOKBodyParamsProxysqlComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetXtraDBClusterOKBodyParamsPxc PXC container parameters.
swagger:model GetXtraDBClusterOKBodyParamsPxc
*/
type GetXtraDBClusterOKBodyParamsPxc struct {

	// compute resources
	ComputeResources *GetXtraDBClusterOKBodyParamsPxcComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this get xtra DB cluster OK body params pxc
func (o *GetXtraDBClusterOKBodyParamsPxc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetXtraDBClusterOKBodyParamsPxc) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getXtraDbClusterOk" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParamsPxc) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParamsPxc) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterOKBodyParamsPxc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetXtraDBClusterOKBodyParamsPxcComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model GetXtraDBClusterOKBodyParamsPxcComputeResources
*/
type GetXtraDBClusterOKBodyParamsPxcComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this get xtra DB cluster OK body params pxc compute resources
func (o *GetXtraDBClusterOKBodyParamsPxcComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParamsPxcComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetXtraDBClusterOKBodyParamsPxcComputeResources) UnmarshalBinary(b []byte) error {
	var res GetXtraDBClusterOKBodyParamsPxcComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}