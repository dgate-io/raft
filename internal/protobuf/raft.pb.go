// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
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

type LogEntry_LogEntryType int32

const (
	LogEntry_LOG_ENTRY_TYPE_NOOP_UNSPECIFIED LogEntry_LogEntryType = 0
	LogEntry_LOG_ENTRY_TYPE_OPERATION        LogEntry_LogEntryType = 1
)

// Enum value maps for LogEntry_LogEntryType.
var (
	LogEntry_LogEntryType_name = map[int32]string{
		0: "LOG_ENTRY_TYPE_NOOP_UNSPECIFIED",
		1: "LOG_ENTRY_TYPE_OPERATION",
	}
	LogEntry_LogEntryType_value = map[string]int32{
		"LOG_ENTRY_TYPE_NOOP_UNSPECIFIED": 0,
		"LOG_ENTRY_TYPE_OPERATION":        1,
	}
)

func (x LogEntry_LogEntryType) Enum() *LogEntry_LogEntryType {
	p := new(LogEntry_LogEntryType)
	*p = x
	return p
}

func (x LogEntry_LogEntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogEntry_LogEntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protobuf_raft_proto_enumTypes[0].Descriptor()
}

func (LogEntry_LogEntryType) Type() protoreflect.EnumType {
	return &file_internal_protobuf_raft_proto_enumTypes[0]
}

func (x LogEntry_LogEntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogEntry_LogEntryType.Descriptor instead.
func (LogEntry_LogEntryType) EnumDescriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{0, 0}
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     uint64                `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Term      uint64                `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Offset    int64                 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Data      []byte                `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	EntryType LogEntry_LogEntryType `protobuf:"varint,5,opt,name=entry_type,json=entryType,proto3,enum=LogEntry_LogEntryType" json:"entry_type,omitempty"`
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

func (x *LogEntry) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *LogEntry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LogEntry) GetEntryType() LogEntry_LogEntryType {
	if x != nil {
		return x.EntryType
	}
	return LogEntry_LOG_ENTRY_TYPE_NOOP_UNSPECIFIED
}

type AppendEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderId     string      `protobuf:"bytes,1,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	Term         uint64      `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	LeaderCommit uint64      `protobuf:"varint,3,opt,name=leader_commit,json=leaderCommit,proto3" json:"leader_commit,omitempty"`
	PrevLogIndex uint64      `protobuf:"varint,4,opt,name=prev_log_index,json=prevLogIndex,proto3" json:"prev_log_index,omitempty"`
	PrevLogTerm  uint64      `protobuf:"varint,5,opt,name=prev_log_term,json=prevLogTerm,proto3" json:"prev_log_term,omitempty"`
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
	Index   uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Success bool   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
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

func (x *AppendEntriesResponse) GetIndex() uint64 {
	if x != nil {
		return x.Index
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

	CandidateId  string `protobuf:"bytes,1,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	Term         uint64 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	LastLogIndex uint64 `protobuf:"varint,3,opt,name=last_log_index,json=lastLogIndex,proto3" json:"last_log_index,omitempty"`
	LastLogTerm  uint64 `protobuf:"varint,4,opt,name=last_log_term,json=lastLogTerm,proto3" json:"last_log_term,omitempty"`
	Prevote      bool   `protobuf:"varint,5,opt,name=prevote,proto3" json:"prevote,omitempty"`
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

func (x *RequestVoteRequest) GetPrevote() bool {
	if x != nil {
		return x.Prevote
	}
	return false
}

type RequestVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=vote_granted,json=voteGranted,proto3" json:"vote_granted,omitempty"`
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
	LastIncludedIndex uint64 `protobuf:"varint,3,opt,name=last_included_index,json=lastIncludedIndex,proto3" json:"last_included_index,omitempty"`
	LastIncludedTerm  uint64 `protobuf:"varint,4,opt,name=last_included_term,json=lastIncludedTerm,proto3" json:"last_included_term,omitempty"`
	Configuration     []byte `protobuf:"bytes,5,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Offset            int64  `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	Data              []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	Done              bool   `protobuf:"varint,8,opt,name=done,proto3" json:"done,omitempty"`
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

func (x *InstallSnapshotRequest) GetConfiguration() []byte {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *InstallSnapshotRequest) GetOffset() int64 {
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

	Term         uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	BytesWritten int64  `protobuf:"varint,2,opt,name=bytes_written,json=bytesWritten,proto3" json:"bytes_written,omitempty"`
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

func (x *InstallSnapshotResponse) GetBytesWritten() int64 {
	if x != nil {
		return x.BytesWritten
	}
	return 0
}

type StorageState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term     uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VotedFor string `protobuf:"bytes,2,opt,name=voted_for,json=votedFor,proto3" json:"voted_for,omitempty"`
}

func (x *StorageState) Reset() {
	*x = StorageState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_raft_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageState) ProtoMessage() {}

func (x *StorageState) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_raft_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageState.ProtoReflect.Descriptor instead.
func (*StorageState) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{7}
}

func (x *StorageState) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *StorageState) GetVotedFor() string {
	if x != nil {
		return x.VotedFor
	}
	return ""
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members  map[string]string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IsVoter  map[string]bool   `protobuf:"bytes,2,rep,name=is_voter,json=isVoter,proto3" json:"is_voter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Index    uint64            `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	LeaderId string            `protobuf:"bytes,4,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_raft_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_raft_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_raft_proto_rawDescGZIP(), []int{8}
}

func (x *Configuration) GetMembers() map[string]string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Configuration) GetIsVoter() map[string]bool {
	if x != nil {
		return x.IsVoter
	}
	return nil
}

func (x *Configuration) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Configuration) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

var File_internal_protobuf_raft_proto protoreflect.FileDescriptor

var file_internal_protobuf_raft_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea,
	0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x35, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x51, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x4c, 0x4f, 0x47, 0x5f,
	0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4f, 0x50, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x4c, 0x4f, 0x47, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x22, 0xdb, 0x01, 0x0a, 0x14,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72,
	0x65, 0x76, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67,
	0x54, 0x65, 0x72, 0x6d, 0x12, 0x23, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x5b, 0x0a, 0x15, 0x41, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x70, 0x72, 0x65, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x4c, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x22, 0x88, 0x02, 0x0a, 0x16, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74,
	0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2c, 0x0a,
	0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x74,
	0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x49,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e,
	0x65, 0x22, 0x52, 0x0a, 0x17, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x12, 0x23, 0x0a, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x62, 0x79, 0x74, 0x65, 0x73, 0x57, 0x72,
	0x69, 0x74, 0x74, 0x65, 0x6e, 0x22, 0x3f, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6f, 0x74,
	0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x6f,
	0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x22, 0xa9, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12,
	0x36, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x73, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x69, 0x73, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x3a, 0x0a, 0x0c, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x49, 0x73, 0x56, 0x6f, 0x74, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x32, 0xcc, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x66, 0x74, 0x12, 0x40, 0x0a, 0x0d, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x17, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x6d, 0x73, 0x61, 0x64, 0x61, 0x69, 0x72, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_internal_protobuf_raft_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_protobuf_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_internal_protobuf_raft_proto_goTypes = []any{
	(LogEntry_LogEntryType)(0),      // 0: LogEntry.LogEntryType
	(*LogEntry)(nil),                // 1: LogEntry
	(*AppendEntriesRequest)(nil),    // 2: AppendEntriesRequest
	(*AppendEntriesResponse)(nil),   // 3: AppendEntriesResponse
	(*RequestVoteRequest)(nil),      // 4: RequestVoteRequest
	(*RequestVoteResponse)(nil),     // 5: RequestVoteResponse
	(*InstallSnapshotRequest)(nil),  // 6: InstallSnapshotRequest
	(*InstallSnapshotResponse)(nil), // 7: InstallSnapshotResponse
	(*StorageState)(nil),            // 8: StorageState
	(*Configuration)(nil),           // 9: Configuration
	nil,                             // 10: Configuration.MembersEntry
	nil,                             // 11: Configuration.IsVoterEntry
}
var file_internal_protobuf_raft_proto_depIdxs = []int32{
	0,  // 0: LogEntry.entry_type:type_name -> LogEntry.LogEntryType
	1,  // 1: AppendEntriesRequest.entries:type_name -> LogEntry
	10, // 2: Configuration.members:type_name -> Configuration.MembersEntry
	11, // 3: Configuration.is_voter:type_name -> Configuration.IsVoterEntry
	2,  // 4: Raft.AppendEntries:input_type -> AppendEntriesRequest
	4,  // 5: Raft.RequestVote:input_type -> RequestVoteRequest
	6,  // 6: Raft.InstallSnapshot:input_type -> InstallSnapshotRequest
	3,  // 7: Raft.AppendEntries:output_type -> AppendEntriesResponse
	5,  // 8: Raft.RequestVote:output_type -> RequestVoteResponse
	7,  // 9: Raft.InstallSnapshot:output_type -> InstallSnapshotResponse
	7,  // [7:10] is the sub-list for method output_type
	4,  // [4:7] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_internal_protobuf_raft_proto_init() }
func file_internal_protobuf_raft_proto_init() {
	if File_internal_protobuf_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protobuf_raft_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_internal_protobuf_raft_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_internal_protobuf_raft_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_internal_protobuf_raft_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
		file_internal_protobuf_raft_proto_msgTypes[4].Exporter = func(v any, i int) any {
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
		file_internal_protobuf_raft_proto_msgTypes[5].Exporter = func(v any, i int) any {
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
		file_internal_protobuf_raft_proto_msgTypes[6].Exporter = func(v any, i int) any {
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
		file_internal_protobuf_raft_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*StorageState); i {
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
		file_internal_protobuf_raft_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Configuration); i {
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
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protobuf_raft_proto_goTypes,
		DependencyIndexes: file_internal_protobuf_raft_proto_depIdxs,
		EnumInfos:         file_internal_protobuf_raft_proto_enumTypes,
		MessageInfos:      file_internal_protobuf_raft_proto_msgTypes,
	}.Build()
	File_internal_protobuf_raft_proto = out.File
	file_internal_protobuf_raft_proto_rawDesc = nil
	file_internal_protobuf_raft_proto_goTypes = nil
	file_internal_protobuf_raft_proto_depIdxs = nil
}
