// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.17.3
// source: api/v1/schema.proto

package apiv1

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

type ActionType int32

const (
	ActionType_UNSPECIFIED ActionType = 0
	ActionType_CREATE      ActionType = 1
	ActionType_READ        ActionType = 2
	ActionType_UPDATE      ActionType = 3
	ActionType_DELETE      ActionType = 4
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "CREATE",
		2: "READ",
		3: "UPDATE",
		4: "DELETE",
	}
	ActionType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"CREATE":      1,
		"READ":        2,
		"UPDATE":      3,
		"DELETE":      4,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_schema_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_api_v1_schema_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_schema_proto_rawDescGZIP(), []int{0}
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of action being performed.
	Type ActionType `protobuf:"varint,1,opt,name=type,proto3,enum=clutch.api.v1.ActionType" json:"type,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_api_v1_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetType() ActionType {
	if x != nil {
		return x.Type
	}
	return ActionType_UNSPECIFIED
}

type Pattern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type URL for the resource.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// A string describing the resource name in terms
	// of message members.
	Pattern string `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`
}

func (x *Pattern) Reset() {
	*x = Pattern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pattern) ProtoMessage() {}

func (x *Pattern) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pattern.ProtoReflect.Descriptor instead.
func (*Pattern) Descriptor() ([]byte, []int) {
	return file_api_v1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Pattern) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *Pattern) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

type Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patterns []*Pattern `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
}

func (x *Identifier) Reset() {
	*x = Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identifier) ProtoMessage() {}

func (x *Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identifier.ProtoReflect.Descriptor instead.
func (*Identifier) Descriptor() ([]byte, []int) {
	return file_api_v1_schema_proto_rawDescGZIP(), []int{2}
}

func (x *Identifier) GetPatterns() []*Pattern {
	if x != nil {
		return x.Patterns
	}
	return nil
}

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field(s) which contain nested ResourceIdentifiers to
	// identify resources contained in the message.
	Fields []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_api_v1_schema_proto_rawDescGZIP(), []int{3}
}

func (x *Reference) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

// Message used to represent redacted messages.
type Redacted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedactedTypeUrl string `protobuf:"bytes,1,opt,name=redacted_type_url,json=redactedTypeUrl,proto3" json:"redacted_type_url,omitempty"`
}

func (x *Redacted) Reset() {
	*x = Redacted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redacted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redacted) ProtoMessage() {}

func (x *Redacted) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redacted.ProtoReflect.Descriptor instead.
func (*Redacted) Descriptor() ([]byte, []int) {
	return file_api_v1_schema_proto_rawDescGZIP(), []int{4}
}

func (x *Redacted) GetRedactedTypeUrl() string {
	if x != nil {
		return x.RedactedTypeUrl
	}
	return ""
}

var File_api_v1_schema_proto protoreflect.FileDescriptor

var file_api_v1_schema_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x22, 0x37, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a,
	0x07, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x40, 0x0a,
	0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x08, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x22,
	0x23, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x22, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x61, 0x63, 0x74, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x64,
	0x61, 0x63, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x2a, 0x4b, 0x0a, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_schema_proto_rawDescOnce sync.Once
	file_api_v1_schema_proto_rawDescData = file_api_v1_schema_proto_rawDesc
)

func file_api_v1_schema_proto_rawDescGZIP() []byte {
	file_api_v1_schema_proto_rawDescOnce.Do(func() {
		file_api_v1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_schema_proto_rawDescData)
	})
	return file_api_v1_schema_proto_rawDescData
}

var file_api_v1_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_schema_proto_goTypes = []interface{}{
	(ActionType)(0),    // 0: clutch.api.v1.ActionType
	(*Action)(nil),     // 1: clutch.api.v1.Action
	(*Pattern)(nil),    // 2: clutch.api.v1.Pattern
	(*Identifier)(nil), // 3: clutch.api.v1.Identifier
	(*Reference)(nil),  // 4: clutch.api.v1.Reference
	(*Redacted)(nil),   // 5: clutch.api.v1.Redacted
}
var file_api_v1_schema_proto_depIdxs = []int32{
	0, // 0: clutch.api.v1.Action.type:type_name -> clutch.api.v1.ActionType
	2, // 1: clutch.api.v1.Identifier.patterns:type_name -> clutch.api.v1.Pattern
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_schema_proto_init() }
func file_api_v1_schema_proto_init() {
	if File_api_v1_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_api_v1_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pattern); i {
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
		file_api_v1_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identifier); i {
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
		file_api_v1_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_api_v1_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redacted); i {
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
			RawDescriptor: file_api_v1_schema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_schema_proto_goTypes,
		DependencyIndexes: file_api_v1_schema_proto_depIdxs,
		EnumInfos:         file_api_v1_schema_proto_enumTypes,
		MessageInfos:      file_api_v1_schema_proto_msgTypes,
	}.Build()
	File_api_v1_schema_proto = out.File
	file_api_v1_schema_proto_rawDesc = nil
	file_api_v1_schema_proto_goTypes = nil
	file_api_v1_schema_proto_depIdxs = nil
}
