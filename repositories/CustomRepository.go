package repositories

import (
	"goApi/model"
)

// CustomInterface 用户接口
type CustomInterface interface {
	// 创建用户
	CreateCustoms()
	// 获得用户名字
	// GetCustom()
}

// CustomRepository 业务结构 也就是这个业务结构需要哪些值
type CustomRepository struct {
	ID         int
	CustomName string
}

// CreateCustoms 创建用户
func (cr *CustomRepository) CreateCustoms() {
	// 将CustomRepository的参数填入模型中，实例化
	custom := model.Custom{
		ID:         cr.ID,
		CustomName: cr.CustomName,
	}
	model.DB.Create(&custom)

}
