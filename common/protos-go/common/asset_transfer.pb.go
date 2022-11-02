// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: common/asset_transfer.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AssetPledge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetDetails    []byte `protobuf:"bytes,1,opt,name=assetDetails,proto3" json:"assetDetails,omitempty"`
	LocalNetworkID  string `protobuf:"bytes,2,opt,name=localNetworkID,proto3" json:"localNetworkID,omitempty"`
	RemoteNetworkID string `protobuf:"bytes,3,opt,name=remoteNetworkID,proto3" json:"remoteNetworkID,omitempty"`
	Recipient       string `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	ExpiryTimeSecs  uint64 `protobuf:"varint,5,opt,name=expiryTimeSecs,proto3" json:"expiryTimeSecs,omitempty"`
}

func (x *AssetPledge) Reset() {
	*x = AssetPledge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_asset_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetPledge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetPledge) ProtoMessage() {}

func (x *AssetPledge) ProtoReflect() protoreflect.Message {
	mi := &file_common_asset_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetPledge.ProtoReflect.Descriptor instead.
func (*AssetPledge) Descriptor() ([]byte, []int) {
	return file_common_asset_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *AssetPledge) GetAssetDetails() []byte {
	if x != nil {
		return x.AssetDetails
	}
	return nil
}

func (x *AssetPledge) GetLocalNetworkID() string {
	if x != nil {
		return x.LocalNetworkID
	}
	return ""
}

func (x *AssetPledge) GetRemoteNetworkID() string {
	if x != nil {
		return x.RemoteNetworkID
	}
	return ""
}

func (x *AssetPledge) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *AssetPledge) GetExpiryTimeSecs() uint64 {
	if x != nil {
		return x.ExpiryTimeSecs
	}
	return 0
}

type AssetClaimStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetDetails     []byte `protobuf:"bytes,1,opt,name=assetDetails,proto3" json:"assetDetails,omitempty"`
	LocalNetworkID   string `protobuf:"bytes,2,opt,name=localNetworkID,proto3" json:"localNetworkID,omitempty"`
	RemoteNetworkID  string `protobuf:"bytes,3,opt,name=remoteNetworkID,proto3" json:"remoteNetworkID,omitempty"`
	Recipient        string `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	ClaimStatus      bool   `protobuf:"varint,5,opt,name=claimStatus,proto3" json:"claimStatus,omitempty"`
	ExpiryTimeSecs   uint64 `protobuf:"varint,6,opt,name=expiryTimeSecs,proto3" json:"expiryTimeSecs,omitempty"`
	ExpirationStatus bool   `protobuf:"varint,7,opt,name=expirationStatus,proto3" json:"expirationStatus,omitempty"`
}

func (x *AssetClaimStatus) Reset() {
	*x = AssetClaimStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_asset_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetClaimStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetClaimStatus) ProtoMessage() {}

func (x *AssetClaimStatus) ProtoReflect() protoreflect.Message {
	mi := &file_common_asset_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetClaimStatus.ProtoReflect.Descriptor instead.
func (*AssetClaimStatus) Descriptor() ([]byte, []int) {
	return file_common_asset_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *AssetClaimStatus) GetAssetDetails() []byte {
	if x != nil {
		return x.AssetDetails
	}
	return nil
}

func (x *AssetClaimStatus) GetLocalNetworkID() string {
	if x != nil {
		return x.LocalNetworkID
	}
	return ""
}

func (x *AssetClaimStatus) GetRemoteNetworkID() string {
	if x != nil {
		return x.RemoteNetworkID
	}
	return ""
}

func (x *AssetClaimStatus) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *AssetClaimStatus) GetClaimStatus() bool {
	if x != nil {
		return x.ClaimStatus
	}
	return false
}

func (x *AssetClaimStatus) GetExpiryTimeSecs() uint64 {
	if x != nil {
		return x.ExpiryTimeSecs
	}
	return 0
}

func (x *AssetClaimStatus) GetExpirationStatus() bool {
	if x != nil {
		return x.ExpirationStatus
	}
	return false
}

var File_common_asset_transfer_proto protoreflect.FileDescriptor

var file_common_asset_transfer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44,
	0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x73,
	0x22, 0x9c, 0x02, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x44, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x63, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x7a, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2d, 0x64,
	0x6c, 0x74, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_asset_transfer_proto_rawDescOnce sync.Once
	file_common_asset_transfer_proto_rawDescData = file_common_asset_transfer_proto_rawDesc
)

func file_common_asset_transfer_proto_rawDescGZIP() []byte {
	file_common_asset_transfer_proto_rawDescOnce.Do(func() {
		file_common_asset_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_asset_transfer_proto_rawDescData)
	})
	return file_common_asset_transfer_proto_rawDescData
}

var file_common_asset_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_asset_transfer_proto_goTypes = []interface{}{
	(*AssetPledge)(nil),      // 0: common.asset_transfer.AssetPledge
	(*AssetClaimStatus)(nil), // 1: common.asset_transfer.AssetClaimStatus
}
var file_common_asset_transfer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_asset_transfer_proto_init() }
func file_common_asset_transfer_proto_init() {
	if File_common_asset_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_asset_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetPledge); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_asset_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetClaimStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_asset_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_asset_transfer_proto_goTypes,
		DependencyIndexes: file_common_asset_transfer_proto_depIdxs,
		MessageInfos:      file_common_asset_transfer_proto_msgTypes,
	}.Build()
	File_common_asset_transfer_proto = out.File
	file_common_asset_transfer_proto_rawDesc = nil
	file_common_asset_transfer_proto_goTypes = nil
	file_common_asset_transfer_proto_depIdxs = nil
}
