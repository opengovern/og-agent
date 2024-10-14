// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.3
// source: pkg/proto/agent.proto

package golang

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OptimizationJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Command      string                 `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Status       string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMessage string                 `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *OptimizationJob) Reset() {
	*x = OptimizationJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptimizationJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptimizationJob) ProtoMessage() {}

func (x *OptimizationJob) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptimizationJob.ProtoReflect.Descriptor instead.
func (*OptimizationJob) Descriptor() ([]byte, []int) {
	return file_pkg_proto_agent_proto_rawDescGZIP(), []int{0}
}

func (x *OptimizationJob) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OptimizationJob) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *OptimizationJob) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OptimizationJob) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *OptimizationJob) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OptimizationJob) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *GetReportRequest) Reset() {
	*x = GetReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportRequest) ProtoMessage() {}

func (x *GetReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportRequest.ProtoReflect.Descriptor instead.
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_agent_proto_rawDescGZIP(), []int{1}
}

func (x *GetReportRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type GetReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report []byte `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *GetReportResponse) Reset() {
	*x = GetReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportResponse) ProtoMessage() {}

func (x *GetReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportResponse.ProtoReflect.Descriptor instead.
func (*GetReportResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_agent_proto_rawDescGZIP(), []int{2}
}

func (x *GetReportResponse) GetReport() []byte {
	if x != nil {
		return x.Report
	}
	return nil
}

type TriggerJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands []string `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *TriggerJobRequest) Reset() {
	*x = TriggerJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerJobRequest) ProtoMessage() {}

func (x *TriggerJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerJobRequest.ProtoReflect.Descriptor instead.
func (*TriggerJobRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_agent_proto_rawDescGZIP(), []int{3}
}

func (x *TriggerJobRequest) GetCommands() []string {
	if x != nil {
		return x.Commands
	}
	return nil
}

type GetLatestJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands []string `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *GetLatestJobsRequest) Reset() {
	*x = GetLatestJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestJobsRequest) ProtoMessage() {}

func (x *GetLatestJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestJobsRequest.ProtoReflect.Descriptor instead.
func (*GetLatestJobsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_agent_proto_rawDescGZIP(), []int{4}
}

func (x *GetLatestJobsRequest) GetCommands() []string {
	if x != nil {
		return x.Commands
	}
	return nil
}

type GetLatestJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs map[string]*OptimizationJob `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetLatestJobsResponse) Reset() {
	*x = GetLatestJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestJobsResponse) ProtoMessage() {}

func (x *GetLatestJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestJobsResponse.ProtoReflect.Descriptor instead.
func (*GetLatestJobsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_agent_proto_rawDescGZIP(), []int{5}
}

func (x *GetLatestJobsResponse) GetJobs() map[string]*OptimizationJob {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type PingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingMessage) Reset() {
	*x = PingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingMessage) ProtoMessage() {}

func (x *PingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingMessage.ProtoReflect.Descriptor instead.
func (*PingMessage) Descriptor() ([]byte, []int) {
	return file_pkg_proto_agent_proto_rawDescGZIP(), []int{6}
}

var File_pkg_proto_agent_proto protoreflect.FileDescriptor

var file_pkg_proto_agent_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x0f, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x2f, 0x0a, 0x11, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x22, 0x32, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4a,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x6a, 0x6f, 0x62, 0x73, 0x1a, 0x58, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xca,
	0x02, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x52, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x2e, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x1b, 0x2e, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x12, 0x21,
	0x2e, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x24, 0x2e, 0x6b,
	0x61, 0x79, 0x74, 0x75, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2d,
	0x69, 0x6f, 0x2f, 0x6b, 0x61, 0x79, 0x74, 0x75, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_agent_proto_rawDescOnce sync.Once
	file_pkg_proto_agent_proto_rawDescData = file_pkg_proto_agent_proto_rawDesc
)

func file_pkg_proto_agent_proto_rawDescGZIP() []byte {
	file_pkg_proto_agent_proto_rawDescOnce.Do(func() {
		file_pkg_proto_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_agent_proto_rawDescData)
	})
	return file_pkg_proto_agent_proto_rawDescData
}

var file_pkg_proto_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_proto_agent_proto_goTypes = []interface{}{
	(*OptimizationJob)(nil),       // 0: kaytu.agent.v1.OptimizationJob
	(*GetReportRequest)(nil),      // 1: kaytu.agent.v1.GetReportRequest
	(*GetReportResponse)(nil),     // 2: kaytu.agent.v1.GetReportResponse
	(*TriggerJobRequest)(nil),     // 3: kaytu.agent.v1.TriggerJobRequest
	(*GetLatestJobsRequest)(nil),  // 4: kaytu.agent.v1.GetLatestJobsRequest
	(*GetLatestJobsResponse)(nil), // 5: kaytu.agent.v1.GetLatestJobsResponse
	(*PingMessage)(nil),           // 6: kaytu.agent.v1.PingMessage
	nil,                           // 7: kaytu.agent.v1.GetLatestJobsResponse.JobsEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_pkg_proto_agent_proto_depIdxs = []int32{
	8, // 0: kaytu.agent.v1.OptimizationJob.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: kaytu.agent.v1.OptimizationJob.updated_at:type_name -> google.protobuf.Timestamp
	7, // 2: kaytu.agent.v1.GetLatestJobsResponse.jobs:type_name -> kaytu.agent.v1.GetLatestJobsResponse.JobsEntry
	0, // 3: kaytu.agent.v1.GetLatestJobsResponse.JobsEntry.value:type_name -> kaytu.agent.v1.OptimizationJob
	1, // 4: kaytu.agent.v1.Agent.GetReport:input_type -> kaytu.agent.v1.GetReportRequest
	6, // 5: kaytu.agent.v1.Agent.Ping:input_type -> kaytu.agent.v1.PingMessage
	3, // 6: kaytu.agent.v1.Agent.TriggerJob:input_type -> kaytu.agent.v1.TriggerJobRequest
	4, // 7: kaytu.agent.v1.Agent.GetLatestJobs:input_type -> kaytu.agent.v1.GetLatestJobsRequest
	2, // 8: kaytu.agent.v1.Agent.GetReport:output_type -> kaytu.agent.v1.GetReportResponse
	6, // 9: kaytu.agent.v1.Agent.Ping:output_type -> kaytu.agent.v1.PingMessage
	9, // 10: kaytu.agent.v1.Agent.TriggerJob:output_type -> google.protobuf.Empty
	5, // 11: kaytu.agent.v1.Agent.GetLatestJobs:output_type -> kaytu.agent.v1.GetLatestJobsResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_proto_agent_proto_init() }
func file_pkg_proto_agent_proto_init() {
	if File_pkg_proto_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptimizationJob); i {
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
		file_pkg_proto_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportRequest); i {
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
		file_pkg_proto_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportResponse); i {
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
		file_pkg_proto_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerJobRequest); i {
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
		file_pkg_proto_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestJobsRequest); i {
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
		file_pkg_proto_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestJobsResponse); i {
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
		file_pkg_proto_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingMessage); i {
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
			RawDescriptor: file_pkg_proto_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_agent_proto_goTypes,
		DependencyIndexes: file_pkg_proto_agent_proto_depIdxs,
		MessageInfos:      file_pkg_proto_agent_proto_msgTypes,
	}.Build()
	File_pkg_proto_agent_proto = out.File
	file_pkg_proto_agent_proto_rawDesc = nil
	file_pkg_proto_agent_proto_goTypes = nil
	file_pkg_proto_agent_proto_depIdxs = nil
}
