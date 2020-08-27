// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agentpb/files.proto

package agentpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	inventorypb "github.com/percona/pmm/api/inventorypb"
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

// DownloadFileChunkRequest is an ServerMessage requesting file chunk download.
type DownloadFileChunkRequest struct {
	FileType             inventorypb.FileType `protobuf:"varint,1,opt,name=file_type,json=fileType,proto3,enum=inventory.FileType" json:"file_type,omitempty"`
	FileId               string               `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Offset               int64                `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	ChunkSize            int64                `protobuf:"varint,4,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DownloadFileChunkRequest) Reset()         { *m = DownloadFileChunkRequest{} }
func (m *DownloadFileChunkRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadFileChunkRequest) ProtoMessage()    {}
func (*DownloadFileChunkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5870d2f7ef940ee3, []int{0}
}

func (m *DownloadFileChunkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadFileChunkRequest.Unmarshal(m, b)
}
func (m *DownloadFileChunkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadFileChunkRequest.Marshal(b, m, deterministic)
}
func (m *DownloadFileChunkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadFileChunkRequest.Merge(m, src)
}
func (m *DownloadFileChunkRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadFileChunkRequest.Size(m)
}
func (m *DownloadFileChunkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadFileChunkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadFileChunkRequest proto.InternalMessageInfo

func (m *DownloadFileChunkRequest) GetFileType() inventorypb.FileType {
	if m != nil {
		return m.FileType
	}
	return inventorypb.FileType_FILE_TYPE_INVALID
}

func (m *DownloadFileChunkRequest) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

func (m *DownloadFileChunkRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DownloadFileChunkRequest) GetChunkSize() int64 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

// DownloadFileChunkResponse is an AgentMessage containing requested file chunk.
type DownloadFileChunkResponse struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	FileSize             int64    `protobuf:"varint,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadFileChunkResponse) Reset()         { *m = DownloadFileChunkResponse{} }
func (m *DownloadFileChunkResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadFileChunkResponse) ProtoMessage()    {}
func (*DownloadFileChunkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5870d2f7ef940ee3, []int{1}
}

func (m *DownloadFileChunkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadFileChunkResponse.Unmarshal(m, b)
}
func (m *DownloadFileChunkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadFileChunkResponse.Marshal(b, m, deterministic)
}
func (m *DownloadFileChunkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadFileChunkResponse.Merge(m, src)
}
func (m *DownloadFileChunkResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadFileChunkResponse.Size(m)
}
func (m *DownloadFileChunkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadFileChunkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadFileChunkResponse proto.InternalMessageInfo

func (m *DownloadFileChunkResponse) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *DownloadFileChunkResponse) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

// DeleteFileRequest is an ServerMessage requesting file deletion.
type DeleteFileRequest struct {
	FileType             inventorypb.FileType `protobuf:"varint,1,opt,name=file_type,json=fileType,proto3,enum=inventory.FileType" json:"file_type,omitempty"`
	FileId               string               `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeleteFileRequest) Reset()         { *m = DeleteFileRequest{} }
func (m *DeleteFileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFileRequest) ProtoMessage()    {}
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5870d2f7ef940ee3, []int{2}
}

func (m *DeleteFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFileRequest.Unmarshal(m, b)
}
func (m *DeleteFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFileRequest.Merge(m, src)
}
func (m *DeleteFileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFileRequest.Size(m)
}
func (m *DeleteFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFileRequest proto.InternalMessageInfo

func (m *DeleteFileRequest) GetFileType() inventorypb.FileType {
	if m != nil {
		return m.FileType
	}
	return inventorypb.FileType_FILE_TYPE_INVALID
}

func (m *DeleteFileRequest) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

// DeleteFileResponse is an AgentMessage confirming file deletion.
type DeleteFileResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFileResponse) Reset()         { *m = DeleteFileResponse{} }
func (m *DeleteFileResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteFileResponse) ProtoMessage()    {}
func (*DeleteFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5870d2f7ef940ee3, []int{3}
}

func (m *DeleteFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFileResponse.Unmarshal(m, b)
}
func (m *DeleteFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFileResponse.Marshal(b, m, deterministic)
}
func (m *DeleteFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFileResponse.Merge(m, src)
}
func (m *DeleteFileResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteFileResponse.Size(m)
}
func (m *DeleteFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFileResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DownloadFileChunkRequest)(nil), "agent.DownloadFileChunkRequest")
	proto.RegisterType((*DownloadFileChunkResponse)(nil), "agent.DownloadFileChunkResponse")
	proto.RegisterType((*DeleteFileRequest)(nil), "agent.DeleteFileRequest")
	proto.RegisterType((*DeleteFileResponse)(nil), "agent.DeleteFileResponse")
}

func init() {
	proto.RegisterFile("agentpb/files.proto", fileDescriptor_5870d2f7ef940ee3)
}

var fileDescriptor_5870d2f7ef940ee3 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x33, 0xf0, 0x51, 0xbe, 0x8e, 0xc6, 0x3f, 0x83, 0x89, 0x95, 0x0d, 0x84, 0xc4, 0xc8,
	0x42, 0x5a, 0xa3, 0x89, 0x0f, 0x80, 0xc4, 0xc4, 0x6d, 0x75, 0x61, 0xdc, 0x90, 0x96, 0xde, 0x96,
	0x09, 0x65, 0x66, 0xec, 0x4c, 0x69, 0xca, 0x03, 0xfa, 0x1a, 0x24, 0xbc, 0x88, 0xa6, 0xd3, 0x8a,
	0x35, 0x61, 0x6b, 0x57, 0x37, 0xf7, 0xfe, 0xee, 0x39, 0x3d, 0x33, 0x83, 0x3b, 0x5e, 0x04, 0x4c,
	0x09, 0xdf, 0x09, 0x69, 0x0c, 0xd2, 0x16, 0x09, 0x57, 0x9c, 0xb4, 0x74, 0xb3, 0x7b, 0x1f, 0x51,
	0x35, 0x4f, 0x7d, 0x7b, 0xc6, 0x97, 0xce, 0x32, 0xa3, 0x6a, 0xc1, 0x33, 0x27, 0xe2, 0x23, 0xcd,
	0x8c, 0x56, 0x5e, 0x4c, 0x03, 0x4f, 0xf1, 0x44, 0x3a, 0xbb, 0xb2, 0x5c, 0xef, 0x9e, 0x53, 0xb6,
	0x02, 0xa6, 0x78, 0x92, 0xff, 0xd6, 0x1d, 0x7c, 0x20, 0x6c, 0x4d, 0x78, 0xc6, 0x62, 0xee, 0x05,
	0x8f, 0x34, 0x86, 0x87, 0x79, 0xca, 0x16, 0x2e, 0xbc, 0xa7, 0x20, 0x15, 0xb9, 0xc1, 0x66, 0xc1,
	0x4e, 0x55, 0x2e, 0xc0, 0x42, 0x7d, 0x34, 0x3c, 0xba, 0xed, 0xd8, 0x3b, 0x25, 0xbb, 0xe0, 0x5f,
	0x72, 0x01, 0xee, 0xff, 0xb0, 0xaa, 0x48, 0x0f, 0xb7, 0xf5, 0x06, 0x0d, 0xac, 0x46, 0x1f, 0x0d,
	0xcd, 0xb1, 0xb1, 0xdd, 0xf4, 0x1a, 0xaf, 0xc8, 0x35, 0x8a, 0xf6, 0x53, 0x40, 0xae, 0xb0, 0xc1,
	0xc3, 0x50, 0x82, 0xb2, 0x9a, 0x7d, 0x34, 0x6c, 0x8e, 0x8f, 0xb7, 0x9b, 0xde, 0xc1, 0xc9, 0xe7,
	0xf7, 0x87, 0xdc, 0x6a, 0x4c, 0x6c, 0x8c, 0x67, 0xc5, 0xbf, 0x4c, 0x25, 0x5d, 0x83, 0xf5, 0x6f,
	0x3f, 0x6c, 0x6a, 0xe4, 0x99, 0xae, 0x61, 0x20, 0xf0, 0xc5, 0x9e, 0x1c, 0x52, 0x70, 0x26, 0x81,
	0x5c, 0xe2, 0x96, 0x26, 0x75, 0x88, 0xc3, 0x52, 0x47, 0xd4, 0x74, 0xca, 0x29, 0xb9, 0xae, 0xf2,
	0x6a, 0xcb, 0xc6, 0x8f, 0x65, 0x1d, 0xd5, 0x59, 0xb5, 0x63, 0x88, 0x4f, 0x27, 0x10, 0x83, 0x82,
	0xc2, 0xef, 0xef, 0x8e, 0x6c, 0x70, 0x86, 0x49, 0xdd, 0xa7, 0x8c, 0x34, 0x36, 0xdf, 0xda, 0xd5,
	0x3b, 0xf1, 0x0d, 0x7d, 0x95, 0x77, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xb4, 0x86, 0x6e,
	0x39, 0x02, 0x00, 0x00,
}