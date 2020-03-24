package model

// Custom 用户模型
type Custom struct {
	ID         int `gorm:"PRIMARY_KEY"`
	CustomName string
}

var res interface{}

// GetNameByID 用ID获取用户name
func GetNameByID(ID int) (Custom, error) {
	var custom Custom
	result := DB.First(&custom, ID)
	return custom, result.Error
}

// SetIDAndName 设置ID and Name
func SetIDAndName(c Custom) error {
	result := DB.Create(c)
	return result.Error
}
