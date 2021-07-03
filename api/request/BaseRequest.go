package request

type Pagination struct {
	Page  int `form:"page" binding:"gte=1"`
	Limit int `form:"limit" binding:"gte=1"`
}
