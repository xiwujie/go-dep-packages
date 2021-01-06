// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/company.proto

package talent

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

// A Company resource represents a company in the service. A company is the
// entity that owns job postings, that is, the hiring entity responsible for
// employing applicants for the job position.
type Company struct {
	// Required during company update.
	//
	// The resource name for a company. This is generated by the service when a
	// company is created.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/foo/tenants/bar/companies/baz".
	//
	// If tenant id is unspecified, the default tenant is used. For
	// example, "projects/foo/companies/bar".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the company, for example, "Google LLC".
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. Client side company identifier, used to uniquely identify the
	// company.
	//
	// The maximum number of allowed characters is 255.
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// The employer's company size.
	Size CompanySize `protobuf:"varint,4,opt,name=size,proto3,enum=google.cloud.talent.v4beta1.CompanySize" json:"size,omitempty"`
	// The street address of the company's main headquarters, which may be
	// different from the job location. The service attempts
	// to geolocate the provided address, and populates a more specific
	// location wherever possible in [DerivedInfo.headquarters_location][google.cloud.talent.v4beta1.Company.DerivedInfo.headquarters_location].
	HeadquartersAddress string `protobuf:"bytes,5,opt,name=headquarters_address,json=headquartersAddress,proto3" json:"headquarters_address,omitempty"`
	// Set to true if it is the hiring agency that post jobs for other
	// employers.
	//
	// Defaults to false if not provided.
	HiringAgency bool `protobuf:"varint,6,opt,name=hiring_agency,json=hiringAgency,proto3" json:"hiring_agency,omitempty"`
	// Equal Employment Opportunity legal disclaimer text to be
	// associated with all jobs, and typically to be displayed in all
	// roles.
	//
	// The maximum number of allowed characters is 500.
	EeoText string `protobuf:"bytes,7,opt,name=eeo_text,json=eeoText,proto3" json:"eeo_text,omitempty"`
	// The URI representing the company's primary web site or home page,
	// for example, "https://www.google.com".
	//
	// The maximum number of allowed characters is 255.
	WebsiteUri string `protobuf:"bytes,8,opt,name=website_uri,json=websiteUri,proto3" json:"website_uri,omitempty"`
	// The URI to employer's career site or careers page on the employer's web
	// site, for example, "https://careers.google.com".
	CareerSiteUri string `protobuf:"bytes,9,opt,name=career_site_uri,json=careerSiteUri,proto3" json:"career_site_uri,omitempty"`
	// A URI that hosts the employer's company logo.
	ImageUri string `protobuf:"bytes,10,opt,name=image_uri,json=imageUri,proto3" json:"image_uri,omitempty"`
	// A list of keys of filterable [Job.custom_attributes][google.cloud.talent.v4beta1.Job.custom_attributes], whose
	// corresponding `string_values` are used in keyword searches. Jobs with
	// `string_values` under these specified field keys are returned if any
	// of the values match the search keyword. Custom field values with
	// parenthesis, brackets and special symbols are not searchable as-is,
	// and those keyword queries must be surrounded by quotes.
	KeywordSearchableJobCustomAttributes []string `protobuf:"bytes,11,rep,name=keyword_searchable_job_custom_attributes,json=keywordSearchableJobCustomAttributes,proto3" json:"keyword_searchable_job_custom_attributes,omitempty"`
	// Output only. Derived details about the company.
	DerivedInfo *Company_DerivedInfo `protobuf:"bytes,12,opt,name=derived_info,json=derivedInfo,proto3" json:"derived_info,omitempty"`
	// Output only. Indicates whether a company is flagged to be suspended from
	// public availability by the service when job content appears suspicious,
	// abusive, or spammy.
	Suspended            bool     `protobuf:"varint,13,opt,name=suspended,proto3" json:"suspended,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Company) Reset()         { *m = Company{} }
func (m *Company) String() string { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()    {}
func (*Company) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbff0b9f9d15156a, []int{0}
}

func (m *Company) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company.Unmarshal(m, b)
}
func (m *Company) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company.Marshal(b, m, deterministic)
}
func (m *Company) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company.Merge(m, src)
}
func (m *Company) XXX_Size() int {
	return xxx_messageInfo_Company.Size(m)
}
func (m *Company) XXX_DiscardUnknown() {
	xxx_messageInfo_Company.DiscardUnknown(m)
}

var xxx_messageInfo_Company proto.InternalMessageInfo

func (m *Company) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Company) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Company) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *Company) GetSize() CompanySize {
	if m != nil {
		return m.Size
	}
	return CompanySize_COMPANY_SIZE_UNSPECIFIED
}

func (m *Company) GetHeadquartersAddress() string {
	if m != nil {
		return m.HeadquartersAddress
	}
	return ""
}

func (m *Company) GetHiringAgency() bool {
	if m != nil {
		return m.HiringAgency
	}
	return false
}

func (m *Company) GetEeoText() string {
	if m != nil {
		return m.EeoText
	}
	return ""
}

func (m *Company) GetWebsiteUri() string {
	if m != nil {
		return m.WebsiteUri
	}
	return ""
}

func (m *Company) GetCareerSiteUri() string {
	if m != nil {
		return m.CareerSiteUri
	}
	return ""
}

func (m *Company) GetImageUri() string {
	if m != nil {
		return m.ImageUri
	}
	return ""
}

func (m *Company) GetKeywordSearchableJobCustomAttributes() []string {
	if m != nil {
		return m.KeywordSearchableJobCustomAttributes
	}
	return nil
}

func (m *Company) GetDerivedInfo() *Company_DerivedInfo {
	if m != nil {
		return m.DerivedInfo
	}
	return nil
}

func (m *Company) GetSuspended() bool {
	if m != nil {
		return m.Suspended
	}
	return false
}

// Derived details about the company.
type Company_DerivedInfo struct {
	// A structured headquarters location of the company, resolved from
	// [Company.headquarters_address][google.cloud.talent.v4beta1.Company.headquarters_address] if provided.
	HeadquartersLocation *Location `protobuf:"bytes,1,opt,name=headquarters_location,json=headquartersLocation,proto3" json:"headquarters_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Company_DerivedInfo) Reset()         { *m = Company_DerivedInfo{} }
func (m *Company_DerivedInfo) String() string { return proto.CompactTextString(m) }
func (*Company_DerivedInfo) ProtoMessage()    {}
func (*Company_DerivedInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbff0b9f9d15156a, []int{0, 0}
}

func (m *Company_DerivedInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company_DerivedInfo.Unmarshal(m, b)
}
func (m *Company_DerivedInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company_DerivedInfo.Marshal(b, m, deterministic)
}
func (m *Company_DerivedInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company_DerivedInfo.Merge(m, src)
}
func (m *Company_DerivedInfo) XXX_Size() int {
	return xxx_messageInfo_Company_DerivedInfo.Size(m)
}
func (m *Company_DerivedInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Company_DerivedInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Company_DerivedInfo proto.InternalMessageInfo

func (m *Company_DerivedInfo) GetHeadquartersLocation() *Location {
	if m != nil {
		return m.HeadquartersLocation
	}
	return nil
}

func init() {
	proto.RegisterType((*Company)(nil), "google.cloud.talent.v4beta1.Company")
	proto.RegisterType((*Company_DerivedInfo)(nil), "google.cloud.talent.v4beta1.Company.DerivedInfo")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/company.proto", fileDescriptor_bbff0b9f9d15156a)
}

var fileDescriptor_bbff0b9f9d15156a = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x4f, 0xd4, 0x4a,
	0x14, 0xc7, 0x53, 0x96, 0x0b, 0xbb, 0xd3, 0xe5, 0xde, 0x64, 0x2e, 0x37, 0x29, 0x70, 0x13, 0x56,
	0x45, 0x52, 0x5f, 0x5a, 0x41, 0x13, 0x13, 0xf5, 0x65, 0xc1, 0x17, 0x8c, 0x31, 0xa4, 0x8b, 0xc6,
	0xf0, 0xd2, 0x4c, 0x3b, 0x87, 0xee, 0x60, 0x3b, 0x53, 0x67, 0xa6, 0xc0, 0x42, 0x30, 0xd1, 0x8f,
	0xe2, 0x87, 0xf1, 0x73, 0xf8, 0xec, 0xa7, 0x30, 0x9d, 0x19, 0xd8, 0x25, 0x21, 0xe8, 0xdb, 0xf4,
	0xff, 0xff, 0x9d, 0x93, 0xd3, 0x99, 0xf3, 0x47, 0x8f, 0x0a, 0x21, 0x8a, 0x12, 0xe2, 0xbc, 0x14,
	0x0d, 0x8d, 0x35, 0x29, 0x81, 0xeb, 0xf8, 0xe4, 0x69, 0x06, 0x9a, 0x6c, 0xc5, 0xb9, 0xa8, 0x6a,
	0xc2, 0x27, 0x51, 0x2d, 0x85, 0x16, 0x78, 0xcd, 0xa2, 0x91, 0x41, 0x23, 0x8b, 0x46, 0x0e, 0x5d,
	0x5d, 0x77, 0x7d, 0x48, 0xcd, 0xe2, 0x23, 0x06, 0x25, 0x4d, 0x33, 0x18, 0x93, 0x13, 0x26, 0xa4,
	0xad, 0x5e, 0x5d, 0x99, 0x01, 0x24, 0x28, 0xd1, 0xc8, 0x1c, 0x9c, 0x15, 0xfe, 0x66, 0x86, 0x4a,
	0x70, 0x47, 0xfe, 0x3f, 0xd3, 0x84, 0x70, 0x2e, 0x34, 0xd1, 0x4c, 0x70, 0x65, 0xdd, 0xfb, 0xdf,
	0x17, 0xd0, 0xe2, 0xae, 0x1d, 0x19, 0x63, 0x34, 0xcf, 0x49, 0x05, 0x81, 0x37, 0xf0, 0xc2, 0x5e,
	0x62, 0xce, 0x78, 0x13, 0xf5, 0x29, 0x53, 0x75, 0x49, 0x26, 0xa9, 0xf1, 0xe6, 0x5a, 0x6f, 0xa7,
	0xf3, 0x63, 0x38, 0x97, 0xf8, 0xce, 0x78, 0xdb, 0x72, 0x1b, 0xc8, 0x87, 0x33, 0x0d, 0x92, 0x93,
	0x32, 0x65, 0x34, 0xe8, 0x4c, 0x31, 0x74, 0xa5, 0xef, 0x51, 0xfc, 0x12, 0xcd, 0x2b, 0x76, 0x0e,
	0xc1, 0xfc, 0xc0, 0x0b, 0xff, 0xde, 0x0e, 0xa3, 0x3b, 0x6e, 0x27, 0x72, 0x53, 0x8d, 0xd8, 0x39,
	0x24, 0xa6, 0x0a, 0x6f, 0xa1, 0xe5, 0x31, 0x10, 0xfa, 0xa9, 0x21, 0x52, 0x83, 0x54, 0x29, 0xa1,
	0x54, 0x82, 0x52, 0xc1, 0x5f, 0x66, 0xde, 0x7f, 0x67, 0xbd, 0xa1, 0xb5, 0xf0, 0x03, 0xb4, 0x34,
	0x66, 0x92, 0xf1, 0x22, 0x25, 0x05, 0xf0, 0x7c, 0x12, 0x2c, 0x0c, 0xbc, 0xb0, 0x9b, 0xf4, 0xad,
	0x38, 0x34, 0x1a, 0x5e, 0x41, 0x5d, 0x00, 0x91, 0x6a, 0x38, 0xd3, 0xc1, 0xa2, 0xe9, 0xb5, 0x08,
	0x20, 0x0e, 0xe0, 0x4c, 0xe3, 0x75, 0xe4, 0x9f, 0x42, 0xa6, 0x98, 0x86, 0xb4, 0x91, 0x2c, 0xe8,
	0x1a, 0x17, 0x39, 0xe9, 0x9d, 0x64, 0x78, 0x13, 0xfd, 0x93, 0x13, 0x09, 0x20, 0xd3, 0x6b, 0xa8,
	0x67, 0xa0, 0x25, 0x2b, 0x8f, 0x1c, 0xb7, 0x86, 0x7a, 0xac, 0x22, 0x85, 0x25, 0x90, 0x21, 0xba,
	0x46, 0x68, 0xcd, 0xf7, 0x28, 0xfc, 0x08, 0x93, 0x53, 0x21, 0x69, 0xaa, 0x80, 0xc8, 0x7c, 0x4c,
	0xb2, 0x12, 0xd2, 0x63, 0x91, 0xa5, 0x79, 0xa3, 0xb4, 0xa8, 0x52, 0xa2, 0xb5, 0x64, 0x59, 0xa3,
	0x41, 0x05, 0xfe, 0xa0, 0x13, 0xf6, 0x92, 0x0d, 0xc7, 0x8f, 0xae, 0xf1, 0xd7, 0x22, 0xdb, 0x35,
	0xf0, 0xf0, 0x9a, 0xc5, 0x1f, 0x50, 0x9f, 0x82, 0x64, 0x27, 0x40, 0x53, 0xc6, 0x8f, 0x44, 0xd0,
	0x1f, 0x78, 0xa1, 0xbf, 0xfd, 0xf8, 0x4f, 0xae, 0x3d, 0x7a, 0x65, 0x0b, 0xf7, 0xf8, 0x91, 0x68,
	0xdf, 0xb1, 0x93, 0xf8, 0x74, 0xaa, 0xe0, 0x7b, 0xa8, 0xa7, 0x1a, 0x55, 0x03, 0xa7, 0x40, 0x83,
	0xa5, 0xf6, 0x4e, 0x2d, 0x34, 0x55, 0x57, 0x19, 0xf2, 0x67, 0x7a, 0xe0, 0x43, 0xf4, 0xdf, 0x8d,
	0xc7, 0x2b, 0x45, 0x6e, 0x16, 0xd1, 0x6c, 0x9b, 0xbf, 0xfd, 0xf0, 0xce, 0xa1, 0xde, 0x38, 0x38,
	0xb9, 0xb1, 0x00, 0x57, 0xea, 0xf3, 0x2f, 0xde, 0xcf, 0xe1, 0x67, 0xb4, 0x76, 0x2c, 0x32, 0xe5,
	0xda, 0x90, 0x9a, 0xa9, 0x28, 0x17, 0x55, 0x7c, 0xb5, 0xdb, 0xcf, 0x6a, 0x29, 0x8e, 0x21, 0xd7,
	0x2a, 0xbe, 0x70, 0xa7, 0xcb, 0x58, 0x03, 0x27, 0xbc, 0x55, 0xec, 0xe1, 0xd2, 0x65, 0x97, 0x81,
	0x8a, 0x2f, 0x5c, 0x8c, 0x2f, 0xf1, 0xe6, 0x2d, 0x85, 0xb7, 0x70, 0x3b, 0x5f, 0x3d, 0xb4, 0x9e,
	0x8b, 0xea, 0xae, 0xdf, 0xd8, 0x59, 0x76, 0xd3, 0x24, 0x2e, 0xcb, 0xfb, 0x6d, 0x04, 0xf7, 0xbd,
	0xc3, 0xa1, 0x2b, 0x2a, 0x44, 0x49, 0x78, 0x11, 0x09, 0x59, 0xc4, 0x05, 0x70, 0x13, 0xd0, 0x78,
	0xfa, 0x3f, 0xb7, 0x66, 0xfd, 0x85, 0xfd, 0xfc, 0x36, 0xd7, 0xd9, 0x3d, 0x18, 0x65, 0x0b, 0xa6,
	0xe6, 0xc9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x85, 0x89, 0x20, 0xa2, 0x04, 0x00, 0x00,
}
