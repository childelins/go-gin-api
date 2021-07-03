package model

import (
	"github.com/childelins/go-gin-api/api/request"
	"github.com/childelins/go-gin-api/global"
)

type Lecturer struct {
	LecturerId int    `gorm:"primaryKey;column:lecturerId" json:"lecturerId"`
	CompanyId  int    `gorm:"column:companyId" json:"companyId"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Avatar     string `json:"avatar"`
	*Model
}

func (l *Lecturer) TableName() string {
	return "lecturers"
}

func (l *Lecturer) GetCount(companyId int, params *request.LecturerList) int {
	var count int64
	db := global.DB.Where("companyId = ?", companyId)
	if len(params.Name) > 0 {
		db.Where("name like ?", "%"+params.Name+"%")
	}

	db.Model(l).Count(&count)
	return int(count)
}

func (l *Lecturer) GetList(companyId int, params *request.LecturerList) ([]*Lecturer, error) {
	var lecturers []*Lecturer
	db := global.DB.Scopes(Paginate(params.Page, params.Limit)).Where("companyId = ?", companyId)
	if len(params.Name) > 0 {
		db.Where("name like ?", "%"+params.Name+"%")
	}

	db.Order("createdAt desc").Find(&lecturers)
	return lecturers, nil
}
