// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: buffer.proto

package buffer

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId     uint64 `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Priority     uint32 `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	CleaningType uint32 `protobuf:"varint,4,opt,name=cleaning_type,json=cleaningType,proto3" json:"cleaning_type,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buffer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_buffer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_buffer_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Request) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Request) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Request) GetCleaningType() uint32 {
	if x != nil {
		return x.CleaningType
	}
	return 0
}

type AppendRequestIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Request `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *AppendRequestIn) Reset() {
	*x = AppendRequestIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buffer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendRequestIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendRequestIn) ProtoMessage() {}

func (x *AppendRequestIn) ProtoReflect() protoreflect.Message {
	mi := &file_buffer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendRequestIn.ProtoReflect.Descriptor instead.
func (*AppendRequestIn) Descriptor() ([]byte, []int) {
	return file_buffer_proto_rawDescGZIP(), []int{1}
}

func (x *AppendRequestIn) GetReq() *Request {
	if x != nil {
		return x.Req
	}
	return nil
}

type AppendRequestOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size   uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Status bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AppendRequestOut) Reset() {
	*x = AppendRequestOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buffer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendRequestOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendRequestOut) ProtoMessage() {}

func (x *AppendRequestOut) ProtoReflect() protoreflect.Message {
	mi := &file_buffer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendRequestOut.ProtoReflect.Descriptor instead.
func (*AppendRequestOut) Descriptor() ([]byte, []int) {
	return file_buffer_proto_rawDescGZIP(), []int{2}
}

func (x *AppendRequestOut) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *AppendRequestOut) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type PopTopIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PopTopIn) Reset() {
	*x = PopTopIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buffer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopTopIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopTopIn) ProtoMessage() {}

func (x *PopTopIn) ProtoReflect() protoreflect.Message {
	mi := &file_buffer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopTopIn.ProtoReflect.Descriptor instead.
func (*PopTopIn) Descriptor() ([]byte, []int) {
	return file_buffer_proto_rawDescGZIP(), []int{3}
}

type PopTopOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Request `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *PopTopOut) Reset() {
	*x = PopTopOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buffer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopTopOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopTopOut) ProtoMessage() {}

func (x *PopTopOut) ProtoReflect() protoreflect.Message {
	mi := &file_buffer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopTopOut.ProtoReflect.Descriptor instead.
func (*PopTopOut) Descriptor() ([]byte, []int) {
	return file_buffer_proto_rawDescGZIP(), []int{4}
}

func (x *PopTopOut) GetReq() *Request {
	if x != nil {
		return x.Req
	}
	return nil
}

type PopBottomIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurReq *Request `protobuf:"bytes,1,opt,name=cur_req,json=curReq,proto3" json:"cur_req,omitempty"`
}

func (x *PopBottomIn) Reset() {
	*x = PopBottomIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buffer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopBottomIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopBottomIn) ProtoMessage() {}

func (x *PopBottomIn) ProtoReflect() protoreflect.Message {
	mi := &file_buffer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopBottomIn.ProtoReflect.Descriptor instead.
func (*PopBottomIn) Descriptor() ([]byte, []int) {
	return file_buffer_proto_rawDescGZIP(), []int{5}
}

func (x *PopBottomIn) GetCurReq() *Request {
	if x != nil {
		return x.CurReq
	}
	return nil
}

type PopBottomOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeclinedReq *Request `protobuf:"bytes,1,opt,name=declined_req,json=declinedReq,proto3" json:"declined_req,omitempty"`
	Status      bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PopBottomOut) Reset() {
	*x = PopBottomOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buffer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopBottomOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopBottomOut) ProtoMessage() {}

func (x *PopBottomOut) ProtoReflect() protoreflect.Message {
	mi := &file_buffer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopBottomOut.ProtoReflect.Descriptor instead.
func (*PopBottomOut) Descriptor() ([]byte, []int) {
	return file_buffer_proto_rawDescGZIP(), []int{6}
}

func (x *PopBottomOut) GetDeclinedReq() *Request {
	if x != nil {
		return x.DeclinedReq
	}
	return nil
}

func (x *PopBottomOut) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_buffer_proto protoreflect.FileDescriptor

var file_buffer_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x22, 0x77, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c,
	0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x34, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x12, 0x21, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x3e, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0a, 0x0a, 0x08, 0x50, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x49,
	0x6e, 0x22, 0x2e, 0x0a, 0x09, 0x50, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x4f, 0x75, 0x74, 0x12, 0x21,
	0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x72, 0x65,
	0x71, 0x22, 0x37, 0x0a, 0x0b, 0x50, 0x6f, 0x70, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x49, 0x6e,
	0x12, 0x28, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x06, 0x63, 0x75, 0x72, 0x52, 0x65, 0x71, 0x22, 0x5a, 0x0a, 0x0c, 0x50, 0x6f,
	0x70, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x4f, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x0c, 0x64, 0x65,
	0x63, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0b, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xba, 0x01, 0x0a, 0x0d, 0x42, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x1a, 0x18, 0x2e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x2d, 0x0a, 0x06,
	0x50, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x12, 0x10, 0x2e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e,
	0x50, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x2e, 0x50, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x4f, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x50,
	0x6f, 0x70, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x12, 0x13, 0x2e, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x2e, 0x50, 0x6f, 0x70, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x1a, 0x14, 0x2e,
	0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x70, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d,
	0x4f, 0x75, 0x74, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x42, 0x61, 0x7a, 0x68, 0x65, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buffer_proto_rawDescOnce sync.Once
	file_buffer_proto_rawDescData = file_buffer_proto_rawDesc
)

func file_buffer_proto_rawDescGZIP() []byte {
	file_buffer_proto_rawDescOnce.Do(func() {
		file_buffer_proto_rawDescData = protoimpl.X.CompressGZIP(file_buffer_proto_rawDescData)
	})
	return file_buffer_proto_rawDescData
}

var file_buffer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_buffer_proto_goTypes = []interface{}{
	(*Request)(nil),          // 0: buffer.Request
	(*AppendRequestIn)(nil),  // 1: buffer.AppendRequestIn
	(*AppendRequestOut)(nil), // 2: buffer.AppendRequestOut
	(*PopTopIn)(nil),         // 3: buffer.PopTopIn
	(*PopTopOut)(nil),        // 4: buffer.PopTopOut
	(*PopBottomIn)(nil),      // 5: buffer.PopBottomIn
	(*PopBottomOut)(nil),     // 6: buffer.PopBottomOut
}
var file_buffer_proto_depIdxs = []int32{
	0, // 0: buffer.AppendRequestIn.req:type_name -> buffer.Request
	0, // 1: buffer.PopTopOut.req:type_name -> buffer.Request
	0, // 2: buffer.PopBottomIn.cur_req:type_name -> buffer.Request
	0, // 3: buffer.PopBottomOut.declined_req:type_name -> buffer.Request
	1, // 4: buffer.BufferService.AppendRequest:input_type -> buffer.AppendRequestIn
	3, // 5: buffer.BufferService.PopTop:input_type -> buffer.PopTopIn
	5, // 6: buffer.BufferService.PopBottom:input_type -> buffer.PopBottomIn
	2, // 7: buffer.BufferService.AppendRequest:output_type -> buffer.AppendRequestOut
	4, // 8: buffer.BufferService.PopTop:output_type -> buffer.PopTopOut
	6, // 9: buffer.BufferService.PopBottom:output_type -> buffer.PopBottomOut
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_buffer_proto_init() }
func file_buffer_proto_init() {
	if File_buffer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buffer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_buffer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendRequestIn); i {
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
		file_buffer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendRequestOut); i {
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
		file_buffer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopTopIn); i {
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
		file_buffer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopTopOut); i {
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
		file_buffer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopBottomIn); i {
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
		file_buffer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopBottomOut); i {
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
			RawDescriptor: file_buffer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buffer_proto_goTypes,
		DependencyIndexes: file_buffer_proto_depIdxs,
		MessageInfos:      file_buffer_proto_msgTypes,
	}.Build()
	File_buffer_proto = out.File
	file_buffer_proto_rawDesc = nil
	file_buffer_proto_goTypes = nil
	file_buffer_proto_depIdxs = nil
}
