// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/query.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ChaincodeQueryResponse returns information about each chaincode that pertains
// to a query in lscc.go, such as GetChaincodes (returns all chaincodes
// instantiated on a channel), and GetInstalledChaincodes (returns all chaincodes
// installed on a peer)
type ChaincodeQueryResponse struct {
	Chaincodes []*ChaincodeInfo `protobuf:"bytes,1,rep,name=chaincodes" json:"chaincodes,omitempty"`
}

func (m *ChaincodeQueryResponse) Reset()                    { *m = ChaincodeQueryResponse{} }
func (m *ChaincodeQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeQueryResponse) ProtoMessage()               {}
func (*ChaincodeQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *ChaincodeQueryResponse) GetChaincodes() []*ChaincodeInfo {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

// ChaincodeInfo contains general information about an installed/instantiated
// chaincode
type ChaincodeInfo struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// the path as specified by the install/instantiate transaction
	Path string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	// the chaincode function upon instantiation and its arguments. This will be
	// blank if the query is returning information about installed chaincodes.
	Input string `protobuf:"bytes,4,opt,name=input" json:"input,omitempty"`
	// the name of the ESCC for this chaincode. This will be
	// blank if the query is returning information about installed chaincodes.
	Escc string `protobuf:"bytes,5,opt,name=escc" json:"escc,omitempty"`
	// the name of the VSCC for this chaincode. This will be
	// blank if the query is returning information about installed chaincodes.
	Vscc string `protobuf:"bytes,6,opt,name=vscc" json:"vscc,omitempty"`
}

func (m *ChaincodeInfo) Reset()                    { *m = ChaincodeInfo{} }
func (m *ChaincodeInfo) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInfo) ProtoMessage()               {}
func (*ChaincodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *ChaincodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChaincodeInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChaincodeInfo) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ChaincodeInfo) GetEscc() string {
	if m != nil {
		return m.Escc
	}
	return ""
}

func (m *ChaincodeInfo) GetVscc() string {
	if m != nil {
		return m.Vscc
	}
	return ""
}

// ChannelQueryResponse returns information about each channel that pertains
// to a query in lscc.go, such as GetChannels (returns all channels for a
// given peer)
type ChannelQueryResponse struct {
	Channels []*ChannelInfo `protobuf:"bytes,1,rep,name=channels" json:"channels,omitempty"`
}

func (m *ChannelQueryResponse) Reset()                    { *m = ChannelQueryResponse{} }
func (m *ChannelQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*ChannelQueryResponse) ProtoMessage()               {}
func (*ChannelQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *ChannelQueryResponse) GetChannels() []*ChannelInfo {
	if m != nil {
		return m.Channels
	}
	return nil
}

// ChannelInfo contains general information about channels
type ChannelInfo struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
}

func (m *ChannelInfo) Reset()                    { *m = ChannelInfo{} }
func (m *ChannelInfo) String() string            { return proto.CompactTextString(m) }
func (*ChannelInfo) ProtoMessage()               {}
func (*ChannelInfo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *ChannelInfo) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*ChaincodeQueryResponse)(nil), "protos.ChaincodeQueryResponse")
	proto.RegisterType((*ChaincodeInfo)(nil), "protos.ChaincodeInfo")
	proto.RegisterType((*ChannelQueryResponse)(nil), "protos.ChannelQueryResponse")
	proto.RegisterType((*ChannelInfo)(nil), "protos.ChannelInfo")
}

func init() { proto.RegisterFile("peer/query.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xa9, 0xfb, 0xa3, 0x3b, 0x43, 0x90, 0x38, 0x25, 0x37, 0xc2, 0xe8, 0xd5, 0x04, 0x69,
	0x40, 0xf1, 0x05, 0xdc, 0x85, 0xec, 0x6a, 0xb8, 0x4b, 0x6f, 0xa4, 0x4d, 0xcf, 0xda, 0xc0, 0x96,
	0xc4, 0xa4, 0x1d, 0xec, 0x29, 0x7c, 0x65, 0x39, 0xc9, 0x3a, 0xba, 0xab, 0x9e, 0xf3, 0xfb, 0x7e,
	0xa1, 0x7c, 0x09, 0xdc, 0x59, 0x44, 0x27, 0x7e, 0x5b, 0x74, 0xc7, 0xcc, 0x3a, 0xd3, 0x18, 0x36,
	0x0e, 0x1f, 0x9f, 0xae, 0xe1, 0x71, 0x59, 0xe7, 0x4a, 0x4b, 0x53, 0xe2, 0x17, 0xe5, 0x1b, 0xf4,
	0xd6, 0x68, 0x8f, 0xec, 0x1d, 0x40, 0x76, 0x89, 0xe7, 0xc9, 0x7c, 0xb0, 0x98, 0xbe, 0x3e, 0xc4,
	0xd3, 0x3e, 0x3b, 0x9f, 0x59, 0xe9, 0xad, 0xd9, 0xf4, 0xc4, 0xf4, 0x2f, 0x81, 0xdb, 0x8b, 0x94,
	0x31, 0x18, 0xea, 0x7c, 0x8f, 0x3c, 0x99, 0x27, 0x8b, 0xc9, 0x26, 0xcc, 0x8c, 0xc3, 0xf5, 0x01,
	0x9d, 0x57, 0x46, 0xf3, 0xab, 0x80, 0xbb, 0x95, 0x6c, 0x9b, 0x37, 0x35, 0x1f, 0x44, 0x9b, 0x66,
	0x36, 0x83, 0x91, 0xd2, 0xb6, 0x6d, 0xf8, 0x30, 0xc0, 0xb8, 0x90, 0x89, 0x5e, 0x4a, 0x3e, 0x8a,
	0x26, 0xcd, 0xc4, 0x0e, 0xc4, 0xc6, 0x91, 0xd1, 0x9c, 0x7e, 0xc2, 0x6c, 0x59, 0xe7, 0x5a, 0xe3,
	0xee, 0xb2, 0xa0, 0x80, 0x1b, 0x19, 0x79, 0x57, 0xef, 0xbe, 0x57, 0x8f, 0x78, 0x28, 0x77, 0x96,
	0xd2, 0x17, 0x98, 0xf6, 0x02, 0xf6, 0x14, 0x2e, 0x88, 0xd6, 0x1f, 0x55, 0x9e, 0xda, 0x4d, 0x4e,
	0x64, 0x55, 0x7e, 0xac, 0x21, 0x35, 0xae, 0xca, 0xea, 0xa3, 0x45, 0xb7, 0xc3, 0xb2, 0x42, 0x97,
	0x6d, 0xf3, 0xc2, 0x29, 0xd9, 0xfd, 0x84, 0xde, 0xe4, 0xfb, 0xb9, 0x52, 0x4d, 0xdd, 0x16, 0x99,
	0x34, 0x7b, 0xd1, 0x53, 0x45, 0x54, 0x45, 0x54, 0x05, 0xa9, 0x45, 0x7c, 0xb2, 0xb7, 0xff, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x56, 0xd1, 0xfe, 0x74, 0xcd, 0x01, 0x00, 0x00,
}
