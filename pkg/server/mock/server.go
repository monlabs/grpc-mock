package mock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/monlabs/grpc-mock/pkg/stub"
)

type StubMatcher interface {
	FindStubs(service, method string, fields map[string]interface{}) []*stub.Stub
}

type Server struct {
	addr    string
	svr     *grpc.Server
	matcher StubMatcher
	wg      sync.WaitGroup
}

func NewServer(addr string, m StubMatcher) *Server {
	return &Server{
		addr:    addr,
		svr:     grpc.NewServer(),
		matcher: m,
	}
}

func (s *Server) RegisterServices(fds []*desc.FileDescriptor) {
	sds := s.createGRPCServiceDesc(fds)
	s.registerServices(sds)
}

func (s *Server) registerServices(sds []*grpc.ServiceDesc) {
	for _, sd := range sds {
		s.svr.RegisterService(sd, nil)
	}
}

func (s *Server) Start() (err error) {
	var lsn net.Listener
	lsn, err = net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("mock server: listen on '%s' failed: %v", s.addr, err)
	}
	log.Infof("mock server starts on %v", lsn.Addr().String())
	reflection.Register(s.svr)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		err = s.svr.Serve(lsn)
	}()
	return
}

func (s *Server) Stop() (err error) {
	done := make(chan int, 1)
	go func() {
		s.svr.GracefulStop()
		close(done)
	}()
	t := time.NewTimer(3 * time.Second)
	select {
	case <-done:
		if !t.Stop() {
			<-t.C
		}
	case <-t.C:
		s.svr.Stop()
		err = errors.New("gracefully stop grpc server timeout")
	}
	return
}

func (s *Server) createUnaryServerHandler(md *desc.MethodDescriptor) func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	return func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
		msgFactory := dynamic.NewMessageFactoryWithDefaults()
		in := msgFactory.NewMessage(md.GetInputType())
		if err := dec(in); err != nil {
			return nil, err
		}

		var (
			out proto.Message
			err error
		)
		data, _ := json.Marshal(in)
		m := make(map[string]interface{})
		json.Unmarshal(data, &m)
		expects := s.matcher.FindStubs(md.GetService().GetFullyQualifiedName(), md.GetName(), m)
		if len(expects) == 0 {
			return nil, status.Error(codes.NotFound, "didn't match any stub")
		}
		expect := expects[0]
		if expect.Out.Data != nil {
			out = msgFactory.NewMessage(md.GetOutputType())
			outData, _ := json.Marshal(expect.Out.Data)
			err = json.Unmarshal(outData, out)
			if err != nil {
				return nil, err
			}
		}

		if expect.Out.Message != "" {
			err = status.Error(codes.Code(expect.Out.Code), expect.Out.Message)
		}
		return out, err
	}
}

func (s *Server) createStreamServerHandler(md *desc.MethodDescriptor) func(srv interface{}, stream grpc.ServerStream) error {
	return func(srv interface{}, stream grpc.ServerStream) error {
		msgFactory := dynamic.NewMessageFactoryWithDefaults()
		in := msgFactory.NewMessage(md.GetInputType())
		if err := stream.RecvMsg(in); err != nil {
			return err
		}

		var (
			out proto.Message
			err error
		)
		data, _ := json.Marshal(in)
		m := make(map[string]interface{})
		json.Unmarshal(data, &m)
		expects := s.matcher.FindStubs(md.GetService().GetFullyQualifiedName(), md.GetName(), m)
		if len(expects) == 0 {
			return status.Error(codes.NotFound, "didn't match any stub")
		}
		expect := expects[0]
		if expect.Out.Data != nil {
			out = msgFactory.NewMessage(md.GetOutputType())
			outData, _ := json.Marshal(expect.Out.Data)
			err = json.Unmarshal(outData, out)
			if err != nil {
				return err
			}

			return stream.SendMsg(out)
		}

		if expect.Out.Message != "" {
			return status.Error(codes.Code(expect.Out.Code), expect.Out.Message)
		}

		return nil
	}
}

func (s *Server) createClientStreamServerHandler(md *desc.MethodDescriptor) func(srv interface{}, stream grpc.ServerStream) error {
	return func(srv interface{}, stream grpc.ServerStream) error {
		for {
			msgFactory := dynamic.NewMessageFactoryWithDefaults()
			in := msgFactory.NewMessage(md.GetInputType())
			err := stream.RecvMsg(in)
			if err == io.EOF {

				var (
					out proto.Message
					err error
				)
				data, _ := json.Marshal(in)
				m := make(map[string]interface{})
				json.Unmarshal(data, &m)
				expects := s.matcher.FindStubs(md.GetService().GetFullyQualifiedName(), md.GetName(), m)
				if len(expects) == 0 {
					return status.Error(codes.NotFound, "didn't match any stub")
				}
				expect := expects[0]
				if expect.Out.Data != nil {
					out = msgFactory.NewMessage(md.GetOutputType())
					outData, _ := json.Marshal(expect.Out.Data)
					err = json.Unmarshal(outData, out)
					if err != nil {
						return err
					}

					return stream.SendMsg(out)
				}

				if expect.Out.Message != "" {
					return status.Error(codes.Code(expect.Out.Code), expect.Out.Message)
				}

			}

			if err != nil {
				return err
			}
		}
	}
}

func (s *Server) createServerStreamServerHandler(md *desc.MethodDescriptor) func(srv interface{}, stream grpc.ServerStream) error {
	return func(srv interface{}, stream grpc.ServerStream) error {
		msgFactory := dynamic.NewMessageFactoryWithDefaults()
		in := msgFactory.NewMessage(md.GetInputType())
		if err := stream.RecvMsg(in); err != nil {
			return err
		}

		var (
			out proto.Message
			err error
		)
		data, _ := json.Marshal(in)
		m := make(map[string]interface{})
		json.Unmarshal(data, &m)
		expects := s.matcher.FindStubs(md.GetService().GetFullyQualifiedName(), md.GetName(), m)
		if len(expects) == 0 {
			return status.Error(codes.NotFound, "didn't match any stub")
		}
		expect := expects[0]
		if expect.Out.Data != nil {
			out = msgFactory.NewMessage(md.GetOutputType())
			outData, _ := json.Marshal(expect.Out.Data)
			err = json.Unmarshal(outData, out)
			if err != nil {
				return err
			}

			return stream.SendMsg(out)
		}

		if expect.Out.Message != "" {
			return status.Error(codes.Code(expect.Out.Code), expect.Out.Message)
		}

		return nil
	}
}

func (s *Server) createBidiStreamServerHandler(md *desc.MethodDescriptor) func(srv interface{}, stream grpc.ServerStream) error {
	return func(srv interface{}, stream grpc.ServerStream) error {
		for {
			msgFactory := dynamic.NewMessageFactoryWithDefaults()
			in := msgFactory.NewMessage(md.GetInputType())
			err := stream.RecvMsg(in)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			var (
				out proto.Message
			)
			data, _ := json.Marshal(in)
			m := make(map[string]interface{})
			json.Unmarshal(data, &m)
			expects := s.matcher.FindStubs(md.GetService().GetFullyQualifiedName(), md.GetName(), m)
			if len(expects) == 0 {
				return status.Error(codes.NotFound, "didn't match any stub")
			}
			expect := expects[0]
			if expect.Out.Data != nil {
				out = msgFactory.NewMessage(md.GetOutputType())
				outData, _ := json.Marshal(expect.Out.Data)
				err = json.Unmarshal(outData, out)
				if err != nil {
					return err
				}

				if err = stream.SendMsg(out); err != nil {
					return err
				}
			}

			if expect.Out.Message != "" {
				return status.Error(codes.Code(expect.Out.Code), expect.Out.Message)
			}
		}
	}
}

func (s *Server) createGRPCServiceDesc(fds []*desc.FileDescriptor) []*grpc.ServiceDesc {
	var gsds []*grpc.ServiceDesc
	for _, fd := range fds {
		for _, sd := range fd.GetServices() {
			gsds = append(gsds, s.createServiceDesc(sd))
		}
	}
	return gsds
}

func (s *Server) createServiceDesc(sd *desc.ServiceDescriptor) *grpc.ServiceDesc {
	gsd := &grpc.ServiceDesc{
		ServiceName: sd.GetFullyQualifiedName(),
		HandlerType: nil,
		Metadata:    sd.GetFile().GetName(),
	}

	gsd.Methods, gsd.Streams = s.createMethodDescs(sd.GetMethods())
	return gsd
}

func (s *Server) createMethodDescs(mds []*desc.MethodDescriptor) ([]grpc.MethodDesc, []grpc.StreamDesc) {
	var methods []grpc.MethodDesc
	var streams []grpc.StreamDesc
	for _, md := range mds {
		if !md.IsClientStreaming() && !md.IsServerStreaming() {
			method := grpc.MethodDesc{
				MethodName: md.GetName(),
				Handler:    s.createUnaryServerHandler(md),
			}
			methods = append(methods, method)
		} else {
			stream := grpc.StreamDesc{
				StreamName:    md.GetName(),
				Handler:       s.createStreamServerHandler(md),
				ServerStreams: md.IsServerStreaming(),
				ClientStreams: md.IsClientStreaming(),
			}
			streams = append(streams, stream)
		}
	}
	return methods, streams
}
