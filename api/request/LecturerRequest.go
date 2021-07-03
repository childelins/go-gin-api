package request

type LecturerList struct {
	Name string `form:"name"`
	// Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
	Pagination
}

type LecturerRequest struct {
	Name   string `form:"name" binding:"required,gte=1,lte=20"`
	Title  string `form:"title" binding:"required,gte=1,lte=30"`
	Avatar string `form:"avatar" binding:"omitempty,url"`
}
