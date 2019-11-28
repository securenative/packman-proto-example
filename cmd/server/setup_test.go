package server

import (
	"fmt"
	. "github.com/securenative/{{{ .PackageName }}}/pkg"
	"google.golang.org/grpc"
	"os"
	"testing"
	"time"
)

var client {{{ .Name }}}Client

func setup() {
	cfg := ParseConfig()

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", cfg.GrpcPort))
	if err != nil {
		panic(err)
	}
	client = New{{{ .Name }}}Client(conn)

	// You can place mocks here if you need to:
	module := NewModule(cfg)
	go func() {
		panic(module.GrpcServer.Start())
	}()
}

func teardown() {

}

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	time.Sleep(1*time.Second)
	os.Exit(m.Run())
}
