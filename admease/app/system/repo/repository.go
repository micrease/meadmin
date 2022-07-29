package repo

import (
	"github.com/micrease/gorme"
	"gorm.io/gorm"
)

//常用的获取数据方法进行封装,继承于gorose
type Repository[T gorme.Model] struct {
	gorme.Repository[T]
}

//默认不使用缓存
func (this *Repository[T]) GetById(id uint) (T, error) {
	return this.Where("id=?", id).First()
}

func (this *Repository[T]) Values(builder *gorm.DB, column string) []any {
	var values []any
	this.QueryWithBuilder(builder).Pluck(column, &values)
	return values
}
