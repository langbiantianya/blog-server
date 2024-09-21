package dto

type PaginationDTO struct {
	Page  int `form:"page" json:"page,omitempty"`
	Limit int `form:"limit" json:"limit,omitempty"`
}

// GetPage 获取分页查询的页码。如果传入的页码参数 p.Page 为 0，则返回默认值 1，否则返回传入的页码值 p.Page。
func (p PaginationDTO) GetPage() int {
	if p.Page == 0 {
		return 1
	} else {
		return p.Page
	}
}

// GetLimit 获取分页查询的每页限制条目数。如果传入的 limit 参数 p.Limit 为 0，则返回默认值 5，否则返回传入的 limit 值 p.Limit。
func (p PaginationDTO) GetLimit() int {
	if p.Limit == 0 {
		return 5
	} else {
		return p.Limit
	}
}

// GetOffset 计算分页查询的偏移量。
// 偏移量的计算基于当前页码和每页限制条目数。
// 如果当前页码为 0，方法会返回默认的第一页偏移量。
// 参数:
//   - p: PaginationDTO 结构体实例，包含分页查询的参数。
//
// 返回值:
//   - int: 分页查询的偏移量。
func (p PaginationDTO) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
