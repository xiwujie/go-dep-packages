// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/customer_extension_setting_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for
// [CustomerExtensionSettingService.GetCustomerExtensionSetting][google.ads.googleads.v3.services.CustomerExtensionSettingService.GetCustomerExtensionSetting].
type GetCustomerExtensionSettingRequest struct {
	// Required. The resource name of the customer extension setting to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerExtensionSettingRequest) Reset()         { *m = GetCustomerExtensionSettingRequest{} }
func (m *GetCustomerExtensionSettingRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerExtensionSettingRequest) ProtoMessage()    {}
func (*GetCustomerExtensionSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ee6b2d7db0d86b, []int{0}
}

func (m *GetCustomerExtensionSettingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerExtensionSettingRequest.Unmarshal(m, b)
}
func (m *GetCustomerExtensionSettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerExtensionSettingRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerExtensionSettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerExtensionSettingRequest.Merge(m, src)
}
func (m *GetCustomerExtensionSettingRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerExtensionSettingRequest.Size(m)
}
func (m *GetCustomerExtensionSettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerExtensionSettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerExtensionSettingRequest proto.InternalMessageInfo

func (m *GetCustomerExtensionSettingRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CustomerExtensionSettingService.MutateCustomerExtensionSettings][google.ads.googleads.v3.services.CustomerExtensionSettingService.MutateCustomerExtensionSettings].
type MutateCustomerExtensionSettingsRequest struct {
	// Required. The ID of the customer whose customer extension settings are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual customer extension
	// settings.
	Operations []*CustomerExtensionSettingOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerExtensionSettingsRequest) Reset() {
	*m = MutateCustomerExtensionSettingsRequest{}
}
func (m *MutateCustomerExtensionSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerExtensionSettingsRequest) ProtoMessage()    {}
func (*MutateCustomerExtensionSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ee6b2d7db0d86b, []int{1}
}

func (m *MutateCustomerExtensionSettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerExtensionSettingsRequest.Unmarshal(m, b)
}
func (m *MutateCustomerExtensionSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerExtensionSettingsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCustomerExtensionSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerExtensionSettingsRequest.Merge(m, src)
}
func (m *MutateCustomerExtensionSettingsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerExtensionSettingsRequest.Size(m)
}
func (m *MutateCustomerExtensionSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerExtensionSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerExtensionSettingsRequest proto.InternalMessageInfo

func (m *MutateCustomerExtensionSettingsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerExtensionSettingsRequest) GetOperations() []*CustomerExtensionSettingOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCustomerExtensionSettingsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCustomerExtensionSettingsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a customer extension setting.
type CustomerExtensionSettingOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CustomerExtensionSettingOperation_Create
	//	*CustomerExtensionSettingOperation_Update
	//	*CustomerExtensionSettingOperation_Remove
	Operation            isCustomerExtensionSettingOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CustomerExtensionSettingOperation) Reset()         { *m = CustomerExtensionSettingOperation{} }
func (m *CustomerExtensionSettingOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerExtensionSettingOperation) ProtoMessage()    {}
func (*CustomerExtensionSettingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ee6b2d7db0d86b, []int{2}
}

func (m *CustomerExtensionSettingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerExtensionSettingOperation.Unmarshal(m, b)
}
func (m *CustomerExtensionSettingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerExtensionSettingOperation.Marshal(b, m, deterministic)
}
func (m *CustomerExtensionSettingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerExtensionSettingOperation.Merge(m, src)
}
func (m *CustomerExtensionSettingOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerExtensionSettingOperation.Size(m)
}
func (m *CustomerExtensionSettingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerExtensionSettingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerExtensionSettingOperation proto.InternalMessageInfo

func (m *CustomerExtensionSettingOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCustomerExtensionSettingOperation_Operation interface {
	isCustomerExtensionSettingOperation_Operation()
}

type CustomerExtensionSettingOperation_Create struct {
	Create *resources.CustomerExtensionSetting `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerExtensionSettingOperation_Update struct {
	Update *resources.CustomerExtensionSetting `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CustomerExtensionSettingOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CustomerExtensionSettingOperation_Create) isCustomerExtensionSettingOperation_Operation() {}

func (*CustomerExtensionSettingOperation_Update) isCustomerExtensionSettingOperation_Operation() {}

func (*CustomerExtensionSettingOperation_Remove) isCustomerExtensionSettingOperation_Operation() {}

func (m *CustomerExtensionSettingOperation) GetOperation() isCustomerExtensionSettingOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CustomerExtensionSettingOperation) GetCreate() *resources.CustomerExtensionSetting {
	if x, ok := m.GetOperation().(*CustomerExtensionSettingOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CustomerExtensionSettingOperation) GetUpdate() *resources.CustomerExtensionSetting {
	if x, ok := m.GetOperation().(*CustomerExtensionSettingOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CustomerExtensionSettingOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CustomerExtensionSettingOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomerExtensionSettingOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomerExtensionSettingOperation_Create)(nil),
		(*CustomerExtensionSettingOperation_Update)(nil),
		(*CustomerExtensionSettingOperation_Remove)(nil),
	}
}

// Response message for a customer extension setting mutate.
type MutateCustomerExtensionSettingsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCustomerExtensionSettingResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MutateCustomerExtensionSettingsResponse) Reset() {
	*m = MutateCustomerExtensionSettingsResponse{}
}
func (m *MutateCustomerExtensionSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerExtensionSettingsResponse) ProtoMessage()    {}
func (*MutateCustomerExtensionSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ee6b2d7db0d86b, []int{3}
}

func (m *MutateCustomerExtensionSettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerExtensionSettingsResponse.Unmarshal(m, b)
}
func (m *MutateCustomerExtensionSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerExtensionSettingsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCustomerExtensionSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerExtensionSettingsResponse.Merge(m, src)
}
func (m *MutateCustomerExtensionSettingsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerExtensionSettingsResponse.Size(m)
}
func (m *MutateCustomerExtensionSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerExtensionSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerExtensionSettingsResponse proto.InternalMessageInfo

func (m *MutateCustomerExtensionSettingsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCustomerExtensionSettingsResponse) GetResults() []*MutateCustomerExtensionSettingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the customer extension setting mutate.
type MutateCustomerExtensionSettingResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerExtensionSettingResult) Reset()         { *m = MutateCustomerExtensionSettingResult{} }
func (m *MutateCustomerExtensionSettingResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerExtensionSettingResult) ProtoMessage()    {}
func (*MutateCustomerExtensionSettingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ee6b2d7db0d86b, []int{4}
}

func (m *MutateCustomerExtensionSettingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerExtensionSettingResult.Unmarshal(m, b)
}
func (m *MutateCustomerExtensionSettingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerExtensionSettingResult.Marshal(b, m, deterministic)
}
func (m *MutateCustomerExtensionSettingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerExtensionSettingResult.Merge(m, src)
}
func (m *MutateCustomerExtensionSettingResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerExtensionSettingResult.Size(m)
}
func (m *MutateCustomerExtensionSettingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerExtensionSettingResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerExtensionSettingResult proto.InternalMessageInfo

func (m *MutateCustomerExtensionSettingResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerExtensionSettingRequest)(nil), "google.ads.googleads.v3.services.GetCustomerExtensionSettingRequest")
	proto.RegisterType((*MutateCustomerExtensionSettingsRequest)(nil), "google.ads.googleads.v3.services.MutateCustomerExtensionSettingsRequest")
	proto.RegisterType((*CustomerExtensionSettingOperation)(nil), "google.ads.googleads.v3.services.CustomerExtensionSettingOperation")
	proto.RegisterType((*MutateCustomerExtensionSettingsResponse)(nil), "google.ads.googleads.v3.services.MutateCustomerExtensionSettingsResponse")
	proto.RegisterType((*MutateCustomerExtensionSettingResult)(nil), "google.ads.googleads.v3.services.MutateCustomerExtensionSettingResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/customer_extension_setting_service.proto", fileDescriptor_56ee6b2d7db0d86b)
}

var fileDescriptor_56ee6b2d7db0d86b = []byte{
	// 798 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xbf, 0x6f, 0xf3, 0x44,
	0x18, 0xc6, 0xce, 0xa7, 0x42, 0x2f, 0xdf, 0x07, 0xd2, 0x21, 0x20, 0xe4, 0x43, 0x6a, 0x30, 0x11,
	0xad, 0x22, 0x64, 0x8b, 0x64, 0xc2, 0x51, 0x91, 0x9c, 0xd2, 0xb4, 0x15, 0x2a, 0xad, 0x1c, 0xd1,
	0x01, 0x45, 0x98, 0x8b, 0x7d, 0x4d, 0xad, 0xda, 0x3e, 0x73, 0x77, 0x8e, 0xa8, 0xaa, 0x2e, 0x08,
	0xb1, 0x30, 0xb2, 0x31, 0x32, 0xb2, 0xf3, 0x4f, 0x74, 0xed, 0xd6, 0xa9, 0x03, 0x13, 0x3b, 0x0b,
	0x13, 0xb2, 0xef, 0x2e, 0x3f, 0x0a, 0xae, 0x91, 0xda, 0xed, 0xf5, 0xbd, 0xcf, 0x3d, 0xef, 0xaf,
	0xe7, 0x5e, 0x19, 0x1c, 0x4c, 0x09, 0x99, 0x46, 0xd8, 0x42, 0x01, 0xb3, 0x84, 0x99, 0x5b, 0xb3,
	0x9e, 0xc5, 0x30, 0x9d, 0x85, 0x3e, 0x66, 0x96, 0x9f, 0x31, 0x4e, 0x62, 0x4c, 0x3d, 0xfc, 0x1d,
	0xc7, 0x09, 0x0b, 0x49, 0xe2, 0x31, 0xcc, 0x79, 0x98, 0x4c, 0x3d, 0x89, 0x31, 0x53, 0x4a, 0x38,
	0x81, 0x2d, 0x71, 0xdf, 0x44, 0x01, 0x33, 0xe7, 0x54, 0xe6, 0xac, 0x67, 0x2a, 0xaa, 0xe6, 0xa0,
	0x2c, 0x18, 0xc5, 0x8c, 0x64, 0xf4, 0xe1, 0x68, 0x22, 0x4a, 0xf3, 0x3d, 0xc5, 0x91, 0x86, 0x16,
	0x4a, 0x12, 0xc2, 0x11, 0x0f, 0x49, 0xc2, 0xa4, 0xf7, 0x9d, 0x25, 0xaf, 0x1f, 0x85, 0x38, 0xe1,
	0xd2, 0xb1, 0xb1, 0xe4, 0x38, 0x0d, 0x71, 0x14, 0x78, 0x13, 0x7c, 0x86, 0x66, 0x21, 0xa1, 0x12,
	0xf0, 0xee, 0x12, 0x40, 0xa5, 0x23, 0x5d, 0xb2, 0x30, 0xab, 0xf8, 0x9a, 0x64, 0xa7, 0x92, 0x20,
	0x46, 0xec, 0xfc, 0x5e, 0x58, 0x9a, 0xfa, 0x16, 0xe3, 0x88, 0x67, 0x32, 0x1f, 0xe3, 0x07, 0x0d,
	0x18, 0x7b, 0x98, 0xef, 0xc8, 0xaa, 0x76, 0x55, 0x51, 0x23, 0x51, 0x93, 0x8b, 0xbf, 0xcd, 0x30,
	0xe3, 0xf0, 0x6b, 0xf0, 0x42, 0xc5, 0xf4, 0x12, 0x14, 0xe3, 0x86, 0xd6, 0xd2, 0xb6, 0xd6, 0x07,
	0x9f, 0xdc, 0x39, 0xfa, 0xdf, 0x4e, 0x0f, 0x7c, 0xbc, 0x68, 0xa7, 0xb4, 0xd2, 0x90, 0x99, 0x3e,
	0x89, 0xad, 0x52, 0xe2, 0xe7, 0x8a, 0xef, 0x0b, 0x14, 0x63, 0xe3, 0x47, 0x1d, 0x7c, 0x78, 0x98,
	0x71, 0xc4, 0x71, 0xd9, 0x05, 0xa6, 0x52, 0x69, 0x83, 0xfa, 0x7c, 0x06, 0x61, 0x20, 0x13, 0xa9,
	0xdd, 0x39, 0xba, 0x0b, 0xd4, 0xf9, 0x41, 0x00, 0xcf, 0x00, 0x20, 0x29, 0xa6, 0xa2, 0xf7, 0x0d,
	0xbd, 0x55, 0xdb, 0xaa, 0x77, 0x77, 0xcc, 0x2a, 0x01, 0x98, 0x65, 0xd1, 0x8f, 0x14, 0x97, 0x8c,
	0xb4, 0xe0, 0x86, 0x9b, 0xe0, 0x8d, 0x14, 0x51, 0x1e, 0xa2, 0xc8, 0x3b, 0x45, 0x61, 0x94, 0x51,
	0xdc, 0xa8, 0xb5, 0xb4, 0xad, 0xd7, 0xdc, 0xd7, 0xe5, 0xf1, 0x50, 0x9c, 0xc2, 0x0f, 0xc0, 0x8b,
	0x19, 0x8a, 0xc2, 0x00, 0x71, 0xec, 0x91, 0x24, 0xba, 0x68, 0x3c, 0x2b, 0x60, 0xcf, 0xd5, 0xe1,
	0x51, 0x12, 0x5d, 0x18, 0xbf, 0xeb, 0xe0, 0xfd, 0xca, 0x24, 0x60, 0x1f, 0xd4, 0xb3, 0xb4, 0x20,
	0xca, 0x67, 0x5c, 0x10, 0xd5, 0xbb, 0x4d, 0x55, 0x9e, 0x92, 0x81, 0x39, 0xcc, 0x65, 0x70, 0x88,
	0xd8, 0xb9, 0x0b, 0x04, 0x3c, 0xb7, 0xe1, 0x97, 0x60, 0xcd, 0xa7, 0x18, 0x71, 0x31, 0xc4, 0x7a,
	0xb7, 0x5f, 0xda, 0x96, 0xb9, 0xea, 0x4b, 0xfb, 0xb2, 0xff, 0x8a, 0x2b, 0xc9, 0x72, 0x5a, 0x11,
	0xa4, 0xa1, 0x3f, 0x09, 0xad, 0x20, 0x83, 0x0d, 0xb0, 0x46, 0x71, 0x4c, 0x66, 0xa2, 0xab, 0xeb,
	0xb9, 0x47, 0x7c, 0x0f, 0xea, 0x60, 0x7d, 0x3e, 0x06, 0xe3, 0x46, 0x03, 0x9b, 0x95, 0x02, 0x62,
	0x29, 0x49, 0x18, 0x86, 0x43, 0xf0, 0xd6, 0xbd, 0x89, 0x79, 0x98, 0x52, 0x42, 0x8b, 0x08, 0xf5,
	0x2e, 0x54, 0x89, 0xd3, 0xd4, 0x37, 0x47, 0xc5, 0x63, 0x71, 0xdf, 0x5c, 0x9d, 0xe5, 0x6e, 0x0e,
	0x87, 0xdf, 0x80, 0x57, 0x29, 0x66, 0x59, 0xc4, 0x95, 0xc0, 0x86, 0xd5, 0x02, 0x7b, 0x38, 0x47,
	0xb7, 0xa0, 0x73, 0x15, 0xad, 0xf1, 0x39, 0x68, 0xff, 0x9f, 0x0b, 0xb9, 0xb4, 0xfe, 0xe3, 0x79,
	0xae, 0xbe, 0xb1, 0xee, 0xcd, 0x33, 0xb0, 0x51, 0xc6, 0x33, 0x12, 0xf9, 0xc1, 0xbf, 0x34, 0xf0,
	0xf2, 0x81, 0x75, 0x00, 0x3f, 0xab, 0xae, 0xb0, 0x7a, 0x9b, 0x34, 0x1f, 0x23, 0x0d, 0x63, 0x74,
	0xeb, 0xac, 0x16, 0xfb, 0xfd, 0xcd, 0x1f, 0x3f, 0xeb, 0xdb, 0xb0, 0x9f, 0xef, 0xe9, 0xcb, 0x15,
	0xcf, 0xb6, 0x5a, 0x08, 0xcc, 0xea, 0xcc, 0x17, 0xf7, 0xbf, 0x74, 0x61, 0x75, 0xae, 0xe0, 0x2f,
	0x3a, 0xd8, 0xa8, 0x90, 0x0f, 0xdc, 0x7f, 0xec, 0x74, 0xd5, 0x0a, 0x6b, 0x1e, 0x3c, 0x01, 0x93,
	0xd0, 0xb2, 0x31, 0xb9, 0x75, 0xde, 0x5e, 0x5a, 0x87, 0x1f, 0x2d, 0x16, 0x53, 0xd1, 0x96, 0x1d,
	0xe3, 0xd3, 0xbc, 0x2d, 0x8b, 0x3e, 0x5c, 0x2e, 0x81, 0xb7, 0x3b, 0x57, 0xe5, 0x5d, 0xb1, 0xe3,
	0x22, 0x03, 0x5b, 0xeb, 0x34, 0x5f, 0x5e, 0x3b, 0x8d, 0xb2, 0x05, 0x3f, 0xf8, 0x49, 0x07, 0x6d,
	0x9f, 0xc4, 0x95, 0x15, 0x0d, 0xda, 0x15, 0xda, 0x3b, 0xce, 0xb7, 0xd6, 0xb1, 0xf6, 0xd5, 0xbe,
	0x64, 0x9a, 0x92, 0x08, 0x25, 0x53, 0x93, 0xd0, 0xa9, 0x35, 0xc5, 0x49, 0xb1, 0xd3, 0xac, 0x45,
	0xec, 0xf2, 0xff, 0x81, 0xbe, 0x32, 0x7e, 0xd5, 0x6b, 0x7b, 0x8e, 0xf3, 0x9b, 0xde, 0xda, 0x13,
	0x84, 0x4e, 0xc0, 0x4c, 0x61, 0xe6, 0xd6, 0x49, 0xcf, 0x94, 0x81, 0xd9, 0xb5, 0x82, 0x8c, 0x9d,
	0x80, 0x8d, 0xe7, 0x90, 0xf1, 0x49, 0x6f, 0xac, 0x20, 0x7f, 0xea, 0x6d, 0x71, 0x6e, 0xdb, 0x4e,
	0xc0, 0x6c, 0x7b, 0x0e, 0xb2, 0xed, 0x93, 0x9e, 0x6d, 0x2b, 0xd8, 0x64, 0xad, 0xc8, 0xb3, 0xf7,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x55, 0xa8, 0xc7, 0xb6, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerExtensionSettingServiceClient is the client API for CustomerExtensionSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerExtensionSettingServiceClient interface {
	// Returns the requested customer extension setting in full detail.
	GetCustomerExtensionSetting(ctx context.Context, in *GetCustomerExtensionSettingRequest, opts ...grpc.CallOption) (*resources.CustomerExtensionSetting, error)
	// Creates, updates, or removes customer extension settings. Operation
	// statuses are returned.
	MutateCustomerExtensionSettings(ctx context.Context, in *MutateCustomerExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCustomerExtensionSettingsResponse, error)
}

type customerExtensionSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerExtensionSettingServiceClient(cc grpc.ClientConnInterface) CustomerExtensionSettingServiceClient {
	return &customerExtensionSettingServiceClient{cc}
}

func (c *customerExtensionSettingServiceClient) GetCustomerExtensionSetting(ctx context.Context, in *GetCustomerExtensionSettingRequest, opts ...grpc.CallOption) (*resources.CustomerExtensionSetting, error) {
	out := new(resources.CustomerExtensionSetting)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerExtensionSettingService/GetCustomerExtensionSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerExtensionSettingServiceClient) MutateCustomerExtensionSettings(ctx context.Context, in *MutateCustomerExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCustomerExtensionSettingsResponse, error) {
	out := new(MutateCustomerExtensionSettingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerExtensionSettingService/MutateCustomerExtensionSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerExtensionSettingServiceServer is the server API for CustomerExtensionSettingService service.
type CustomerExtensionSettingServiceServer interface {
	// Returns the requested customer extension setting in full detail.
	GetCustomerExtensionSetting(context.Context, *GetCustomerExtensionSettingRequest) (*resources.CustomerExtensionSetting, error)
	// Creates, updates, or removes customer extension settings. Operation
	// statuses are returned.
	MutateCustomerExtensionSettings(context.Context, *MutateCustomerExtensionSettingsRequest) (*MutateCustomerExtensionSettingsResponse, error)
}

// UnimplementedCustomerExtensionSettingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerExtensionSettingServiceServer struct {
}

func (*UnimplementedCustomerExtensionSettingServiceServer) GetCustomerExtensionSetting(ctx context.Context, req *GetCustomerExtensionSettingRequest) (*resources.CustomerExtensionSetting, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCustomerExtensionSetting not implemented")
}
func (*UnimplementedCustomerExtensionSettingServiceServer) MutateCustomerExtensionSettings(ctx context.Context, req *MutateCustomerExtensionSettingsRequest) (*MutateCustomerExtensionSettingsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCustomerExtensionSettings not implemented")
}

func RegisterCustomerExtensionSettingServiceServer(s *grpc.Server, srv CustomerExtensionSettingServiceServer) {
	s.RegisterService(&_CustomerExtensionSettingService_serviceDesc, srv)
}

func _CustomerExtensionSettingService_GetCustomerExtensionSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerExtensionSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerExtensionSettingServiceServer).GetCustomerExtensionSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerExtensionSettingService/GetCustomerExtensionSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerExtensionSettingServiceServer).GetCustomerExtensionSetting(ctx, req.(*GetCustomerExtensionSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerExtensionSettingService_MutateCustomerExtensionSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerExtensionSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerExtensionSettingServiceServer).MutateCustomerExtensionSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerExtensionSettingService/MutateCustomerExtensionSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerExtensionSettingServiceServer).MutateCustomerExtensionSettings(ctx, req.(*MutateCustomerExtensionSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerExtensionSettingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CustomerExtensionSettingService",
	HandlerType: (*CustomerExtensionSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerExtensionSetting",
			Handler:    _CustomerExtensionSettingService_GetCustomerExtensionSetting_Handler,
		},
		{
			MethodName: "MutateCustomerExtensionSettings",
			Handler:    _CustomerExtensionSettingService_MutateCustomerExtensionSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/customer_extension_setting_service.proto",
}
