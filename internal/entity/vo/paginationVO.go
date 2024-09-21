package vo

type PaginationVO[T any] struct {
	Page  int   `json:"page,omitempty"`
	Limit int   `json:"limit,omitempty"`
	Size  int64 `json:"size,omitempty"`
	Data  T     `json:"data,omitempty"`
}
