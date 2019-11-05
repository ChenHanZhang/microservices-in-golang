// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vessel/vessel.proto

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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd0828b40ed5ef5, []int{0}
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

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

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
	return fileDescriptor_2bd0828b40ed5ef5, []int{1}
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

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd0828b40ed5ef5, []int{2}
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

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

func init() { proto.RegisterFile("vessel/vessel.proto", fileDescriptor_2bd0828b40ed5ef5) }

var fileDescriptor_2bd0828b40ed5ef5 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6a, 0xc2, 0x40,
	0x10, 0xed, 0x46, 0x8d, 0x71, 0x8a, 0x3d, 0x8c, 0x97, 0xad, 0xad, 0x10, 0x72, 0xf2, 0x94, 0x82,
	0xd2, 0x0f, 0x28, 0x05, 0xa1, 0x3d, 0xc6, 0xd2, 0x1e, 0x65, 0x4d, 0xa6, 0x76, 0xc0, 0x64, 0x43,
	0x36, 0x44, 0xfb, 0x37, 0xfd, 0x81, 0xfe, 0x63, 0x61, 0xb7, 0x5a, 0x2c, 0xa2, 0xa7, 0x9d, 0x79,
	0xf3, 0xe6, 0xf1, 0xe6, 0xb1, 0x30, 0x68, 0xc8, 0x18, 0x5a, 0xdf, 0xb9, 0x27, 0x2e, 0x2b, 0x5d,
	0x6b, 0x1c, 0xac, 0x74, 0x9c, 0x73, 0x5a, 0xe9, 0xd8, 0x54, 0x4d, 0xec, 0x46, 0xd1, 0x97, 0x00,
	0xff, 0xd5, 0x96, 0x78, 0x05, 0x1e, 0x67, 0x52, 0x84, 0x62, 0xdc, 0x4b, 0x3c, 0xce, 0x70, 0x08,
	0x41, 0xaa, 0x4a, 0x95, 0x72, 0xfd, 0x29, 0xbd, 0x50, 0x8c, 0x3b, 0xc9, 0xbe, 0xc7, 0x11, 0x40,
	0xae, 0xb6, 0x8b, 0x0d, 0xf1, 0xea, 0xa3, 0x96, 0x2d, 0x3b, 0xed, 0xe5, 0x6a, 0xfb, 0x66, 0x01,
	0x44, 0x68, 0x17, 0x2a, 0x27, 0xd9, 0xb6, 0x62, 0xb6, 0xc6, 0x5b, 0xe8, 0xa9, 0x46, 0xf1, 0x5a,
	0x2d, 0xd7, 0x24, 0x3b, 0xa1, 0x18, 0x07, 0xc9, 0x1f, 0x80, 0xd7, 0x10, 0xe8, 0x4d, 0x41, 0xd5,
	0x82, 0x33, 0xe9, 0xdb, 0xad, 0xae, 0xed, 0x9f, 0xb2, 0xe8, 0x19, 0xfa, 0xf3, 0x92, 0x52, 0x7e,
	0xe7, 0x54, 0xd5, 0xac, 0x8b, 0x03, 0x63, 0xe2, 0xa4, 0x31, 0xef, 0x9f, 0xb1, 0xa8, 0x81, 0x20,
	0x21, 0x53, 0xea, 0xc2, 0x10, 0x4e, 0xc1, 0x77, 0x21, 0x58, 0x91, 0xcb, 0xc9, 0x4d, 0x7c, 0x24,
	0xa0, 0xd8, 0x85, 0x93, 0xfc, 0x52, 0xf1, 0x1e, 0xba, 0xae, 0x32, 0xd2, 0x0b, 0x5b, 0xe7, 0xb6,
	0x76, 0xdc, 0xc9, 0xb7, 0x80, 0xbe, 0xc3, 0xe6, 0x54, 0x35, 0x9c, 0x12, 0xbe, 0x40, 0x7f, 0xc6,
	0x45, 0xf6, 0xb0, 0x4f, 0x20, 0x3a, 0x2a, 0x74, 0x70, 0xf9, 0x70, 0x74, 0x94, 0xb3, 0xbb, 0x28,
	0xba, 0xc0, 0x19, 0xf8, 0x8f, 0x15, 0xa9, 0x9a, 0xf0, 0x94, 0xaf, 0xb3, 0x3a, 0x4b, 0xdf, 0x7e,
	0x99, 0xe9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x8d, 0x46, 0x9b, 0x49, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
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
	FindAvailable(context.Context, *Specification, *Response) error
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
