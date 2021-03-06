// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proccessor_message.proto

package pb

import (
	fmt "fmt"
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

type CPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NumberCores          uint32   `protobuf:"varint,3,opt,name=number_cores,json=numberCores,proto3" json:"number_cores,omitempty"`
	NumberThreads        uint32   `protobuf:"varint,4,opt,name=number_threads,json=numberThreads,proto3" json:"number_threads,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,5,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,6,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd5b3a1be58f3bef, []int{0}
}

func (m *CPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPU.Unmarshal(m, b)
}
func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
}
func (m *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(m, src)
}
func (m *CPU) XXX_Size() int {
	return xxx_messageInfo_CPU.Size(m)
}
func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CPU) GetNumberCores() uint32 {
	if m != nil {
		return m.NumberCores
	}
	return 0
}

func (m *CPU) GetNumberThreads() uint32 {
	if m != nil {
		return m.NumberThreads
	}
	return 0
}

func (m *CPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *CPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

type GPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,3,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,4,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	Memory               *Memory  `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GPU) Reset()         { *m = GPU{} }
func (m *GPU) String() string { return proto.CompactTextString(m) }
func (*GPU) ProtoMessage()    {}
func (*GPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd5b3a1be58f3bef, []int{1}
}

func (m *GPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GPU.Unmarshal(m, b)
}
func (m *GPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GPU.Marshal(b, m, deterministic)
}
func (m *GPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPU.Merge(m, src)
}
func (m *GPU) XXX_Size() int {
	return xxx_messageInfo_GPU.Size(m)
}
func (m *GPU) XXX_DiscardUnknown() {
	xxx_messageInfo_GPU.DiscardUnknown(m)
}

var xxx_messageInfo_GPU proto.InternalMessageInfo

func (m *GPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *GPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *GPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

func (m *GPU) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterType((*CPU)(nil), "proto.CPU")
	proto.RegisterType((*GPU)(nil), "proto.GPU")
}

func init() {
	proto.RegisterFile("proccessor_message.proto", fileDescriptor_cd5b3a1be58f3bef)
}

var fileDescriptor_cd5b3a1be58f3bef = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x4f,
	0x4e, 0x4e, 0x2d, 0x2e, 0xce, 0x2f, 0x8a, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x22, 0xb9, 0xa9, 0xb9, 0xf9, 0x45, 0x95,
	0xa8, 0x92, 0x4a, 0xab, 0x19, 0xb9, 0x98, 0x9d, 0x03, 0x42, 0x85, 0x44, 0xb8, 0x58, 0x93, 0x8a,
	0x12, 0xf3, 0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0xb0, 0x20, 0x98, 0x2d, 0xa4, 0xc8, 0xc5, 0x93, 0x57, 0x9a,
	0x9b, 0x94, 0x5a, 0x14, 0x9f, 0x9c, 0x5f, 0x94, 0x5a, 0x2c, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1b,
	0xc4, 0x0d, 0x11, 0x73, 0x06, 0x09, 0x09, 0xa9, 0x72, 0xf1, 0x41, 0x95, 0x94, 0x64, 0x14, 0xa5,
	0x26, 0xa6, 0x14, 0x4b, 0xb0, 0x80, 0x15, 0xf1, 0x42, 0x44, 0x43, 0x20, 0x82, 0x42, 0xe2, 0x5c,
	0xec, 0xb9, 0x99, 0x79, 0xf1, 0xe9, 0x19, 0x55, 0x12, 0xac, 0x0a, 0x8c, 0x1a, 0x8c, 0x41, 0x6c,
	0xb9, 0x99, 0x79, 0xee, 0x19, 0x55, 0x60, 0x89, 0xc4, 0x0a, 0xb0, 0x04, 0x1b, 0x54, 0x22, 0xb1,
	0xc2, 0x3d, 0xa3, 0x4a, 0xa9, 0x83, 0x91, 0x8b, 0xd9, 0x9d, 0x24, 0xd7, 0x22, 0xd9, 0xc1, 0x8c,
	0xcb, 0x0e, 0x16, 0x64, 0x3b, 0x84, 0x54, 0xb9, 0xd8, 0x20, 0x21, 0x05, 0x76, 0x14, 0xb7, 0x11,
	0x2f, 0x24, 0xa4, 0xf4, 0x7c, 0xc1, 0x82, 0x41, 0x50, 0x49, 0x27, 0x96, 0x28, 0xa6, 0x82, 0xa4,
	0x24, 0x36, 0xb0, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x6d, 0x79, 0xbd, 0x7e, 0x01,
	0x00, 0x00,
}
