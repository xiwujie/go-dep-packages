// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/ad_group_ad_rotation_mode.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The possible ad rotation modes of an ad group.
type AdGroupAdRotationModeEnum_AdGroupAdRotationMode int32

const (
	// The ad rotation mode has not been specified.
	AdGroupAdRotationModeEnum_UNSPECIFIED AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupAdRotationModeEnum_UNKNOWN AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 1
	// Optimize ad group ads based on clicks or conversions.
	AdGroupAdRotationModeEnum_OPTIMIZE AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 2
	// Rotate evenly forever.
	AdGroupAdRotationModeEnum_ROTATE_FOREVER AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 3
)

var AdGroupAdRotationModeEnum_AdGroupAdRotationMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPTIMIZE",
	3: "ROTATE_FOREVER",
}

var AdGroupAdRotationModeEnum_AdGroupAdRotationMode_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"OPTIMIZE":       2,
	"ROTATE_FOREVER": 3,
}

func (x AdGroupAdRotationModeEnum_AdGroupAdRotationMode) String() string {
	return proto.EnumName(AdGroupAdRotationModeEnum_AdGroupAdRotationMode_name, int32(x))
}

func (AdGroupAdRotationModeEnum_AdGroupAdRotationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2286469973ba5e1d, []int{0, 0}
}

// Container for enum describing possible ad rotation modes of ads within an
// ad group.
type AdGroupAdRotationModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAdRotationModeEnum) Reset()         { *m = AdGroupAdRotationModeEnum{} }
func (m *AdGroupAdRotationModeEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdRotationModeEnum) ProtoMessage()    {}
func (*AdGroupAdRotationModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_2286469973ba5e1d, []int{0}
}

func (m *AdGroupAdRotationModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdRotationModeEnum.Unmarshal(m, b)
}
func (m *AdGroupAdRotationModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdRotationModeEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupAdRotationModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdRotationModeEnum.Merge(m, src)
}
func (m *AdGroupAdRotationModeEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdRotationModeEnum.Size(m)
}
func (m *AdGroupAdRotationModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdRotationModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdRotationModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode", AdGroupAdRotationModeEnum_AdGroupAdRotationMode_name, AdGroupAdRotationModeEnum_AdGroupAdRotationMode_value)
	proto.RegisterType((*AdGroupAdRotationModeEnum)(nil), "google.ads.googleads.v3.enums.AdGroupAdRotationModeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/ad_group_ad_rotation_mode.proto", fileDescriptor_2286469973ba5e1d)
}

var fileDescriptor_2286469973ba5e1d = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x5d, 0x07, 0x2a, 0x99, 0x68, 0x29, 0x78, 0x70, 0xb8, 0xc3, 0xf6, 0x00, 0xe9, 0xa1,
	0xb7, 0x88, 0x87, 0x4c, 0xb3, 0x51, 0x64, 0x6d, 0xa9, 0x5b, 0x07, 0xa3, 0x50, 0xa2, 0x29, 0x61,
	0xb0, 0xe6, 0x5f, 0x9a, 0x6e, 0x0f, 0xe4, 0xd1, 0x47, 0xf1, 0x51, 0xbc, 0xf8, 0x0a, 0xd2, 0x66,
	0xdb, 0x69, 0x7a, 0x09, 0x1f, 0xf9, 0xfe, 0xbf, 0x2f, 0xff, 0x7c, 0xe8, 0x51, 0x02, 0xc8, 0x4d,
	0xee, 0x72, 0xa1, 0x5d, 0x23, 0x1b, 0xb5, 0xf3, 0xdc, 0x5c, 0x6d, 0x0b, 0xed, 0x72, 0x91, 0xc9,
	0x0a, 0xb6, 0x65, 0xc6, 0x45, 0x56, 0x41, 0xcd, 0xeb, 0x35, 0xa8, 0xac, 0x00, 0x91, 0xe3, 0xb2,
	0x82, 0x1a, 0x9c, 0x81, 0x61, 0x30, 0x17, 0x1a, 0x1f, 0x71, 0xbc, 0xf3, 0x70, 0x8b, 0xf7, 0xef,
	0x0f, 0xe9, 0xe5, 0xda, 0xe5, 0x4a, 0xed, 0x03, 0xb4, 0x81, 0x47, 0x35, 0xba, 0xa3, 0x62, 0xda,
	0xc4, 0x53, 0x11, 0xef, 0xbd, 0x19, 0x88, 0x9c, 0xa9, 0x6d, 0x31, 0x5a, 0xa2, 0xdb, 0x93, 0xa6,
	0x73, 0x83, 0x7a, 0x8b, 0xe0, 0x35, 0x62, 0x4f, 0xfe, 0xc4, 0x67, 0xcf, 0xf6, 0x99, 0xd3, 0x43,
	0x17, 0x8b, 0xe0, 0x25, 0x08, 0x97, 0x81, 0xdd, 0x71, 0xae, 0xd0, 0x65, 0x18, 0xcd, 0xfd, 0x99,
	0xbf, 0x62, 0xb6, 0xe5, 0x38, 0xe8, 0x3a, 0x0e, 0xe7, 0x74, 0xce, 0xb2, 0x49, 0x18, 0xb3, 0x84,
	0xc5, 0x76, 0x77, 0xfc, 0xd3, 0x41, 0xc3, 0x77, 0x28, 0xf0, 0xbf, 0x9b, 0x8f, 0xfb, 0x27, 0x1f,
	0x8f, 0x9a, 0xbd, 0xa3, 0xce, 0x6a, 0xbc, 0x87, 0x25, 0x6c, 0xb8, 0x92, 0x18, 0x2a, 0xe9, 0xca,
	0x5c, 0xb5, 0xbf, 0x3a, 0xb4, 0x58, 0xae, 0xf5, 0x1f, 0xa5, 0x3e, 0xb4, 0xe7, 0x87, 0xd5, 0x9d,
	0x52, 0xfa, 0x69, 0x0d, 0xa6, 0x26, 0x8a, 0x0a, 0x8d, 0x8d, 0x6c, 0x54, 0xe2, 0xe1, 0xa6, 0x05,
	0xfd, 0x75, 0xf0, 0x53, 0x2a, 0x74, 0x7a, 0xf4, 0xd3, 0xc4, 0x4b, 0x5b, 0xff, 0xdb, 0x1a, 0x9a,
	0x4b, 0x42, 0xa8, 0xd0, 0x84, 0x1c, 0x27, 0x08, 0x49, 0x3c, 0x42, 0xda, 0x99, 0xb7, 0xf3, 0x76,
	0x31, 0xef, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x79, 0x2f, 0xf6, 0xec, 0x01, 0x00, 0x00,
}
