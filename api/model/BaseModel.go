package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/childelins/go-gin-api/global"
)

type Model struct {
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func All(t interface{}, cond map[string]interface{}) {
	global.DB.Where(cond).Find(t)
}

func Find(t interface{}, cond map[string]interface{}) {
	global.DB.Where(cond).First(t)
}

func Create(t interface{}) {
	global.DB.Create(t)
}

func Update(t interface{}, columns map[string]interface{}) {
	global.DB.Model(t).Updates(columns)
}

func Delete(t interface{}, cond map[string]interface{}) {
	global.DB.Where(cond).Delete(t)
}

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		/*
			page, _ := strconv.Atoi(c.Query("page"))
			if page <= 0 {
				page = 1
			}

			limit, _ := strconv.Atoi(c.Query("limit"))
			if limit <= 0 {
				limit = 10
			}
			if limit > 100 {
				limit = 100
			}
		*/

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
