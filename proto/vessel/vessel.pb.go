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
	Created              bool      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
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

func init() { proto.RegisterFile("vessel/vessel.proto", fileDescriptor_2bd0828b40ed5ef5) }

var fileDescriptor_2bd0828b40ed5ef5 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6a, 0xfa, 0x40,
	0x10, 0xc7, 0x7f, 0x1b, 0x35, 0xc6, 0xf9, 0x61, 0x0f, 0xe3, 0x65, 0x6b, 0x2b, 0x84, 0x9c, 0x3c,
	0xa5, 0xa0, 0xf4, 0x01, 0x4a, 0x41, 0x68, 0x8f, 0xb1, 0xb4, 0x47, 0x59, 0x37, 0x53, 0x3b, 0x60,
	0xb2, 0x21, 0x09, 0xd1, 0xbe, 0x45, 0x1f, 0xa1, 0x2f, 0xd0, 0x77, 0x2c, 0xec, 0x56, 0x8b, 0x45,
	0xf4, 0x94, 0xf9, 0xf3, 0x9d, 0x99, 0x6f, 0x3e, 0x2c, 0x0c, 0x1a, 0xaa, 0x2a, 0x5a, 0xdf, 0xb8,
	0x4f, 0x5c, 0x94, 0xa6, 0x36, 0x38, 0x58, 0x99, 0x38, 0x63, 0x5d, 0x9a, 0xb8, 0x2a, 0x9b, 0xd8,
	0xb5, 0xa2, 0x4f, 0x01, 0xfe, 0xb3, 0x0d, 0xf1, 0x02, 0x3c, 0x4e, 0xa5, 0x08, 0xc5, 0xb8, 0x97,
	0x78, 0x9c, 0xe2, 0x10, 0x02, 0xad, 0x0a, 0xa5, 0xb9, 0x7e, 0x97, 0x5e, 0x28, 0xc6, 0x9d, 0x64,
	0x9f, 0xe3, 0x08, 0x20, 0x53, 0xdb, 0xc5, 0x86, 0x78, 0xf5, 0x56, 0xcb, 0x96, 0xed, 0xf6, 0x32,
	0xb5, 0x7d, 0xb1, 0x05, 0x44, 0x68, 0xe7, 0x2a, 0x23, 0xd9, 0xb6, 0xcb, 0x6c, 0x8c, 0xd7, 0xd0,
	0x53, 0x8d, 0xe2, 0xb5, 0x5a, 0xae, 0x49, 0x76, 0x42, 0x31, 0x0e, 0x92, 0xdf, 0x02, 0x5e, 0x42,
	0x60, 0x36, 0x39, 0x95, 0x0b, 0x4e, 0xa5, 0x6f, 0xa7, 0xba, 0x36, 0x7f, 0x48, 0xa3, 0x47, 0xe8,
	0xcf, 0x0b, 0xd2, 0xfc, 0xca, 0x5a, 0xd5, 0x6c, 0xf2, 0x03, 0x63, 0xe2, 0xa4, 0x31, 0xef, 0x8f,
	0xb1, 0xe8, 0x43, 0x40, 0x90, 0x50, 0x55, 0x98, 0xbc, 0x22, 0x9c, 0x82, 0xef, 0x28, 0xd8, 0x2d,
	0xff, 0x27, 0x57, 0xf1, 0x11, 0x42, 0xb1, 0xa3, 0x93, 0xfc, 0x48, 0xf1, 0x16, 0xba, 0x2e, 0xaa,
	0xa4, 0x17, 0xb6, 0xce, 0x4d, 0xed, 0xb4, 0x28, 0xa1, 0xab, 0x4b, 0x52, 0x35, 0xa5, 0x96, 0x56,
	0x90, 0xec, 0xd2, 0xc9, 0x97, 0x80, 0xbe, 0x53, 0xcf, 0xa9, 0x6c, 0x58, 0x13, 0x3e, 0x41, 0x7f,
	0xc6, 0x79, 0x7a, 0xb7, 0x87, 0x13, 0x1d, 0x3d, 0x71, 0x00, 0x65, 0x38, 0x3a, 0xaa, 0xd9, 0xfd,
	0x6b, 0xf4, 0x0f, 0x67, 0xe0, 0xdf, 0xdb, 0x93, 0x78, 0xca, 0xf1, 0xd9, 0x3d, 0x4b, 0xdf, 0xbe,
	0xa6, 0xe9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xba, 0xa9, 0xf4, 0x64, 0x02, 0x00, 0x00,
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
