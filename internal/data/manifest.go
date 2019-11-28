package data

type Repository interface {
	Find() (interface{}, error)
	Save(id string, model interface{}) error
}