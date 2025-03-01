// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: pkg/proto/remoteworker/remoteworker.proto

package remoteworker

import (
	v2 "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SynchronizeRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	WorkerId           map[string]string      `protobuf:"bytes,1,rep,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	InstanceNamePrefix string                 `protobuf:"bytes,2,opt,name=instance_name_prefix,json=instanceNamePrefix,proto3" json:"instance_name_prefix,omitempty"`
	Platform           *v2.Platform           `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	SizeClass          uint32                 `protobuf:"varint,5,opt,name=size_class,json=sizeClass,proto3" json:"size_class,omitempty"`
	CurrentState       *CurrentState          `protobuf:"bytes,4,opt,name=current_state,json=currentState,proto3" json:"current_state,omitempty"`
	PreferBeingIdle    bool                   `protobuf:"varint,6,opt,name=prefer_being_idle,json=preferBeingIdle,proto3" json:"prefer_being_idle,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *SynchronizeRequest) Reset() {
	*x = SynchronizeRequest{}
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SynchronizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchronizeRequest) ProtoMessage() {}

func (x *SynchronizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchronizeRequest.ProtoReflect.Descriptor instead.
func (*SynchronizeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_remoteworker_remoteworker_proto_rawDescGZIP(), []int{0}
}

func (x *SynchronizeRequest) GetWorkerId() map[string]string {
	if x != nil {
		return x.WorkerId
	}
	return nil
}

func (x *SynchronizeRequest) GetInstanceNamePrefix() string {
	if x != nil {
		return x.InstanceNamePrefix
	}
	return ""
}

func (x *SynchronizeRequest) GetPlatform() *v2.Platform {
	if x != nil {
		return x.Platform
	}
	return nil
}

func (x *SynchronizeRequest) GetSizeClass() uint32 {
	if x != nil {
		return x.SizeClass
	}
	return 0
}

func (x *SynchronizeRequest) GetCurrentState() *CurrentState {
	if x != nil {
		return x.CurrentState
	}
	return nil
}

func (x *SynchronizeRequest) GetPreferBeingIdle() bool {
	if x != nil {
		return x.PreferBeingIdle
	}
	return false
}

type CurrentState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to WorkerState:
	//
	//	*CurrentState_Idle
	//	*CurrentState_Executing_
	WorkerState   isCurrentState_WorkerState `protobuf_oneof:"worker_state"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CurrentState) Reset() {
	*x = CurrentState{}
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrentState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentState) ProtoMessage() {}

func (x *CurrentState) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentState.ProtoReflect.Descriptor instead.
func (*CurrentState) Descriptor() ([]byte, []int) {
	return file_pkg_proto_remoteworker_remoteworker_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentState) GetWorkerState() isCurrentState_WorkerState {
	if x != nil {
		return x.WorkerState
	}
	return nil
}

func (x *CurrentState) GetIdle() *emptypb.Empty {
	if x != nil {
		if x, ok := x.WorkerState.(*CurrentState_Idle); ok {
			return x.Idle
		}
	}
	return nil
}

func (x *CurrentState) GetExecuting() *CurrentState_Executing {
	if x != nil {
		if x, ok := x.WorkerState.(*CurrentState_Executing_); ok {
			return x.Executing
		}
	}
	return nil
}

type isCurrentState_WorkerState interface {
	isCurrentState_WorkerState()
}

type CurrentState_Idle struct {
	Idle *emptypb.Empty `protobuf:"bytes,1,opt,name=idle,proto3,oneof"`
}

type CurrentState_Executing_ struct {
	Executing *CurrentState_Executing `protobuf:"bytes,2,opt,name=executing,proto3,oneof"`
}

func (*CurrentState_Idle) isCurrentState_WorkerState() {}

func (*CurrentState_Executing_) isCurrentState_WorkerState() {}

type SynchronizeResponse struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	NextSynchronizationAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=next_synchronization_at,json=nextSynchronizationAt,proto3" json:"next_synchronization_at,omitempty"`
	DesiredState          *DesiredState          `protobuf:"bytes,2,opt,name=desired_state,json=desiredState,proto3" json:"desired_state,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *SynchronizeResponse) Reset() {
	*x = SynchronizeResponse{}
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SynchronizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchronizeResponse) ProtoMessage() {}

func (x *SynchronizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchronizeResponse.ProtoReflect.Descriptor instead.
func (*SynchronizeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_remoteworker_remoteworker_proto_rawDescGZIP(), []int{2}
}

func (x *SynchronizeResponse) GetNextSynchronizationAt() *timestamppb.Timestamp {
	if x != nil {
		return x.NextSynchronizationAt
	}
	return nil
}

func (x *SynchronizeResponse) GetDesiredState() *DesiredState {
	if x != nil {
		return x.DesiredState
	}
	return nil
}

type DesiredState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to WorkerState:
	//
	//	*DesiredState_Idle
	//	*DesiredState_Executing_
	WorkerState   isDesiredState_WorkerState `protobuf_oneof:"worker_state"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DesiredState) Reset() {
	*x = DesiredState{}
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DesiredState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesiredState) ProtoMessage() {}

func (x *DesiredState) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesiredState.ProtoReflect.Descriptor instead.
func (*DesiredState) Descriptor() ([]byte, []int) {
	return file_pkg_proto_remoteworker_remoteworker_proto_rawDescGZIP(), []int{3}
}

func (x *DesiredState) GetWorkerState() isDesiredState_WorkerState {
	if x != nil {
		return x.WorkerState
	}
	return nil
}

func (x *DesiredState) GetIdle() *emptypb.Empty {
	if x != nil {
		if x, ok := x.WorkerState.(*DesiredState_Idle); ok {
			return x.Idle
		}
	}
	return nil
}

func (x *DesiredState) GetExecuting() *DesiredState_Executing {
	if x != nil {
		if x, ok := x.WorkerState.(*DesiredState_Executing_); ok {
			return x.Executing
		}
	}
	return nil
}

type isDesiredState_WorkerState interface {
	isDesiredState_WorkerState()
}

type DesiredState_Idle struct {
	Idle *emptypb.Empty `protobuf:"bytes,1,opt,name=idle,proto3,oneof"`
}

type DesiredState_Executing_ struct {
	Executing *DesiredState_Executing `protobuf:"bytes,2,opt,name=executing,proto3,oneof"`
}

func (*DesiredState_Idle) isDesiredState_WorkerState() {}

func (*DesiredState_Executing_) isDesiredState_WorkerState() {}

type CurrentState_Executing struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	ActionDigest *v2.Digest             `protobuf:"bytes,1,opt,name=action_digest,json=actionDigest,proto3" json:"action_digest,omitempty"`
	// Types that are valid to be assigned to ExecutionState:
	//
	//	*CurrentState_Executing_Started
	//	*CurrentState_Executing_FetchingInputs
	//	*CurrentState_Executing_Running
	//	*CurrentState_Executing_UploadingOutputs
	//	*CurrentState_Executing_Completed
	ExecutionState isCurrentState_Executing_ExecutionState `protobuf_oneof:"execution_state"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CurrentState_Executing) Reset() {
	*x = CurrentState_Executing{}
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrentState_Executing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentState_Executing) ProtoMessage() {}

func (x *CurrentState_Executing) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentState_Executing.ProtoReflect.Descriptor instead.
func (*CurrentState_Executing) Descriptor() ([]byte, []int) {
	return file_pkg_proto_remoteworker_remoteworker_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CurrentState_Executing) GetActionDigest() *v2.Digest {
	if x != nil {
		return x.ActionDigest
	}
	return nil
}

func (x *CurrentState_Executing) GetExecutionState() isCurrentState_Executing_ExecutionState {
	if x != nil {
		return x.ExecutionState
	}
	return nil
}

func (x *CurrentState_Executing) GetStarted() *emptypb.Empty {
	if x != nil {
		if x, ok := x.ExecutionState.(*CurrentState_Executing_Started); ok {
			return x.Started
		}
	}
	return nil
}

func (x *CurrentState_Executing) GetFetchingInputs() *emptypb.Empty {
	if x != nil {
		if x, ok := x.ExecutionState.(*CurrentState_Executing_FetchingInputs); ok {
			return x.FetchingInputs
		}
	}
	return nil
}

func (x *CurrentState_Executing) GetRunning() *emptypb.Empty {
	if x != nil {
		if x, ok := x.ExecutionState.(*CurrentState_Executing_Running); ok {
			return x.Running
		}
	}
	return nil
}

func (x *CurrentState_Executing) GetUploadingOutputs() *emptypb.Empty {
	if x != nil {
		if x, ok := x.ExecutionState.(*CurrentState_Executing_UploadingOutputs); ok {
			return x.UploadingOutputs
		}
	}
	return nil
}

func (x *CurrentState_Executing) GetCompleted() *v2.ExecuteResponse {
	if x != nil {
		if x, ok := x.ExecutionState.(*CurrentState_Executing_Completed); ok {
			return x.Completed
		}
	}
	return nil
}

type isCurrentState_Executing_ExecutionState interface {
	isCurrentState_Executing_ExecutionState()
}

type CurrentState_Executing_Started struct {
	Started *emptypb.Empty `protobuf:"bytes,2,opt,name=started,proto3,oneof"`
}

type CurrentState_Executing_FetchingInputs struct {
	FetchingInputs *emptypb.Empty `protobuf:"bytes,3,opt,name=fetching_inputs,json=fetchingInputs,proto3,oneof"`
}

type CurrentState_Executing_Running struct {
	Running *emptypb.Empty `protobuf:"bytes,4,opt,name=running,proto3,oneof"`
}

type CurrentState_Executing_UploadingOutputs struct {
	UploadingOutputs *emptypb.Empty `protobuf:"bytes,5,opt,name=uploading_outputs,json=uploadingOutputs,proto3,oneof"`
}

type CurrentState_Executing_Completed struct {
	Completed *v2.ExecuteResponse `protobuf:"bytes,6,opt,name=completed,proto3,oneof"`
}

func (*CurrentState_Executing_Started) isCurrentState_Executing_ExecutionState() {}

func (*CurrentState_Executing_FetchingInputs) isCurrentState_Executing_ExecutionState() {}

func (*CurrentState_Executing_Running) isCurrentState_Executing_ExecutionState() {}

func (*CurrentState_Executing_UploadingOutputs) isCurrentState_Executing_ExecutionState() {}

func (*CurrentState_Executing_Completed) isCurrentState_Executing_ExecutionState() {}

type DesiredState_Executing struct {
	state              protoimpl.MessageState  `protogen:"open.v1"`
	ActionDigest       *v2.Digest              `protobuf:"bytes,1,opt,name=action_digest,json=actionDigest,proto3" json:"action_digest,omitempty"`
	Action             *v2.Action              `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	QueuedTimestamp    *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=queued_timestamp,json=queuedTimestamp,proto3" json:"queued_timestamp,omitempty"`
	AuxiliaryMetadata  []*anypb.Any            `protobuf:"bytes,6,rep,name=auxiliary_metadata,json=auxiliaryMetadata,proto3" json:"auxiliary_metadata,omitempty"`
	InstanceNameSuffix string                  `protobuf:"bytes,7,opt,name=instance_name_suffix,json=instanceNameSuffix,proto3" json:"instance_name_suffix,omitempty"`
	W3CTraceContext    map[string]string       `protobuf:"bytes,8,rep,name=w3c_trace_context,json=w3cTraceContext,proto3" json:"w3c_trace_context,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DigestFunction     v2.DigestFunction_Value `protobuf:"varint,9,opt,name=digest_function,json=digestFunction,proto3,enum=build.bazel.remote.execution.v2.DigestFunction_Value" json:"digest_function,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *DesiredState_Executing) Reset() {
	*x = DesiredState_Executing{}
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DesiredState_Executing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesiredState_Executing) ProtoMessage() {}

func (x *DesiredState_Executing) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesiredState_Executing.ProtoReflect.Descriptor instead.
func (*DesiredState_Executing) Descriptor() ([]byte, []int) {
	return file_pkg_proto_remoteworker_remoteworker_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DesiredState_Executing) GetActionDigest() *v2.Digest {
	if x != nil {
		return x.ActionDigest
	}
	return nil
}

func (x *DesiredState_Executing) GetAction() *v2.Action {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *DesiredState_Executing) GetQueuedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.QueuedTimestamp
	}
	return nil
}

func (x *DesiredState_Executing) GetAuxiliaryMetadata() []*anypb.Any {
	if x != nil {
		return x.AuxiliaryMetadata
	}
	return nil
}

func (x *DesiredState_Executing) GetInstanceNameSuffix() string {
	if x != nil {
		return x.InstanceNameSuffix
	}
	return ""
}

func (x *DesiredState_Executing) GetW3CTraceContext() map[string]string {
	if x != nil {
		return x.W3CTraceContext
	}
	return nil
}

func (x *DesiredState_Executing) GetDigestFunction() v2.DigestFunction_Value {
	if x != nil {
		return x.DigestFunction
	}
	return v2.DigestFunction_Value(0)
}

var File_pkg_proto_remoteworker_remoteworker_proto protoreflect.FileDescriptor

var file_pkg_proto_remoteworker_remoteworker_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x1a, 0x36, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c,
	0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x03, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x09, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x45, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62,
	0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x73, 0x69, 0x7a, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x49, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x5f,
	0x62, 0x65, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x42, 0x65, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x6c,
	0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd5,
	0x04, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x2c, 0x0a, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x12, 0x4e, 0x0a,
	0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0xb6, 0x03,
	0x0a, 0x09, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x4c, 0x0a, 0x0d, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x48, 0x00, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x41, 0x0a,
	0x0f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00,
	0x52, 0x0e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x12, 0x32, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x07, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x12, 0x45, 0x0a, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x10, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x50, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x11, 0x0a,
	0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x42, 0x0e, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63, 0x68,
	0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x17, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x15, 0x6e, 0x65, 0x78,
	0x74, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x74, 0x12, 0x49, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x0c, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x98, 0x06,
	0x0a, 0x0c, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x09,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x48,
	0x00, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0xf9, 0x04, 0x0a,
	0x09, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x4c, 0x0a, 0x0d, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x10, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x43, 0x0a, 0x12, 0x61, 0x75, 0x78, 0x69, 0x6c, 0x69, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x11, 0x61, 0x75, 0x78, 0x69, 0x6c, 0x69, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x6f, 0x0a, 0x11, 0x77, 0x33, 0x63, 0x5f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x69,
	0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x57, 0x33, 0x63, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x77, 0x33, 0x63, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x5e, 0x0a, 0x0f, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x42, 0x0a, 0x14, 0x57, 0x33, 0x63, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x03,
	0x10, 0x04, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x42, 0x0e, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0x78, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x66, 0x0a, 0x0b, 0x53, 0x79,
	0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72,
	0x6e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_pkg_proto_remoteworker_remoteworker_proto_rawDescOnce sync.Once
	file_pkg_proto_remoteworker_remoteworker_proto_rawDescData []byte
)

func file_pkg_proto_remoteworker_remoteworker_proto_rawDescGZIP() []byte {
	file_pkg_proto_remoteworker_remoteworker_proto_rawDescOnce.Do(func() {
		file_pkg_proto_remoteworker_remoteworker_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_remoteworker_remoteworker_proto_rawDesc), len(file_pkg_proto_remoteworker_remoteworker_proto_rawDesc)))
	})
	return file_pkg_proto_remoteworker_remoteworker_proto_rawDescData
}

var file_pkg_proto_remoteworker_remoteworker_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_proto_remoteworker_remoteworker_proto_goTypes = []any{
	(*SynchronizeRequest)(nil),     // 0: buildbarn.remoteworker.SynchronizeRequest
	(*CurrentState)(nil),           // 1: buildbarn.remoteworker.CurrentState
	(*SynchronizeResponse)(nil),    // 2: buildbarn.remoteworker.SynchronizeResponse
	(*DesiredState)(nil),           // 3: buildbarn.remoteworker.DesiredState
	nil,                            // 4: buildbarn.remoteworker.SynchronizeRequest.WorkerIdEntry
	(*CurrentState_Executing)(nil), // 5: buildbarn.remoteworker.CurrentState.Executing
	(*DesiredState_Executing)(nil), // 6: buildbarn.remoteworker.DesiredState.Executing
	nil,                            // 7: buildbarn.remoteworker.DesiredState.Executing.W3cTraceContextEntry
	(*v2.Platform)(nil),            // 8: build.bazel.remote.execution.v2.Platform
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
	(*v2.Digest)(nil),              // 11: build.bazel.remote.execution.v2.Digest
	(*v2.ExecuteResponse)(nil),     // 12: build.bazel.remote.execution.v2.ExecuteResponse
	(*v2.Action)(nil),              // 13: build.bazel.remote.execution.v2.Action
	(*anypb.Any)(nil),              // 14: google.protobuf.Any
	(v2.DigestFunction_Value)(0),   // 15: build.bazel.remote.execution.v2.DigestFunction.Value
}
var file_pkg_proto_remoteworker_remoteworker_proto_depIdxs = []int32{
	4,  // 0: buildbarn.remoteworker.SynchronizeRequest.worker_id:type_name -> buildbarn.remoteworker.SynchronizeRequest.WorkerIdEntry
	8,  // 1: buildbarn.remoteworker.SynchronizeRequest.platform:type_name -> build.bazel.remote.execution.v2.Platform
	1,  // 2: buildbarn.remoteworker.SynchronizeRequest.current_state:type_name -> buildbarn.remoteworker.CurrentState
	9,  // 3: buildbarn.remoteworker.CurrentState.idle:type_name -> google.protobuf.Empty
	5,  // 4: buildbarn.remoteworker.CurrentState.executing:type_name -> buildbarn.remoteworker.CurrentState.Executing
	10, // 5: buildbarn.remoteworker.SynchronizeResponse.next_synchronization_at:type_name -> google.protobuf.Timestamp
	3,  // 6: buildbarn.remoteworker.SynchronizeResponse.desired_state:type_name -> buildbarn.remoteworker.DesiredState
	9,  // 7: buildbarn.remoteworker.DesiredState.idle:type_name -> google.protobuf.Empty
	6,  // 8: buildbarn.remoteworker.DesiredState.executing:type_name -> buildbarn.remoteworker.DesiredState.Executing
	11, // 9: buildbarn.remoteworker.CurrentState.Executing.action_digest:type_name -> build.bazel.remote.execution.v2.Digest
	9,  // 10: buildbarn.remoteworker.CurrentState.Executing.started:type_name -> google.protobuf.Empty
	9,  // 11: buildbarn.remoteworker.CurrentState.Executing.fetching_inputs:type_name -> google.protobuf.Empty
	9,  // 12: buildbarn.remoteworker.CurrentState.Executing.running:type_name -> google.protobuf.Empty
	9,  // 13: buildbarn.remoteworker.CurrentState.Executing.uploading_outputs:type_name -> google.protobuf.Empty
	12, // 14: buildbarn.remoteworker.CurrentState.Executing.completed:type_name -> build.bazel.remote.execution.v2.ExecuteResponse
	11, // 15: buildbarn.remoteworker.DesiredState.Executing.action_digest:type_name -> build.bazel.remote.execution.v2.Digest
	13, // 16: buildbarn.remoteworker.DesiredState.Executing.action:type_name -> build.bazel.remote.execution.v2.Action
	10, // 17: buildbarn.remoteworker.DesiredState.Executing.queued_timestamp:type_name -> google.protobuf.Timestamp
	14, // 18: buildbarn.remoteworker.DesiredState.Executing.auxiliary_metadata:type_name -> google.protobuf.Any
	7,  // 19: buildbarn.remoteworker.DesiredState.Executing.w3c_trace_context:type_name -> buildbarn.remoteworker.DesiredState.Executing.W3cTraceContextEntry
	15, // 20: buildbarn.remoteworker.DesiredState.Executing.digest_function:type_name -> build.bazel.remote.execution.v2.DigestFunction.Value
	0,  // 21: buildbarn.remoteworker.OperationQueue.Synchronize:input_type -> buildbarn.remoteworker.SynchronizeRequest
	2,  // 22: buildbarn.remoteworker.OperationQueue.Synchronize:output_type -> buildbarn.remoteworker.SynchronizeResponse
	22, // [22:23] is the sub-list for method output_type
	21, // [21:22] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_pkg_proto_remoteworker_remoteworker_proto_init() }
func file_pkg_proto_remoteworker_remoteworker_proto_init() {
	if File_pkg_proto_remoteworker_remoteworker_proto != nil {
		return
	}
	file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[1].OneofWrappers = []any{
		(*CurrentState_Idle)(nil),
		(*CurrentState_Executing_)(nil),
	}
	file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[3].OneofWrappers = []any{
		(*DesiredState_Idle)(nil),
		(*DesiredState_Executing_)(nil),
	}
	file_pkg_proto_remoteworker_remoteworker_proto_msgTypes[5].OneofWrappers = []any{
		(*CurrentState_Executing_Started)(nil),
		(*CurrentState_Executing_FetchingInputs)(nil),
		(*CurrentState_Executing_Running)(nil),
		(*CurrentState_Executing_UploadingOutputs)(nil),
		(*CurrentState_Executing_Completed)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_remoteworker_remoteworker_proto_rawDesc), len(file_pkg_proto_remoteworker_remoteworker_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_remoteworker_remoteworker_proto_goTypes,
		DependencyIndexes: file_pkg_proto_remoteworker_remoteworker_proto_depIdxs,
		MessageInfos:      file_pkg_proto_remoteworker_remoteworker_proto_msgTypes,
	}.Build()
	File_pkg_proto_remoteworker_remoteworker_proto = out.File
	file_pkg_proto_remoteworker_remoteworker_proto_goTypes = nil
	file_pkg_proto_remoteworker_remoteworker_proto_depIdxs = nil
}
