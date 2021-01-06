// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/change_status.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// Describes the status of returned resource.
type ChangeStatus struct {
	// Output only. The resource name of the change status.
	// Change status resource names have the form:
	//
	// `customers/{customer_id}/changeStatus/{change_status_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Time at which the most recent change has occurred on this resource.
	LastChangeDateTime *wrappers.StringValue `protobuf:"bytes,3,opt,name=last_change_date_time,json=lastChangeDateTime,proto3" json:"last_change_date_time,omitempty"`
	// Output only. Represents the type of the changed resource. This dictates what fields
	// will be set. For example, for AD_GROUP, campaign and ad_group fields will
	// be set.
	ResourceType enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType `protobuf:"varint,4,opt,name=resource_type,json=resourceType,proto3,enum=google.ads.googleads.v2.enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType" json:"resource_type,omitempty"`
	// Output only. The Campaign affected by this change.
	Campaign *wrappers.StringValue `protobuf:"bytes,5,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Output only. The AdGroup affected by this change.
	AdGroup *wrappers.StringValue `protobuf:"bytes,6,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Output only. Represents the status of the changed resource.
	ResourceStatus enums.ChangeStatusOperationEnum_ChangeStatusOperation `protobuf:"varint,8,opt,name=resource_status,json=resourceStatus,proto3,enum=google.ads.googleads.v2.enums.ChangeStatusOperationEnum_ChangeStatusOperation" json:"resource_status,omitempty"`
	// Output only. The AdGroupAd affected by this change.
	AdGroupAd *wrappers.StringValue `protobuf:"bytes,9,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	// Output only. The AdGroupCriterion affected by this change.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,10,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// Output only. The CampaignCriterion affected by this change.
	CampaignCriterion *wrappers.StringValue `protobuf:"bytes,11,opt,name=campaign_criterion,json=campaignCriterion,proto3" json:"campaign_criterion,omitempty"`
	// Output only. The Feed affected by this change.
	Feed *wrappers.StringValue `protobuf:"bytes,12,opt,name=feed,proto3" json:"feed,omitempty"`
	// Output only. The FeedItem affected by this change.
	FeedItem *wrappers.StringValue `protobuf:"bytes,13,opt,name=feed_item,json=feedItem,proto3" json:"feed_item,omitempty"`
	// Output only. The AdGroupFeed affected by this change.
	AdGroupFeed *wrappers.StringValue `protobuf:"bytes,14,opt,name=ad_group_feed,json=adGroupFeed,proto3" json:"ad_group_feed,omitempty"`
	// Output only. The CampaignFeed affected by this change.
	CampaignFeed *wrappers.StringValue `protobuf:"bytes,15,opt,name=campaign_feed,json=campaignFeed,proto3" json:"campaign_feed,omitempty"`
	// Output only. The AdGroupBidModifier affected by this change.
	AdGroupBidModifier   *wrappers.StringValue `protobuf:"bytes,16,opt,name=ad_group_bid_modifier,json=adGroupBidModifier,proto3" json:"ad_group_bid_modifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ChangeStatus) Reset()         { *m = ChangeStatus{} }
func (m *ChangeStatus) String() string { return proto.CompactTextString(m) }
func (*ChangeStatus) ProtoMessage()    {}
func (*ChangeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc122555564308c, []int{0}
}

func (m *ChangeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatus.Unmarshal(m, b)
}
func (m *ChangeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatus.Marshal(b, m, deterministic)
}
func (m *ChangeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatus.Merge(m, src)
}
func (m *ChangeStatus) XXX_Size() int {
	return xxx_messageInfo_ChangeStatus.Size(m)
}
func (m *ChangeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatus proto.InternalMessageInfo

func (m *ChangeStatus) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ChangeStatus) GetLastChangeDateTime() *wrappers.StringValue {
	if m != nil {
		return m.LastChangeDateTime
	}
	return nil
}

func (m *ChangeStatus) GetResourceType() enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType {
	if m != nil {
		return m.ResourceType
	}
	return enums.ChangeStatusResourceTypeEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *ChangeStatus) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *ChangeStatus) GetResourceStatus() enums.ChangeStatusOperationEnum_ChangeStatusOperation {
	if m != nil {
		return m.ResourceStatus
	}
	return enums.ChangeStatusOperationEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetAdGroupAd() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupAd
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *ChangeStatus) GetCampaignCriterion() *wrappers.StringValue {
	if m != nil {
		return m.CampaignCriterion
	}
	return nil
}

func (m *ChangeStatus) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *ChangeStatus) GetFeedItem() *wrappers.StringValue {
	if m != nil {
		return m.FeedItem
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupFeed() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupFeed
	}
	return nil
}

func (m *ChangeStatus) GetCampaignFeed() *wrappers.StringValue {
	if m != nil {
		return m.CampaignFeed
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupBidModifier() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupBidModifier
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeStatus)(nil), "google.ads.googleads.v2.resources.ChangeStatus")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/change_status.proto", fileDescriptor_edc122555564308c)
}

var fileDescriptor_edc122555564308c = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4d, 0x4f, 0xe3, 0x46,
	0x18, 0xc7, 0x15, 0x42, 0x69, 0x32, 0x49, 0x78, 0x19, 0x09, 0xc9, 0x45, 0xb4, 0x04, 0x4a, 0xda,
	0xf0, 0x66, 0x0b, 0xa3, 0x5e, 0xcc, 0xa1, 0x72, 0x68, 0x8b, 0x5a, 0xa9, 0x14, 0x99, 0x28, 0x95,
	0x50, 0x2a, 0x6b, 0xe2, 0x99, 0x98, 0x91, 0x62, 0x8f, 0xe5, 0x17, 0x2a, 0x44, 0xe9, 0x5e, 0xf6,
	0x93, 0xec, 0x71, 0x3f, 0xca, 0x7e, 0x0a, 0xce, 0x5c, 0xf6, 0xbe, 0x87, 0xd5, 0xca, 0xe3, 0xf1,
	0x24, 0x80, 0xb2, 0x31, 0xd2, 0x9e, 0x78, 0x3c, 0xf3, 0x3c, 0xff, 0xdf, 0xf3, 0x7f, 0xc6, 0xf1,
	0x00, 0x7e, 0x72, 0x19, 0x73, 0x47, 0x44, 0x43, 0x38, 0xd2, 0xb2, 0x30, 0x8d, 0xae, 0x75, 0x2d,
	0x24, 0x11, 0x4b, 0x42, 0x87, 0x44, 0x9a, 0x73, 0x85, 0x7c, 0x97, 0xd8, 0x51, 0x8c, 0xe2, 0x24,
	0x52, 0x83, 0x90, 0xc5, 0x0c, 0x6e, 0x66, 0xb9, 0x2a, 0xc2, 0x91, 0x2a, 0xcb, 0xd4, 0x6b, 0x5d,
	0x95, 0x65, 0x6b, 0xc7, 0xd3, 0x94, 0x89, 0x9f, 0x78, 0x4f, 0x54, 0x6d, 0x16, 0x90, 0x10, 0xc5,
	0x94, 0xf9, 0x99, 0xfe, 0xda, 0xcf, 0x2f, 0x29, 0xce, 0x99, 0x76, 0x7c, 0x13, 0x10, 0x21, 0xb0,
	0x91, 0x0b, 0x04, 0x54, 0x1b, 0x52, 0x32, 0xc2, 0xf6, 0x80, 0x5c, 0xa1, 0x6b, 0xca, 0x42, 0x91,
	0xf0, 0xcd, 0x44, 0x42, 0x2e, 0x20, 0xb6, 0xbe, 0x13, 0x5b, 0xfc, 0x69, 0x90, 0x0c, 0xb5, 0x7f,
	0x43, 0x14, 0x04, 0x24, 0x14, 0xe6, 0xd7, 0xd6, 0x27, 0x4a, 0x91, 0xef, 0xb3, 0x98, 0x77, 0x2e,
	0x76, 0xb7, 0xde, 0xd7, 0x41, 0xfd, 0x84, 0xf7, 0x77, 0xc1, 0xdb, 0x83, 0x16, 0x68, 0xc8, 0x0e,
	0x7d, 0xe4, 0x11, 0xa5, 0xd4, 0x2c, 0xb5, 0xab, 0x9d, 0x83, 0x7b, 0xb3, 0xfc, 0xc1, 0xfc, 0x11,
	0xb4, 0xc6, 0xf3, 0x13, 0x51, 0x40, 0x23, 0xd5, 0x61, 0x9e, 0x36, 0xa9, 0x62, 0xd5, 0x73, 0x8d,
	0x33, 0xe4, 0x11, 0xd8, 0x05, 0xab, 0x23, 0x14, 0xc5, 0xb6, 0x18, 0x04, 0x46, 0x31, 0xb1, 0x63,
	0xea, 0x11, 0xa5, 0xdc, 0x2c, 0xb5, 0x6b, 0xfa, 0xba, 0x90, 0x52, 0x73, 0x0b, 0xea, 0x45, 0x1c,
	0x52, 0xdf, 0xed, 0xa1, 0x51, 0x42, 0x3a, 0xe5, 0x7b, 0xb3, 0x6c, 0xc1, 0xb4, 0x3e, 0x23, 0xfc,
	0x82, 0x62, 0xd2, 0xa5, 0x1e, 0x81, 0xff, 0x4f, 0x74, 0x9a, 0xce, 0x52, 0x99, 0x6f, 0x96, 0xda,
	0x8b, 0x7a, 0x57, 0x9d, 0x76, 0xda, 0xfc, 0x34, 0xd4, 0x47, 0x7d, 0x8a, 0xfa, 0xee, 0x4d, 0x40,
	0x7e, 0xf5, 0x13, 0x6f, 0xea, 0x66, 0xd6, 0x85, 0x74, 0x95, 0x2e, 0x41, 0x07, 0x54, 0x1c, 0xe4,
	0x05, 0x88, 0xba, 0xbe, 0xf2, 0x55, 0x01, 0x23, 0x3b, 0x7c, 0x84, 0xdf, 0x83, 0xcd, 0xe9, 0x23,
	0x14, 0x72, 0x96, 0x14, 0x86, 0x08, 0x54, 0x10, 0xb6, 0xdd, 0x90, 0x25, 0x81, 0xb2, 0x50, 0x00,
	0xd2, 0xe6, 0x90, 0x2d, 0xd0, 0x9c, 0x0a, 0x31, 0xf1, 0x69, 0xaa, 0x66, 0x7d, 0x8d, 0xb2, 0x00,
	0xfe, 0x07, 0x96, 0xe4, 0x1c, 0xb3, 0x77, 0x54, 0xa9, 0xf0, 0x49, 0x9e, 0xbd, 0x60, 0x92, 0x7f,
	0xe5, 0x3f, 0x89, 0x67, 0x63, 0x94, 0x3b, 0xd9, 0x0c, 0x17, 0x73, 0x96, 0x78, 0xdf, 0xae, 0x40,
	0x2d, 0x37, 0x68, 0x23, 0xac, 0x54, 0x0b, 0x78, 0xdc, 0xe5, 0x1e, 0xb7, 0xc1, 0xd6, 0x2c, 0x8f,
	0x26, 0xb6, 0xaa, 0x28, 0x0f, 0xe1, 0x1d, 0x80, 0x92, 0xe4, 0x84, 0x34, 0x26, 0x21, 0x65, 0xbe,
	0x02, 0x0a, 0x00, 0x0f, 0x39, 0x70, 0x0f, 0xec, 0xcc, 0x02, 0x9e, 0xe4, 0xb2, 0xd6, 0x32, 0x7a,
	0xb2, 0x02, 0x5f, 0x01, 0x98, 0x9f, 0xea, 0x04, 0xbe, 0x56, 0x00, 0xaf, 0x73, 0xfc, 0x3e, 0xd8,
	0x9d, 0xf9, 0xe2, 0x8c, 0xf9, 0x2b, 0xce, 0xd3, 0x25, 0xf8, 0x37, 0x98, 0x1f, 0x12, 0x82, 0x95,
	0x7a, 0x01, 0x64, 0x8b, 0x23, 0x37, 0xc0, 0xb7, 0x53, 0x91, 0xbf, 0x11, 0x82, 0x2d, 0x2e, 0x08,
	0x31, 0xa8, 0xa6, 0x7f, 0x6d, 0x1a, 0x13, 0x4f, 0x69, 0x7c, 0x91, 0x5f, 0x42, 0xaa, 0xfe, 0x7b,
	0x4c, 0x3c, 0xab, 0x32, 0x14, 0x11, 0xf4, 0x41, 0x43, 0x1e, 0x1f, 0xf7, 0xb1, 0x58, 0x80, 0xb4,
	0xcf, 0x49, 0x3f, 0x80, 0xed, 0x59, 0x27, 0xc7, 0xed, 0xd4, 0xd0, 0xf8, 0x01, 0x06, 0xa0, 0x21,
	0xcf, 0x8b, 0xf3, 0x96, 0x0a, 0xf0, 0x0a, 0x7c, 0x26, 0x85, 0x26, 0x07, 0xd6, 0x9d, 0x89, 0x27,
	0xf8, 0xba, 0x04, 0x56, 0xa5, 0xc5, 0x01, 0xc5, 0xb6, 0xc7, 0x30, 0x1d, 0x52, 0x12, 0x2a, 0xcb,
	0x05, 0xd0, 0x47, 0x1c, 0x7d, 0x00, 0xf6, 0x66, 0x59, 0xed, 0x50, 0xfc, 0xa7, 0x10, 0xb6, 0x20,
	0x7a, 0xb6, 0x66, 0xfc, 0xf3, 0x60, 0x5e, 0x16, 0xfc, 0xce, 0xc3, 0x43, 0x27, 0x89, 0x62, 0xe6,
	0x91, 0x30, 0xd2, 0x6e, 0xf3, 0xf0, 0x4e, 0x5c, 0x78, 0x59, 0x8a, 0x76, 0xfb, 0xe8, 0xfa, 0xbb,
	0xeb, 0x7c, 0x2c, 0x81, 0x96, 0xc3, 0x3c, 0x75, 0xe6, 0x9d, 0xdc, 0x59, 0x99, 0x44, 0x9d, 0xa7,
	0x9e, 0xcf, 0x4b, 0x97, 0x7f, 0x88, 0x3a, 0x97, 0x8d, 0x90, 0xef, 0xaa, 0x2c, 0x74, 0x35, 0x97,
	0xf8, 0x7c, 0x22, 0xda, 0xb8, 0xd3, 0xcf, 0xfc, 0x87, 0x70, 0x2c, 0xa3, 0x37, 0x73, 0xe5, 0x53,
	0xd3, 0x7c, 0x3b, 0xb7, 0x79, 0x9a, 0x49, 0x9a, 0x38, 0x52, 0xb3, 0x30, 0x8d, 0x7a, 0xba, 0x9a,
	0xdf, 0x00, 0xd1, 0xbb, 0x3c, 0xa7, 0x6f, 0xe2, 0xa8, 0x2f, 0x73, 0xfa, 0x3d, 0xbd, 0x2f, 0x73,
	0x1e, 0xe6, 0x5a, 0xd9, 0x86, 0x61, 0x98, 0x38, 0x32, 0x0c, 0x99, 0x65, 0x18, 0x3d, 0xdd, 0x30,
	0x64, 0xde, 0x60, 0x81, 0x37, 0x7b, 0xf4, 0x29, 0x00, 0x00, 0xff, 0xff, 0x37, 0xab, 0x8a, 0xa7,
	0xcd, 0x08, 0x00, 0x00,
}