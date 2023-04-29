// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/login.proto

package proto

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

type ClientAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Port      int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ClientAddress) Reset() {
	*x = ClientAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAddress) ProtoMessage() {}

func (x *ClientAddress) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAddress.ProtoReflect.Descriptor instead.
func (*ClientAddress) Descriptor() ([]byte, []int) {
	return file_proto_login_proto_rawDescGZIP(), []int{0}
}

func (x *ClientAddress) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *ClientAddress) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type UserContact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string         `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientAddress *ClientAddress `protobuf:"bytes,2,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
}

func (x *UserContact) Reset() {
	*x = UserContact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserContact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserContact) ProtoMessage() {}

func (x *UserContact) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserContact.ProtoReflect.Descriptor instead.
func (*UserContact) Descriptor() ([]byte, []int) {
	return file_proto_login_proto_rawDescGZIP(), []int{1}
}

func (x *UserContact) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserContact) GetClientAddress() *ClientAddress {
	if x != nil {
		return x.ClientAddress
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string         `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string         `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ClientAddress *ClientAddress `protobuf:"bytes,3,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetClientAddress() *ClientAddress {
	if x != nil {
		return x.ClientAddress
	}
	return nil
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error     *Error            `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Message   string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	UserId    string            `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Follows   []*UserContact    `protobuf:"bytes,4,rep,name=follows,proto3" json:"follows,omitempty"`
	Following []*UserContact    `protobuf:"bytes,5,rep,name=following,proto3" json:"following,omitempty"`
	Blocked   []string          `protobuf:"bytes,6,rep,name=blocked,proto3" json:"blocked,omitempty"`
	Settings  map[string]string `protobuf:"bytes,7,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_login_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *LoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginResponse) GetFollows() []*UserContact {
	if x != nil {
		return x.Follows
	}
	return nil
}

func (x *LoginResponse) GetFollowing() []*UserContact {
	if x != nil {
		return x.Following
	}
	return nil
}

func (x *LoginResponse) GetBlocked() []string {
	if x != nil {
		return x.Blocked
	}
	return nil
}

func (x *LoginResponse) GetSettings() map[string]string {
	if x != nil {
		return x.Settings
	}
	return nil
}

var File_proto_login_proto protoreflect.FileDescriptor

var file_proto_login_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42,
	0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x63, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x3b, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xdd, 0x02,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x73, 0x12, 0x30, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x09, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x12, 0x3e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x42, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x6f, 0x79, 0x6c, 0x65, 0x31, 0x39, 0x37, 0x34, 0x2f, 0x73, 0x65, 0x77, 0x73, 0x68, 0x75,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_login_proto_rawDescOnce sync.Once
	file_proto_login_proto_rawDescData = file_proto_login_proto_rawDesc
)

func file_proto_login_proto_rawDescGZIP() []byte {
	file_proto_login_proto_rawDescOnce.Do(func() {
		file_proto_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_login_proto_rawDescData)
	})
	return file_proto_login_proto_rawDescData
}

var file_proto_login_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_login_proto_goTypes = []interface{}{
	(*ClientAddress)(nil), // 0: proto.ClientAddress
	(*UserContact)(nil),   // 1: proto.UserContact
	(*LoginRequest)(nil),  // 2: proto.LoginRequest
	(*LoginResponse)(nil), // 3: proto.LoginResponse
	nil,                   // 4: proto.LoginResponse.SettingsEntry
	(*Error)(nil),         // 5: proto.Error
}
var file_proto_login_proto_depIdxs = []int32{
	0, // 0: proto.UserContact.client_address:type_name -> proto.ClientAddress
	0, // 1: proto.LoginRequest.client_address:type_name -> proto.ClientAddress
	5, // 2: proto.LoginResponse.error:type_name -> proto.Error
	1, // 3: proto.LoginResponse.follows:type_name -> proto.UserContact
	1, // 4: proto.LoginResponse.following:type_name -> proto.UserContact
	4, // 5: proto.LoginResponse.settings:type_name -> proto.LoginResponse.SettingsEntry
	2, // 6: proto.LoginService.Login:input_type -> proto.LoginRequest
	3, // 7: proto.LoginService.Login:output_type -> proto.LoginResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_login_proto_init() }
func file_proto_login_proto_init() {
	if File_proto_login_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAddress); i {
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
		file_proto_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserContact); i {
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
		file_proto_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_proto_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_proto_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_login_proto_goTypes,
		DependencyIndexes: file_proto_login_proto_depIdxs,
		MessageInfos:      file_proto_login_proto_msgTypes,
	}.Build()
	File_proto_login_proto = out.File
	file_proto_login_proto_rawDesc = nil
	file_proto_login_proto_goTypes = nil
	file_proto_login_proto_depIdxs = nil
}
