// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: internal/protobuf/raft.proto

package protobuf

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

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Term  uint64 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{0}
}

func (x *LogEntry) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LogEntry) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *LogEntry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AppendEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderId     string      `protobuf:"bytes,1,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	Term         uint64      `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	LeaderCommit uint64      `protobuf:"varint,3,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
	PrevLogIndex uint64      `protobuf:"varint,4,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	PrevLogTerm  uint64      `protobuf:"varint,5,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	Entries      []*LogEntry `protobuf:"bytes,6,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *AppendEntriesRequest) Reset() {
	*x = AppendEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesRequest) ProtoMessage() {}

func (x *AppendEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesRequest.ProtoReflect.Descriptor instead.
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{1}
}

func (x *AppendEntriesRequest) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *AppendEntriesRequest) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesRequest) GetLeaderCommit() uint64 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

func (x *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntriesRequest) GetEntries() []*LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type AppendEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AppendEntriesResponse) Reset() {
	*x = AppendEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesResponse) ProtoMessage() {}

func (x *AppendEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesResponse.ProtoReflect.Descriptor instead.
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{2}
}

func (x *AppendEntriesResponse) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RequestVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidateId  string `protobuf:"bytes,1,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
	Term         uint64 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	LastLogIndex uint64 `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm  uint64 `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
}

func (x *RequestVoteRequest) Reset() {
	*x = RequestVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_raft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteRequest) ProtoMessage() {}

func (x *RequestVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_raft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteRequest.ProtoReflect.Descriptor instead.
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{3}
}

func (x *RequestVoteRequest) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

func (x *RequestVoteRequest) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteRequest) GetLastLogIndex() uint64 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RequestVoteRequest) GetLastLogTerm() uint64 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

type RequestVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (x *RequestVoteResponse) Reset() {
	*x = RequestVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_raft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteResponse) ProtoMessage() {}

func (x *RequestVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_raft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteResponse.ProtoReflect.Descriptor instead.
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{4}
}

func (x *RequestVoteResponse) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type InstallSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term              uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Leader            string `protobuf:"bytes,2,opt,name=leader,proto3" json:"leader,omitempty"`
	LastIncludedIndex uint64 `protobuf:"varint,3,opt,name=lastIncludedIndex,proto3" json:"lastIncludedIndex,omitempty"`
	LastIncludedTerm  uint64 `protobuf:"varint,4,opt,name=lastIncludedTerm,proto3" json:"lastIncludedTerm,omitempty"`
	Offset            uint64 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	Data              []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Done              bool   `protobuf:"varint,7,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *InstallSnapshotRequest) Reset() {
	*x = InstallSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_raft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallSnapshotRequest) ProtoMessage() {}

func (x *InstallSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_raft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallSnapshotRequest.ProtoReflect.Descriptor instead.
func (*InstallSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{5}
}

func (x *InstallSnapshotRequest) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *InstallSnapshotRequest) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

func (x *InstallSnapshotRequest) GetLastIncludedIndex() uint64 {
	if x != nil {
		return x.LastIncludedIndex
	}
	return 0
}

func (x *InstallSnapshotRequest) GetLastIncludedTerm() uint64 {
	if x != nil {
		return x.LastIncludedTerm
	}
	return 0
}

func (x *InstallSnapshotRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *InstallSnapshotRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InstallSnapshotRequest) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type InstallSnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *InstallSnapshotResponse) Reset() {
	*x = InstallSnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_raft_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallSnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallSnapshotResponse) ProtoMessage() {}

func (x *InstallSnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_raft_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallSnapshotResponse.ProtoReflect.Descriptor instead.
func (*InstallSnapshotResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{6}
}

func (x *InstallSnapshotResponse) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

var File_internal_protobuf_raft_proto protoreflect.FileDescriptor

var file_internal_protobuf_raft_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48,
	0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd5, 0x01, 0x0a, 0x14, 0x41, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x72, 0x65,
	0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x65,
	0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x23, 0x0a, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c,
	0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x45, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x4b, 0x0a, 0x13, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x22, 0xde, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x49,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x10,
	0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x54, 0x65, 0x72, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x2d, 0x0a, 0x17, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x32, 0xcc, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x66, 0x74,
	0x12, 0x40, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x15, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x12, 0x17, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6d, 0x73, 0x61, 0x64, 0x61, 0x69, 0x72, 0x2f, 0x72, 0x61,
	0x66, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protobuf_raft_proto_rawDescOnce sync.Once
	file_internal_protobuf_raft_proto_rawDescData = file_internal_protobuf_raft_proto_rawDesc
)

func file_internal_protobuf_raft_proto_rawDescGZIP() []byte {
	file_internal_protobuf_raft_proto_rawDescOnce.Do(func() {
		file_internal_protobuf_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protobuf_raft_proto_rawDescData)
	})
	return file_internal_protobuf_raft_proto_rawDescData
}

var file_internal_protobuf_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_protobuf_raft_proto_goTypes = []interface{}{
	(*LogEntry)(nil),                // 0: LogEntry
	(*AppendEntriesRequest)(nil),    // 1: AppendEntriesRequest
	(*AppendEntriesResponse)(nil),   // 2: AppendEntriesResponse
	(*RequestVoteRequest)(nil),      // 3: RequestVoteRequest
	(*RequestVoteResponse)(nil),     // 4: RequestVoteResponse
	(*InstallSnapshotRequest)(nil),  // 5: InstallSnapshotRequest
	(*InstallSnapshotResponse)(nil), // 6: InstallSnapshotResponse
}
var file_internal_protobuf_raft_proto_depIdxs = []int32{
	0, // 0: AppendEntriesRequest.entries:type_name -> LogEntry
	1, // 1: Raft.AppendEntries:input_type -> AppendEntriesRequest
	3, // 2: Raft.RequestVote:input_type -> RequestVoteRequest
	5, // 3: Raft.InstallSnapshot:input_type -> InstallSnapshotRequest
	2, // 4: Raft.AppendEntries:output_type -> AppendEntriesResponse
	4, // 5: Raft.RequestVote:output_type -> RequestVoteResponse
	6, // 6: Raft.InstallSnapshot:output_type -> InstallSnapshotResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_protobuf_raft_proto_init() }
func file_internal_protobuf_raft_proto_init() {
	if File_internal_protobuf_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protobuf_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_internal_protobuf_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesRequest); i {
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
		file_internal_protobuf_raft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesResponse); i {
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
		file_internal_protobuf_raft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteRequest); i {
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
		file_internal_protobuf_raft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteResponse); i {
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
		file_internal_protobuf_raft_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallSnapshotRequest); i {
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
		file_internal_protobuf_raft_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallSnapshotResponse); i {
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
			RawDescriptor: file_internal_protobuf_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protobuf_raft_proto_goTypes,
		DependencyIndexes: file_internal_protobuf_raft_proto_depIdxs,
		MessageInfos:      file_internal_protobuf_raft_proto_msgTypes,
	}.Build()
	File_internal_protobuf_raft_proto = out.File
	file_internal_protobuf_raft_proto_rawDesc = nil
	file_internal_protobuf_raft_proto_goTypes = nil
	file_internal_protobuf_raft_proto_depIdxs = nil
}
