package repositories

type IBaseRepository[T any] interface {
	FindAll() []T
	Create(entity T) (T, error)
	//Find(id uint) T
	//Update(entity T) (T, error)
	//Delete(id uint) error

}
