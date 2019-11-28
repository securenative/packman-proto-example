package server

import (
	"context"
	. "github.com/securenative/{{{ .PackageName }}}/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)
{{{ range $k, $v := .Methods }}}
func TestIntegration_{{{ $k }}}(t *testing.T) {
	input := &{{{ $v.Input.Name }}}{
		{{{- range $tk, $tv := $v.Input.Fields }}}
		{{{ $tv.Name }}}: nil,
		{{{- end }}}
	}

	output, err := client.{{{ $k }}}(context.TODO(), input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	{{{- range $tk, $tv := $v.Output.Fields }}}
	assert.EqualValues(t, "PLACE VALUE FOR {{{ $tv.Name }}}", output.{{{ $tv.Name }}})
	{{{- end }}}
}
{{{ end }}}