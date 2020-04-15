package database

type NotFoundErr struct{}

func NewNotFoundErr() NotFoundErr {
	return NotFoundErr{}
}

func (n NotFoundErr) Error() string {
	return "entity not found"
}
