// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deviceQueue.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EnqueueDeviceQueueItemRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Random reference (used on ack notification).
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
	// Is an ACK required from the node.
	Confirmed bool `protobuf:"varint,3,opt,name=confirmed" json:"confirmed,omitempty"`
	// FPort used (must be >0)
	FPort uint32 `protobuf:"varint,4,opt,name=fPort" json:"fPort,omitempty"`
	// Base64 encoded data.
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EnqueueDeviceQueueItemRequest) Reset()                    { *m = EnqueueDeviceQueueItemRequest{} }
func (m *EnqueueDeviceQueueItemRequest) String() string            { return proto.CompactTextString(m) }
func (*EnqueueDeviceQueueItemRequest) ProtoMessage()               {}
func (*EnqueueDeviceQueueItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EnqueueDeviceQueueItemRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *EnqueueDeviceQueueItemRequest) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *EnqueueDeviceQueueItemRequest) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *EnqueueDeviceQueueItemRequest) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *EnqueueDeviceQueueItemRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type EnqueueDeviceQueueItemResponse struct {
}

func (m *EnqueueDeviceQueueItemResponse) Reset()                    { *m = EnqueueDeviceQueueItemResponse{} }
func (m *EnqueueDeviceQueueItemResponse) String() string            { return proto.CompactTextString(m) }
func (*EnqueueDeviceQueueItemResponse) ProtoMessage()               {}
func (*EnqueueDeviceQueueItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type FlushDeviceQueueRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *FlushDeviceQueueRequest) Reset()                    { *m = FlushDeviceQueueRequest{} }
func (m *FlushDeviceQueueRequest) String() string            { return proto.CompactTextString(m) }
func (*FlushDeviceQueueRequest) ProtoMessage()               {}
func (*FlushDeviceQueueRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *FlushDeviceQueueRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

type FlushDeviceQueueResponse struct {
}

func (m *FlushDeviceQueueResponse) Reset()                    { *m = FlushDeviceQueueResponse{} }
func (m *FlushDeviceQueueResponse) String() string            { return proto.CompactTextString(m) }
func (*FlushDeviceQueueResponse) ProtoMessage()               {}
func (*FlushDeviceQueueResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type DeviceQueueItem struct {
	// Hex encoded DevEUI of the device.
	DevEUI string `protobuf:"bytes,2,opt,name=devEUI" json:"devEUI,omitempty"`
	// Random reference (used on ack notification).
	Reference string `protobuf:"bytes,3,opt,name=reference" json:"reference,omitempty"`
	// Is an ACK required from the device.
	Confirmed bool `protobuf:"varint,4,opt,name=confirmed" json:"confirmed,omitempty"`
	// FPort used (must be >0).
	FPort uint32 `protobuf:"varint,6,opt,name=fPort" json:"fPort,omitempty"`
	// Base64 encoded data.
	Data []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	// FCnt of the queue item.
	FCnt uint32 `protobuf:"varint,8,opt,name=fCnt" json:"fCnt,omitempty"`
}

func (m *DeviceQueueItem) Reset()                    { *m = DeviceQueueItem{} }
func (m *DeviceQueueItem) String() string            { return proto.CompactTextString(m) }
func (*DeviceQueueItem) ProtoMessage()               {}
func (*DeviceQueueItem) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *DeviceQueueItem) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *DeviceQueueItem) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *DeviceQueueItem) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *DeviceQueueItem) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *DeviceQueueItem) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DeviceQueueItem) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

type ListDeviceQueueItemsRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *ListDeviceQueueItemsRequest) Reset()                    { *m = ListDeviceQueueItemsRequest{} }
func (m *ListDeviceQueueItemsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDeviceQueueItemsRequest) ProtoMessage()               {}
func (*ListDeviceQueueItemsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ListDeviceQueueItemsRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

type ListDeviceQueueItemsResponse struct {
	Items []*DeviceQueueItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListDeviceQueueItemsResponse) Reset()                    { *m = ListDeviceQueueItemsResponse{} }
func (m *ListDeviceQueueItemsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDeviceQueueItemsResponse) ProtoMessage()               {}
func (*ListDeviceQueueItemsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *ListDeviceQueueItemsResponse) GetItems() []*DeviceQueueItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*EnqueueDeviceQueueItemRequest)(nil), "api.EnqueueDeviceQueueItemRequest")
	proto.RegisterType((*EnqueueDeviceQueueItemResponse)(nil), "api.EnqueueDeviceQueueItemResponse")
	proto.RegisterType((*FlushDeviceQueueRequest)(nil), "api.FlushDeviceQueueRequest")
	proto.RegisterType((*FlushDeviceQueueResponse)(nil), "api.FlushDeviceQueueResponse")
	proto.RegisterType((*DeviceQueueItem)(nil), "api.DeviceQueueItem")
	proto.RegisterType((*ListDeviceQueueItemsRequest)(nil), "api.ListDeviceQueueItemsRequest")
	proto.RegisterType((*ListDeviceQueueItemsResponse)(nil), "api.ListDeviceQueueItemsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeviceQueue service

type DeviceQueueClient interface {
	// Enqueue adds the given item to the device-queue.
	Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error)
	// Flush flushes the downlink device-queue.
	Flush(ctx context.Context, in *FlushDeviceQueueRequest, opts ...grpc.CallOption) (*FlushDeviceQueueResponse, error)
	// List lists the items in the device-queue.
	List(ctx context.Context, in *ListDeviceQueueItemsRequest, opts ...grpc.CallOption) (*ListDeviceQueueItemsResponse, error)
}

type deviceQueueClient struct {
	cc *grpc.ClientConn
}

func NewDeviceQueueClient(cc *grpc.ClientConn) DeviceQueueClient {
	return &deviceQueueClient{cc}
}

func (c *deviceQueueClient) Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error) {
	out := new(EnqueueDeviceQueueItemResponse)
	err := grpc.Invoke(ctx, "/api.DeviceQueue/Enqueue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceQueueClient) Flush(ctx context.Context, in *FlushDeviceQueueRequest, opts ...grpc.CallOption) (*FlushDeviceQueueResponse, error) {
	out := new(FlushDeviceQueueResponse)
	err := grpc.Invoke(ctx, "/api.DeviceQueue/Flush", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceQueueClient) List(ctx context.Context, in *ListDeviceQueueItemsRequest, opts ...grpc.CallOption) (*ListDeviceQueueItemsResponse, error) {
	out := new(ListDeviceQueueItemsResponse)
	err := grpc.Invoke(ctx, "/api.DeviceQueue/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceQueue service

type DeviceQueueServer interface {
	// Enqueue adds the given item to the device-queue.
	Enqueue(context.Context, *EnqueueDeviceQueueItemRequest) (*EnqueueDeviceQueueItemResponse, error)
	// Flush flushes the downlink device-queue.
	Flush(context.Context, *FlushDeviceQueueRequest) (*FlushDeviceQueueResponse, error)
	// List lists the items in the device-queue.
	List(context.Context, *ListDeviceQueueItemsRequest) (*ListDeviceQueueItemsResponse, error)
}

func RegisterDeviceQueueServer(s *grpc.Server, srv DeviceQueueServer) {
	s.RegisterService(&_DeviceQueue_serviceDesc, srv)
}

func _DeviceQueue_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueDeviceQueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceQueue/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServer).Enqueue(ctx, req.(*EnqueueDeviceQueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceQueue_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushDeviceQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceQueue/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServer).Flush(ctx, req.(*FlushDeviceQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceQueue_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceQueueItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceQueue/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServer).List(ctx, req.(*ListDeviceQueueItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceQueue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceQueue",
	HandlerType: (*DeviceQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _DeviceQueue_Enqueue_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _DeviceQueue_Flush_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceQueue_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deviceQueue.proto",
}

func init() { proto.RegisterFile("deviceQueue.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcf, 0xea, 0xd3, 0x40,
	0x10, 0x66, 0xf3, 0xa7, 0xbf, 0x5f, 0xa7, 0x8a, 0xb8, 0x14, 0x5d, 0xd2, 0x54, 0xd6, 0x14, 0x24,
	0xf4, 0xd0, 0x60, 0xc5, 0x8b, 0x57, 0xad, 0x50, 0xf1, 0xa0, 0x01, 0x1f, 0x20, 0x36, 0x9b, 0xba,
	0xd0, 0xee, 0xa6, 0xd9, 0x4d, 0x0f, 0x8a, 0x17, 0x5f, 0xc1, 0xb3, 0x27, 0x1f, 0xc9, 0x57, 0x10,
	0x9f, 0x43, 0xb2, 0x09, 0xa4, 0xd6, 0x26, 0xbd, 0x65, 0x67, 0xbe, 0x6f, 0xbe, 0x99, 0x6f, 0x26,
	0x70, 0x3f, 0x65, 0x47, 0xbe, 0x61, 0xef, 0x4b, 0x56, 0xb2, 0x45, 0x5e, 0x48, 0x2d, 0xb1, 0x9d,
	0xe4, 0xdc, 0xf3, 0xb7, 0x52, 0x6e, 0x77, 0x2c, 0x4a, 0x72, 0x1e, 0x25, 0x42, 0x48, 0x9d, 0x68,
	0x2e, 0x85, 0xaa, 0x21, 0xc1, 0x0f, 0x04, 0xd3, 0x95, 0x38, 0x54, 0xa4, 0x57, 0x2d, 0x7f, 0xad,
	0xd9, 0x3e, 0x66, 0x87, 0x92, 0x29, 0x8d, 0x1f, 0xc0, 0x20, 0x65, 0xc7, 0xd5, 0x87, 0x35, 0x41,
	0x14, 0x85, 0xc3, 0xb8, 0x79, 0x61, 0x1f, 0x86, 0x05, 0xcb, 0x58, 0xc1, 0xc4, 0x86, 0x11, 0xcb,
	0xa4, 0xda, 0x40, 0x95, 0xdd, 0x48, 0x91, 0xf1, 0x62, 0xcf, 0x52, 0x62, 0x53, 0x14, 0xde, 0xc6,
	0x6d, 0x00, 0x8f, 0xc1, 0xcd, 0xde, 0xc9, 0x42, 0x13, 0x87, 0xa2, 0xf0, 0x6e, 0x5c, 0x3f, 0x30,
	0x06, 0x27, 0x4d, 0x74, 0x42, 0x5c, 0x8a, 0xc2, 0x3b, 0xb1, 0xf9, 0x0e, 0x28, 0x3c, 0xea, 0x6a,
	0x4f, 0xe5, 0x52, 0x28, 0x16, 0x3c, 0x85, 0x87, 0xaf, 0x77, 0xa5, 0xfa, 0x74, 0x92, 0xbf, 0xd2,
	0x7a, 0xe0, 0x01, 0xf9, 0x9f, 0xd2, 0x94, 0xfb, 0x89, 0xe0, 0xde, 0x99, 0xd4, 0x49, 0x1d, 0xab,
	0xdb, 0x02, 0xbb, 0xd7, 0x02, 0xa7, 0xd3, 0x82, 0xc1, 0x25, 0x0b, 0x6e, 0x5a, 0x0b, 0xaa, 0x58,
	0xf6, 0x52, 0x68, 0x72, 0x6b, 0x80, 0xe6, 0x3b, 0x78, 0x0e, 0x93, 0xb7, 0x5c, 0xe9, 0xb3, 0x46,
	0xd5, 0xb5, 0xc1, 0xdf, 0x80, 0x7f, 0x99, 0x56, 0x0f, 0x8f, 0xe7, 0xe0, 0xf2, 0x2a, 0x40, 0x10,
	0xb5, 0xc3, 0xd1, 0x72, 0xbc, 0x48, 0x72, 0xbe, 0x38, 0x37, 0xbe, 0x86, 0x2c, 0xff, 0x58, 0x30,
	0x3a, 0x49, 0xe1, 0xcf, 0x70, 0xd3, 0x6c, 0x0a, 0x07, 0x86, 0xd7, 0x7b, 0x56, 0xde, 0xac, 0x17,
	0xd3, 0x2c, 0xe3, 0xc9, 0xb7, 0x5f, 0xbf, 0xbf, 0x5b, 0x34, 0x98, 0x98, 0xeb, 0xad, 0x0f, 0x5c,
	0x45, 0x5f, 0xea, 0x69, 0xbe, 0x46, 0x86, 0xfb, 0x02, 0xcd, 0x31, 0x07, 0xd7, 0x2c, 0x14, 0xfb,
	0xa6, 0x6a, 0xc7, 0x3d, 0x78, 0xd3, 0x8e, 0x6c, 0xa3, 0x36, 0x33, 0x6a, 0xd3, 0x79, 0x9f, 0x1a,
	0xce, 0xc1, 0xa9, 0x2c, 0xc4, 0xd4, 0xd4, 0xea, 0x59, 0x82, 0xf7, 0xb8, 0x07, 0xf1, 0xaf, 0x22,
	0xee, 0x53, 0xfc, 0x38, 0x30, 0x7f, 0xea, 0xb3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x02, 0x00,
	0xe7, 0xf2, 0xe1, 0x03, 0x00, 0x00,
}
