// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: envoy/admin/v3/config_dump.proto

package adminv3

import (
	_ "github.com/cncf/xds/go/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The :ref:`/config_dump <operations_admin_interface_config_dump>` admin endpoint uses this wrapper
// message to maintain and serve arbitrary configuration information from any component in Envoy.
type ConfigDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This list is serialized and dumped in its entirety at the
	// :ref:`/config_dump <operations_admin_interface_config_dump>` endpoint.
	//
	// The following configurations are currently supported and will be dumped in the order given
	// below:
	//
	// * “bootstrap“: :ref:`BootstrapConfigDump <envoy_v3_api_msg_admin.v3.BootstrapConfigDump>`
	// * “clusters“: :ref:`ClustersConfigDump <envoy_v3_api_msg_admin.v3.ClustersConfigDump>`
	// * “ecds_filter_http“: :ref:`EcdsConfigDump <envoy_v3_api_msg_admin.v3.EcdsConfigDump>`
	// * “ecds_filter_quic_listener“: :ref:`EcdsConfigDump <envoy_v3_api_msg_admin.v3.EcdsConfigDump>`
	// * “ecds_filter_tcp_listener“: :ref:`EcdsConfigDump <envoy_v3_api_msg_admin.v3.EcdsConfigDump>`
	// * “endpoints“:  :ref:`EndpointsConfigDump <envoy_v3_api_msg_admin.v3.EndpointsConfigDump>`
	// * “listeners“: :ref:`ListenersConfigDump <envoy_v3_api_msg_admin.v3.ListenersConfigDump>`
	// * “scoped_routes“: :ref:`ScopedRoutesConfigDump <envoy_v3_api_msg_admin.v3.ScopedRoutesConfigDump>`
	// * “routes“:  :ref:`RoutesConfigDump <envoy_v3_api_msg_admin.v3.RoutesConfigDump>`
	// * “secrets“:  :ref:`SecretsConfigDump <envoy_v3_api_msg_admin.v3.SecretsConfigDump>`
	//
	// EDS Configuration will only be dumped by using parameter “?include_eds“
	//
	// Currently ECDS is supported in HTTP and listener filters. Note, ECDS configuration for
	// either HTTP or listener filter will only be dumped if it is actually configured.
	//
	// You can filter output with the resource and mask query parameters.
	// See :ref:`/config_dump?resource={} <operations_admin_interface_config_dump_by_resource>`,
	// :ref:`/config_dump?mask={} <operations_admin_interface_config_dump_by_mask>`,
	// or :ref:`/config_dump?resource={},mask={}
	// <operations_admin_interface_config_dump_by_resource_and_mask>` for more information.
	Configs []*anypb.Any `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *ConfigDump) Reset() {
	*x = ConfigDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigDump) ProtoMessage() {}

func (x *ConfigDump) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigDump.ProtoReflect.Descriptor instead.
func (*ConfigDump) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_config_dump_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigDump) GetConfigs() []*anypb.Any {
	if x != nil {
		return x.Configs
	}
	return nil
}

// This message describes the bootstrap configuration that Envoy was started with. This includes
// any CLI overrides that were merged. Bootstrap configuration information can be used to recreate
// the static portions of an Envoy configuration by reusing the output as the bootstrap
// configuration for another Envoy.
type BootstrapConfigDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bootstrap *v3.Bootstrap `protobuf:"bytes,1,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	// The timestamp when the BootstrapConfig was last updated.
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *BootstrapConfigDump) Reset() {
	*x = BootstrapConfigDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapConfigDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapConfigDump) ProtoMessage() {}

func (x *BootstrapConfigDump) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapConfigDump.ProtoReflect.Descriptor instead.
func (*BootstrapConfigDump) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_config_dump_proto_rawDescGZIP(), []int{1}
}

func (x *BootstrapConfigDump) GetBootstrap() *v3.Bootstrap {
	if x != nil {
		return x.Bootstrap
	}
	return nil
}

func (x *BootstrapConfigDump) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

// Envoys SDS implementation fills this message with all secrets fetched dynamically via SDS.
type SecretsConfigDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The statically loaded secrets.
	StaticSecrets []*SecretsConfigDump_StaticSecret `protobuf:"bytes,1,rep,name=static_secrets,json=staticSecrets,proto3" json:"static_secrets,omitempty"`
	// The dynamically loaded active secrets. These are secrets that are available to service
	// clusters or listeners.
	DynamicActiveSecrets []*SecretsConfigDump_DynamicSecret `protobuf:"bytes,2,rep,name=dynamic_active_secrets,json=dynamicActiveSecrets,proto3" json:"dynamic_active_secrets,omitempty"`
	// The dynamically loaded warming secrets. These are secrets that are currently undergoing
	// warming in preparation to service clusters or listeners.
	DynamicWarmingSecrets []*SecretsConfigDump_DynamicSecret `protobuf:"bytes,3,rep,name=dynamic_warming_secrets,json=dynamicWarmingSecrets,proto3" json:"dynamic_warming_secrets,omitempty"`
}

func (x *SecretsConfigDump) Reset() {
	*x = SecretsConfigDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsConfigDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsConfigDump) ProtoMessage() {}

func (x *SecretsConfigDump) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsConfigDump.ProtoReflect.Descriptor instead.
func (*SecretsConfigDump) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_config_dump_proto_rawDescGZIP(), []int{2}
}

func (x *SecretsConfigDump) GetStaticSecrets() []*SecretsConfigDump_StaticSecret {
	if x != nil {
		return x.StaticSecrets
	}
	return nil
}

func (x *SecretsConfigDump) GetDynamicActiveSecrets() []*SecretsConfigDump_DynamicSecret {
	if x != nil {
		return x.DynamicActiveSecrets
	}
	return nil
}

func (x *SecretsConfigDump) GetDynamicWarmingSecrets() []*SecretsConfigDump_DynamicSecret {
	if x != nil {
		return x.DynamicWarmingSecrets
	}
	return nil
}

// DynamicSecret contains secret information fetched via SDS.
// [#next-free-field: 7]
type SecretsConfigDump_DynamicSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name assigned to the secret.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// This is the per-resource version information.
	VersionInfo string `protobuf:"bytes,2,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The timestamp when the secret was last updated.
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	// The actual secret information.
	// Security sensitive information is redacted (replaced with "[redacted]") for
	// private keys and passwords in TLS certificates.
	Secret *anypb.Any `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	// Set if the last update failed, cleared after the next successful update.
	// The *error_state* field contains the rejected version of this particular
	// resource along with the reason and timestamp. For successfully updated or
	// acknowledged resource, this field should be empty.
	// [#not-implemented-hide:]
	ErrorState *UpdateFailureState `protobuf:"bytes,5,opt,name=error_state,json=errorState,proto3" json:"error_state,omitempty"`
	// The client status of this resource.
	// [#not-implemented-hide:]
	ClientStatus ClientResourceStatus `protobuf:"varint,6,opt,name=client_status,json=clientStatus,proto3,enum=envoy.admin.v3.ClientResourceStatus" json:"client_status,omitempty"`
}

func (x *SecretsConfigDump_DynamicSecret) Reset() {
	*x = SecretsConfigDump_DynamicSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsConfigDump_DynamicSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsConfigDump_DynamicSecret) ProtoMessage() {}

func (x *SecretsConfigDump_DynamicSecret) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsConfigDump_DynamicSecret.ProtoReflect.Descriptor instead.
func (*SecretsConfigDump_DynamicSecret) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_config_dump_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SecretsConfigDump_DynamicSecret) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecretsConfigDump_DynamicSecret) GetVersionInfo() string {
	if x != nil {
		return x.VersionInfo
	}
	return ""
}

func (x *SecretsConfigDump_DynamicSecret) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *SecretsConfigDump_DynamicSecret) GetSecret() *anypb.Any {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *SecretsConfigDump_DynamicSecret) GetErrorState() *UpdateFailureState {
	if x != nil {
		return x.ErrorState
	}
	return nil
}

func (x *SecretsConfigDump_DynamicSecret) GetClientStatus() ClientResourceStatus {
	if x != nil {
		return x.ClientStatus
	}
	return ClientResourceStatus_UNKNOWN
}

// StaticSecret specifies statically loaded secret in bootstrap.
type SecretsConfigDump_StaticSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name assigned to the secret.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The timestamp when the secret was last updated.
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	// The actual secret information.
	// Security sensitive information is redacted (replaced with "[redacted]") for
	// private keys and passwords in TLS certificates.
	Secret *anypb.Any `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *SecretsConfigDump_StaticSecret) Reset() {
	*x = SecretsConfigDump_StaticSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsConfigDump_StaticSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsConfigDump_StaticSecret) ProtoMessage() {}

func (x *SecretsConfigDump_StaticSecret) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_config_dump_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsConfigDump_StaticSecret.ProtoReflect.Descriptor instead.
func (*SecretsConfigDump_StaticSecret) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_config_dump_proto_rawDescGZIP(), []int{2, 1}
}

func (x *SecretsConfigDump_StaticSecret) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecretsConfigDump_StaticSecret) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *SecretsConfigDump_StaticSecret) GetSecret() *anypb.Any {
	if x != nil {
		return x.Secret
	}
	return nil
}

var File_envoy_admin_v3_config_dump_proto protoreflect.FileDescriptor

var file_envoy_admin_v3_config_dump_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x33, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75,
	0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x3a, 0x25, 0x9a, 0xc5, 0x88, 0x1e, 0x20, 0x0a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x22, 0xc8, 0x01, 0x0a, 0x13, 0x42, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d,
	0x70, 0x12, 0x42, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x76, 0x33,
	0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x3a, 0x2e, 0x9a, 0xc5, 0x88, 0x1e, 0x29, 0x0a, 0x27, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x44, 0x75, 0x6d, 0x70, 0x22, 0xb7, 0x07, 0x0a, 0x11, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x55, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x44, 0x75, 0x6d, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x12, 0x65, 0x0a, 0x16, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x44, 0x75, 0x6d, 0x70, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x14, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x67, 0x0a, 0x17, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x2e, 0x44, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x15, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x57, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x1a, 0xff, 0x02, 0x0a, 0x0d, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x49, 0x0a, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x3a, 0x9a, 0xc5, 0x88, 0x1e, 0x35, 0x0a, 0x33,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x44, 0x75, 0x6d, 0x70, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x1a, 0xca, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x3a, 0x39, 0x9a, 0xc5, 0x88, 0x1e, 0x34, 0x0a, 0x32, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44,
	0x75, 0x6d, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x3a, 0x2c, 0x9a, 0xc5, 0x88, 0x1e, 0x27, 0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x42, 0xc6,
	0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x42, 0x0f, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x33, 0xa2,
	0x02, 0x03, 0x45, 0x41, 0x58, 0xaa, 0x02, 0x0e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x0e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x1a, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x3a, 0x3a, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_admin_v3_config_dump_proto_rawDescOnce sync.Once
	file_envoy_admin_v3_config_dump_proto_rawDescData = file_envoy_admin_v3_config_dump_proto_rawDesc
)

func file_envoy_admin_v3_config_dump_proto_rawDescGZIP() []byte {
	file_envoy_admin_v3_config_dump_proto_rawDescOnce.Do(func() {
		file_envoy_admin_v3_config_dump_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_admin_v3_config_dump_proto_rawDescData)
	})
	return file_envoy_admin_v3_config_dump_proto_rawDescData
}

var file_envoy_admin_v3_config_dump_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_admin_v3_config_dump_proto_goTypes = []interface{}{
	(*ConfigDump)(nil),                      // 0: envoy.admin.v3.ConfigDump
	(*BootstrapConfigDump)(nil),             // 1: envoy.admin.v3.BootstrapConfigDump
	(*SecretsConfigDump)(nil),               // 2: envoy.admin.v3.SecretsConfigDump
	(*SecretsConfigDump_DynamicSecret)(nil), // 3: envoy.admin.v3.SecretsConfigDump.DynamicSecret
	(*SecretsConfigDump_StaticSecret)(nil),  // 4: envoy.admin.v3.SecretsConfigDump.StaticSecret
	(*anypb.Any)(nil),                       // 5: google.protobuf.Any
	(*v3.Bootstrap)(nil),                    // 6: envoy.config.bootstrap.v3.Bootstrap
	(*timestamppb.Timestamp)(nil),           // 7: google.protobuf.Timestamp
	(*UpdateFailureState)(nil),              // 8: envoy.admin.v3.UpdateFailureState
	(ClientResourceStatus)(0),               // 9: envoy.admin.v3.ClientResourceStatus
}
var file_envoy_admin_v3_config_dump_proto_depIdxs = []int32{
	5,  // 0: envoy.admin.v3.ConfigDump.configs:type_name -> google.protobuf.Any
	6,  // 1: envoy.admin.v3.BootstrapConfigDump.bootstrap:type_name -> envoy.config.bootstrap.v3.Bootstrap
	7,  // 2: envoy.admin.v3.BootstrapConfigDump.last_updated:type_name -> google.protobuf.Timestamp
	4,  // 3: envoy.admin.v3.SecretsConfigDump.static_secrets:type_name -> envoy.admin.v3.SecretsConfigDump.StaticSecret
	3,  // 4: envoy.admin.v3.SecretsConfigDump.dynamic_active_secrets:type_name -> envoy.admin.v3.SecretsConfigDump.DynamicSecret
	3,  // 5: envoy.admin.v3.SecretsConfigDump.dynamic_warming_secrets:type_name -> envoy.admin.v3.SecretsConfigDump.DynamicSecret
	7,  // 6: envoy.admin.v3.SecretsConfigDump.DynamicSecret.last_updated:type_name -> google.protobuf.Timestamp
	5,  // 7: envoy.admin.v3.SecretsConfigDump.DynamicSecret.secret:type_name -> google.protobuf.Any
	8,  // 8: envoy.admin.v3.SecretsConfigDump.DynamicSecret.error_state:type_name -> envoy.admin.v3.UpdateFailureState
	9,  // 9: envoy.admin.v3.SecretsConfigDump.DynamicSecret.client_status:type_name -> envoy.admin.v3.ClientResourceStatus
	7,  // 10: envoy.admin.v3.SecretsConfigDump.StaticSecret.last_updated:type_name -> google.protobuf.Timestamp
	5,  // 11: envoy.admin.v3.SecretsConfigDump.StaticSecret.secret:type_name -> google.protobuf.Any
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_envoy_admin_v3_config_dump_proto_init() }
func file_envoy_admin_v3_config_dump_proto_init() {
	if File_envoy_admin_v3_config_dump_proto != nil {
		return
	}
	file_envoy_admin_v3_config_dump_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_admin_v3_config_dump_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigDump); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_admin_v3_config_dump_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapConfigDump); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_admin_v3_config_dump_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsConfigDump); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_admin_v3_config_dump_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsConfigDump_DynamicSecret); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_admin_v3_config_dump_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsConfigDump_StaticSecret); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_admin_v3_config_dump_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_admin_v3_config_dump_proto_goTypes,
		DependencyIndexes: file_envoy_admin_v3_config_dump_proto_depIdxs,
		MessageInfos:      file_envoy_admin_v3_config_dump_proto_msgTypes,
	}.Build()
	File_envoy_admin_v3_config_dump_proto = out.File
	file_envoy_admin_v3_config_dump_proto_rawDesc = nil
	file_envoy_admin_v3_config_dump_proto_goTypes = nil
	file_envoy_admin_v3_config_dump_proto_depIdxs = nil
}
