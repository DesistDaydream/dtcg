package database

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 代码来源: https://github.com/xusenlin/gin-pagination

type Pagination[T any] struct {
	Count     int64 `json:"count"`
	PageSize  int   `json:"page_size"`
	PageNum   int   `json:"page_current"`
	PageTotal int   `json:"page_total"`
	Data      []T   `json:"data"`
	query     *gorm.DB
	ctx       *gin.Context
}

func NewPagination[T any](model T, c *gin.Context) *Pagination[T] {
	// TODO: 把 c 传进来这样好么？这代码是在数据库相关的目录中，好像不应该与 gin 强关联吧？
	// 而且写测试用例的时候也需要模拟 http 数据，比较繁琐。
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	return &Pagination[T]{
		PageSize: pageSize,
		PageNum:  pageNum,
		ctx:      c,
		// 类型约束，让 Pagination.Data 的类型变为 T，以便外部调用后，可以直接使用
		Data:  make([]T, 0),
		query: DB.Model(model),
	}
}

func computeTotalPage(total int64, pageSize int) int {
	totalPage := int(total) / pageSize
	if int(total)%pageSize != 0 {
		totalPage++
	}
	return totalPage
}

func (p *Pagination[T]) Query() error {
	p.query.Count(&p.Count)
	err := p.query.Offset((p.PageNum - 1) * p.PageSize).Limit(p.PageSize).Find(&p.Data).Error
	if err != nil {
		return err
	}
	p.PageTotal = computeTotalPage(p.Count, p.PageSize)
	return nil
}
