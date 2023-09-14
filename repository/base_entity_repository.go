package repository

import (
	"fmt"
	"go-middleware-recipe/database"
)

type Repository[T any] interface {
	Create(data T) (T, error)
	Read(id int) (T, error)
	Update(data T) error
	Delete(id int) error
	Count() (int64, error)
}

type Impl[T any] struct {
}

func (r *Impl[T]) Create(data T) (T, error) {
	tx := database.DBConn.Create(&data)

	if tx.Error != nil {
		return data, tx.Error
	}

	return data, nil
}

func (r *Impl[T]) Read(id int) (T, error) {
	var entity T
	tx := database.DBConn.First(&entity, id)
	return entity, tx.Error
}

func (r *Impl[T]) Update(data T) error {
	fmt.Println("RepositoryImpl Update")
	return nil
}

func (r *Impl[T]) Delete(id int) error {
	fmt.Println("RepositoryImpl Delete")
	return nil
}

func (r *Impl[T]) Count() (int64, error) {
	// TODO: Is there really really no way to do this without creating an instance of T?
	var entity T
	var count int64
	tx := database.DBConn.Model(entity).Count(&count)
	return count, tx.Error
}
