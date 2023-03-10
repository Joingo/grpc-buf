// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/demo/v1/demo.proto

package demov1

import (
	datetime "google.golang.org/genproto/googleapis/type/datetime"
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

type IMPORTANCE int32

const (
	IMPORTANCE_IMPORTANCE_UNSPECIFIED IMPORTANCE = 0
	IMPORTANCE_IMPORTANCE_LOW         IMPORTANCE = 1
	IMPORTANCE_IMPORTANCE_MEDIUM      IMPORTANCE = 2
	IMPORTANCE_IMPORTANCE_HIGH        IMPORTANCE = 3
	IMPORTANCE_IMPORTANCE_EMERGENCY   IMPORTANCE = 4
)

// Enum value maps for IMPORTANCE.
var (
	IMPORTANCE_name = map[int32]string{
		0: "IMPORTANCE_UNSPECIFIED",
		1: "IMPORTANCE_LOW",
		2: "IMPORTANCE_MEDIUM",
		3: "IMPORTANCE_HIGH",
		4: "IMPORTANCE_EMERGENCY",
	}
	IMPORTANCE_value = map[string]int32{
		"IMPORTANCE_UNSPECIFIED": 0,
		"IMPORTANCE_LOW":         1,
		"IMPORTANCE_MEDIUM":      2,
		"IMPORTANCE_HIGH":        3,
		"IMPORTANCE_EMERGENCY":   4,
	}
)

func (x IMPORTANCE) Enum() *IMPORTANCE {
	p := new(IMPORTANCE)
	*p = x
	return p
}

func (x IMPORTANCE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IMPORTANCE) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_demo_v1_demo_proto_enumTypes[0].Descriptor()
}

func (IMPORTANCE) Type() protoreflect.EnumType {
	return &file_proto_demo_v1_demo_proto_enumTypes[0]
}

func (x IMPORTANCE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IMPORTANCE.Descriptor instead.
func (IMPORTANCE) EnumDescriptor() ([]byte, []int) {
	return file_proto_demo_v1_demo_proto_rawDescGZIP(), []int{0}
}

type GetDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt *datetime.DateTime `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GetDemoRequest) Reset() {
	*x = GetDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_demo_v1_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoRequest) ProtoMessage() {}

func (x *GetDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_demo_v1_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoRequest.ProtoReflect.Descriptor instead.
func (*GetDemoRequest) Descriptor() ([]byte, []int) {
	return file_proto_demo_v1_demo_proto_rawDescGZIP(), []int{0}
}

func (x *GetDemoRequest) GetCreatedAt() *datetime.DateTime {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetDemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt *datetime.DateTime `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Reply     *Reply             `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *GetDemoResponse) Reset() {
	*x = GetDemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_demo_v1_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoResponse) ProtoMessage() {}

func (x *GetDemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_demo_v1_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoResponse.ProtoReflect.Descriptor instead.
func (*GetDemoResponse) Descriptor() ([]byte, []int) {
	return file_proto_demo_v1_demo_proto_rawDescGZIP(), []int{1}
}

func (x *GetDemoResponse) GetCreatedAt() *datetime.DateTime {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetDemoResponse) GetReply() *Reply {
	if x != nil {
		return x.Reply
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority IMPORTANCE `protobuf:"varint,1,opt,name=priority,proto3,enum=proto.demo.v1.IMPORTANCE" json:"priority,omitempty"`
	From     string     `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To       string     `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Subject  string     `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Body     string     `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_demo_v1_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_demo_v1_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_proto_demo_v1_demo_proto_rawDescGZIP(), []int{2}
}

func (x *Reply) GetPriority() IMPORTANCE {
	if x != nil {
		return x.Priority
	}
	return IMPORTANCE_IMPORTANCE_UNSPECIFIED
}

func (x *Reply) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Reply) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Reply) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Reply) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_proto_demo_v1_demo_proto protoreflect.FileDescriptor

var file_proto_demo_v1_demo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x73, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x4d, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x2a, 0x82, 0x01, 0x0a, 0x0a, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54,
	0x41, 0x4e, 0x43, 0x45, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4e,
	0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x4c,
	0x4f, 0x57, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4e,
	0x43, 0x45, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x49,
	0x4d, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x03,
	0x12, 0x18, 0x0a, 0x14, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x45,
	0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x04, 0x32, 0x59, 0x0a, 0x0b, 0x44, 0x65,
	0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xaa, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x44, 0x65, 0x6d,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x6f, 0x69, 0x6e, 0x67, 0x6f, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x6d, 0x6f, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x50, 0x44, 0x58, 0xaa, 0x02, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6d,
	0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x44, 0x65, 0x6d,
	0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x44, 0x65, 0x6d,
	0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x44, 0x65, 0x6d, 0x6f, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_demo_v1_demo_proto_rawDescOnce sync.Once
	file_proto_demo_v1_demo_proto_rawDescData = file_proto_demo_v1_demo_proto_rawDesc
)

func file_proto_demo_v1_demo_proto_rawDescGZIP() []byte {
	file_proto_demo_v1_demo_proto_rawDescOnce.Do(func() {
		file_proto_demo_v1_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_demo_v1_demo_proto_rawDescData)
	})
	return file_proto_demo_v1_demo_proto_rawDescData
}

var file_proto_demo_v1_demo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_demo_v1_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_demo_v1_demo_proto_goTypes = []interface{}{
	(IMPORTANCE)(0),           // 0: proto.demo.v1.IMPORTANCE
	(*GetDemoRequest)(nil),    // 1: proto.demo.v1.GetDemoRequest
	(*GetDemoResponse)(nil),   // 2: proto.demo.v1.GetDemoResponse
	(*Reply)(nil),             // 3: proto.demo.v1.Reply
	(*datetime.DateTime)(nil), // 4: google.type.DateTime
}
var file_proto_demo_v1_demo_proto_depIdxs = []int32{
	4, // 0: proto.demo.v1.GetDemoRequest.created_at:type_name -> google.type.DateTime
	4, // 1: proto.demo.v1.GetDemoResponse.created_at:type_name -> google.type.DateTime
	3, // 2: proto.demo.v1.GetDemoResponse.reply:type_name -> proto.demo.v1.Reply
	0, // 3: proto.demo.v1.Reply.priority:type_name -> proto.demo.v1.IMPORTANCE
	1, // 4: proto.demo.v1.DemoService.GetDemo:input_type -> proto.demo.v1.GetDemoRequest
	2, // 5: proto.demo.v1.DemoService.GetDemo:output_type -> proto.demo.v1.GetDemoResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_demo_v1_demo_proto_init() }
func file_proto_demo_v1_demo_proto_init() {
	if File_proto_demo_v1_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_demo_v1_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDemoRequest); i {
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
		file_proto_demo_v1_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDemoResponse); i {
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
		file_proto_demo_v1_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_proto_demo_v1_demo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_demo_v1_demo_proto_goTypes,
		DependencyIndexes: file_proto_demo_v1_demo_proto_depIdxs,
		EnumInfos:         file_proto_demo_v1_demo_proto_enumTypes,
		MessageInfos:      file_proto_demo_v1_demo_proto_msgTypes,
	}.Build()
	File_proto_demo_v1_demo_proto = out.File
	file_proto_demo_v1_demo_proto_rawDesc = nil
	file_proto_demo_v1_demo_proto_goTypes = nil
	file_proto_demo_v1_demo_proto_depIdxs = nil
}
