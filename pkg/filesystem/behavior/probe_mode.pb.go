// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: filesystem/behavior/probe_mode.proto

package behavior

import (
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

// ProbeMode specifies the mode for filesystem probing.
type ProbeMode int32

const (
	// ProbeMode_ProbeModeDefault represents an unspecified probe mode. It
	// should be converted to one of the following values based on the desired
	// default behavior.
	ProbeMode_ProbeModeDefault ProbeMode = 0
	// ProbeMode_ProbeModeProbe specifies that filesystem behavior should be
	// determined using temporary files or, if possible, a "fast-path" mechanism
	// (such as filesystem format detection) that provides quick but certain
	// determination of filesystem behavior.
	ProbeMode_ProbeModeProbe ProbeMode = 1
	// ProbeMode_ProbeModeAssume specifies that filesystem behavior should be
	// assumed based on the underlying platform. This is not as accurate as
	// ProbeMode_ProbeModeProbe.
	ProbeMode_ProbeModeAssume ProbeMode = 2
)

// Enum value maps for ProbeMode.
var (
	ProbeMode_name = map[int32]string{
		0: "ProbeModeDefault",
		1: "ProbeModeProbe",
		2: "ProbeModeAssume",
	}
	ProbeMode_value = map[string]int32{
		"ProbeModeDefault": 0,
		"ProbeModeProbe":   1,
		"ProbeModeAssume":  2,
	}
)

func (x ProbeMode) Enum() *ProbeMode {
	p := new(ProbeMode)
	*p = x
	return p
}

func (x ProbeMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProbeMode) Descriptor() protoreflect.EnumDescriptor {
	return file_filesystem_behavior_probe_mode_proto_enumTypes[0].Descriptor()
}

func (ProbeMode) Type() protoreflect.EnumType {
	return &file_filesystem_behavior_probe_mode_proto_enumTypes[0]
}

func (x ProbeMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProbeMode.Descriptor instead.
func (ProbeMode) EnumDescriptor() ([]byte, []int) {
	return file_filesystem_behavior_probe_mode_proto_rawDescGZIP(), []int{0}
}

var File_filesystem_behavior_probe_mode_proto protoreflect.FileDescriptor

var file_filesystem_behavior_probe_mode_proto_rawDesc = []byte{
	0x0a, 0x24, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2a, 0x4a, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x10, 0x02, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67,
	0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filesystem_behavior_probe_mode_proto_rawDescOnce sync.Once
	file_filesystem_behavior_probe_mode_proto_rawDescData = file_filesystem_behavior_probe_mode_proto_rawDesc
)

func file_filesystem_behavior_probe_mode_proto_rawDescGZIP() []byte {
	file_filesystem_behavior_probe_mode_proto_rawDescOnce.Do(func() {
		file_filesystem_behavior_probe_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_filesystem_behavior_probe_mode_proto_rawDescData)
	})
	return file_filesystem_behavior_probe_mode_proto_rawDescData
}

var file_filesystem_behavior_probe_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_filesystem_behavior_probe_mode_proto_goTypes = []any{
	(ProbeMode)(0), // 0: behavior.ProbeMode
}
var file_filesystem_behavior_probe_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_filesystem_behavior_probe_mode_proto_init() }
func file_filesystem_behavior_probe_mode_proto_init() {
	if File_filesystem_behavior_probe_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filesystem_behavior_probe_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filesystem_behavior_probe_mode_proto_goTypes,
		DependencyIndexes: file_filesystem_behavior_probe_mode_proto_depIdxs,
		EnumInfos:         file_filesystem_behavior_probe_mode_proto_enumTypes,
	}.Build()
	File_filesystem_behavior_probe_mode_proto = out.File
	file_filesystem_behavior_probe_mode_proto_rawDesc = nil
	file_filesystem_behavior_probe_mode_proto_goTypes = nil
	file_filesystem_behavior_probe_mode_proto_depIdxs = nil
}
