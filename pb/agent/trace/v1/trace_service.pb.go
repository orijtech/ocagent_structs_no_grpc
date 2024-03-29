// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opencensus/proto/agent/trace/v1/trace_service.proto

package v1

import (
	context "context"
	fmt "fmt"

	v1 "github.com/orijtech/ocagent_structs_no_grpc/pb/agent/common/v1"
	v12 "github.com/orijtech/ocagent_structs_no_grpc/pb/resource/v1"
	v11 "github.com/orijtech/ocagent_structs_no_grpc/pb/trace/v1"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CurrentLibraryConfig struct {
	// This is required only in the first message on the stream or if the
	// previous sent CurrentLibraryConfig message has a different Node (e.g.
	// when the same RPC is used to configure multiple Applications).
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Current configuration.
	Config               *v11.TraceConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CurrentLibraryConfig) Reset()         { *m = CurrentLibraryConfig{} }
func (m *CurrentLibraryConfig) String() string { return proto.CompactTextString(m) }
func (*CurrentLibraryConfig) ProtoMessage()    {}
func (*CurrentLibraryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7027f99caf7ac6a5, []int{0}
}

func (m *CurrentLibraryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentLibraryConfig.Unmarshal(m, b)
}
func (m *CurrentLibraryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentLibraryConfig.Marshal(b, m, deterministic)
}
func (m *CurrentLibraryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentLibraryConfig.Merge(m, src)
}
func (m *CurrentLibraryConfig) XXX_Size() int {
	return xxx_messageInfo_CurrentLibraryConfig.Size(m)
}
func (m *CurrentLibraryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentLibraryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentLibraryConfig proto.InternalMessageInfo

func (m *CurrentLibraryConfig) GetNode() *v1.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *CurrentLibraryConfig) GetConfig() *v11.TraceConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type UpdatedLibraryConfig struct {
	// This field is ignored when the RPC is used to configure only one Application.
	// This is required only in the first message on the stream or if the
	// previous sent UpdatedLibraryConfig message has a different Node.
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Requested updated configuration.
	Config               *v11.TraceConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdatedLibraryConfig) Reset()         { *m = UpdatedLibraryConfig{} }
func (m *UpdatedLibraryConfig) String() string { return proto.CompactTextString(m) }
func (*UpdatedLibraryConfig) ProtoMessage()    {}
func (*UpdatedLibraryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7027f99caf7ac6a5, []int{1}
}

func (m *UpdatedLibraryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatedLibraryConfig.Unmarshal(m, b)
}
func (m *UpdatedLibraryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatedLibraryConfig.Marshal(b, m, deterministic)
}
func (m *UpdatedLibraryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatedLibraryConfig.Merge(m, src)
}
func (m *UpdatedLibraryConfig) XXX_Size() int {
	return xxx_messageInfo_UpdatedLibraryConfig.Size(m)
}
func (m *UpdatedLibraryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatedLibraryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatedLibraryConfig proto.InternalMessageInfo

func (m *UpdatedLibraryConfig) GetNode() *v1.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *UpdatedLibraryConfig) GetConfig() *v11.TraceConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type ExportTraceServiceRequest struct {
	// This is required only in the first message on the stream or if the
	// previous sent ExportTraceServiceRequest message has a different Node (e.g.
	// when the same RPC is used to send Spans from multiple Applications).
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// A list of Spans that belong to the last received Node.
	Spans []*v11.Span `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	// The resource for the spans in this message that do not have an explicit
	// resource set.
	// If unset, the most recently set resource in the RPC stream applies. It is
	// valid to never be set within a stream, e.g. when no resource info is known.
	Resource             *v12.Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportTraceServiceRequest) Reset()         { *m = ExportTraceServiceRequest{} }
func (m *ExportTraceServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportTraceServiceRequest) ProtoMessage()    {}
func (*ExportTraceServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7027f99caf7ac6a5, []int{2}
}

func (m *ExportTraceServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportTraceServiceRequest.Unmarshal(m, b)
}
func (m *ExportTraceServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportTraceServiceRequest.Marshal(b, m, deterministic)
}
func (m *ExportTraceServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportTraceServiceRequest.Merge(m, src)
}
func (m *ExportTraceServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ExportTraceServiceRequest.Size(m)
}
func (m *ExportTraceServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportTraceServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportTraceServiceRequest proto.InternalMessageInfo

func (m *ExportTraceServiceRequest) GetNode() *v1.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *ExportTraceServiceRequest) GetSpans() []*v11.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *ExportTraceServiceRequest) GetResource() *v12.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type ExportTraceServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportTraceServiceResponse) Reset()         { *m = ExportTraceServiceResponse{} }
func (m *ExportTraceServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportTraceServiceResponse) ProtoMessage()    {}
func (*ExportTraceServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7027f99caf7ac6a5, []int{3}
}

func (m *ExportTraceServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportTraceServiceResponse.Unmarshal(m, b)
}
func (m *ExportTraceServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportTraceServiceResponse.Marshal(b, m, deterministic)
}
func (m *ExportTraceServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportTraceServiceResponse.Merge(m, src)
}
func (m *ExportTraceServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ExportTraceServiceResponse.Size(m)
}
func (m *ExportTraceServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportTraceServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportTraceServiceResponse proto.InternalMessageInfo

var fileDescriptor_7027f99caf7ac6a5 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcf, 0xaa, 0xd4, 0x30,
	0x14, 0xc6, 0x4d, 0xaf, 0x16, 0xc9, 0x75, 0x63, 0x71, 0x51, 0x8b, 0x30, 0x97, 0x82, 0x32, 0xa0,
	0x4d, 0xed, 0x5c, 0xee, 0xe6, 0x0a, 0x82, 0x33, 0x08, 0x2e, 0x44, 0x2f, 0x1d, 0x75, 0xe1, 0x66,
	0xe8, 0xb4, 0xc7, 0xda, 0xc5, 0x24, 0x31, 0x49, 0x8b, 0x82, 0x7b, 0xf7, 0x2e, 0x7c, 0x03, 0x5f,
	0xc8, 0xc7, 0xf0, 0x29, 0xa4, 0x39, 0x9d, 0x3f, 0x3a, 0x53, 0x0b, 0xba, 0xb9, 0xbb, 0x43, 0xf3,
	0xfd, 0xbe, 0xf3, 0x25, 0x39, 0x29, 0x3d, 0x15, 0x12, 0x78, 0x0e, 0x5c, 0xd7, 0x3a, 0x96, 0x4a,
	0x18, 0x11, 0x67, 0x25, 0x70, 0x13, 0x1b, 0x95, 0xe5, 0x10, 0x37, 0x09, 0x16, 0x0b, 0x0d, 0xaa,
	0xa9, 0x72, 0x60, 0x56, 0xe2, 0x8d, 0xb6, 0x10, 0x7e, 0x61, 0x16, 0x62, 0x56, 0xcb, 0x9a, 0x24,
	0x88, 0x7a, 0x5c, 0x73, 0xb1, 0x5a, 0x09, 0xde, 0xda, 0x62, 0x85, 0x74, 0x70, 0x7f, 0x4f, 0xae,
	0x40, 0x8b, 0x5a, 0x61, 0x82, 0x75, 0xdd, 0x89, 0xef, 0xee, 0x89, 0x7f, 0xcf, 0xda, 0xc9, 0x1e,
	0x0c, 0xc8, 0x16, 0xb9, 0xe0, 0xef, 0xaa, 0x12, 0xd5, 0xe1, 0x57, 0x42, 0x6f, 0xcd, 0x6a, 0xa5,
	0x80, 0x9b, 0xe7, 0xd5, 0x52, 0x65, 0xea, 0xd3, 0xcc, 0x2e, 0x7b, 0xe7, 0xf4, 0x2a, 0x17, 0x05,
	0xf8, 0xe4, 0x84, 0x8c, 0x8f, 0x27, 0xf7, 0x58, 0xcf, 0xce, 0xbb, 0xed, 0x34, 0x09, 0x7b, 0x21,
	0x0a, 0x48, 0x2d, 0xe3, 0x3d, 0xa6, 0x2e, 0x36, 0xf1, 0x9d, 0x3e, 0x7a, 0x7d, 0x62, 0xec, 0x55,
	0x5b, 0x60, 0xcf, 0xb4, 0xa3, 0x6c, 0xa8, 0xd7, 0xb2, 0xc8, 0x0c, 0x14, 0x97, 0x27, 0xd4, 0x0f,
	0x42, 0x6f, 0x3f, 0xfd, 0x28, 0x85, 0x32, 0x76, 0x75, 0x8e, 0x83, 0x91, 0xc2, 0x87, 0x1a, 0xb4,
	0xf9, 0xaf, 0x64, 0x67, 0xf4, 0x9a, 0x96, 0x19, 0xd7, 0xbe, 0x73, 0x72, 0x34, 0x3e, 0x9e, 0x8c,
	0xfe, 0x12, 0x6c, 0x2e, 0x33, 0x9e, 0xa2, 0xda, 0x9b, 0xd2, 0xeb, 0xeb, 0x09, 0xf1, 0x8f, 0xfa,
	0xda, 0x6e, 0x66, 0xa8, 0x49, 0x58, 0xda, 0xd5, 0xe9, 0x86, 0x0b, 0xef, 0xd0, 0xe0, 0xd0, 0x9e,
	0xb4, 0x14, 0x5c, 0xc3, 0xe4, 0x9b, 0x43, 0x6f, 0xec, 0x2e, 0x78, 0x9f, 0xa9, 0xdb, 0xdd, 0xc4,
	0x19, 0x1b, 0x78, 0x0a, 0xec, 0xd0, 0x54, 0x05, 0xc3, 0xd8, 0xa1, 0x7b, 0x0f, 0xaf, 0x8c, 0xc9,
	0x43, 0xe2, 0x7d, 0x21, 0xd4, 0xc5, 0xb4, 0xde, 0xf9, 0xa0, 0x4f, 0xef, 0x55, 0x05, 0x8f, 0xfe,
	0x89, 0xc5, 0x23, 0xc1, 0x24, 0xd3, 0xef, 0x84, 0x86, 0x95, 0x18, 0xf2, 0x99, 0xde, 0xdc, 0xb5,
	0xb8, 0x68, 0x15, 0x17, 0xe4, 0xed, 0xb3, 0xb2, 0x32, 0xef, 0xeb, 0x65, 0x3b, 0x0a, 0x31, 0xc2,
	0x51, 0xc5, 0xb5, 0x51, 0xf5, 0x0a, 0xb8, 0xc9, 0x4c, 0x25, 0x78, 0xbc, 0xf5, 0x8d, 0xf0, 0x05,
	0x97, 0xc0, 0xa3, 0xf2, 0xcf, 0x3f, 0xd4, 0x4f, 0x67, 0xf4, 0x52, 0x02, 0x9f, 0x61, 0x00, 0x6b,
	0xcf, 0x9e, 0xd8, 0x00, 0xb6, 0x2d, 0x7b, 0x93, 0x2c, 0x5d, 0x8b, 0x9f, 0xfe, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x65, 0x76, 0xd7, 0xb9, 0xed, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
