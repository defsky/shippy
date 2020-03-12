// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package go_micro_srv_vessel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

// 每条货轮的信息
type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwerId               string   `protobuf:"bytes,6,opt,name=ower_id,json=owerId,proto3" json:"ower_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwerId() string {
	if m != nil {
		return m.OwerId
	}
	return ""
}

// 等待运送的货物
type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

// 货轮装得下的话
// 返回的多条货轮信息
type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=Created,proto3" json:"Created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716) }

var fileDescriptor_04ef66862bb50716 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xfa, 0x40,
	0x10, 0xc7, 0x7f, 0x5b, 0xa0, 0x94, 0xf9, 0x05, 0x0f, 0xe3, 0xc1, 0x15, 0x25, 0x69, 0x7a, 0xe2,
	0x54, 0x13, 0x88, 0x0f, 0x60, 0x4c, 0x48, 0xf4, 0x58, 0x8c, 0x1e, 0xc9, 0xd2, 0x8e, 0x38, 0x09,
	0xed, 0x36, 0x6d, 0x53, 0xf0, 0x2d, 0x7c, 0x02, 0xdf, 0xc0, 0x77, 0x34, 0xee, 0x02, 0x06, 0xd3,
	0xc0, 0xa9, 0xf3, 0xe7, 0x3b, 0x33, 0xdf, 0x7e, 0xb2, 0x70, 0x99, 0x17, 0xba, 0xd2, 0x37, 0x35,
	0x95, 0x25, 0xad, 0xb6, 0x9f, 0xd0, 0xd4, 0xf0, 0x7c, 0xa9, 0xc3, 0x94, 0xe3, 0x42, 0x87, 0x65,
	0x51, 0x87, 0xb6, 0x15, 0x7c, 0x0a, 0x70, 0x9f, 0x4d, 0x88, 0x67, 0xe0, 0x70, 0x22, 0x85, 0x2f,
	0x46, 0xbd, 0xc8, 0xe1, 0x04, 0x07, 0xe0, 0xc5, 0x2a, 0x57, 0x31, 0x57, 0xef, 0xd2, 0xf1, 0xc5,
	0xa8, 0x13, 0xed, 0x73, 0x1c, 0x02, 0xa4, 0x6a, 0x33, 0x5f, 0x13, 0x2f, 0xdf, 0x2a, 0xd9, 0x32,
	0xdd, 0x5e, 0xaa, 0x36, 0x2f, 0xa6, 0x80, 0x08, 0xed, 0x4c, 0xa5, 0x24, 0xdb, 0x66, 0x99, 0x89,
	0xf1, 0x1a, 0x7a, 0xaa, 0x56, 0xbc, 0x52, 0x8b, 0x15, 0xc9, 0x8e, 0x2f, 0x46, 0x5e, 0xf4, 0x5b,
	0xc0, 0x0b, 0xe8, 0xea, 0x35, 0x15, 0x73, 0x4e, 0xa4, 0x6b, 0x86, 0xdc, 0x9f, 0xf4, 0x21, 0x09,
	0x1e, 0xa1, 0x3f, 0xcb, 0x29, 0xe6, 0x57, 0x8e, 0x55, 0xc5, 0x3a, 0x3b, 0xb0, 0x25, 0x8e, 0xda,
	0x72, 0xfe, 0xd8, 0x0a, 0x3e, 0x04, 0x78, 0x11, 0x95, 0xb9, 0xce, 0x4a, 0xc2, 0x09, 0xb8, 0x96,
	0x81, 0xd9, 0xf2, 0x7f, 0x7c, 0x15, 0x36, 0xf0, 0x09, 0x2d, 0x9b, 0x68, 0x2b, 0xc5, 0x5b, 0xe8,
	0xda, 0xa8, 0x94, 0x8e, 0xdf, 0x3a, 0x35, 0xb5, 0xd3, 0xa2, 0x84, 0xee, 0x7d, 0x41, 0xaa, 0xa2,
	0xc4, 0xb0, 0xf2, 0xa2, 0x5d, 0x3a, 0xfe, 0x12, 0xd0, 0xb7, 0xea, 0x19, 0x15, 0x35, 0xc7, 0x84,
	0x4f, 0xd0, 0x9f, 0x72, 0x96, 0xdc, 0xed, 0xd1, 0x04, 0x8d, 0x27, 0x0e, 0xa0, 0x0c, 0x86, 0x8d,
	0x9a, 0xdd, 0xbf, 0x06, 0xff, 0x70, 0x0a, 0xae, 0x3d, 0x89, 0xc7, 0x1c, 0x9f, 0xdc, 0xb3, 0x70,
	0xcd, 0x5b, 0x9a, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x70, 0xc9, 0x72, 0xab, 0x68, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	// 检查是否有能运送货物的轮船
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	// 创建货轮
	Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	// 检查是否有能运送货物的轮船
	FindAvailable(context.Context, *Specification, *Response) error
	// 创建货轮
	Create(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *VesselService) Create(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.Create(ctx, in, out)
}
