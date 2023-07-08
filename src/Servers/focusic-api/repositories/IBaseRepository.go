package repositories

type IBaseRepository[T any] interface {
	Create(entity T) (T, error)
	FindAll() ([]T, error)
}
