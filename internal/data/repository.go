package data

type repositoryImpl struct {
}

func NewRepositoryImpl() *repositoryImpl {
	return &repositoryImpl{}
}

func (this *repositoryImpl) Find() (interface{}, error) {
	panic("implement me")
}

func (this *repositoryImpl) Save(id string, model interface{}) error {
	panic("implement me")
}

