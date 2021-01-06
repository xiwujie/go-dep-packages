// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_parameter.proto

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

// An ad parameter that is used to update numeric values (such as prices or
// inventory levels) in any text line of an ad (including URLs). There can
// be a maximum of two AdParameters per ad group criterion. (One with
// parameter_index = 1 and one with parameter_index = 2.)
// In the ad the parameters are referenced by a placeholder of the form
// "{param#:value}". E.g. "{param1:$17}"
type AdParameter struct {
	// Immutable. The resource name of the ad parameter.
	// Ad parameter resource names have the form:
	//
	// `customers/{customer_id}/adParameters/{ad_group_id}~{criterion_id}~{parameter_index}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The ad group criterion that this ad parameter belongs to.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// Immutable. The unique index of this ad parameter. Must be either 1 or 2.
	ParameterIndex *wrappers.Int64Value `protobuf:"bytes,3,opt,name=parameter_index,json=parameterIndex,proto3" json:"parameter_index,omitempty"`
	// Numeric value to insert into the ad text. The following restrictions
	//  apply:
	//  - Can use comma or period as a separator, with an optional period or
	//    comma (respectively) for fractional values. For example, 1,000,000.00
	//    and 2.000.000,10 are valid.
	//  - Can be prepended or appended with a currency symbol. For example,
	//    $99.99 is valid.
	//  - Can be prepended or appended with a currency code. For example, 99.99USD
	//    and EUR200 are valid.
	//  - Can use '%'. For example, 1.0% and 1,0% are valid.
	//  - Can use plus or minus. For example, -10.99 and 25+ are valid.
	//  - Can use '/' between two numbers. For example 4/1 and 0.95/0.45 are
	//    valid.
	InsertionText        *wrappers.StringValue `protobuf:"bytes,4,opt,name=insertion_text,json=insertionText,proto3" json:"insertion_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdParameter) Reset()         { *m = AdParameter{} }
func (m *AdParameter) String() string { return proto.CompactTextString(m) }
func (*AdParameter) ProtoMessage()    {}
func (*AdParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_80d3f7afedc49596, []int{0}
}

func (m *AdParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdParameter.Unmarshal(m, b)
}
func (m *AdParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdParameter.Marshal(b, m, deterministic)
}
func (m *AdParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdParameter.Merge(m, src)
}
func (m *AdParameter) XXX_Size() int {
	return xxx_messageInfo_AdParameter.Size(m)
}
func (m *AdParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_AdParameter.DiscardUnknown(m)
}

var xxx_messageInfo_AdParameter proto.InternalMessageInfo

func (m *AdParameter) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdParameter) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *AdParameter) GetParameterIndex() *wrappers.Int64Value {
	if m != nil {
		return m.ParameterIndex
	}
	return nil
}

func (m *AdParameter) GetInsertionText() *wrappers.StringValue {
	if m != nil {
		return m.InsertionText
	}
	return nil
}

func init() {
	proto.RegisterType((*AdParameter)(nil), "google.ads.googleads.v1.resources.AdParameter")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_parameter.proto", fileDescriptor_80d3f7afedc49596)
}

var fileDescriptor_80d3f7afedc49596 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x6a, 0x13, 0x41,
	0x18, 0x67, 0x13, 0x2b, 0xb8, 0xb5, 0xb5, 0xec, 0x29, 0xd6, 0xa2, 0xad, 0x58, 0xa9, 0x28, 0x33,
	0xae, 0x16, 0x0f, 0xeb, 0x69, 0xd2, 0x43, 0x89, 0x07, 0x89, 0xab, 0x04, 0xd1, 0xc0, 0x32, 0xd9,
	0xf9, 0xba, 0x0e, 0x64, 0x67, 0x96, 0x99, 0xd9, 0x18, 0x28, 0x79, 0x19, 0x8f, 0x3e, 0x8a, 0xf8,
	0x10, 0x3d, 0xf7, 0x11, 0x04, 0x41, 0xf6, 0xcf, 0x4c, 0x16, 0x85, 0xd6, 0xdb, 0x2f, 0xf9, 0x7e,
	0xff, 0xbe, 0xd9, 0xcf, 0x3f, 0xce, 0xa4, 0xcc, 0xe6, 0x80, 0x29, 0xd3, 0xb8, 0x81, 0x15, 0x5a,
	0x84, 0x58, 0x81, 0x96, 0xa5, 0x4a, 0x41, 0x63, 0xca, 0x92, 0x82, 0x2a, 0x9a, 0x83, 0x01, 0x85,
	0x0a, 0x25, 0x8d, 0x0c, 0x0e, 0x1a, 0x2a, 0xa2, 0x4c, 0x23, 0xa7, 0x42, 0x8b, 0x10, 0x39, 0xd5,
	0xee, 0x03, 0x6b, 0x5c, 0x70, 0x7c, 0xc6, 0x61, 0xce, 0x92, 0x19, 0x7c, 0xa1, 0x0b, 0x2e, 0x5b,
	0x8f, 0xdd, 0xbb, 0x1d, 0x82, 0x95, 0xb5, 0xa3, 0xfb, 0xed, 0xa8, 0xfe, 0x35, 0x2b, 0xcf, 0xf0,
	0x57, 0x45, 0x8b, 0x02, 0x94, 0x6e, 0xe7, 0x7b, 0x1d, 0x29, 0x15, 0x42, 0x1a, 0x6a, 0xb8, 0x14,
	0xed, 0xf4, 0xe1, 0xcf, 0xbe, 0xbf, 0x49, 0xd8, 0xd8, 0x56, 0x0e, 0xde, 0xf9, 0x5b, 0xd6, 0x3f,
	0x11, 0x34, 0x87, 0x81, 0xb7, 0xef, 0x1d, 0xdd, 0x1a, 0x3e, 0xbb, 0x20, 0x1b, 0xbf, 0xc8, 0x63,
	0xff, 0xd1, 0x7a, 0x81, 0x16, 0x15, 0x5c, 0xa3, 0x54, 0xe6, 0xb8, 0x63, 0x12, 0xdf, 0xb6, 0x16,
	0x6f, 0x69, 0x0e, 0xc1, 0xca, 0x0f, 0x28, 0x4b, 0x32, 0x25, 0xcb, 0x22, 0x49, 0x15, 0x37, 0xa0,
	0xb8, 0x14, 0x83, 0xde, 0xbe, 0x77, 0xb4, 0xf9, 0x62, 0xaf, 0xb5, 0x41, 0xb6, 0x3d, 0x7a, 0x6f,
	0x14, 0x17, 0xd9, 0x84, 0xce, 0x4b, 0x18, 0x86, 0x75, 0xea, 0x53, 0xff, 0xc9, 0x15, 0xa9, 0xa7,
	0x95, 0xef, 0x89, 0xb5, 0x8d, 0x77, 0xe8, 0x5f, 0xff, 0x04, 0x23, 0xff, 0x8e, 0xfb, 0x22, 0x09,
	0x17, 0x0c, 0x96, 0x83, 0x7e, 0x9d, 0x7d, 0xef, 0x9f, 0xec, 0x91, 0x30, 0xaf, 0x8e, 0x9b, 0xe8,
	0xfe, 0x05, 0xd9, 0x88, 0xb7, 0x9d, 0x70, 0x54, 0xe9, 0x82, 0x13, 0x7f, 0x9b, 0x0b, 0x0d, 0xaa,
	0x7a, 0xc0, 0xc4, 0xc0, 0xd2, 0x0c, 0x6e, 0x5c, 0xbf, 0x45, 0xbc, 0xe5, 0x34, 0x1f, 0x60, 0x69,
	0xa2, 0xcf, 0x97, 0xe4, 0xe3, 0xff, 0xbd, 0x63, 0xf0, 0x3c, 0x2d, 0xb5, 0x91, 0x39, 0x28, 0x8d,
	0xcf, 0x2d, 0x5c, 0x61, 0xba, 0x66, 0x68, 0x7c, 0xde, 0x3d, 0xb8, 0xd5, 0xf0, 0xb7, 0xe7, 0x1f,
	0xa6, 0x32, 0x47, 0xd7, 0x9e, 0xdc, 0x70, 0xa7, 0x13, 0x34, 0xae, 0x6a, 0x8f, 0xbd, 0x4f, 0x6f,
	0x5a, 0x59, 0x26, 0xe7, 0x54, 0x64, 0x48, 0xaa, 0x0c, 0x67, 0x20, 0xea, 0xa5, 0xf0, 0xba, 0xe6,
	0x15, 0xe7, 0xff, 0xda, 0xa1, 0x6f, 0xbd, 0xfe, 0x29, 0x21, 0xdf, 0x7b, 0x07, 0xa7, 0x8d, 0x25,
	0x61, 0x1a, 0x35, 0xb0, 0x42, 0x93, 0x10, 0xc5, 0x96, 0xf9, 0xc3, 0x72, 0xa6, 0x84, 0xe9, 0xa9,
	0xe3, 0x4c, 0x27, 0xe1, 0xd4, 0x71, 0x2e, 0x7b, 0x87, 0xcd, 0x20, 0x8a, 0x08, 0xd3, 0x51, 0xe4,
	0x58, 0x51, 0x34, 0x09, 0xa3, 0xc8, 0xf1, 0x66, 0x37, 0xeb, 0xb2, 0x2f, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xd4, 0xf6, 0xf2, 0x89, 0xaa, 0x03, 0x00, 0x00,
}
