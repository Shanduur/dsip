// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: messaging.proto

package messaging

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StatusCode int32

const (
	StatusCode_Unknown StatusCode = 0
	StatusCode_Ok      StatusCode = 1
	StatusCode_Failed  StatusCode = 2
	StatusCode_Busy    StatusCode = 3
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
		3: "Busy",
	}
	StatusCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
		"Busy":    3,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_messaging_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_messaging_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{0}
}

type Handshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	BlobSize int32 `protobuf:"varint,2,opt,name=BlobSize,proto3" json:"BlobSize,omitempty"`
}

func (x *Handshake) Reset() {
	*x = Handshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake) ProtoMessage() {}

func (x *Handshake) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake.ProtoReflect.Descriptor instead.
func (*Handshake) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{0}
}

func (x *Handshake) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Handshake) GetBlobSize() int32 {
	if x != nil {
		return x.BlobSize
	}
	return 0
}

type HandshakeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string     `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code    StatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=messaging.StatusCode" json:"Code,omitempty"`
}

func (x *HandshakeStatus) Reset() {
	*x = HandshakeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeStatus) ProtoMessage() {}

func (x *HandshakeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeStatus.ProtoReflect.Descriptor instead.
func (*HandshakeStatus) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{1}
}

func (x *HandshakeStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HandshakeStatus) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_Unknown
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{2}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type UploadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string     `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code    StatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=messaging.StatusCode" json:"Code,omitempty"`
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{3}
}

func (x *UploadStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadStatus) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_Unknown
}

var File_messaging_proto protoreflect.FileDescriptor

var file_messaging_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x3f, 0x0a, 0x09,
	0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x56, 0x0a,
	0x0f, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x21, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x37, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x75, 0x73, 0x79, 0x10, 0x03, 0x32, 0x8d, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x48, 0x61, 0x6e,
	0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x1a, 0x1a, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68,
	0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messaging_proto_rawDescOnce sync.Once
	file_messaging_proto_rawDescData = file_messaging_proto_rawDesc
)

func file_messaging_proto_rawDescGZIP() []byte {
	file_messaging_proto_rawDescOnce.Do(func() {
		file_messaging_proto_rawDescData = protoimpl.X.CompressGZIP(file_messaging_proto_rawDescData)
	})
	return file_messaging_proto_rawDescData
}

var file_messaging_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messaging_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messaging_proto_goTypes = []interface{}{
	(StatusCode)(0),         // 0: messaging.StatusCode
	(*Handshake)(nil),       // 1: messaging.Handshake
	(*HandshakeStatus)(nil), // 2: messaging.HandshakeStatus
	(*Chunk)(nil),           // 3: messaging.Chunk
	(*UploadStatus)(nil),    // 4: messaging.UploadStatus
}
var file_messaging_proto_depIdxs = []int32{
	0, // 0: messaging.HandshakeStatus.Code:type_name -> messaging.StatusCode
	0, // 1: messaging.UploadStatus.Code:type_name -> messaging.StatusCode
	1, // 2: messaging.PuploadService.Handshaker:input_type -> messaging.Handshake
	3, // 3: messaging.PuploadService.Uploader:input_type -> messaging.Chunk
	2, // 4: messaging.PuploadService.Handshaker:output_type -> messaging.HandshakeStatus
	4, // 5: messaging.PuploadService.Uploader:output_type -> messaging.UploadStatus
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_messaging_proto_init() }
func file_messaging_proto_init() {
	if File_messaging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messaging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake); i {
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
		file_messaging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeStatus); i {
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
		file_messaging_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_messaging_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadStatus); i {
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
			RawDescriptor: file_messaging_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messaging_proto_goTypes,
		DependencyIndexes: file_messaging_proto_depIdxs,
		EnumInfos:         file_messaging_proto_enumTypes,
		MessageInfos:      file_messaging_proto_msgTypes,
	}.Build()
	File_messaging_proto = out.File
	file_messaging_proto_rawDesc = nil
	file_messaging_proto_goTypes = nil
	file_messaging_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PuploadServiceClient is the client API for PuploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PuploadServiceClient interface {
	Handshaker(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (*HandshakeStatus, error)
	Uploader(ctx context.Context, opts ...grpc.CallOption) (PuploadService_UploaderClient, error)
}

type puploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPuploadServiceClient(cc grpc.ClientConnInterface) PuploadServiceClient {
	return &puploadServiceClient{cc}
}

func (c *puploadServiceClient) Handshaker(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (*HandshakeStatus, error) {
	out := new(HandshakeStatus)
	err := c.cc.Invoke(ctx, "/messaging.PuploadService/Handshaker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *puploadServiceClient) Uploader(ctx context.Context, opts ...grpc.CallOption) (PuploadService_UploaderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PuploadService_serviceDesc.Streams[0], "/messaging.PuploadService/Uploader", opts...)
	if err != nil {
		return nil, err
	}
	x := &puploadServiceUploaderClient{stream}
	return x, nil
}

type PuploadService_UploaderClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type puploadServiceUploaderClient struct {
	grpc.ClientStream
}

func (x *puploadServiceUploaderClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *puploadServiceUploaderClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PuploadServiceServer is the server API for PuploadService service.
type PuploadServiceServer interface {
	Handshaker(context.Context, *Handshake) (*HandshakeStatus, error)
	Uploader(PuploadService_UploaderServer) error
}

// UnimplementedPuploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPuploadServiceServer struct {
}

func (*UnimplementedPuploadServiceServer) Handshaker(context.Context, *Handshake) (*HandshakeStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshaker not implemented")
}
func (*UnimplementedPuploadServiceServer) Uploader(PuploadService_UploaderServer) error {
	return status.Errorf(codes.Unimplemented, "method Uploader not implemented")
}

func RegisterPuploadServiceServer(s *grpc.Server, srv PuploadServiceServer) {
	s.RegisterService(&_PuploadService_serviceDesc, srv)
}

func _PuploadService_Handshaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Handshake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PuploadServiceServer).Handshaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.PuploadService/Handshaker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PuploadServiceServer).Handshaker(ctx, req.(*Handshake))
	}
	return interceptor(ctx, in, info, handler)
}

func _PuploadService_Uploader_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PuploadServiceServer).Uploader(&puploadServiceUploaderServer{stream})
}

type PuploadService_UploaderServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type puploadServiceUploaderServer struct {
	grpc.ServerStream
}

func (x *puploadServiceUploaderServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *puploadServiceUploaderServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PuploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messaging.PuploadService",
	HandlerType: (*PuploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshaker",
			Handler:    _PuploadService_Handshaker_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Uploader",
			Handler:       _PuploadService_Uploader_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "messaging.proto",
}
