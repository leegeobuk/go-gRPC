// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screen.proto

package pc

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

type Screen_Panel int32

const (
	Screen_UNKNOWN Screen_Panel = 0
	Screen_IPS     Screen_Panel = 1
	Screen_OLED    Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKNOWN",
	1: "IPS",
	2: "OLED",
}

var Screen_Panel_value = map[string]int32{
	"UNKNOWN": 0,
	"IPS":     1,
	"OLED":    2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}

func (Screen_Panel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b79f3b69cf8c7563, []int{0, 0}
}

type Screen struct {
	SizeInch             float32            `protobuf:"fixed32,1,opt,name=size_inch,json=sizeInch,proto3" json:"size_inch,omitempty"`
	Resolution           *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Panel                Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=pc.Screen_Panel" json:"panel,omitempty"`
	Multitouch           bool               `protobuf:"varint,4,opt,name=multitouch,proto3" json:"multitouch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Screen) Reset()         { *m = Screen{} }
func (m *Screen) String() string { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()    {}
func (*Screen) Descriptor() ([]byte, []int) {
	return fileDescriptor_b79f3b69cf8c7563, []int{0}
}

func (m *Screen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen.Unmarshal(m, b)
}
func (m *Screen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen.Marshal(b, m, deterministic)
}
func (m *Screen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen.Merge(m, src)
}
func (m *Screen) XXX_Size() int {
	return xxx_messageInfo_Screen.Size(m)
}
func (m *Screen) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen.DiscardUnknown(m)
}

var xxx_messageInfo_Screen proto.InternalMessageInfo

func (m *Screen) GetSizeInch() float32 {
	if m != nil {
		return m.SizeInch
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKNOWN
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screen_Resolution) Reset()         { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()    {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_b79f3b69cf8c7563, []int{0, 0}
}

func (m *Screen_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen_Resolution.Unmarshal(m, b)
}
func (m *Screen_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen_Resolution.Marshal(b, m, deterministic)
}
func (m *Screen_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen_Resolution.Merge(m, src)
}
func (m *Screen_Resolution) XXX_Size() int {
	return xxx_messageInfo_Screen_Resolution.Size(m)
}
func (m *Screen_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Screen_Resolution proto.InternalMessageInfo

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("pc.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
	proto.RegisterType((*Screen)(nil), "pc.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "pc.Screen.Resolution")
}

func init() {
	proto.RegisterFile("screen.proto", fileDescriptor_b79f3b69cf8c7563)
}

var fileDescriptor_b79f3b69cf8c7563 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x4d, 0xba, 0x76, 0xf5, 0x9b, 0x93, 0xf2, 0xa1, 0x52, 0x14, 0xa4, 0xec, 0xa0, 0x3d,
	0x55, 0x98, 0x78, 0xf1, 0x28, 0x7a, 0x18, 0x4a, 0x37, 0x32, 0x44, 0xf0, 0x22, 0x1a, 0x83, 0x09,
	0xd4, 0x24, 0xb4, 0x29, 0x82, 0xff, 0x81, 0xff, 0xb5, 0x2c, 0x11, 0xd7, 0xe3, 0xfb, 0xbd, 0xef,
	0xf1, 0x1e, 0x1f, 0xec, 0x75, 0xbc, 0x15, 0x42, 0x57, 0xb6, 0x35, 0xce, 0x20, 0xb5, 0x7c, 0xf6,
	0x43, 0x21, 0x59, 0x7b, 0x88, 0x27, 0xb0, 0xdb, 0xa9, 0x6f, 0xf1, 0xa2, 0x34, 0x97, 0x39, 0x29,
	0x48, 0x49, 0x59, 0xba, 0x01, 0x0b, 0xcd, 0x25, 0x5e, 0x01, 0xb4, 0xa2, 0x33, 0x4d, 0xef, 0x94,
	0xd1, 0x39, 0x2d, 0x48, 0x39, 0x99, 0x1f, 0x56, 0x96, 0x57, 0x21, 0x5c, 0xb1, 0x7f, 0x93, 0x0d,
	0x0e, 0xf1, 0x0c, 0x62, 0xfb, 0xaa, 0x45, 0x93, 0x47, 0x05, 0x29, 0xf7, 0xe7, 0xd9, 0x20, 0xb1,
	0xda, 0x70, 0x16, 0x6c, 0x3c, 0x05, 0xf8, 0xec, 0x1b, 0xa7, 0x9c, 0xe9, 0xb9, 0xcc, 0x47, 0x05,
	0x29, 0x53, 0x36, 0x20, 0xc7, 0xd7, 0x00, 0xdb, 0x06, 0x3c, 0x80, 0xf8, 0x4b, 0xbd, 0xbb, 0xb0,
	0x72, 0xca, 0x82, 0xc0, 0x23, 0x48, 0xa4, 0x50, 0x1f, 0xd2, 0xf9, 0x79, 0x53, 0xf6, 0xa7, 0x66,
	0xe7, 0x10, 0xfb, 0x2e, 0x9c, 0xc0, 0xf8, 0xb1, 0xbe, 0xaf, 0x97, 0x4f, 0x75, 0xb6, 0x83, 0x63,
	0x88, 0x16, 0xab, 0x75, 0x46, 0x30, 0x85, 0xd1, 0xf2, 0xe1, 0xee, 0x36, 0xa3, 0x37, 0xf1, 0x73,
	0x74, 0x61, 0xf9, 0x5b, 0xe2, 0xbf, 0x73, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xdc, 0x58,
	0x71, 0x2d, 0x01, 0x00, 0x00,
}
