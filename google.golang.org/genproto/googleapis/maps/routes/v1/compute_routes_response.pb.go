// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/maps/routes/v1/compute_routes_response.proto

package routes

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// ComputeRoutes the response message.
type ComputeRoutesResponse struct {
	// Contains an array of computed routes (up to three) when you specify
	// compute_alternatives_routes, and contains just one route when you don't.
	// When this array contains multiple entries, the first one is the most
	// recommended route. If the array is empty, then it means no route could be
	// found.
	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// In some cases when the server is not able to compute the route results with
	// all of the input preferences, it may fallback to using a different way of
	// computation. When fallback mode is used, this field contains detailed info
	// about the fallback response. Otherwise this field is unset.
	FallbackInfo         *FallbackInfo `protobuf:"bytes,2,opt,name=fallback_info,json=fallbackInfo,proto3" json:"fallback_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ComputeRoutesResponse) Reset()         { *m = ComputeRoutesResponse{} }
func (m *ComputeRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeRoutesResponse) ProtoMessage()    {}
func (*ComputeRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1609c8dc8046a099, []int{0}
}

func (m *ComputeRoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeRoutesResponse.Unmarshal(m, b)
}
func (m *ComputeRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeRoutesResponse.Marshal(b, m, deterministic)
}
func (m *ComputeRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeRoutesResponse.Merge(m, src)
}
func (m *ComputeRoutesResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeRoutesResponse.Size(m)
}
func (m *ComputeRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeRoutesResponse proto.InternalMessageInfo

func (m *ComputeRoutesResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *ComputeRoutesResponse) GetFallbackInfo() *FallbackInfo {
	if m != nil {
		return m.FallbackInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ComputeRoutesResponse)(nil), "google.maps.routes.v1.ComputeRoutesResponse")
}

func init() {
	proto.RegisterFile("google/maps/routes/v1/compute_routes_response.proto", fileDescriptor_1609c8dc8046a099)
}

var fileDescriptor_1609c8dc8046a099 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0x4d, 0x2c, 0x28, 0xd6, 0x2f, 0xca, 0x2f, 0x2d, 0x49, 0x2d, 0xd6, 0x2f,
	0x33, 0xd4, 0x4f, 0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49, 0x8d, 0x87, 0x88, 0xc4, 0x17, 0xa5, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x42, 0x34, 0xe9,
	0x81, 0x34, 0xe9, 0x41, 0x94, 0xe8, 0x95, 0x19, 0x4a, 0x69, 0x62, 0x37, 0x2b, 0x2d, 0x31, 0x27,
	0x27, 0x29, 0x31, 0x39, 0x3b, 0x3e, 0x33, 0x2f, 0x2d, 0x1f, 0x62, 0x82, 0x94, 0x22, 0x76, 0xa5,
	0x60, 0x16, 0x44, 0x89, 0xd2, 0x74, 0x46, 0x2e, 0x51, 0x67, 0x88, 0x33, 0x82, 0xc0, 0x0a, 0x82,
	0xa0, 0x8e, 0x10, 0x32, 0xe1, 0x62, 0x83, 0x68, 0x91, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92,
	0xd1, 0xc3, 0xea, 0x1e, 0x3d, 0xb0, 0xb6, 0x20, 0xa8, 0x5a, 0x21, 0x0f, 0x2e, 0x5e, 0x14, 0x97,
	0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x29, 0xe3, 0xd0, 0xec, 0x06, 0x55, 0xeb, 0x99, 0x97,
	0x96, 0x1f, 0xc4, 0x93, 0x86, 0xc4, 0x73, 0xda, 0xc0, 0xc8, 0x25, 0x99, 0x9c, 0x9f, 0x8b, 0x5d,
	0xa3, 0x93, 0x14, 0x56, 0x47, 0x07, 0x80, 0xfc, 0x14, 0xc0, 0x18, 0x65, 0x0d, 0xd5, 0x94, 0x9e,
	0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f, 0x9e, 0x9a, 0x07, 0xf6, 0xb1, 0x3e,
	0x44, 0x2a, 0xb1, 0x20, 0xb3, 0x18, 0x2d, 0x5c, 0xac, 0x21, 0xac, 0x1f, 0x8c, 0x8c, 0x8b, 0x98,
	0x58, 0xdc, 0x7d, 0x83, 0x82, 0x57, 0x31, 0x89, 0xba, 0x43, 0xcc, 0xf1, 0x05, 0x59, 0x0e, 0xb1,
	0x4a, 0x2f, 0xcc, 0xf0, 0x14, 0x4c, 0x3c, 0x06, 0x24, 0x1e, 0x03, 0x11, 0x8f, 0x09, 0x33, 0x4c,
	0x62, 0x03, 0xdb, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x33, 0x65, 0xa6, 0xef, 0x01,
	0x00, 0x00,
}