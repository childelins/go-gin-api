package controller

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/childelins/go-gin-api/api/model"
	"github.com/childelins/go-gin-api/api/request"
	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/pkg/app"
	"github.com/childelins/go-gin-api/pkg/errcode"
	"github.com/childelins/go-gin-api/proto"
)

type Lecturer struct{}

type LecturerResp struct {
	LecturerId int    `json:"lecturerId"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Title      string `json:"title"`
	CreatedAt  string `json:"createdAt"`
}

func (l *Lecturer) List(c *gin.Context) {
	// companyId := 13
	params := &request.LecturerList{}
	valid, err := app.Validate(c, params)
	if !valid {
		global.Logger.Errorf("app.Validate err: %v", err)
		app.ToErrorResponse(c, errcode.InvalidParams.WithError(err))
		return
	}

	// 限流
	/*
		e, b := sentinel.Entry("lecturer-list", sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			app.ToErrorResponse(c, errcode.TooManyRequests)
			return
		}
	*/
	pCtx := c.MustGet("ctx").(context.Context)
	resp, err := global.LecturerSrvClient.GetLecturerList(pCtx, &proto.LecturerListRequest{
		Page:  uint32(params.Page),
		Limit: uint32(params.Limit),
		Name:  params.Name,
	})
	if err != nil {
		global.Logger.Errorf("GetLecturerList err: %v", err)
		app.ToErrorResponse(c, errcode.ServerError)
		return
	}

	//e.Exit()

	data := make([]interface{}, 0)
	for _, value := range resp.Data {
		item := &LecturerResp{
			LecturerId: int(value.LecturerId),
			Name:       value.Name,
			Avatar:     value.Avatar,
			Title:      value.Title,
			CreatedAt:  value.CreatedAt,
		}

		data = append(data, item)
	}

	app.ToResponse(c, gin.H{
		"list":  data,
		"total": resp.Total,
	})
	return

	// parentSpan := opentracing.SpanFromContext(c.Request.Context())
	/*
		parentSpan := c.MustGet("X-Span").(opentracing.Span)
		ctx := opentracing.ContextWithSpan(context.Background(), parentSpan)
		global.DB = global.DB.WithContext(ctx)
	*/

	/*
		lecturerModel := &model.Lecturer{}
		lecturers, err := lecturerModel.GetList(companyId, params)
		if err != nil {
			app.ToErrorResponse(c, errcode.ServerError)
		}
		count := lecturerModel.GetCount(companyId, params)

		app.ToResponse(c, gin.H{
			"list":  lecturers,
			"total": count,
		})
		return
	*/
}

func (l *Lecturer) Create(c *gin.Context) {
	params := &request.LecturerRequest{}
	valid, err := app.Validate(c, params)
	if !valid {
		global.Logger.Errorf("app.Validate err: %v", err)
		app.ToErrorResponse(c, errcode.InvalidParams.WithError(err))
		return
	}

	companyId := 13
	lecturerModel := &model.Lecturer{
		CompanyId: companyId,
		Name:      params.Name,
		Title:     params.Title,
		Avatar:    params.Avatar,
	}
	model.Create(lecturerModel)
	app.ToResponse(c, gin.H{
		"lecturerId": lecturerModel.LecturerId,
	})
	return
}

func (l *Lecturer) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	params := &request.LecturerRequest{}
	valid, err := app.Validate(c, params)
	if !valid {
		global.Logger.Errorf("app.Validate err: %v", err)
		app.ToErrorResponse(c, errcode.InvalidParams.WithError(err))
		return
	}

	companyId := 13
	lecturerModel := &model.Lecturer{}
	model.Find(lecturerModel, map[string]interface{}{
		"lecturerId": id,
		"companyId":  companyId,
	})
	model.Update(lecturerModel, map[string]interface{}{
		"name":   params.Name,
		"title":  params.Title,
		"avatar": params.Avatar,
	})
	app.ToResponse(c, nil)
	return
}

func (l *Lecturer) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	companyId := 13
	lecturerModel := &model.Lecturer{}
	model.Delete(lecturerModel, map[string]interface{}{
		"lecturerId": id,
		"companyId":  companyId,
	})
	app.ToResponse(c, nil)
	return
}
