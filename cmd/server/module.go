package server

import (
	"github.com/securenative/{{{ .PackageName }}}/internal/data"
	"github.com/securenative/{{{ .PackageName }}}/internal/business"
)

type Module struct {
	Config Config
	GrpcServer *GrpcServer
}

func NewModule(cfg Config) *Module {
	// This is the place where you should put all your initialization logic, database connections, migrations, etc...
	// Use panic to shutdown the all app if module cannot be initialized
	// If the caller of NewModule gets a module it means that the module has been initialized properly
	repository := data.NewRepositoryImpl()
	service := business.NewServiceImpl(repository)

	grpcServer := NewGrpcServer(cfg, service)

	return &Module{
		Config:         cfg,
		GrpcServer:     grpcServer,
	}
}