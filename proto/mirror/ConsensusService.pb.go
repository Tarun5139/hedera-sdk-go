//-
// ‌
// Hedera Mirror Node
// ​
// Copyright (C) 2019 Hedera Hashgraph, LLC
// ​
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ‍

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.0.0
// source: proto/mirror/ConsensusService.proto

package mirror

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/hashgraph/hedera-sdk-go/v2/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ConsensusTopicQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicID *proto1.TopicID `protobuf:"bytes,1,opt,name=topicID,proto3" json:"topicID,omitempty"` // A required topic ID to retrieve messages for.
	// Include messages which reached consensus on or after this time. Defaults to current time if not set.
	ConsensusStartTime *proto1.Timestamp `protobuf:"bytes,2,opt,name=consensusStartTime,proto3" json:"consensusStartTime,omitempty"`
	// Include messages which reached consensus before this time. If not set it will receive indefinitely.
	ConsensusEndTime *proto1.Timestamp `protobuf:"bytes,3,opt,name=consensusEndTime,proto3" json:"consensusEndTime,omitempty"`
	// The maximum number of messages to receive before stopping. If not set or set to zero it will return messages
	// indefinitely.
	Limit uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ConsensusTopicQuery) Reset() {
	*x = ConsensusTopicQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mirror_ConsensusService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusTopicQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusTopicQuery) ProtoMessage() {}

func (x *ConsensusTopicQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mirror_ConsensusService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusTopicQuery.ProtoReflect.Descriptor instead.
func (*ConsensusTopicQuery) Descriptor() ([]byte, []int) {
	return file_proto_mirror_ConsensusService_proto_rawDescGZIP(), []int{0}
}

func (x *ConsensusTopicQuery) GetTopicID() *proto1.TopicID {
	if x != nil {
		return x.TopicID
	}
	return nil
}

func (x *ConsensusTopicQuery) GetConsensusStartTime() *proto1.Timestamp {
	if x != nil {
		return x.ConsensusStartTime
	}
	return nil
}

func (x *ConsensusTopicQuery) GetConsensusEndTime() *proto1.Timestamp {
	if x != nil {
		return x.ConsensusEndTime
	}
	return nil
}

func (x *ConsensusTopicQuery) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ConsensusTopicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsensusTimestamp *proto1.Timestamp `protobuf:"bytes,1,opt,name=consensusTimestamp,proto3" json:"consensusTimestamp,omitempty"` // The time at which the transaction reached consensus
	// The message body originally in the ConsensusSubmitMessageTransactionBody. Message size will be less than 6KiB.
	Message            []byte                            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RunningHash        []byte                            `protobuf:"bytes,3,opt,name=runningHash,proto3" json:"runningHash,omitempty"`                // The running hash (SHA384) of every message.
	SequenceNumber     uint64                            `protobuf:"varint,4,opt,name=sequenceNumber,proto3" json:"sequenceNumber,omitempty"`         // Starts at 1 for first submitted message. Incremented on each submitted message.
	RunningHashVersion uint64                            `protobuf:"varint,5,opt,name=runningHashVersion,proto3" json:"runningHashVersion,omitempty"` // Version of the SHA-384 digest used to update the running hash.
	ChunkInfo          *proto1.ConsensusMessageChunkInfo `protobuf:"bytes,6,opt,name=chunkInfo,proto3" json:"chunkInfo,omitempty"`                    // Optional information of the current chunk in a fragmented message.
}

func (x *ConsensusTopicResponse) Reset() {
	*x = ConsensusTopicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mirror_ConsensusService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusTopicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusTopicResponse) ProtoMessage() {}

func (x *ConsensusTopicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mirror_ConsensusService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusTopicResponse.ProtoReflect.Descriptor instead.
func (*ConsensusTopicResponse) Descriptor() ([]byte, []int) {
	return file_proto_mirror_ConsensusService_proto_rawDescGZIP(), []int{1}
}

func (x *ConsensusTopicResponse) GetConsensusTimestamp() *proto1.Timestamp {
	if x != nil {
		return x.ConsensusTimestamp
	}
	return nil
}

func (x *ConsensusTopicResponse) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ConsensusTopicResponse) GetRunningHash() []byte {
	if x != nil {
		return x.RunningHash
	}
	return nil
}

func (x *ConsensusTopicResponse) GetSequenceNumber() uint64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

func (x *ConsensusTopicResponse) GetRunningHashVersion() uint64 {
	if x != nil {
		return x.RunningHashVersion
	}
	return 0
}

func (x *ConsensusTopicResponse) GetChunkInfo() *proto1.ConsensusMessageChunkInfo {
	if x != nil {
		return x.ChunkInfo
	}
	return nil
}

var File_proto_mirror_ConsensusService_proto protoreflect.FileDescriptor

var file_proto_mirror_ConsensusService_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x2e, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x28, 0x0a,
	0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x44, 0x52, 0x07,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x44, 0x12, 0x40, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xae, 0x02,
	0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x12, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e,
	0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x8d,
	0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x2e, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x2e, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x53,
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mirror_ConsensusService_proto_rawDescOnce sync.Once
	file_proto_mirror_ConsensusService_proto_rawDescData = file_proto_mirror_ConsensusService_proto_rawDesc
)

func file_proto_mirror_ConsensusService_proto_rawDescGZIP() []byte {
	file_proto_mirror_ConsensusService_proto_rawDescOnce.Do(func() {
		file_proto_mirror_ConsensusService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mirror_ConsensusService_proto_rawDescData)
	})
	return file_proto_mirror_ConsensusService_proto_rawDescData
}

var file_proto_mirror_ConsensusService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_mirror_ConsensusService_proto_goTypes = []interface{}{
	(*ConsensusTopicQuery)(nil),              // 0: com.hedera.mirror.api.proto.ConsensusTopicQuery
	(*ConsensusTopicResponse)(nil),           // 1: com.hedera.mirror.api.proto.ConsensusTopicResponse
	(*proto1.TopicID)(nil),                   // 2: proto.TopicID
	(*proto1.Timestamp)(nil),                 // 3: proto.Timestamp
	(*proto1.ConsensusMessageChunkInfo)(nil), // 4: proto.ConsensusMessageChunkInfo
}
var file_proto_mirror_ConsensusService_proto_depIdxs = []int32{
	2, // 0: com.hedera.mirror.api.proto.ConsensusTopicQuery.topicID:type_name -> proto.TopicID
	3, // 1: com.hedera.mirror.api.proto.ConsensusTopicQuery.consensusStartTime:type_name -> proto.Timestamp
	3, // 2: com.hedera.mirror.api.proto.ConsensusTopicQuery.consensusEndTime:type_name -> proto.Timestamp
	3, // 3: com.hedera.mirror.api.proto.ConsensusTopicResponse.consensusTimestamp:type_name -> proto.Timestamp
	4, // 4: com.hedera.mirror.api.proto.ConsensusTopicResponse.chunkInfo:type_name -> proto.ConsensusMessageChunkInfo
	0, // 5: com.hedera.mirror.api.proto.ConsensusService.subscribeTopic:input_type -> com.hedera.mirror.api.proto.ConsensusTopicQuery
	1, // 6: com.hedera.mirror.api.proto.ConsensusService.subscribeTopic:output_type -> com.hedera.mirror.api.proto.ConsensusTopicResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_mirror_ConsensusService_proto_init() }
func file_proto_mirror_ConsensusService_proto_init() {
	if File_proto_mirror_ConsensusService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mirror_ConsensusService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusTopicQuery); i {
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
		file_proto_mirror_ConsensusService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusTopicResponse); i {
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
			RawDescriptor: file_proto_mirror_ConsensusService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mirror_ConsensusService_proto_goTypes,
		DependencyIndexes: file_proto_mirror_ConsensusService_proto_depIdxs,
		MessageInfos:      file_proto_mirror_ConsensusService_proto_msgTypes,
	}.Build()
	File_proto_mirror_ConsensusService_proto = out.File
	file_proto_mirror_ConsensusService_proto_rawDesc = nil
	file_proto_mirror_ConsensusService_proto_goTypes = nil
	file_proto_mirror_ConsensusService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConsensusServiceClient is the client API for ConsensusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsensusServiceClient interface {
	SubscribeTopic(ctx context.Context, in *ConsensusTopicQuery, opts ...grpc.CallOption) (ConsensusService_SubscribeTopicClient, error)
}

type consensusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsensusServiceClient(cc grpc.ClientConnInterface) ConsensusServiceClient {
	return &consensusServiceClient{cc}
}

func (c *consensusServiceClient) SubscribeTopic(ctx context.Context, in *ConsensusTopicQuery, opts ...grpc.CallOption) (ConsensusService_SubscribeTopicClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConsensusService_serviceDesc.Streams[0], "/com.hedera.mirror.api.proto.ConsensusService/subscribeTopic", opts...)
	if err != nil {
		return nil, err
	}
	x := &consensusServiceSubscribeTopicClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConsensusService_SubscribeTopicClient interface {
	Recv() (*ConsensusTopicResponse, error)
	grpc.ClientStream
}

type consensusServiceSubscribeTopicClient struct {
	grpc.ClientStream
}

func (x *consensusServiceSubscribeTopicClient) Recv() (*ConsensusTopicResponse, error) {
	m := new(ConsensusTopicResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConsensusServiceServer is the server API for ConsensusService service.
type ConsensusServiceServer interface {
	SubscribeTopic(*ConsensusTopicQuery, ConsensusService_SubscribeTopicServer) error
}

// UnimplementedConsensusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConsensusServiceServer struct {
}

func (*UnimplementedConsensusServiceServer) SubscribeTopic(*ConsensusTopicQuery, ConsensusService_SubscribeTopicServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTopic not implemented")
}

func RegisterConsensusServiceServer(s *grpc.Server, srv ConsensusServiceServer) {
	s.RegisterService(&_ConsensusService_serviceDesc, srv)
}

func _ConsensusService_SubscribeTopic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsensusTopicQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsensusServiceServer).SubscribeTopic(m, &consensusServiceSubscribeTopicServer{stream})
}

type ConsensusService_SubscribeTopicServer interface {
	Send(*ConsensusTopicResponse) error
	grpc.ServerStream
}

type consensusServiceSubscribeTopicServer struct {
	grpc.ServerStream
}

func (x *consensusServiceSubscribeTopicServer) Send(m *ConsensusTopicResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ConsensusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.hedera.mirror.api.proto.ConsensusService",
	HandlerType: (*ConsensusServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "subscribeTopic",
			Handler:       _ConsensusService_SubscribeTopic_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/mirror/ConsensusService.proto",
}
