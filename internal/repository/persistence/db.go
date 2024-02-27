package persistence

type Repository[T any] interface {
	Store(T) error
	Get(string) (*T, error)
	GetAll() ([]T, error)
}
