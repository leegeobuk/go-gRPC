// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop.proto

package pc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Laptop struct {
	ID       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages,proto3" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight               isLaptop_Weight      `protobuf_oneof:"weight"`
	PriceUsd             float64              `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear          uint32               `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}
func (*Laptop) Descriptor() ([]byte, []int) {
	return fileDescriptor_28a7e4886f546705, []int{0}
}

func (m *Laptop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laptop.Unmarshal(m, b)
}
func (m *Laptop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laptop.Marshal(b, m, deterministic)
}
func (m *Laptop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laptop.Merge(m, src)
}
func (m *Laptop) XXX_Size() int {
	return xxx_messageInfo_Laptop.Size(m)
}
func (m *Laptop) XXX_DiscardUnknown() {
	xxx_messageInfo_Laptop.DiscardUnknown(m)
}

var xxx_messageInfo_Laptop proto.InternalMessageInfo

func (m *Laptop) GetId() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetRam() *Memory {
	if m != nil {
		return m.Ram
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLb) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLb() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Laptop) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
}

type CreateLaptopRequest struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopRequest) Reset()         { *m = CreateLaptopRequest{} }
func (m *CreateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopRequest) ProtoMessage()    {}
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_28a7e4886f546705, []int{1}
}

func (m *CreateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopRequest.Unmarshal(m, b)
}
func (m *CreateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *CreateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopRequest.Merge(m, src)
}
func (m *CreateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopRequest.Size(m)
}
func (m *CreateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopRequest proto.InternalMessageInfo

func (m *CreateLaptopRequest) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopResponse) Reset()         { *m = CreateLaptopResponse{} }
func (m *CreateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopResponse) ProtoMessage()    {}
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28a7e4886f546705, []int{2}
}

func (m *CreateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopResponse.Unmarshal(m, b)
}
func (m *CreateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *CreateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopResponse.Merge(m, src)
}
func (m *CreateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopResponse.Size(m)
}
func (m *CreateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopResponse proto.InternalMessageInfo

func (m *CreateLaptopResponse) GetId() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*Laptop)(nil), "pc.Laptop")
	proto.RegisterType((*CreateLaptopRequest)(nil), "pc.CreateLaptopRequest")
	proto.RegisterType((*CreateLaptopResponse)(nil), "pc.CreateLaptopResponse")
}

func init() {
	proto.RegisterFile("laptop.proto", fileDescriptor_28a7e4886f546705)
}

var fileDescriptor_28a7e4886f546705 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0xd3, 0x76, 0xe9, 0x6b, 0x5a, 0x24, 0x33, 0x09, 0xd3, 0x81, 0x08, 0x3d, 0x40,
	0x4e, 0xa9, 0x34, 0x4e, 0x3b, 0x6e, 0x3b, 0x80, 0xb4, 0x21, 0x4d, 0x2e, 0x3d, 0xc0, 0xa5, 0x72,
	0x92, 0x47, 0x88, 0xd6, 0xd4, 0xc6, 0x76, 0x40, 0xfd, 0x37, 0xf9, 0x8b, 0x50, 0x6c, 0x67, 0xec,
	0xc7, 0xcd, 0xef, 0xf3, 0xfd, 0x5a, 0xef, 0x27, 0xc4, 0x3b, 0x2e, 0x8d, 0x90, 0x99, 0x54, 0xc2,
	0x08, 0x32, 0x90, 0xc5, 0xe2, 0xb9, 0x54, 0xa2, 0x40, 0xad, 0x85, 0x72, 0x70, 0x11, 0x37, 0xd8,
	0x08, 0x75, 0xf0, 0xd1, 0x4c, 0x1b, 0xa1, 0x78, 0x85, 0xbd, 0xa8, 0x0b, 0x85, 0xb8, 0xf7, 0xd1,
	0xfc, 0x16, 0x0f, 0xb9, 0xe0, 0xaa, 0xf4, 0xf1, 0xdb, 0x4a, 0x88, 0x6a, 0x87, 0x2b, 0x1b, 0xe5,
	0xed, 0x8f, 0x95, 0xa9, 0x1b, 0xd4, 0x86, 0x37, 0x3e, 0xe1, 0xf2, 0x6f, 0x08, 0xe3, 0x6b, 0x5b,
	0x01, 0x99, 0xc3, 0xa0, 0x2e, 0x69, 0x90, 0x04, 0xe9, 0x84, 0x0d, 0xea, 0x92, 0x1c, 0xc3, 0x28,
	0x57, 0x7c, 0x5f, 0xd2, 0x81, 0x45, 0x2e, 0x20, 0x04, 0x86, 0x7b, 0xde, 0x20, 0x0d, 0x2d, 0xb4,
	0x6f, 0xf2, 0x0a, 0xc2, 0x42, 0xb6, 0x74, 0x98, 0x04, 0xe9, 0xf4, 0xf4, 0x28, 0x93, 0x45, 0x76,
	0x79, 0xb3, 0x61, 0x1d, 0x23, 0xaf, 0x21, 0x54, 0xbc, 0xa1, 0x23, 0x2b, 0x41, 0x27, 0x7d, 0xb1,
	0xcd, 0xb0, 0x0e, 0x93, 0x13, 0x18, 0x56, 0xb2, 0xd5, 0x74, 0x9c, 0x84, 0xfd, 0xcf, 0x4f, 0x37,
	0x1b, 0x66, 0x21, 0xf9, 0x00, 0x91, 0x6f, 0x55, 0xd3, 0x23, 0x6b, 0x98, 0x76, 0x86, 0xb5, 0x63,
	0xec, 0x4e, 0x24, 0x4b, 0x18, 0xbb, 0x21, 0xd0, 0xe8, 0x7f, 0x9a, 0xb5, 0x25, 0xcc, 0x2b, 0x24,
	0x85, 0xa8, 0x1f, 0x0d, 0x9d, 0x58, 0x57, 0xdc, 0xb9, 0xae, 0x3c, 0x63, 0x77, 0x2a, 0x79, 0x03,
	0x93, 0x3f, 0x58, 0x57, 0x3f, 0xcd, 0xf6, 0xb6, 0xa2, 0x90, 0x04, 0x69, 0xf0, 0xf9, 0x19, 0x8b,
	0x1c, 0xba, 0xaa, 0xee, 0xc9, 0xbb, 0x9c, 0x4e, 0x1f, 0xca, 0xd7, 0x39, 0x39, 0x81, 0x89, 0x54,
	0x75, 0x81, 0xdb, 0x56, 0x97, 0x34, 0xee, 0x64, 0x16, 0x59, 0xb0, 0xd1, 0x25, 0x79, 0x07, 0xb1,
	0xc2, 0x1d, 0x72, 0x8d, 0xdb, 0x03, 0x72, 0x45, 0x67, 0x49, 0x90, 0xce, 0xd8, 0xd4, 0xb3, 0x6f,
	0xc8, 0x15, 0x39, 0x03, 0x68, 0x65, 0xc9, 0x0d, 0x96, 0x5b, 0x6e, 0xe8, 0xdc, 0x56, 0xba, 0xc8,
	0xdc, 0x16, 0xb3, 0x7e, 0x8b, 0xd9, 0xd7, 0x7e, 0x8b, 0x6c, 0xe2, 0xdd, 0xe7, 0xe6, 0x22, 0x82,
	0xb1, 0x2b, 0x63, 0x79, 0x06, 0x2f, 0x2e, 0x15, 0x72, 0x83, 0x6e, 0xb3, 0x0c, 0x7f, 0xb5, 0xa8,
	0x4d, 0x37, 0x27, 0x77, 0x6c, 0x76, 0xc9, 0x7e, 0x4e, 0xde, 0xe2, 0x95, 0xe5, 0x7b, 0x38, 0x7e,
	0xf8, 0x55, 0x4b, 0xb1, 0xd7, 0xf8, 0xf8, 0x38, 0x4e, 0x19, 0xcc, 0x9c, 0x63, 0x8d, 0xea, 0x77,
	0x5d, 0x20, 0x39, 0x87, 0xf8, 0xfe, 0x47, 0xf2, 0xd2, 0x9e, 0xc1, 0xd3, 0x2a, 0x16, 0xf4, 0xa9,
	0xe0, 0x72, 0x5c, 0x8c, 0xbe, 0x87, 0x2b, 0x59, 0xe4, 0x63, 0xdb, 0xe6, 0xc7, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb0, 0x25, 0xab, 0x41, 0x1a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
}

type laptopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaptopServiceClient(cc grpc.ClientConnInterface) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/pc.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LaptopServiceServer is the server API for LaptopService service.
type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
}

// UnimplementedLaptopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLaptopServiceServer struct {
}

func (*UnimplementedLaptopServiceServer) CreateLaptop(ctx context.Context, req *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}

func RegisterLaptopServiceServer(s *grpc.Server, srv LaptopServiceServer) {
	s.RegisterService(&_LaptopService_serviceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pc.LaptopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LaptopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pc.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "laptop.proto",
}
