package repo

import "blog-server/internal/entity/vo"

type IRepo[T any] interface {
	Info(uint) (*T, error)
	List(any) (*[]T, error)
	Add(T) error
	Delete(uint) error
}

type IRepoPage[T any, DTO any, VO any] interface {
	Info(uint) (*T, error)
	Find(DTO) (*vo.PaginationVO[[]VO], error)
	Add(T) error
	Delete(uint) error
}
