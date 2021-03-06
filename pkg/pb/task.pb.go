// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateTaskRequest struct {
	X          *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=_" json:"_,omitempty"`
	JobId      *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	NodeId     *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Target     *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	TaskAction *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Directive  *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
}

func (m *CreateTaskRequest) Reset()                    { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()               {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *CreateTaskRequest) GetX() *google_protobuf.StringValue {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *CreateTaskRequest) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateTaskRequest) GetNodeId() *google_protobuf.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *CreateTaskRequest) GetTarget() *google_protobuf.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *CreateTaskRequest) GetTaskAction() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *CreateTaskRequest) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId  *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *CreateTaskResponse) Reset()                    { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()               {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *CreateTaskResponse) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *CreateTaskResponse) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

type Task struct {
	TaskId     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId      *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	TaskAction *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Status     *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	ErrorCode  *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	Directive  *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
	Executor   *google_protobuf.StringValue `protobuf:"bytes,7,opt,name=executor" json:"executor,omitempty"`
	Owner      *google_protobuf.StringValue `protobuf:"bytes,8,opt,name=owner" json:"owner,omitempty"`
	Target     *google_protobuf.StringValue `protobuf:"bytes,9,opt,name=target" json:"target,omitempty"`
	NodeId     *google_protobuf.StringValue `protobuf:"bytes,10,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	CreateTime *google_protobuf1.Timestamp  `protobuf:"bytes,11,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime *google_protobuf1.Timestamp  `protobuf:"bytes,12,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *Task) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *Task) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Task) GetTaskAction() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *Task) GetStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Task) GetErrorCode() *google_protobuf.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Task) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Task) GetExecutor() *google_protobuf.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Task) GetOwner() *google_protobuf.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Task) GetTarget() *google_protobuf.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Task) GetNodeId() *google_protobuf.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Task) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeTasksRequest struct {
	TaskId   []string                     `protobuf:"bytes,1,rep,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId    []string                     `protobuf:"bytes,2,rep,name=job_id,json=jobId" json:"job_id,omitempty"`
	Executor *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=executor" json:"executor,omitempty"`
	Status   []string                     `protobuf:"bytes,4,rep,name=status" json:"status,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	// default is 0
	Offset uint32 `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeTasksRequest) Reset()                    { *m = DescribeTasksRequest{} }
func (m *DescribeTasksRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeTasksRequest) ProtoMessage()               {}
func (*DescribeTasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *DescribeTasksRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *DescribeTasksRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeTasksRequest) GetExecutor() *google_protobuf.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *DescribeTasksRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeTasksRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeTasksRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeTasksResponse struct {
	TotalCount uint32  `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	TaskSet    []*Task `protobuf:"bytes,2,rep,name=task_set,json=taskSet" json:"task_set,omitempty"`
}

func (m *DescribeTasksResponse) Reset()                    { *m = DescribeTasksResponse{} }
func (m *DescribeTasksResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeTasksResponse) ProtoMessage()               {}
func (*DescribeTasksResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *DescribeTasksResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeTasksResponse) GetTaskSet() []*Task {
	if m != nil {
		return m.TaskSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "openpitrix.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "openpitrix.CreateTaskResponse")
	proto.RegisterType((*Task)(nil), "openpitrix.Task")
	proto.RegisterType((*DescribeTasksRequest)(nil), "openpitrix.DescribeTasksRequest")
	proto.RegisterType((*DescribeTasksResponse)(nil), "openpitrix.DescribeTasksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskManager service

type TaskManagerClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error)
}

type taskManagerClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagerClient(cc *grpc.ClientConn) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/CreateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error) {
	out := new(DescribeTasksResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/DescribeTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskManager service

type TaskManagerServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	DescribeTasks(context.Context, *DescribeTasksRequest) (*DescribeTasksResponse, error)
}

func RegisterTaskManagerServer(s *grpc.Server, srv TaskManagerServer) {
	s.RegisterService(&_TaskManager_serviceDesc, srv)
}

func _TaskManager_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_DescribeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).DescribeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/DescribeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).DescribeTasks(ctx, req.(*DescribeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManager_CreateTask_Handler,
		},
		{
			MethodName: "DescribeTasks",
			Handler:    _TaskManager_DescribeTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto.RegisterFile("task.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xfe, 0x6d, 0x4b, 0x0b, 0x7d, 0xf7, 0x57, 0x83, 0x13, 0xd0, 0x4d, 0x83, 0x50, 0x37, 0x1e,
	0x10, 0x64, 0x17, 0x0a, 0x26, 0xa6, 0xc4, 0x43, 0xc5, 0x83, 0x3d, 0x18, 0x49, 0x41, 0x0f, 0x5e,
	0x9a, 0xe9, 0xee, 0x74, 0x99, 0x52, 0x76, 0x96, 0x99, 0x29, 0xe5, 0x68, 0x3c, 0x98, 0x78, 0xc5,
	0xcf, 0xe2, 0x57, 0xf0, 0x0b, 0x78, 0xf4, 0xea, 0x07, 0x31, 0x33, 0xbb, 0xa5, 0x2b, 0x05, 0xdd,
	0x6a, 0xe2, 0x69, 0x33, 0xef, 0x3c, 0xcf, 0xbc, 0xff, 0x9e, 0xf7, 0x5d, 0x00, 0x89, 0xc5, 0xb1,
	0x13, 0x71, 0x26, 0x19, 0x02, 0x16, 0x91, 0x30, 0xa2, 0x92, 0xd3, 0xf3, 0xca, 0x72, 0xc0, 0x58,
	0xd0, 0x27, 0xae, 0xbe, 0xe9, 0x0c, 0xba, 0xee, 0x90, 0xe3, 0x28, 0x22, 0x5c, 0xc4, 0xd8, 0xca,
	0xca, 0xd5, 0x7b, 0x49, 0x4f, 0x88, 0x90, 0xf8, 0x24, 0x4a, 0x00, 0x4b, 0x09, 0x00, 0x47, 0xd4,
	0xc5, 0x61, 0xc8, 0x24, 0x96, 0x94, 0x85, 0x23, 0xfa, 0x23, 0xfd, 0xf1, 0x36, 0x02, 0x12, 0x6e,
	0x88, 0x21, 0x0e, 0x02, 0xc2, 0x5d, 0x16, 0x69, 0xc4, 0x24, 0xda, 0xfe, 0x96, 0x83, 0xdb, 0x7b,
	0x9c, 0x60, 0x49, 0x0e, 0xb1, 0x38, 0x6e, 0x91, 0xd3, 0x01, 0x11, 0x12, 0x3d, 0x04, 0xa3, 0x6d,
	0x19, 0x55, 0x63, 0xd5, 0xac, 0x2d, 0x39, 0xb1, 0x37, 0x67, 0x14, 0x8e, 0x73, 0x20, 0x39, 0x0d,
	0x83, 0x37, 0xb8, 0x3f, 0x20, 0xad, 0xff, 0xd0, 0x36, 0x14, 0x7b, 0xac, 0xd3, 0xa6, 0xbe, 0x95,
	0xcb, 0x80, 0x2f, 0xf4, 0x58, 0xa7, 0xe9, 0xa3, 0xc7, 0x30, 0x1b, 0x32, 0x9f, 0x28, 0x56, 0x3e,
	0x03, 0xab, 0xa8, 0xc0, 0x4d, 0x1f, 0xed, 0x40, 0x51, 0x62, 0x1e, 0x10, 0x69, 0xcd, 0x64, 0x61,
	0xc5, 0x58, 0xf4, 0x14, 0x4c, 0xd5, 0x89, 0x36, 0xf6, 0x54, 0xe2, 0x56, 0x21, 0x03, 0x55, 0xb7,
	0xae, 0xa1, 0xf1, 0xa8, 0x0e, 0x25, 0x9f, 0x72, 0xe2, 0x49, 0x7a, 0x46, 0xac, 0x62, 0x06, 0xf2,
	0x18, 0x6e, 0xbf, 0x33, 0x00, 0xa5, 0xab, 0x2b, 0x22, 0x16, 0x0a, 0xa2, 0xd2, 0xd7, 0x11, 0x51,
	0x3f, 0x53, 0x91, 0x8b, 0x0a, 0xdc, 0xf4, 0xff, 0xa8, 0xd4, 0xf6, 0xe7, 0x02, 0xcc, 0x28, 0xe7,
	0xff, 0xd2, 0xe9, 0xd5, 0x92, 0xe7, 0xa7, 0x2c, 0xf9, 0x0e, 0x14, 0x85, 0xc4, 0x72, 0x20, 0xb2,
	0xf5, 0x39, 0xc6, 0xa2, 0x5d, 0x00, 0xc2, 0x39, 0xe3, 0x6d, 0x8f, 0xf9, 0xe4, 0xc6, 0x36, 0xbf,
	0x6e, 0x86, 0x72, 0xbb, 0x96, 0x74, 0x4a, 0xe3, 0xf7, 0x98, 0x4f, 0xfe, 0xa6, 0xcb, 0xe8, 0x09,
	0xcc, 0x91, 0x73, 0xe2, 0x0d, 0x24, 0xe3, 0xd6, 0x6c, 0x06, 0xea, 0x25, 0x1a, 0xd5, 0xa0, 0xc0,
	0x86, 0x21, 0xe1, 0xd6, 0x5c, 0x96, 0xda, 0x6a, 0x68, 0x6a, 0x08, 0x4a, 0x53, 0x0c, 0x41, 0x6a,
	0xe2, 0x60, 0x8a, 0x89, 0xdb, 0x05, 0xd3, 0xd3, 0xfa, 0x6d, 0xab, 0x25, 0x64, 0x99, 0x9a, 0x5a,
	0x99, 0xa0, 0x1e, 0x8e, 0x36, 0x54, 0x0b, 0x62, 0xb8, 0x32, 0x28, 0x72, 0xdc, 0x9a, 0x98, 0xfc,
	0xff, 0xef, 0xc9, 0x31, 0x5c, 0x19, 0xec, 0x2f, 0x06, 0x2c, 0x3c, 0x27, 0xc2, 0xe3, 0xb4, 0xa3,
	0x87, 0x47, 0x8c, 0x76, 0xd3, 0xdd, 0xb4, 0x8e, 0xf3, 0xab, 0xa5, 0x4b, 0xa5, 0x2e, 0xa6, 0x94,
	0xaa, 0xec, 0x89, 0x16, 0xd3, 0xdd, 0xc9, 0x4f, 0xd5, 0x9d, 0x3b, 0x29, 0x19, 0x6a, 0x47, 0x89,
	0xd0, 0x16, 0xa0, 0xd0, 0xa7, 0x27, 0x54, 0x6a, 0x8d, 0x95, 0x5b, 0xf1, 0x41, 0xa1, 0x59, 0xb7,
	0x2b, 0x88, 0xd4, 0xf2, 0x29, 0xb7, 0x92, 0x93, 0x4d, 0x60, 0xf1, 0x4a, 0x1e, 0xc9, 0x16, 0x58,
	0x01, 0x53, 0x32, 0x89, 0xfb, 0x6d, 0x8f, 0x0d, 0x42, 0xa9, 0x87, 0xb2, 0xdc, 0x02, 0x6d, 0xda,
	0x53, 0x16, 0xb4, 0x0e, 0x73, 0x3a, 0x53, 0xf5, 0xa6, 0x4a, 0xc9, 0xac, 0xcd, 0x3b, 0xe3, 0xff,
	0x88, 0xa3, 0x57, 0x8a, 0xae, 0xc5, 0x01, 0x91, 0xb5, 0x8f, 0x39, 0x30, 0x95, 0xe5, 0x25, 0x0e,
	0x71, 0x40, 0x38, 0x3a, 0x05, 0x18, 0x6f, 0x1e, 0x74, 0x2f, 0x4d, 0x9c, 0xd8, 0xf7, 0x95, 0xe5,
	0x9b, 0xae, 0xe3, 0x50, 0xed, 0x07, 0x17, 0x8d, 0x32, 0x4a, 0x94, 0x50, 0x55, 0x1e, 0xdf, 0x7f,
	0xfd, 0xfe, 0x29, 0x77, 0xcb, 0x2e, 0xb9, 0x67, 0x5b, 0xae, 0x3a, 0x8b, 0xba, 0xb1, 0x86, 0x3e,
	0x18, 0x50, 0xfe, 0x29, 0x55, 0x54, 0x4d, 0xbf, 0x7b, 0x5d, 0x37, 0x2b, 0xf7, 0x7f, 0x81, 0x48,
	0x9c, 0x6f, 0x5e, 0x34, 0x96, 0x50, 0xc5, 0x4f, 0xee, 0xb4, 0x7b, 0x51, 0x1d, 0x52, 0x79, 0x54,
	0xed, 0xd2, 0xbe, 0x24, 0x5c, 0xc7, 0x62, 0xa2, 0x71, 0x2c, 0xcf, 0xce, 0x2f, 0x1a, 0xa7, 0xe8,
	0x05, 0xa0, 0x57, 0x11, 0x09, 0xf7, 0xf5, 0xd3, 0xd5, 0x7d, 0xce, 0x7a, 0xc4, 0x93, 0xf6, 0xfa,
	0x75, 0x56, 0xb4, 0x78, 0x24, 0x65, 0x24, 0xea, 0xae, 0x9b, 0x0a, 0x86, 0xb2, 0x5a, 0x61, 0xd3,
	0xd9, 0x74, 0xb6, 0xd6, 0x0c, 0xa3, 0x36, 0x8f, 0xa3, 0xa8, 0x4f, 0x3d, 0xfd, 0xeb, 0x74, 0x7b,
	0x82, 0x85, 0xf5, 0x09, 0xcb, 0xdb, 0x5c, 0xd4, 0xe9, 0x14, 0xb5, 0xa4, 0xb6, 0x7f, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x10, 0xfb, 0x18, 0x94, 0xfc, 0x07, 0x00, 0x00,
}
