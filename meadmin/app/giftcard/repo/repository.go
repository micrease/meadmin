package repo

import (
	"github.com/micrease/gorme"
)

// 常用的获取数据方法进行封装,继承于gorose
type Repository[T gorme.Model] struct {
	gorme.Repository[T]
}

// 默认不使用缓存
func (this *Repository[T]) FindById(id uint) (T, error) {
	return this.Where("id=?", id).First()
}
