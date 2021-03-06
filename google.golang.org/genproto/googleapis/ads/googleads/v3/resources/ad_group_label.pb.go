// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/ad_group_label.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A relationship between an ad group and a label.
type AdGroupLabel struct {
	// Immutable. The resource name of the ad group label.
	// Ad group label resource names have the form:
	// `customers/{customer_id}/adGroupLabels/{ad_group_id}~{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The ad group to which the label is attached.
	AdGroup *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Immutable. The label assigned to the ad group.
	Label                *wrappers.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdGroupLabel) Reset()         { *m = AdGroupLabel{} }
func (m *AdGroupLabel) String() string { return proto.CompactTextString(m) }
func (*AdGroupLabel) ProtoMessage()    {}
func (*AdGroupLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc4fee60b3e79c3f, []int{0}
}

func (m *AdGroupLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupLabel.Unmarshal(m, b)
}
func (m *AdGroupLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupLabel.Marshal(b, m, deterministic)
}
func (m *AdGroupLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupLabel.Merge(m, src)
}
func (m *AdGroupLabel) XXX_Size() int {
	return xxx_messageInfo_AdGroupLabel.Size(m)
}
func (m *AdGroupLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupLabel.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupLabel proto.InternalMessageInfo

func (m *AdGroupLabel) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupLabel) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupLabel) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*AdGroupLabel)(nil), "google.ads.googleads.v3.resources.AdGroupLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/ad_group_label.proto", fileDescriptor_cc4fee60b3e79c3f)
}

var fileDescriptor_cc4fee60b3e79c3f = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x8b, 0xd4, 0x30,
	0x1c, 0xa5, 0x1d, 0xd6, 0x3f, 0x71, 0x3d, 0xd8, 0xd3, 0xb8, 0x2c, 0x6b, 0x77, 0x61, 0x74, 0x2e,
	0x26, 0x60, 0xc1, 0x43, 0x3c, 0xa5, 0x97, 0x01, 0x11, 0x59, 0x2a, 0xf4, 0xb0, 0x14, 0x4a, 0xda,
	0x64, 0x63, 0xa1, 0x6d, 0x4a, 0xd2, 0x8e, 0x87, 0x65, 0x2f, 0x7e, 0x14, 0x8f, 0x7e, 0x14, 0x3f,
	0x85, 0xe7, 0xf9, 0x08, 0x1e, 0x44, 0xda, 0x34, 0x9d, 0x7a, 0xd0, 0x99, 0xdb, 0x0b, 0xbf, 0xf7,
	0xde, 0xef, 0xbd, 0xf0, 0x03, 0x6f, 0x85, 0x94, 0xa2, 0xe4, 0x88, 0x32, 0x8d, 0x0c, 0xec, 0xd1,
	0x36, 0x40, 0x8a, 0x6b, 0xd9, 0xa9, 0x9c, 0x6b, 0x44, 0x59, 0x2a, 0x94, 0xec, 0x9a, 0xb4, 0xa4,
	0x19, 0x2f, 0x61, 0xa3, 0x64, 0x2b, 0xbd, 0x4b, 0x43, 0x86, 0x94, 0x69, 0x38, 0xe9, 0xe0, 0x36,
	0x80, 0x93, 0xee, 0xec, 0x85, 0xb5, 0x6e, 0x0a, 0x74, 0x5b, 0xf0, 0x92, 0xa5, 0x19, 0xff, 0x4c,
	0xb7, 0x85, 0x54, 0xc6, 0xe3, 0xec, 0xf9, 0x8c, 0x60, 0x65, 0xe3, 0xe8, 0x62, 0x1c, 0x0d, 0xaf,
	0xac, 0xbb, 0x45, 0x5f, 0x14, 0x6d, 0x1a, 0xae, 0xf4, 0x38, 0x3f, 0x9f, 0x49, 0x69, 0x5d, 0xcb,
	0x96, 0xb6, 0x85, 0xac, 0xc7, 0xe9, 0xd5, 0xd7, 0x05, 0x38, 0x25, 0x6c, 0xd3, 0x87, 0xfe, 0xd0,
	0x67, 0xf6, 0x22, 0xf0, 0xd4, 0x2e, 0x48, 0x6b, 0x5a, 0xf1, 0xa5, 0xe3, 0x3b, 0xeb, 0xc7, 0xe1,
	0xeb, 0x9f, 0xe4, 0xe4, 0x17, 0x79, 0x05, 0x56, 0xfb, 0x06, 0x23, 0x6a, 0x0a, 0x0d, 0x73, 0x59,
	0xa1, 0xb9, 0x4b, 0x74, 0x6a, 0x3d, 0x3e, 0xd2, 0x8a, 0x7b, 0x14, 0x3c, 0xb2, 0x3f, 0xb3, 0x74,
	0x7d, 0x67, 0xfd, 0xe4, 0xcd, 0xf9, 0xa8, 0x86, 0x36, 0x35, 0xfc, 0xd4, 0xaa, 0xa2, 0x16, 0x31,
	0x2d, 0x3b, 0x1e, 0xae, 0x87, 0x65, 0x57, 0xc0, 0x3f, 0xb4, 0x2c, 0x7a, 0x48, 0x0d, 0xf0, 0x6e,
	0xc0, 0xc9, 0xf0, 0xe7, 0xcb, 0xc5, 0x11, 0xfe, 0x2f, 0x07, 0x7f, 0x1f, 0x5c, 0xfc, 0xd3, 0xdf,
	0xb4, 0x30, 0x96, 0x38, 0xdd, 0x91, 0xe4, 0xc8, 0xe2, 0x5e, 0x90, 0x77, 0xba, 0x95, 0x15, 0x57,
	0x1a, 0xdd, 0x59, 0x78, 0x8f, 0xe8, 0x8c, 0xa2, 0xd1, 0xdd, 0xdf, 0x67, 0x72, 0x1f, 0xfe, 0x76,
	0xc0, 0x2a, 0x97, 0x15, 0x3c, 0x78, 0x28, 0xe1, 0xb3, 0xf9, 0xb2, 0xeb, 0xbe, 0xdb, 0xb5, 0x73,
	0xf3, 0x7e, 0xd4, 0x09, 0x59, 0xd2, 0x5a, 0x40, 0xa9, 0x04, 0x12, 0xbc, 0x1e, 0x9a, 0xa3, 0x7d,
	0xd6, 0xff, 0xdc, 0xed, 0xbb, 0x09, 0x7d, 0x73, 0x17, 0x1b, 0x42, 0xbe, 0xbb, 0x97, 0x1b, 0x63,
	0x49, 0x98, 0x86, 0x06, 0xf6, 0x28, 0x0e, 0x60, 0x64, 0x99, 0x3f, 0x2c, 0x27, 0x21, 0x4c, 0x27,
	0x13, 0x27, 0x89, 0x83, 0x64, 0xe2, 0xec, 0xdc, 0x95, 0x19, 0x60, 0x4c, 0x98, 0xc6, 0x78, 0x62,
	0x61, 0x1c, 0x07, 0x18, 0x4f, 0xbc, 0xec, 0xc1, 0x10, 0x36, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x0f, 0x28, 0x03, 0xd2, 0x63, 0x03, 0x00, 0x00,
}
