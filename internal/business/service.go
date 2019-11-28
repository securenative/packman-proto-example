package business


import (
	"errors"
	"fmt"
	"github.com/securenative/{{{ .PackageName }}}/internal/data"
	. "github.com/securenative/{{{ .PackageName }}}/pkg"
)

type ServiceImpl struct {
	repository data.Repository
}

func NewServiceImpl(repository data.Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

{{{ range $k, $v :=.Methods }}}
func (this *ServiceImpl) {{{ $k }}}(input *{{{ $v.Input.Name }}}) (*{{{ $v.Output.Name }}}, error) {
	model, err := this.repository.Find()
	if err != nil {
		fmt.Println(model)
	}

	return nil, errors.New("implement me")
}
{{{ end }}}