// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: envoy/admin/v3/memory.proto

package adminv3

import (
	_ "github.com/cncf/xds/go/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Proto representation of the internal memory consumption of an Envoy instance. These represent
// values extracted from an internal TCMalloc instance. For more information, see the section of the
// docs entitled ["Generic Tcmalloc Status"](https://gperftools.github.io/gperftools/tcmalloc.html).
// [#next-free-field: 7]
type Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of bytes allocated by the heap for Envoy. This is an alias for
	// “generic.current_allocated_bytes“.
	Allocated uint64 `protobuf:"varint,1,opt,name=allocated,proto3" json:"allocated,omitempty"`
	// The number of bytes reserved by the heap but not necessarily allocated. This is an alias for
	// “generic.heap_size“.
	HeapSize uint64 `protobuf:"varint,2,opt,name=heap_size,json=heapSize,proto3" json:"heap_size,omitempty"`
	// The number of bytes in free, unmapped pages in the page heap. These bytes always count towards
	// virtual memory usage, and depending on the OS, typically do not count towards physical memory
	// usage. This is an alias for “tcmalloc.pageheap_unmapped_bytes“.
	PageheapUnmapped uint64 `protobuf:"varint,3,opt,name=pageheap_unmapped,json=pageheapUnmapped,proto3" json:"pageheap_unmapped,omitempty"`
	// The number of bytes in free, mapped pages in the page heap. These bytes always count towards
	// virtual memory usage, and unless the underlying memory is swapped out by the OS, they also
	// count towards physical memory usage. This is an alias for “tcmalloc.pageheap_free_bytes“.
	PageheapFree uint64 `protobuf:"varint,4,opt,name=pageheap_free,json=pageheapFree,proto3" json:"pageheap_free,omitempty"`
	// The amount of memory used by the TCMalloc thread caches (for small objects). This is an alias
	// for “tcmalloc.current_total_thread_cache_bytes“.
	TotalThreadCache uint64 `protobuf:"varint,5,opt,name=total_thread_cache,json=totalThreadCache,proto3" json:"total_thread_cache,omitempty"`
	// The number of bytes of the physical memory usage by the allocator. This is an alias for
	// “generic.total_physical_bytes“.
	TotalPhysicalBytes uint64 `protobuf:"varint,6,opt,name=total_physical_bytes,json=totalPhysicalBytes,proto3" json:"total_physical_bytes,omitempty"`
}

func (x *Memory) Reset() {
	*x = Memory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_memory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memory) ProtoMessage() {}

func (x *Memory) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_memory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memory.ProtoReflect.Descriptor instead.
func (*Memory) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_memory_proto_rawDescGZIP(), []int{0}
}

func (x *Memory) GetAllocated() uint64 {
	if x != nil {
		return x.Allocated
	}
	return 0
}

func (x *Memory) GetHeapSize() uint64 {
	if x != nil {
		return x.HeapSize
	}
	return 0
}

func (x *Memory) GetPageheapUnmapped() uint64 {
	if x != nil {
		return x.PageheapUnmapped
	}
	return 0
}

func (x *Memory) GetPageheapFree() uint64 {
	if x != nil {
		return x.PageheapFree
	}
	return 0
}

func (x *Memory) GetTotalThreadCache() uint64 {
	if x != nil {
		return x.TotalThreadCache
	}
	return 0
}

func (x *Memory) GetTotalPhysicalBytes() uint64 {
	if x != nil {
		return x.TotalPhysicalBytes
	}
	return 0
}

var File_envoy_admin_v3_memory_proto protoreflect.FileDescriptor

var file_envoy_admin_v3_memory_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x98, 0x02, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x70,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x68, 0x65, 0x61,
	0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x61, 0x67, 0x65, 0x68, 0x65, 0x61,
	0x70, 0x5f, 0x75, 0x6e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x10, 0x70, 0x61, 0x67, 0x65, 0x68, 0x65, 0x61, 0x70, 0x55, 0x6e, 0x6d, 0x61, 0x70, 0x70,
	0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x67, 0x65, 0x68, 0x65, 0x61, 0x70, 0x5f, 0x66,
	0x72, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x68,
	0x65, 0x61, 0x70, 0x46, 0x72, 0x65, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63,
	0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x3a, 0x21, 0x9a, 0xc5, 0x88, 0x1e, 0x1c, 0x0a, 0x1a,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0xc2, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x3b,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x33, 0xa2, 0x02, 0x03, 0x45, 0x41, 0x58, 0xaa, 0x02, 0x0e,
	0x45, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x33, 0xca, 0x02,
	0x0e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x33, 0xe2,
	0x02, 0x1a, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x33,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45,
	0x6e, 0x76, 0x6f, 0x79, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_admin_v3_memory_proto_rawDescOnce sync.Once
	file_envoy_admin_v3_memory_proto_rawDescData = file_envoy_admin_v3_memory_proto_rawDesc
)

func file_envoy_admin_v3_memory_proto_rawDescGZIP() []byte {
	file_envoy_admin_v3_memory_proto_rawDescOnce.Do(func() {
		file_envoy_admin_v3_memory_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_admin_v3_memory_proto_rawDescData)
	})
	return file_envoy_admin_v3_memory_proto_rawDescData
}

var file_envoy_admin_v3_memory_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_admin_v3_memory_proto_goTypes = []interface{}{
	(*Memory)(nil), // 0: envoy.admin.v3.Memory
}
var file_envoy_admin_v3_memory_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_admin_v3_memory_proto_init() }
func file_envoy_admin_v3_memory_proto_init() {
	if File_envoy_admin_v3_memory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_admin_v3_memory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memory); i {
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
			RawDescriptor: file_envoy_admin_v3_memory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_admin_v3_memory_proto_goTypes,
		DependencyIndexes: file_envoy_admin_v3_memory_proto_depIdxs,
		MessageInfos:      file_envoy_admin_v3_memory_proto_msgTypes,
	}.Build()
	File_envoy_admin_v3_memory_proto = out.File
	file_envoy_admin_v3_memory_proto_rawDesc = nil
	file_envoy_admin_v3_memory_proto_goTypes = nil
	file_envoy_admin_v3_memory_proto_depIdxs = nil
}
