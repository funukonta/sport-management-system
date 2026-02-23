package dto

type PaginationQuery struct {
	Page     int `form:"page" binding:"omitempty,min=1"`
	PageSize int `form:"page_size" binding:"omitempty,min=1,max=100"`
}

type PaginationMeta struct {
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (pq *PaginationQuery) GetPage() int {
	if pq.Page == 0 {
		return 1
	}
	return pq.Page
}

func (pq *PaginationQuery) GetLimit() int {
	if pq.PageSize == 0 {
		return 10
	}
	return pq.PageSize
}

func (pq *PaginationQuery) GetOffset() int {
	return (pq.GetPage() - 1) * pq.GetLimit()
}
