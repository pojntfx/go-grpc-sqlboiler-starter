// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todos.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// TodoChangeType is a change type
type TodoChangeType int32

const (
	TodoChangeType_CREATE TodoChangeType = 0
	TodoChangeType_UPDATE TodoChangeType = 1
	TodoChangeType_DELETE TodoChangeType = 2
)

var TodoChangeType_name = map[int32]string{
	0: "CREATE",
	1: "UPDATE",
	2: "DELETE",
}

var TodoChangeType_value = map[string]int32{
	"CREATE": 0,
	"UPDATE": 1,
	"DELETE": 2,
}

func (x TodoChangeType) String() string {
	return proto.EnumName(TodoChangeType_name, int32(x))
}

func (TodoChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{0}
}

// NewTodo is a new todo
type NewTodo struct {
	Title                string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewTodo) Reset()         { *m = NewTodo{} }
func (m *NewTodo) String() string { return proto.CompactTextString(m) }
func (*NewTodo) ProtoMessage()    {}
func (*NewTodo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{0}
}

func (m *NewTodo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTodo.Unmarshal(m, b)
}
func (m *NewTodo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTodo.Marshal(b, m, deterministic)
}
func (m *NewTodo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTodo.Merge(m, src)
}
func (m *NewTodo) XXX_Size() int {
	return xxx_messageInfo_NewTodo.Size(m)
}
func (m *NewTodo) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTodo.DiscardUnknown(m)
}

var xxx_messageInfo_NewTodo proto.InternalMessageInfo

func (m *NewTodo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NewTodo) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// Todo is a todo
type Todo struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
	Index                int64    `protobuf:"varint,4,opt,name=Index,proto3" json:"Index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{1}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Todo) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// TodoList is a list of todos
type TodoList struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=Todos,proto3" json:"Todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoList) Reset()         { *m = TodoList{} }
func (m *TodoList) String() string { return proto.CompactTextString(m) }
func (*TodoList) ProtoMessage()    {}
func (*TodoList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{2}
}

func (m *TodoList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoList.Unmarshal(m, b)
}
func (m *TodoList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoList.Marshal(b, m, deterministic)
}
func (m *TodoList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoList.Merge(m, src)
}
func (m *TodoList) XXX_Size() int {
	return xxx_messageInfo_TodoList.Size(m)
}
func (m *TodoList) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoList.DiscardUnknown(m)
}

var xxx_messageInfo_TodoList proto.InternalMessageInfo

func (m *TodoList) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

// TodoID is a todo's ID
type TodoID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoID) Reset()         { *m = TodoID{} }
func (m *TodoID) String() string { return proto.CompactTextString(m) }
func (*TodoID) ProtoMessage()    {}
func (*TodoID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{3}
}

func (m *TodoID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoID.Unmarshal(m, b)
}
func (m *TodoID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoID.Marshal(b, m, deterministic)
}
func (m *TodoID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoID.Merge(m, src)
}
func (m *TodoID) XXX_Size() int {
	return xxx_messageInfo_TodoID.Size(m)
}
func (m *TodoID) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoID.DiscardUnknown(m)
}

var xxx_messageInfo_TodoID proto.InternalMessageInfo

func (m *TodoID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

// TodoReorder is a request to reorder a todo
type TodoReorder struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoReorder) Reset()         { *m = TodoReorder{} }
func (m *TodoReorder) String() string { return proto.CompactTextString(m) }
func (*TodoReorder) ProtoMessage()    {}
func (*TodoReorder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{4}
}

func (m *TodoReorder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoReorder.Unmarshal(m, b)
}
func (m *TodoReorder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoReorder.Marshal(b, m, deterministic)
}
func (m *TodoReorder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoReorder.Merge(m, src)
}
func (m *TodoReorder) XXX_Size() int {
	return xxx_messageInfo_TodoReorder.Size(m)
}
func (m *TodoReorder) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoReorder.DiscardUnknown(m)
}

var xxx_messageInfo_TodoReorder proto.InternalMessageInfo

func (m *TodoReorder) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TodoReorder) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

// TodoChange is a change with the changed todo
type TodoChange struct {
	Type                 TodoChangeType `protobuf:"varint,1,opt,name=type,proto3,enum=todos.TodoChangeType" json:"type,omitempty"`
	Todo                 *Todo          `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TodoChange) Reset()         { *m = TodoChange{} }
func (m *TodoChange) String() string { return proto.CompactTextString(m) }
func (*TodoChange) ProtoMessage()    {}
func (*TodoChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{5}
}

func (m *TodoChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoChange.Unmarshal(m, b)
}
func (m *TodoChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoChange.Marshal(b, m, deterministic)
}
func (m *TodoChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoChange.Merge(m, src)
}
func (m *TodoChange) XXX_Size() int {
	return xxx_messageInfo_TodoChange.Size(m)
}
func (m *TodoChange) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoChange.DiscardUnknown(m)
}

var xxx_messageInfo_TodoChange proto.InternalMessageInfo

func (m *TodoChange) GetType() TodoChangeType {
	if m != nil {
		return m.Type
	}
	return TodoChangeType_CREATE
}

func (m *TodoChange) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterEnum("todos.TodoChangeType", TodoChangeType_name, TodoChangeType_value)
	proto.RegisterType((*NewTodo)(nil), "todos.NewTodo")
	proto.RegisterType((*Todo)(nil), "todos.Todo")
	proto.RegisterType((*TodoList)(nil), "todos.TodoList")
	proto.RegisterType((*TodoID)(nil), "todos.TodoID")
	proto.RegisterType((*TodoReorder)(nil), "todos.TodoReorder")
	proto.RegisterType((*TodoChange)(nil), "todos.TodoChange")
}

func init() {
	proto.RegisterFile("todos.proto", fileDescriptor_a7703e18c08e8710)
}

var fileDescriptor_a7703e18c08e8710 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xed, 0x6e, 0xda, 0x40,
	0x10, 0xac, 0x3f, 0x69, 0xd7, 0xaa, 0x4b, 0x4f, 0x2d, 0xb2, 0xe8, 0x8f, 0x52, 0xab, 0x95, 0x28,
	0x52, 0x4d, 0x05, 0xed, 0x03, 0x00, 0xb6, 0x2a, 0x4b, 0xa8, 0x89, 0x1c, 0x13, 0x45, 0xf9, 0x07,
	0xf1, 0x42, 0x90, 0x08, 0x67, 0xd9, 0x87, 0x12, 0x3f, 0x5b, 0x5e, 0x2e, 0xba, 0x3b, 0x23, 0x6c,
	0x48, 0xf2, 0xcb, 0xbb, 0xb3, 0xb3, 0xe3, 0xbd, 0x19, 0xb0, 0x18, 0x4d, 0x68, 0xee, 0xa5, 0x19,
	0x65, 0x94, 0x18, 0xa2, 0x69, 0x7f, 0x59, 0x51, 0xba, 0xda, 0x60, 0x5f, 0x80, 0x8b, 0xdd, 0xb2,
	0x8f, 0x77, 0x29, 0x2b, 0x24, 0xc7, 0x1d, 0x42, 0xe3, 0x3f, 0xde, 0xc7, 0x34, 0xa1, 0xe4, 0x13,
	0x18, 0xf1, 0x9a, 0x6d, 0xd0, 0x51, 0x3a, 0x4a, 0xf7, 0x5d, 0x24, 0x1b, 0x42, 0x40, 0x1f, 0xd3,
	0xa4, 0x70, 0x54, 0x01, 0x8a, 0xda, 0xbd, 0x04, 0x5d, 0x6c, 0xd8, 0xa0, 0x86, 0xbe, 0xa0, 0x6b,
	0x91, 0x1a, 0xfa, 0x07, 0x05, 0xf5, 0x39, 0x05, 0xed, 0xa0, 0xc0, 0x99, 0xe1, 0x36, 0xc1, 0x07,
	0x47, 0x17, 0xcb, 0xb2, 0x71, 0x7f, 0xc1, 0x5b, 0xae, 0x3b, 0x5d, 0xe7, 0x8c, 0x7c, 0x03, 0x83,
	0xd7, 0xb9, 0xa3, 0x74, 0xb4, 0xae, 0x35, 0xb0, 0x3c, 0xf9, 0x32, 0x8e, 0x45, 0x72, 0xe2, 0x3a,
	0x60, 0xf2, 0x22, 0xf4, 0x8f, 0x0f, 0x71, 0xff, 0x82, 0x25, 0x88, 0x48, 0xb3, 0x04, 0xb3, 0x93,
	0x3b, 0x5b, 0x60, 0x9e, 0x2d, 0x97, 0x39, 0x32, 0x71, 0xa8, 0x16, 0x95, 0x9d, 0x7b, 0x05, 0xc0,
	0xd7, 0x26, 0xb7, 0xf3, 0xed, 0x0a, 0xc9, 0x4f, 0xd0, 0x59, 0x91, 0x4a, 0x3b, 0xec, 0xc1, 0xe7,
	0xca, 0x01, 0x92, 0x10, 0x17, 0x29, 0x46, 0x82, 0x42, 0xbe, 0x82, 0xce, 0xa7, 0x42, 0xee, 0xe8,
	0x56, 0x31, 0xe8, 0xfd, 0x01, 0xbb, 0xbe, 0x48, 0x00, 0xcc, 0x49, 0x14, 0x8c, 0xe2, 0xa0, 0xf9,
	0x86, 0xd7, 0xb3, 0x73, 0x9f, 0xd7, 0x0a, 0xaf, 0xfd, 0x60, 0x1a, 0xc4, 0x41, 0x53, 0x1d, 0x3c,
	0xaa, 0xa5, 0x09, 0xe4, 0x07, 0x98, 0x93, 0x0c, 0xe7, 0x0c, 0x89, 0x5d, 0x8a, 0x97, 0xa9, 0xb5,
	0xab, 0x3f, 0x23, 0x7d, 0xd0, 0x85, 0x79, 0x2d, 0x4f, 0x66, 0xee, 0xed, 0x33, 0xf7, 0x02, 0x9e,
	0x79, 0xfb, 0x43, 0x85, 0x5c, 0xba, 0xac, 0xfd, 0x43, 0x46, 0xde, 0x57, 0xf0, 0xd0, 0xaf, 0x6b,
	0xba, 0x60, 0xce, 0xd2, 0x84, 0xff, 0xba, 0x0a, 0xd7, 0x39, 0xdf, 0xc1, 0xf4, 0x71, 0x83, 0x0c,
	0x5f, 0x55, 0xea, 0x41, 0x63, 0x9f, 0x08, 0xa9, 0x5a, 0x24, 0xb1, 0x3a, 0x77, 0x04, 0xe4, 0x62,
	0xb7, 0xc8, 0x6f, 0xb2, 0xf5, 0x02, 0xe3, 0xd2, 0xb7, 0xfc, 0xc5, 0x77, 0x7d, 0x3c, 0x09, 0xe7,
	0xb7, 0x32, 0x6e, 0x5c, 0x1b, 0x92, 0x66, 0x8a, 0xcf, 0xf0, 0x29, 0x00, 0x00, 0xff, 0xff, 0x39,
	0x58, 0x13, 0x81, 0x1d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodosClient is the client API for Todos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodosClient interface {
	// Create creates a todo
	Create(ctx context.Context, in *NewTodo, opts ...grpc.CallOption) (*Todo, error)
	// List lists all todos
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TodoList, error)
	// Get gets one todo
	Get(ctx context.Context, in *TodoID, opts ...grpc.CallOption) (*Todo, error)
	// Update updates one todo
	// Index can't be updated with the `Update` rpc, use the `Reorder` rpc
	Update(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	// Delete deletes one todo
	Delete(ctx context.Context, in *TodoID, opts ...grpc.CallOption) (*Todo, error)
	// Reorder reorders a todo
	Reorder(ctx context.Context, in *TodoReorder, opts ...grpc.CallOption) (*Todo, error)
	// SubscribeToChanges subscribes to all changes
	SubscribeToChanges(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Todos_SubscribeToChangesClient, error)
}

type todosClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosClient(cc grpc.ClientConnInterface) TodosClient {
	return &todosClient{cc}
}

func (c *todosClient) Create(ctx context.Context, in *NewTodo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todos.Todos/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TodoList, error) {
	out := new(TodoList)
	err := c.cc.Invoke(ctx, "/todos.Todos/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Get(ctx context.Context, in *TodoID, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todos.Todos/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Update(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todos.Todos/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Delete(ctx context.Context, in *TodoID, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todos.Todos/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Reorder(ctx context.Context, in *TodoReorder, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todos.Todos/Reorder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) SubscribeToChanges(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Todos_SubscribeToChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Todos_serviceDesc.Streams[0], "/todos.Todos/SubscribeToChanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &todosSubscribeToChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Todos_SubscribeToChangesClient interface {
	Recv() (*TodoChange, error)
	grpc.ClientStream
}

type todosSubscribeToChangesClient struct {
	grpc.ClientStream
}

func (x *todosSubscribeToChangesClient) Recv() (*TodoChange, error) {
	m := new(TodoChange)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodosServer is the server API for Todos service.
type TodosServer interface {
	// Create creates a todo
	Create(context.Context, *NewTodo) (*Todo, error)
	// List lists all todos
	List(context.Context, *empty.Empty) (*TodoList, error)
	// Get gets one todo
	Get(context.Context, *TodoID) (*Todo, error)
	// Update updates one todo
	// Index can't be updated with the `Update` rpc, use the `Reorder` rpc
	Update(context.Context, *Todo) (*Todo, error)
	// Delete deletes one todo
	Delete(context.Context, *TodoID) (*Todo, error)
	// Reorder reorders a todo
	Reorder(context.Context, *TodoReorder) (*Todo, error)
	// SubscribeToChanges subscribes to all changes
	SubscribeToChanges(*empty.Empty, Todos_SubscribeToChangesServer) error
}

// UnimplementedTodosServer can be embedded to have forward compatible implementations.
type UnimplementedTodosServer struct {
}

func (*UnimplementedTodosServer) Create(ctx context.Context, req *NewTodo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTodosServer) List(ctx context.Context, req *empty.Empty) (*TodoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTodosServer) Get(ctx context.Context, req *TodoID) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTodosServer) Update(ctx context.Context, req *Todo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTodosServer) Delete(ctx context.Context, req *TodoID) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedTodosServer) Reorder(ctx context.Context, req *TodoReorder) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reorder not implemented")
}
func (*UnimplementedTodosServer) SubscribeToChanges(req *empty.Empty, srv Todos_SubscribeToChangesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToChanges not implemented")
}

func RegisterTodosServer(s *grpc.Server, srv TodosServer) {
	s.RegisterService(&_Todos_serviceDesc, srv)
}

func _Todos_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTodo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Create(ctx, req.(*NewTodo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Get(ctx, req.(*TodoID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Update(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Delete(ctx, req.(*TodoID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Reorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoReorder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Reorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/Reorder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Reorder(ctx, req.(*TodoReorder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_SubscribeToChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodosServer).SubscribeToChanges(m, &todosSubscribeToChangesServer{stream})
}

type Todos_SubscribeToChangesServer interface {
	Send(*TodoChange) error
	grpc.ServerStream
}

type todosSubscribeToChangesServer struct {
	grpc.ServerStream
}

func (x *todosSubscribeToChangesServer) Send(m *TodoChange) error {
	return x.ServerStream.SendMsg(m)
}

var _Todos_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todos.Todos",
	HandlerType: (*TodosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Todos_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Todos_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Todos_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Todos_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Todos_Delete_Handler,
		},
		{
			MethodName: "Reorder",
			Handler:    _Todos_Reorder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToChanges",
			Handler:       _Todos_SubscribeToChanges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todos.proto",
}
