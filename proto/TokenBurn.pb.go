// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.0.0
// source: proto/TokenBurn.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//
//Burns tokens from the Token's treasury Account. If no Supply Key is defined, the transaction will resolve to TOKEN_HAS_NO_SUPPLY_KEY.
//The operation decreases the Total Supply of the Token. Total supply cannot go below zero.
//The amount provided must be in the lowest denomination possible. Example:
//Token A has 2 decimals. In order to burn 100 tokens, one must provide amount of 10000. In order to burn 100.55 tokens, one must provide amount of 10055.
type TokenBurnTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  *TokenID `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`    // The token for which to burn tokens. If token does not exist, transaction results in INVALID_TOKEN_ID
	Amount uint64   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"` // The amount to burn from the Treasury Account. Amount must be a positive non-zero number, not bigger than the token balance of the treasury account (0; balance], represented in the lowest denomination.
}

func (x *TokenBurnTransactionBody) Reset() {
	*x = TokenBurnTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_TokenBurn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenBurnTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBurnTransactionBody) ProtoMessage() {}

func (x *TokenBurnTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_TokenBurn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBurnTransactionBody.ProtoReflect.Descriptor instead.
func (*TokenBurnTransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_TokenBurn_proto_rawDescGZIP(), []int{0}
}

func (x *TokenBurnTransactionBody) GetToken() *TokenID {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *TokenBurnTransactionBody) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_proto_TokenBurn_proto protoreflect.FileDescriptor

var file_proto_TokenBurn_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x72,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x18, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x75, 0x72, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x44, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x4b, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_TokenBurn_proto_rawDescOnce sync.Once
	file_proto_TokenBurn_proto_rawDescData = file_proto_TokenBurn_proto_rawDesc
)

func file_proto_TokenBurn_proto_rawDescGZIP() []byte {
	file_proto_TokenBurn_proto_rawDescOnce.Do(func() {
		file_proto_TokenBurn_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_TokenBurn_proto_rawDescData)
	})
	return file_proto_TokenBurn_proto_rawDescData
}

var file_proto_TokenBurn_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_TokenBurn_proto_goTypes = []interface{}{
	(*TokenBurnTransactionBody)(nil), // 0: proto.TokenBurnTransactionBody
	(*TokenID)(nil),                  // 1: proto.TokenID
}
var file_proto_TokenBurn_proto_depIdxs = []int32{
	1, // 0: proto.TokenBurnTransactionBody.token:type_name -> proto.TokenID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_TokenBurn_proto_init() }
func file_proto_TokenBurn_proto_init() {
	if File_proto_TokenBurn_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_TokenBurn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenBurnTransactionBody); i {
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
			RawDescriptor: file_proto_TokenBurn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_TokenBurn_proto_goTypes,
		DependencyIndexes: file_proto_TokenBurn_proto_depIdxs,
		MessageInfos:      file_proto_TokenBurn_proto_msgTypes,
	}.Build()
	File_proto_TokenBurn_proto = out.File
	file_proto_TokenBurn_proto_rawDesc = nil
	file_proto_TokenBurn_proto_goTypes = nil
	file_proto_TokenBurn_proto_depIdxs = nil
}
