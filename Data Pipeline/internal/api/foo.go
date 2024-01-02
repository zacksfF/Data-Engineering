package api

import (
	"reflect"
	"sync"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Interim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SepalLength []float64 `protobuf:"fixed64,1,rep,packed,name=SepalLength,proto3" json:"SepalLength,omitempty"`
	SepalWidth  []float64 `protobuf:"fixed64,2,rep,packed,name=SepalWidth,proto3" json:"SepalWidth,omitempty"`
	PetalLength []float64 `protobuf:"fixed64,3,rep,packed,name=PetalLength,proto3" json:"PetalLength,omitempty"`
	PetalWidth  []float64 `protobuf:"fixed64,4,rep,packed,name=PetalWidth,proto3" json:"PetalWidth,omitempty"`
	Species     []string  `protobuf:"bytes,5,rep,name=Species,proto3" json:"Species,omitempty"`
}

func (x *Interim) Reset() {
	*x = Interim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_interim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interim) ProtoMessage() {}

func (x *Interim) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_interim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interim.ProtoReflect.Descriptor instead.
func (*Interim) Descriptor() ([]byte, []int) {
	return file_internal_api_interim_proto_rawDescGZIP(), []int{0}
}

func (x *Interim) GetSepalLength() []float64 {
	if x != nil {
		return x.SepalLength
	}
	return nil
}

func (x *Interim) GetSepalWidth() []float64 {
	if x != nil {
		return x.SepalWidth
	}
	return nil
}

func (x *Interim) GetPetalLength() []float64 {
	if x != nil {
		return x.PetalLength
	}
	return nil
}

func (x *Interim) GetPetalWidth() []float64 {
	if x != nil {
		return x.PetalWidth
	}
	return nil
}

func (x *Interim) GetSpecies() []string {
	if x != nil {
		return x.Species
	}
	return nil
}

var File_internal_api_interim_proto protoreflect.FileDescriptor

var file_internal_api_interim_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a,
	0x07, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x69, 0x6d, 0x12, 0x24, 0x0a, 0x0b, 0x53, 0x65, 0x70, 0x61,
	0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10,
	0x01, 0x52, 0x0b, 0x53, 0x65, 0x70, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x22,
	0x0a, 0x0a, 0x53, 0x65, 0x70, 0x61, 0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x01, 0x42, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x53, 0x65, 0x70, 0x61, 0x6c, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x24, 0x0a, 0x0b, 0x50, 0x65, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x50, 0x65, 0x74,
	0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0a, 0x50, 0x65, 0x74, 0x61,
	0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10, 0x01,
	0x52, 0x0a, 0x50, 0x65, 0x74, 0x61, 0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_api_interim_proto_rawDescOnce sync.Once
	file_internal_api_interim_proto_rawDescData = file_internal_api_interim_proto_rawDesc
)

func file_internal_api_interim_proto_rawDescGZIP() []byte {
	file_internal_api_interim_proto_rawDescOnce.Do(func() {
		file_internal_api_interim_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_interim_proto_rawDescData)
	})
	return file_internal_api_interim_proto_rawDescData
}

var file_internal_api_interim_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_api_interim_proto_goTypes = []interface{}{
	(*Interim)(nil), // 0: Interim
}
var file_internal_api_interim_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_api_interim_proto_init() }
func file_internal_api_interim_proto_init() {
	if File_internal_api_interim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_api_interim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interim); i {
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
			RawDescriptor: file_internal_api_interim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_api_interim_proto_goTypes,
		DependencyIndexes: file_internal_api_interim_proto_depIdxs,
		MessageInfos:      file_internal_api_interim_proto_msgTypes,
	}.Build()
	File_internal_api_interim_proto = out.File
	file_internal_api_interim_proto_rawDesc = nil
	file_internal_api_interim_proto_goTypes = nil
	file_internal_api_interim_proto_depIdxs = nil
}
