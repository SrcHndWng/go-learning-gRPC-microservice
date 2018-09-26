// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo-service.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Taks we have to do
type ToDo struct {
	// Unique integer identifier of the todo task
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Title of the task
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Detail description of the todo task
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Date and time to remind the todo task
	Reminder             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ToDo) Reset()         { *m = ToDo{} }
func (m *ToDo) String() string { return proto.CompactTextString(m) }
func (*ToDo) ProtoMessage()    {}
func (*ToDo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{0}
}
func (m *ToDo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToDo.Unmarshal(m, b)
}
func (m *ToDo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToDo.Marshal(b, m, deterministic)
}
func (dst *ToDo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToDo.Merge(dst, src)
}
func (m *ToDo) XXX_Size() int {
	return xxx_messageInfo_ToDo.Size(m)
}
func (m *ToDo) XXX_DiscardUnknown() {
	xxx_messageInfo_ToDo.DiscardUnknown(m)
}

var xxx_messageInfo_ToDo proto.InternalMessageInfo

func (m *ToDo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ToDo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToDo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToDo) GetReminder() *timestamp.Timestamp {
	if m != nil {
		return m.Reminder
	}
	return nil
}

// Request data to create new todo task
type CreateRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity to add
	ToDo                 *ToDo    `protobuf:"bytes,3,opt,name=toDo,proto3" json:"toDo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{1}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Response that contains data for created todo task
type CreateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	// ID of created task
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// Message: Add V2 API
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{2}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Request data to read todo task
type ReadRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier of the todo task
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{3}
}
func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (dst *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(dst, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains todo task data specified in by ID request
type ReadResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity read by ID
	ToDo *ToDo `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
	// Message: Add V2 API
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{4}
}
func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (dst *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(dst, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

func (m *ReadResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Request data to update todo task
type UpdateRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity to update
	ToDo                 *ToDo    `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{5}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Contains status of update operation
type UpdateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have beed updated
	// Equals 1 in case of succesfull update
	Updated int64 `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	// Message: Add V2 API
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{6}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *UpdateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Request data to delete todo task
type DeleteRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier of the todo task to delete
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{7}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains status of delete operation
type DeleteResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have beed deleted
	// Equals 1 in case of succesfull delete
	Deleted int64 `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Message: Add V2 API
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{8}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func (m *DeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Request data to read all todo task
type ReadAllRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{9}
}
func (m *ReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequest.Unmarshal(m, b)
}
func (m *ReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequest.Marshal(b, m, deterministic)
}
func (dst *ReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequest.Merge(dst, src)
}
func (m *ReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequest.Size(m)
}
func (m *ReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequest proto.InternalMessageInfo

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

// Contains list of all todo tasks
type ReadAllResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// List of all todo tasks
	ToDos []*ToDo `protobuf:"bytes,2,rep,name=toDos,proto3" json:"toDos,omitempty"`
	// Message: Add V2 API
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_service_d4e3697a2636aead, []int{10}
}
func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (dst *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(dst, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetToDos() []*ToDo {
	if m != nil {
		return m.ToDos
	}
	return nil
}

func (m *ReadAllResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ToDo)(nil), "v2.ToDo")
	proto.RegisterType((*CreateRequest)(nil), "v2.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v2.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "v2.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "v2.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "v2.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "v2.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "v2.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "v2.DeleteResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "v2.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "v2.ReadAllResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToDoServiceClient interface {
	// Create new todo task
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read todo task
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update todo task
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete todo task
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Read all todo tasks
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type toDoServiceClient struct {
	cc *grpc.ClientConn
}

func NewToDoServiceClient(cc *grpc.ClientConn) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v2.ToDoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v2.ToDoService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/v2.ToDoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/v2.ToDoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/v2.ToDoService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceServer is the server API for ToDoService service.
type ToDoServiceServer interface {
	// Create new todo task
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read todo task
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update todo task
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete todo task
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Read all todo tasks
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

func RegisterToDoServiceServer(s *grpc.Server, srv ToDoServiceServer) {
	s.RegisterService(&_ToDoService_serviceDesc, srv)
}

func _ToDoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.ToDoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.ToDoService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.ToDoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.ToDoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.ToDoService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToDoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v2.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ToDoService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ToDoService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ToDoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ToDoService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _ToDoService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo-service.proto",
}

func init() { proto.RegisterFile("todo-service.proto", fileDescriptor_todo_service_d4e3697a2636aead) }

var fileDescriptor_todo_service_d4e3697a2636aead = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe5, 0xa4, 0x6b, 0xb7, 0x17, 0x9a, 0x0d, 0xc3, 0x21, 0x8a, 0x10, 0x44, 0x3e, 0x95,
	0x03, 0x89, 0x08, 0x12, 0xd7, 0x09, 0xd1, 0x23, 0xa7, 0x30, 0x26, 0x2e, 0x1c, 0xb2, 0xf9, 0x51,
	0x59, 0x4a, 0xeb, 0x10, 0xbb, 0xfd, 0x08, 0x7c, 0x6d, 0x90, 0xed, 0xb8, 0x6b, 0xb2, 0xb5, 0xda,
	0x2d, 0x7e, 0x79, 0xef, 0xff, 0x7e, 0xfe, 0xbf, 0x67, 0xa0, 0x5a, 0x72, 0xf9, 0x41, 0x61, 0xb7,
	0x13, 0xf7, 0x98, 0xb7, 0x9d, 0xd4, 0x92, 0x06, 0xbb, 0x32, 0x7d, 0xb7, 0x92, 0x72, 0xd5, 0x60,
	0x61, 0x23, 0x77, 0xdb, 0xdf, 0x85, 0x16, 0x6b, 0x54, 0xba, 0x5e, 0xb7, 0x2e, 0x89, 0xfd, 0x25,
	0x30, 0xb9, 0x91, 0x4b, 0x49, 0x63, 0x08, 0x04, 0x4f, 0x48, 0x46, 0x16, 0x61, 0x15, 0x08, 0x4e,
	0x5f, 0xc3, 0x99, 0x16, 0xba, 0xc1, 0x24, 0xc8, 0xc8, 0xe2, 0xa2, 0x72, 0x07, 0x9a, 0x41, 0xc4,
	0x51, 0xdd, 0x77, 0xa2, 0xd5, 0x42, 0x6e, 0x92, 0xd0, 0xfe, 0x3b, 0x0c, 0xd1, 0xcf, 0x70, 0xde,
	0xe1, 0x5a, 0x6c, 0x38, 0x76, 0xc9, 0x24, 0x23, 0x8b, 0xa8, 0x4c, 0x73, 0x07, 0x91, 0x7b, 0x88,
	0xfc, 0xc6, 0x43, 0x54, 0xfb, 0x5c, 0x76, 0x0d, 0xf3, 0xaf, 0x1d, 0xd6, 0x1a, 0x2b, 0xfc, 0xb3,
	0x45, 0xa5, 0xe9, 0x15, 0x84, 0x75, 0x2b, 0xfa, 0xf6, 0xe6, 0x93, 0xbe, 0x81, 0x89, 0x96, 0x4b,
	0x69, 0xbb, 0x46, 0xe5, 0x79, 0xbe, 0x2b, 0x73, 0x83, 0x5e, 0xd9, 0x28, 0xfb, 0x06, 0xb1, 0x17,
	0x50, 0xad, 0xdc, 0x28, 0x7c, 0x42, 0xc1, 0x5d, 0x32, 0xdc, 0x5f, 0x32, 0x81, 0xd9, 0x1a, 0x95,
	0xaa, 0x57, 0x68, 0x59, 0x2f, 0x2a, 0x7f, 0x64, 0x05, 0x44, 0x15, 0xd6, 0x7c, 0x04, 0x43, 0xc6,
	0x52, 0x81, 0x97, 0x62, 0x3f, 0xe1, 0x85, 0x2b, 0x18, 0x36, 0x27, 0x8f, 0xf1, 0x83, 0xa7, 0xf0,
	0x0f, 0x51, 0xc2, 0x21, 0xca, 0x35, 0xcc, 0x7f, 0xb4, 0xfc, 0xb1, 0x33, 0xcf, 0x95, 0x66, 0xb7,
	0x10, 0x7b, 0x81, 0xa3, 0x70, 0x09, 0xcc, 0xb6, 0x36, 0xc7, 0xdf, 0xc9, 0x1f, 0x4f, 0x80, 0x7d,
	0x84, 0xf9, 0x12, 0x1b, 0x3c, 0x05, 0x36, 0x76, 0xe9, 0x16, 0x62, 0x5f, 0x72, 0x0a, 0x85, 0xdb,
	0x9c, 0x3d, 0x4a, 0x7f, 0x3c, 0x81, 0xc2, 0x20, 0x36, 0xee, 0x7f, 0x69, 0x9a, 0xa3, 0x2c, 0xec,
	0x17, 0x5c, 0xee, 0x73, 0x8e, 0x36, 0x7f, 0x0b, 0x67, 0xc6, 0x33, 0x95, 0x04, 0x59, 0x38, 0xb0,
	0xd2, 0x85, 0x8f, 0x23, 0x94, 0xff, 0x08, 0x44, 0x26, 0xf3, 0xbb, 0x7b, 0x84, 0xb4, 0x80, 0xa9,
	0xdb, 0x47, 0xfa, 0xd2, 0x88, 0x0c, 0x96, 0x3b, 0xa5, 0x87, 0xa1, 0x1e, 0xe6, 0x3d, 0x4c, 0x0c,
	0x1f, 0xbd, 0x34, 0xff, 0x0e, 0x96, 0x2f, 0xbd, 0x7a, 0x08, 0xf4, 0xa9, 0x05, 0x4c, 0xdd, 0x44,
	0x9d, 0xf6, 0x60, 0x3d, 0x9c, 0xf6, 0x68, 0xe0, 0x05, 0x4c, 0x9d, 0xef, 0xae, 0x60, 0x30, 0x36,
	0x57, 0x30, 0x1a, 0x4b, 0x09, 0xb3, 0xde, 0x2c, 0x4a, 0x7d, 0xfb, 0x07, 0x77, 0xd3, 0x57, 0x83,
	0x98, 0xab, 0xb9, 0x9b, 0xda, 0x07, 0xfe, 0xe9, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xb1,
	0xc5, 0x30, 0x8d, 0x04, 0x00, 0x00,
}
