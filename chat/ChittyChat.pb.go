//
//OBS Charlotte - i protofilen har jeg lavet et Mock timestamp, som bare er en int. Når du skal
//implementerer den del skal du generere en ny .pb-fil med kommandoen (i den rigtige mappe hvor
//proto-filen er):
//
//- protoc --go_out=plugins=grpc:chat ChittyChat.proto
//
//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ChittyChat.proto
//kommandoen laver to filer - det er kun ChittyChat.pb-filen der skal beholdes - den anden fil hedder ..._grpc.pb.
//
//Den laver nok pbfilen i en forkert mappe, men den skal bare rykkes til den rigtige chat-mappe og
//overskrive den der er, så kan du slette den "nye" generede mappe. Herfra skal koden ændres, der
//hvor der vil være fejl til at modtage en timestamp istedet for en en int som nu. J
//eg efter lader kommentarer hvor det er med ordne "TIMECHAR".
//
//SLET DENNE KOMMENTAR INDEN AFLEVERING!!!!

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: ChittyChat.proto

package chat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
//Publish message - the reason for sending both id and name, is to confirm server data, and to prepare
//client for a change username feature.
type ClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId         int32  `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	UserName         string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Msg              string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	LamportTimestamp int32  `protobuf:"varint,4,opt,name=lamport_timestamp,json=lamportTimestamp,proto3" json:"lamport_timestamp,omitempty"` // TIMECHAR: Turn this into a timestamp object = 4;
}

func (x *ClientMessage) Reset() {
	*x = ClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChittyChat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMessage) ProtoMessage() {}

func (x *ClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ChittyChat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMessage.ProtoReflect.Descriptor instead.
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return file_ChittyChat_proto_rawDescGZIP(), []int{0}
}

func (x *ClientMessage) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ClientMessage) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ClientMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ClientMessage) GetLamportTimestamp() int32 {
	if x != nil {
		return x.LamportTimestamp
	}
	return 0
}

//
//Message from the user to the server -> chatroom.
type ChatRoomMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg              string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	LamportTimestamp int32  `protobuf:"varint,2,opt,name=lamport_timestamp,json=lamportTimestamp,proto3" json:"lamport_timestamp,omitempty"` // TIMECHAR: Turn this into a timestamp object= 2;
	Username         string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ChatRoomMessages) Reset() {
	*x = ChatRoomMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChittyChat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomMessages) ProtoMessage() {}

func (x *ChatRoomMessages) ProtoReflect() protoreflect.Message {
	mi := &file_ChittyChat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoomMessages.ProtoReflect.Descriptor instead.
func (*ChatRoomMessages) Descriptor() ([]byte, []int) {
	return file_ChittyChat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatRoomMessages) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ChatRoomMessages) GetLamportTimestamp() int32 {
	if x != nil {
		return x.LamportTimestamp
	}
	return 0
}

func (x *ChatRoomMessages) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

//
//String describing what the server tried to do. Enum with a status - succes/failure/invalid.
type StatusMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Status    Status `protobuf:"varint,2,opt,name=status,proto3,enum=main.Status" json:"status,omitempty"`
	NewId     *int32 `protobuf:"varint,3,opt,name=NewId,proto3,oneof" json:"NewId,omitempty"`
}

func (x *StatusMessage) Reset() {
	*x = StatusMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChittyChat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusMessage) ProtoMessage() {}

func (x *StatusMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ChittyChat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusMessage.ProtoReflect.Descriptor instead.
func (*StatusMessage) Descriptor() ([]byte, []int) {
	return file_ChittyChat_proto_rawDescGZIP(), []int{2}
}

func (x *StatusMessage) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *StatusMessage) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INVALID
}

func (x *StatusMessage) GetNewId() int32 {
	if x != nil && x.NewId != nil {
		return *x.NewId
	}
	return 0
}

//
//Username for subscribing/cennecting to a server. The id is optional and there to help the server
//avoid dublicates and change usernames.
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   *int32 `protobuf:"varint,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChittyChat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChittyChat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_ChittyChat_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

var File_ChittyChat_proto protoreflect.FileDescriptor

var file_ChittyChat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2b,
	0x0a, 0x11, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6c, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x6d, 0x0a, 0x10, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x78, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x19, 0x0a, 0x05, 0x4e, 0x65, 0x77, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x05, 0x4e, 0x65, 0x77, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x4e,
	0x65, 0x77, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x32, 0xf0, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChittyChat_proto_rawDescOnce sync.Once
	file_ChittyChat_proto_rawDescData = file_ChittyChat_proto_rawDesc
)

func file_ChittyChat_proto_rawDescGZIP() []byte {
	file_ChittyChat_proto_rawDescOnce.Do(func() {
		file_ChittyChat_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChittyChat_proto_rawDescData)
	})
	return file_ChittyChat_proto_rawDescData
}

var file_ChittyChat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ChittyChat_proto_goTypes = []interface{}{
	(*ClientMessage)(nil),    // 0: main.ClientMessage
	(*ChatRoomMessages)(nil), // 1: main.ChatRoomMessages
	(*StatusMessage)(nil),    // 2: main.StatusMessage
	(*UserInfo)(nil),         // 3: main.UserInfo
	(Status)(0),              // 4: main.Status
	(*emptypb.Empty)(nil),    // 5: google.protobuf.Empty
}
var file_ChittyChat_proto_depIdxs = []int32{
	4, // 0: main.StatusMessage.status:type_name -> main.Status
	0, // 1: main.ChittyChatService.Publish:input_type -> main.ClientMessage
	5, // 2: main.ChittyChatService.Broadcast:input_type -> google.protobuf.Empty
	3, // 3: main.ChittyChatService.Connect:input_type -> main.UserInfo
	3, // 4: main.ChittyChatService.Disconnect:input_type -> main.UserInfo
	2, // 5: main.ChittyChatService.Publish:output_type -> main.StatusMessage
	1, // 6: main.ChittyChatService.Broadcast:output_type -> main.ChatRoomMessages
	2, // 7: main.ChittyChatService.Connect:output_type -> main.StatusMessage
	2, // 8: main.ChittyChatService.Disconnect:output_type -> main.StatusMessage
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChittyChat_proto_init() }
func file_ChittyChat_proto_init() {
	if File_ChittyChat_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChittyChat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMessage); i {
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
		file_ChittyChat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoomMessages); i {
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
		file_ChittyChat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusMessage); i {
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
		file_ChittyChat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
	file_ChittyChat_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_ChittyChat_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChittyChat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ChittyChat_proto_goTypes,
		DependencyIndexes: file_ChittyChat_proto_depIdxs,
		MessageInfos:      file_ChittyChat_proto_msgTypes,
	}.Build()
	File_ChittyChat_proto = out.File
	file_ChittyChat_proto_rawDesc = nil
	file_ChittyChat_proto_goTypes = nil
	file_ChittyChat_proto_depIdxs = nil
}
