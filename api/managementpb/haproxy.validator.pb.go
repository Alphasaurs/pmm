// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/haproxy.proto

package managementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddHAProxyRequest) Validate() error {
	if this.AddNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AddNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AddNode", err)
		}
	}
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if !(this.ListenPort > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ListenPort", fmt.Errorf(`value '%v' must be greater than '0'`, this.ListenPort))
	}
	if !(this.ListenPort < 65536) {
		return github_com_mwitkow_go_proto_validators.FieldError("ListenPort", fmt.Errorf(`value '%v' must be less than '65536'`, this.ListenPort))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddHAProxyResponse) Validate() error {
	if this.Service != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Service); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Service", err)
		}
	}
	if this.HaproxyExporter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HaproxyExporter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HaproxyExporter", err)
		}
	}
	return nil
}
