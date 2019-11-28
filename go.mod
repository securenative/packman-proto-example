module github.com/securenative/{{{ .PackageName }}}

go 1.13

require (
	github.com/caarlos0/env latest
	google.golang.org/grpc latest
	github.com/stretchr/testify latest
	github.com/grpc-ecosystem/go-grpc-middleware latest
	github.com/grpc-ecosystem/go-grpc-prometheus latest
	github.com/kazegusuri/grpc-panic-handler latest
)
