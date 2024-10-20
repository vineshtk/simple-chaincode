// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msp/msp_principal.proto

package msp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type MSPPrincipal_Classification int32

const (
	MSPPrincipal_ROLE MSPPrincipal_Classification = 0
	// one of a member of MSP network, and the one of an
	// administrator of an MSP network
	MSPPrincipal_ORGANIZATION_UNIT MSPPrincipal_Classification = 1
	// groupping of entities, per MSP affiliation
	// E.g., this can well be represented by an MSP's
	// Organization unit
	MSPPrincipal_IDENTITY MSPPrincipal_Classification = 2
	// identity
	MSPPrincipal_ANONYMITY MSPPrincipal_Classification = 3
	// an identity to be anonymous or nominal.
	MSPPrincipal_COMBINED MSPPrincipal_Classification = 4
)

var MSPPrincipal_Classification_name = map[int32]string{
	0: "ROLE",
	1: "ORGANIZATION_UNIT",
	2: "IDENTITY",
	3: "ANONYMITY",
	4: "COMBINED",
}

var MSPPrincipal_Classification_value = map[string]int32{
	"ROLE":              0,
	"ORGANIZATION_UNIT": 1,
	"IDENTITY":          2,
	"ANONYMITY":         3,
	"COMBINED":          4,
}

func (x MSPPrincipal_Classification) String() string {
	return proto.EnumName(MSPPrincipal_Classification_name, int32(x))
}

func (MSPPrincipal_Classification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_82e08b7ead29bd48, []int{0, 0}
}

type MSPRole_MSPRoleType int32

const (
	MSPRole_MEMBER  MSPRole_MSPRoleType = 0
	MSPRole_ADMIN   MSPRole_MSPRoleType = 1
	MSPRole_CLIENT  MSPRole_MSPRoleType = 2
	MSPRole_PEER    MSPRole_MSPRoleType = 3
	MSPRole_ORDERER MSPRole_MSPRoleType = 4
)

var MSPRole_MSPRoleType_name = map[int32]string{
	0: "MEMBER",
	1: "ADMIN",
	2: "CLIENT",
	3: "PEER",
	4: "ORDERER",
}

var MSPRole_MSPRoleType_value = map[string]int32{
	"MEMBER":  0,
	"ADMIN":   1,
	"CLIENT":  2,
	"PEER":    3,
	"ORDERER": 4,
}

func (x MSPRole_MSPRoleType) String() string {
	return proto.EnumName(MSPRole_MSPRoleType_name, int32(x))
}

func (MSPRole_MSPRoleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_82e08b7ead29bd48, []int{2, 0}
}

type MSPIdentityAnonymity_MSPIdentityAnonymityType int32

const (
	MSPIdentityAnonymity_NOMINAL   MSPIdentityAnonymity_MSPIdentityAnonymityType = 0
	MSPIdentityAnonymity_ANONYMOUS MSPIdentityAnonymity_MSPIdentityAnonymityType = 1
)

var MSPIdentityAnonymity_MSPIdentityAnonymityType_name = map[int32]string{
	0: "NOMINAL",
	1: "ANONYMOUS",
}

var MSPIdentityAnonymity_MSPIdentityAnonymityType_value = map[string]int32{
	"NOMINAL":   0,
	"ANONYMOUS": 1,
}

func (x MSPIdentityAnonymity_MSPIdentityAnonymityType) String() string {
	return proto.EnumName(MSPIdentityAnonymity_MSPIdentityAnonymityType_name, int32(x))
}

func (MSPIdentityAnonymity_MSPIdentityAnonymityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_82e08b7ead29bd48, []int{3, 0}
}

// MSPPrincipal aims to represent an MSP-centric set of identities.
// In particular, this structure allows for definition of
//   - a group of identities that are member of the same MSP
//   - a group of identities that are member of the same organization unit
//     in the same MSP
//   - a group of identities that are administering a specific MSP
//   - a specific identity
//
// Expressing these groups is done given two fields of the fields below
//   - Classification, that defines the type of classification of identities
//     in an MSP this principal would be defined on; Classification can take
//     three values:
//     (i)  ByMSPRole: that represents a classification of identities within
//     MSP based on one of the two pre-defined MSP rules, "member" and "admin"
//     (ii) ByOrganizationUnit: that represents a classification of identities
//     within MSP based on the organization unit an identity belongs to
//     (iii)ByIdentity that denotes that MSPPrincipal is mapped to a single
//     identity/certificate; this would mean that the Principal bytes
//     message
type MSPPrincipal struct {
	// Classification describes the way that one should process
	// Principal. An Classification value of "ByOrganizationUnit" reflects
	// that "Principal" contains the name of an organization this MSP
	// handles. A Classification value "ByIdentity" means that
	// "Principal" contains a specific identity. Default value
	// denotes that Principal contains one of the groups by
	// default supported by all MSPs ("admin" or "member").
	PrincipalClassification MSPPrincipal_Classification `protobuf:"varint,1,opt,name=principal_classification,json=principalClassification,proto3,enum=common.MSPPrincipal_Classification" json:"principal_classification,omitempty"`
	// Principal completes the policy principal definition. For the default
	// principal types, Principal can be either "Admin" or "Member".
	// For the ByOrganizationUnit/ByIdentity values of Classification,
	// PolicyPrincipal acquires its value from an organization unit or
	// identity, respectively.
	// For the Combined Classification type, the Principal is a marshalled
	// CombinedPrincipal.
	Principal            []byte   `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MSPPrincipal) Reset()         { *m = MSPPrincipal{} }
func (m *MSPPrincipal) String() string { return proto.CompactTextString(m) }
func (*MSPPrincipal) ProtoMessage()    {}
func (*MSPPrincipal) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e08b7ead29bd48, []int{0}
}

func (m *MSPPrincipal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MSPPrincipal.Unmarshal(m, b)
}
func (m *MSPPrincipal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MSPPrincipal.Marshal(b, m, deterministic)
}
func (m *MSPPrincipal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MSPPrincipal.Merge(m, src)
}
func (m *MSPPrincipal) XXX_Size() int {
	return xxx_messageInfo_MSPPrincipal.Size(m)
}
func (m *MSPPrincipal) XXX_DiscardUnknown() {
	xxx_messageInfo_MSPPrincipal.DiscardUnknown(m)
}

var xxx_messageInfo_MSPPrincipal proto.InternalMessageInfo

func (m *MSPPrincipal) GetPrincipalClassification() MSPPrincipal_Classification {
	if m != nil {
		return m.PrincipalClassification
	}
	return MSPPrincipal_ROLE
}

func (m *MSPPrincipal) GetPrincipal() []byte {
	if m != nil {
		return m.Principal
	}
	return nil
}

// OrganizationUnit governs the organization of the Principal
// field of a policy principal when a specific organization unity members
// are to be defined within a policy principal.
type OrganizationUnit struct {
	// MSPIdentifier represents the identifier of the MSP this organization unit
	// refers to
	MspIdentifier string `protobuf:"bytes,1,opt,name=msp_identifier,json=mspIdentifier,proto3" json:"msp_identifier,omitempty"`
	// OrganizationUnitIdentifier defines the organizational unit under the
	// MSP identified with MSPIdentifier
	OrganizationalUnitIdentifier string `protobuf:"bytes,2,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier,proto3" json:"organizational_unit_identifier,omitempty"`
	// CertifiersIdentifier is the hash of certificates chain of trust
	// related to this organizational unit
	CertifiersIdentifier []byte   `protobuf:"bytes,3,opt,name=certifiers_identifier,json=certifiersIdentifier,proto3" json:"certifiers_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrganizationUnit) Reset()         { *m = OrganizationUnit{} }
func (m *OrganizationUnit) String() string { return proto.CompactTextString(m) }
func (*OrganizationUnit) ProtoMessage()    {}
func (*OrganizationUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e08b7ead29bd48, []int{1}
}

func (m *OrganizationUnit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationUnit.Unmarshal(m, b)
}
func (m *OrganizationUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationUnit.Marshal(b, m, deterministic)
}
func (m *OrganizationUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationUnit.Merge(m, src)
}
func (m *OrganizationUnit) XXX_Size() int {
	return xxx_messageInfo_OrganizationUnit.Size(m)
}
func (m *OrganizationUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationUnit.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationUnit proto.InternalMessageInfo

func (m *OrganizationUnit) GetMspIdentifier() string {
	if m != nil {
		return m.MspIdentifier
	}
	return ""
}

func (m *OrganizationUnit) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func (m *OrganizationUnit) GetCertifiersIdentifier() []byte {
	if m != nil {
		return m.CertifiersIdentifier
	}
	return nil
}

// MSPRole governs the organization of the Principal
// field of an MSPPrincipal when it aims to define one of the
// two dedicated roles within an MSP: Admin and Members.
type MSPRole struct {
	// MSPIdentifier represents the identifier of the MSP this principal
	// refers to
	MspIdentifier string `protobuf:"bytes,1,opt,name=msp_identifier,json=mspIdentifier,proto3" json:"msp_identifier,omitempty"`
	// MSPRoleType defines which of the available, pre-defined MSP-roles
	// an identiy should posess inside the MSP with identifier MSPidentifier
	Role                 MSPRole_MSPRoleType `protobuf:"varint,2,opt,name=role,proto3,enum=common.MSPRole_MSPRoleType" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MSPRole) Reset()         { *m = MSPRole{} }
func (m *MSPRole) String() string { return proto.CompactTextString(m) }
func (*MSPRole) ProtoMessage()    {}
func (*MSPRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e08b7ead29bd48, []int{2}
}

func (m *MSPRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MSPRole.Unmarshal(m, b)
}
func (m *MSPRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MSPRole.Marshal(b, m, deterministic)
}
func (m *MSPRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MSPRole.Merge(m, src)
}
func (m *MSPRole) XXX_Size() int {
	return xxx_messageInfo_MSPRole.Size(m)
}
func (m *MSPRole) XXX_DiscardUnknown() {
	xxx_messageInfo_MSPRole.DiscardUnknown(m)
}

var xxx_messageInfo_MSPRole proto.InternalMessageInfo

func (m *MSPRole) GetMspIdentifier() string {
	if m != nil {
		return m.MspIdentifier
	}
	return ""
}

func (m *MSPRole) GetRole() MSPRole_MSPRoleType {
	if m != nil {
		return m.Role
	}
	return MSPRole_MEMBER
}

// MSPIdentityAnonymity can be used to enforce an identity to be anonymous or nominal.
type MSPIdentityAnonymity struct {
	AnonymityType        MSPIdentityAnonymity_MSPIdentityAnonymityType `protobuf:"varint,1,opt,name=anonymity_type,json=anonymityType,proto3,enum=common.MSPIdentityAnonymity_MSPIdentityAnonymityType" json:"anonymity_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *MSPIdentityAnonymity) Reset()         { *m = MSPIdentityAnonymity{} }
func (m *MSPIdentityAnonymity) String() string { return proto.CompactTextString(m) }
func (*MSPIdentityAnonymity) ProtoMessage()    {}
func (*MSPIdentityAnonymity) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e08b7ead29bd48, []int{3}
}

func (m *MSPIdentityAnonymity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MSPIdentityAnonymity.Unmarshal(m, b)
}
func (m *MSPIdentityAnonymity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MSPIdentityAnonymity.Marshal(b, m, deterministic)
}
func (m *MSPIdentityAnonymity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MSPIdentityAnonymity.Merge(m, src)
}
func (m *MSPIdentityAnonymity) XXX_Size() int {
	return xxx_messageInfo_MSPIdentityAnonymity.Size(m)
}
func (m *MSPIdentityAnonymity) XXX_DiscardUnknown() {
	xxx_messageInfo_MSPIdentityAnonymity.DiscardUnknown(m)
}

var xxx_messageInfo_MSPIdentityAnonymity proto.InternalMessageInfo

func (m *MSPIdentityAnonymity) GetAnonymityType() MSPIdentityAnonymity_MSPIdentityAnonymityType {
	if m != nil {
		return m.AnonymityType
	}
	return MSPIdentityAnonymity_NOMINAL
}

// CombinedPrincipal governs the organization of the Principal
// field of a policy principal when principal_classification has
// indicated that a combined form of principals is required
type CombinedPrincipal struct {
	// Principals refer to combined principals
	Principals           []*MSPPrincipal `protobuf:"bytes,1,rep,name=principals,proto3" json:"principals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CombinedPrincipal) Reset()         { *m = CombinedPrincipal{} }
func (m *CombinedPrincipal) String() string { return proto.CompactTextString(m) }
func (*CombinedPrincipal) ProtoMessage()    {}
func (*CombinedPrincipal) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e08b7ead29bd48, []int{4}
}

func (m *CombinedPrincipal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CombinedPrincipal.Unmarshal(m, b)
}
func (m *CombinedPrincipal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CombinedPrincipal.Marshal(b, m, deterministic)
}
func (m *CombinedPrincipal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CombinedPrincipal.Merge(m, src)
}
func (m *CombinedPrincipal) XXX_Size() int {
	return xxx_messageInfo_CombinedPrincipal.Size(m)
}
func (m *CombinedPrincipal) XXX_DiscardUnknown() {
	xxx_messageInfo_CombinedPrincipal.DiscardUnknown(m)
}

var xxx_messageInfo_CombinedPrincipal proto.InternalMessageInfo

func (m *CombinedPrincipal) GetPrincipals() []*MSPPrincipal {
	if m != nil {
		return m.Principals
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.MSPPrincipal_Classification", MSPPrincipal_Classification_name, MSPPrincipal_Classification_value)
	proto.RegisterEnum("common.MSPRole_MSPRoleType", MSPRole_MSPRoleType_name, MSPRole_MSPRoleType_value)
	proto.RegisterEnum("common.MSPIdentityAnonymity_MSPIdentityAnonymityType", MSPIdentityAnonymity_MSPIdentityAnonymityType_name, MSPIdentityAnonymity_MSPIdentityAnonymityType_value)
	proto.RegisterType((*MSPPrincipal)(nil), "common.MSPPrincipal")
	proto.RegisterType((*OrganizationUnit)(nil), "common.OrganizationUnit")
	proto.RegisterType((*MSPRole)(nil), "common.MSPRole")
	proto.RegisterType((*MSPIdentityAnonymity)(nil), "common.MSPIdentityAnonymity")
	proto.RegisterType((*CombinedPrincipal)(nil), "common.CombinedPrincipal")
}

func init() { proto.RegisterFile("msp/msp_principal.proto", fileDescriptor_82e08b7ead29bd48) }

var fileDescriptor_82e08b7ead29bd48 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5f, 0x6b, 0xdb, 0x3e,
	0x14, 0xad, 0x93, 0xfc, 0xda, 0xe6, 0xe6, 0x0f, 0xaa, 0x48, 0x69, 0xe0, 0x57, 0x46, 0xf0, 0x36,
	0x08, 0x8c, 0x3a, 0x90, 0x6e, 0x7b, 0x77, 0x12, 0x53, 0x0c, 0xb1, 0x1c, 0x14, 0xe7, 0xa1, 0xa5,
	0x2c, 0x38, 0x8e, 0x92, 0x0a, 0x6c, 0xcb, 0xd8, 0xee, 0x83, 0xf7, 0x91, 0xc6, 0x1e, 0xf7, 0xa9,
	0xf6, 0x29, 0x86, 0xed, 0x26, 0x51, 0xb6, 0x0e, 0xf6, 0x64, 0xce, 0xbd, 0xe7, 0x1c, 0x1f, 0x49,
	0xf7, 0xc2, 0x55, 0x90, 0x44, 0x83, 0x20, 0x89, 0x96, 0x51, 0xcc, 0x43, 0x8f, 0x47, 0xae, 0xaf,
	0x45, 0xb1, 0x48, 0x05, 0x3e, 0xf5, 0x44, 0x10, 0x88, 0x50, 0xfd, 0xa9, 0x40, 0xd3, 0x9a, 0xcf,
	0x66, 0xbb, 0x36, 0xfe, 0x02, 0xdd, 0x3d, 0x77, 0xe9, 0xf9, 0x6e, 0x92, 0xf0, 0x0d, 0xf7, 0xdc,
	0x94, 0x8b, 0xb0, 0xab, 0xf4, 0x94, 0x7e, 0x7b, 0xf8, 0x56, 0x2b, 0xb5, 0x9a, 0xac, 0xd3, 0xc6,
	0x47, 0x54, 0x7a, 0xb5, 0x37, 0x39, 0x6e, 0xe0, 0x6b, 0xa8, 0xef, 0x5b, 0xdd, 0x4a, 0x4f, 0xe9,
	0x37, 0xe9, 0xa1, 0xa0, 0x3e, 0x42, 0xfb, 0x37, 0xfe, 0x39, 0xd4, 0xa8, 0x3d, 0x35, 0xd0, 0x09,
	0xbe, 0x84, 0x0b, 0x9b, 0xde, 0xe9, 0xc4, 0x7c, 0xd0, 0x1d, 0xd3, 0x26, 0xcb, 0x05, 0x31, 0x1d,
	0xa4, 0xe0, 0x26, 0x9c, 0x9b, 0x13, 0x83, 0x38, 0xa6, 0x73, 0x8f, 0x2a, 0xb8, 0x05, 0x75, 0x9d,
	0xd8, 0xe4, 0xde, 0xca, 0x61, 0x35, 0x6f, 0x8e, 0x6d, 0x6b, 0x64, 0x12, 0x63, 0x82, 0x6a, 0xea,
	0x0f, 0x05, 0x90, 0x1d, 0x6f, 0xdd, 0x90, 0x7f, 0x2d, 0xcc, 0x17, 0x21, 0x4f, 0xf1, 0x7b, 0x68,
	0xe7, 0x17, 0xc4, 0xd7, 0x2c, 0x4c, 0xf9, 0x86, 0xb3, 0xb8, 0x38, 0x66, 0x9d, 0xb6, 0x82, 0x24,
	0x32, 0xf7, 0x45, 0x3c, 0x81, 0x37, 0x42, 0x92, 0xba, 0xfe, 0xf2, 0x39, 0xe4, 0xa9, 0x2c, 0xab,
	0x14, 0xb2, 0xeb, 0x63, 0x56, 0xfe, 0x0b, 0xc9, 0xe5, 0x16, 0x2e, 0x3d, 0x16, 0x97, 0x20, 0x91,
	0xc5, 0xd5, 0xe2, 0x26, 0x3a, 0x87, 0xe6, 0x41, 0xa4, 0x7e, 0x53, 0xe0, 0xcc, 0x9a, 0xcf, 0xa8,
	0xf0, 0xd9, 0xbf, 0xa6, 0x1d, 0x40, 0x2d, 0x16, 0x3e, 0x2b, 0x32, 0xb5, 0x87, 0xff, 0x4b, 0x2f,
	0x96, 0xbb, 0xec, 0xbe, 0x4e, 0x16, 0x31, 0x5a, 0x10, 0xd5, 0x3b, 0x68, 0x48, 0x45, 0x0c, 0x70,
	0x6a, 0x19, 0xd6, 0xc8, 0xa0, 0xe8, 0x04, 0xd7, 0xe1, 0x3f, 0x7d, 0x62, 0x99, 0x04, 0x29, 0x79,
	0x79, 0x3c, 0x35, 0x0d, 0xe2, 0xa0, 0x4a, 0xfe, 0x30, 0x33, 0xc3, 0xa0, 0xa8, 0x8a, 0x1b, 0x70,
	0x66, 0xd3, 0x89, 0x41, 0x0d, 0x8a, 0x6a, 0xea, 0x77, 0x05, 0x3a, 0xd6, 0x7c, 0x56, 0x66, 0x49,
	0x33, 0x3d, 0x14, 0x61, 0x16, 0xf0, 0x34, 0xc3, 0x8f, 0xd0, 0x76, 0x77, 0x60, 0x99, 0x66, 0x11,
	0x7b, 0x19, 0xa7, 0x4f, 0x52, 0xb8, 0x3f, 0x54, 0xaf, 0x16, 0x8b, 0xd8, 0x2d, 0x57, 0x86, 0xea,
	0x67, 0xe8, 0xfe, 0x8d, 0x9a, 0xe7, 0x23, 0xb6, 0x65, 0x12, 0x7d, 0x8a, 0x4e, 0x0e, 0x03, 0x62,
	0x2f, 0xe6, 0x48, 0x51, 0x4d, 0xb8, 0x18, 0x8b, 0x60, 0xc5, 0x43, 0xb6, 0x3e, 0xec, 0xc0, 0x47,
	0x80, 0xfd, 0x48, 0x26, 0x5d, 0xa5, 0x57, 0xed, 0x37, 0x86, 0x9d, 0xd7, 0xa6, 0x9e, 0x4a, 0xbc,
	0xd1, 0x1c, 0xde, 0x89, 0x78, 0xab, 0x3d, 0x65, 0x11, 0x8b, 0x7d, 0xb6, 0xde, 0xb2, 0x58, 0xdb,
	0xb8, 0xab, 0x98, 0x7b, 0xe5, 0xca, 0x25, 0x2f, 0x06, 0x0f, 0x1f, 0xb6, 0x3c, 0x7d, 0x7a, 0x5e,
	0xe5, 0x70, 0x20, 0x91, 0x07, 0x25, 0xf9, 0xa6, 0x24, 0xdf, 0x6c, 0x45, 0xbe, 0xb7, 0xab, 0xd3,
	0x02, 0xde, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x94, 0x15, 0xa2, 0xc9, 0x03, 0x00, 0x00,
}
