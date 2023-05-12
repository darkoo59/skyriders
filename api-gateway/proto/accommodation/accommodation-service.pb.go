// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.4
// source: accommodation/accommodation-service.proto

package accommodation

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AM_GetAll_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AM_GetAll_Request) Reset() {
	*x = AM_GetAll_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_accommodation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_GetAll_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_GetAll_Request) ProtoMessage() {}

func (x *AM_GetAll_Request) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_accommodation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_GetAll_Request.ProtoReflect.Descriptor instead.
func (*AM_GetAll_Request) Descriptor() ([]byte, []int) {
	return file_accommodation_accommodation_service_proto_rawDescGZIP(), []int{0}
}

type AM_GetAll_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AM_GetAll_Response) Reset() {
	*x = AM_GetAll_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_accommodation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_GetAll_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_GetAll_Response) ProtoMessage() {}

func (x *AM_GetAll_Response) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_accommodation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_GetAll_Response.ProtoReflect.Descriptor instead.
func (*AM_GetAll_Response) Descriptor() ([]byte, []int) {
	return file_accommodation_accommodation_service_proto_rawDescGZIP(), []int{1}
}

func (x *AM_GetAll_Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type AM_Create_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Benefits  string `protobuf:"bytes,2,opt,name=benefits,proto3" json:"benefits,omitempty"`
	MinGuests int32  `protobuf:"varint,3,opt,name=minGuests,proto3" json:"minGuests,omitempty"`
	MaxGuests int32  `protobuf:"varint,4,opt,name=maxGuests,proto3" json:"maxGuests,omitempty"`
}

func (x *AM_Create_Request) Reset() {
	*x = AM_Create_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_accommodation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_Create_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_Create_Request) ProtoMessage() {}

func (x *AM_Create_Request) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_accommodation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_Create_Request.ProtoReflect.Descriptor instead.
func (*AM_Create_Request) Descriptor() ([]byte, []int) {
	return file_accommodation_accommodation_service_proto_rawDescGZIP(), []int{2}
}

func (x *AM_Create_Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AM_Create_Request) GetBenefits() string {
	if x != nil {
		return x.Benefits
	}
	return ""
}

func (x *AM_Create_Request) GetMinGuests() int32 {
	if x != nil {
		return x.MinGuests
	}
	return 0
}

func (x *AM_Create_Request) GetMaxGuests() int32 {
	if x != nil {
		return x.MaxGuests
	}
	return 0
}

type AM_Create_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AM_Create_Response) Reset() {
	*x = AM_Create_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_accommodation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_Create_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_Create_Response) ProtoMessage() {}

func (x *AM_Create_Response) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_accommodation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_Create_Response.ProtoReflect.Descriptor instead.
func (*AM_Create_Response) Descriptor() ([]byte, []int) {
	return file_accommodation_accommodation_service_proto_rawDescGZIP(), []int{3}
}

func (x *AM_Create_Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_accommodation_accommodation_service_proto protoreflect.FileDescriptor

var file_accommodation_accommodation_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x4d, 0x5f,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28,
	0x0a, 0x12, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7f, 0x0a, 0x11, 0x41, 0x4d, 0x5f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x41, 0x4d, 0x5f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0xaf, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x5f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e,
	0x41, 0x4d, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x41, 0x4d, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01,
	0x2a, 0x22, 0x0e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accommodation_accommodation_service_proto_rawDescOnce sync.Once
	file_accommodation_accommodation_service_proto_rawDescData = file_accommodation_accommodation_service_proto_rawDesc
)

func file_accommodation_accommodation_service_proto_rawDescGZIP() []byte {
	file_accommodation_accommodation_service_proto_rawDescOnce.Do(func() {
		file_accommodation_accommodation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_accommodation_accommodation_service_proto_rawDescData)
	})
	return file_accommodation_accommodation_service_proto_rawDescData
}

var file_accommodation_accommodation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_accommodation_accommodation_service_proto_goTypes = []interface{}{
	(*AM_GetAll_Request)(nil),  // 0: AM_GetAll_Request
	(*AM_GetAll_Response)(nil), // 1: AM_GetAll_Response
	(*AM_Create_Request)(nil),  // 2: AM_Create_Request
	(*AM_Create_Response)(nil), // 3: AM_Create_Response
}
var file_accommodation_accommodation_service_proto_depIdxs = []int32{
	0, // 0: AccommodationService.GetAll:input_type -> AM_GetAll_Request
	2, // 1: AccommodationService.Create:input_type -> AM_Create_Request
	1, // 2: AccommodationService.GetAll:output_type -> AM_GetAll_Response
	3, // 3: AccommodationService.Create:output_type -> AM_Create_Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accommodation_accommodation_service_proto_init() }
func file_accommodation_accommodation_service_proto_init() {
	if File_accommodation_accommodation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accommodation_accommodation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_GetAll_Request); i {
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
		file_accommodation_accommodation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_GetAll_Response); i {
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
		file_accommodation_accommodation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_Create_Request); i {
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
		file_accommodation_accommodation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_Create_Response); i {
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
			RawDescriptor: file_accommodation_accommodation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accommodation_accommodation_service_proto_goTypes,
		DependencyIndexes: file_accommodation_accommodation_service_proto_depIdxs,
		MessageInfos:      file_accommodation_accommodation_service_proto_msgTypes,
	}.Build()
	File_accommodation_accommodation_service_proto = out.File
	file_accommodation_accommodation_service_proto_rawDesc = nil
	file_accommodation_accommodation_service_proto_goTypes = nil
	file_accommodation_accommodation_service_proto_depIdxs = nil
}
