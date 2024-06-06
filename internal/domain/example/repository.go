package example

type ExampleRepository interface {
	Store(e *Example) error
	FindByID(id string) (*Example, error)
}
