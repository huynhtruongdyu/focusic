package repositories

type IBaseRepository[T any] interface {
	Create(entity T) (T, error)
	//Update(entity T) (T, error)
	//Delete(id uint) error
	//FindAll() []T
	//Find(id uint) T
}
