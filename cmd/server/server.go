package server

import (
	context "context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	panichandler "github.com/kazegusuri/grpc-panic-handler"
	"github.com/securenative/{{{ .PackageName }}}/internal/business"
	. "github.com/securenative/{{{ .PackageName }}}/pkg"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	Config Config
	server *grpc.Server
}

func NewGrpcServer(config Config, service business.Service) *GrpcServer {
	impl := newServerImpl(service)
	server := initGrpcServer(impl)
	return &GrpcServer{Config: config, server: server}
}

func (this *GrpcServer) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", this.Config.GrpcPort))
	if err != nil {
		return err
	}
	return this.server.Serve(listener)
}

func initGrpcServer(impl *serverImpl) *grpc.Server {
	unaryInterceptors := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		panichandler.UnaryPanicHandler,
		grpc_prometheus.UnaryServerInterceptor,
	))

	streamInterceptors := grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_prometheus.StreamServerInterceptor,
	))

	var options []grpc.ServerOption
	options = append(options, unaryInterceptors, streamInterceptors)
	server := grpc.NewServer(options...)
	Register{{{ .Name }}}Server(server, impl)

	return server
}

type serverImpl struct {
	service business.Service
}

func newServerImpl(service business.Service) *serverImpl {
	return &serverImpl{service: service}
}

{{{- range $k, $v := .Methods }}}
func (this* serverImpl) {{{ $k }}}(ctx context.Context, input *{{{ $v.Input.Name }}}) (*{{{ $v.Output.Name }}}, error) {
	return this.service.{{{ $k }}}(input)
}
{{{ end }}}
