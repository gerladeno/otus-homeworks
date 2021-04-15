// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package eventspb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventsHandlerClient is the client API for EventsHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventsHandlerClient interface {
	ListEvents(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Events, error)
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Id, error)
	UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*empty.Empty, error)
	ReadEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Event, error)
}

type eventsHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsHandlerClient(cc grpc.ClientConnInterface) EventsHandlerClient {
	return &eventsHandlerClient{cc}
}

func (c *eventsHandlerClient) ListEvents(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, "/eventspb.EventsHandler/ListEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsHandlerClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/eventspb.EventsHandler/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsHandlerClient) UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/eventspb.EventsHandler/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsHandlerClient) DeleteEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/eventspb.EventsHandler/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsHandlerClient) ReadEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/eventspb.EventsHandler/ReadEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsHandlerServer is the server API for EventsHandler service.
// All implementations should embed UnimplementedEventsHandlerServer
// for forward compatibility
type EventsHandlerServer interface {
	ListEvents(context.Context, *empty.Empty) (*Events, error)
	CreateEvent(context.Context, *Event) (*Id, error)
	UpdateEvent(context.Context, *Event) (*empty.Empty, error)
	DeleteEvent(context.Context, *Id) (*empty.Empty, error)
	ReadEvent(context.Context, *Id) (*Event, error)
}

// UnimplementedEventsHandlerServer should be embedded to have forward compatible implementations.
type UnimplementedEventsHandlerServer struct {
}

func (UnimplementedEventsHandlerServer) ListEvents(context.Context, *empty.Empty) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (UnimplementedEventsHandlerServer) CreateEvent(context.Context, *Event) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedEventsHandlerServer) UpdateEvent(context.Context, *Event) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEventsHandlerServer) DeleteEvent(context.Context, *Id) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventsHandlerServer) ReadEvent(context.Context, *Id) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEvent not implemented")
}

// UnsafeEventsHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventsHandlerServer will
// result in compilation errors.
type UnsafeEventsHandlerServer interface {
	mustEmbedUnimplementedEventsHandlerServer()
}

func RegisterEventsHandlerServer(s grpc.ServiceRegistrar, srv EventsHandlerServer) {
	s.RegisterService(&EventsHandler_ServiceDesc, srv)
}

func _EventsHandler_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsHandlerServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventspb.EventsHandler/ListEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsHandlerServer).ListEvents(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsHandler_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsHandlerServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventspb.EventsHandler/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsHandlerServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsHandler_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsHandlerServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventspb.EventsHandler/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsHandlerServer).UpdateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsHandler_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsHandlerServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventspb.EventsHandler/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsHandlerServer).DeleteEvent(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsHandler_ReadEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsHandlerServer).ReadEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventspb.EventsHandler/ReadEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsHandlerServer).ReadEvent(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// EventsHandler_ServiceDesc is the grpc.ServiceDesc for EventsHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventsHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eventspb.EventsHandler",
	HandlerType: (*EventsHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEvents",
			Handler:    _EventsHandler_ListEvents_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _EventsHandler_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _EventsHandler_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _EventsHandler_DeleteEvent_Handler,
		},
		{
			MethodName: "ReadEvent",
			Handler:    _EventsHandler_ReadEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events.proto",
}