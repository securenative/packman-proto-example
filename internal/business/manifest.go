package business

import (
	. "github.com/securenative/{{{ .PackageName }}}/pkg"
)

type Service interface {
{{{- range $k, $v :=.Methods }}}
	{{{ $k }}}(input *{{{ $v.Input.Name }}}) (*{{{ $v.Output.Name }}}, error)
{{{- end }}}
}