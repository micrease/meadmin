package repo

import (
	"github.com/micrease/gorme"
	"meadmin/system/datasource"
)

// 常用的获取数据方法进行封装,继承于gorose
type Repository[T gorme.Model] struct {
	gorme.Repository[T]
}

func (this *Repository[T]) initialize() *Repository[T] {
	this.SetDB(datasource.GetDB())
	this.NewQuery()
	return this
}

// 默认不使用缓存
func (this *Repository[T]) GetById(id uint) (T, error) {
	return this.Where("id=?", id).First()
}
