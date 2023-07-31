package models

type User struct { //gorm 标签指定了表格 User 中各字段的数据类型以及其他特殊属性
	UserID   uint64 `gorm:"primaryKey" json:"user_id"`                // 指定一个列作为主键
	Name     string `gorm:"not null" json:"name" validate:"required"` // not null 表示该字段不能为空，validate 表示必须有
	Email    string `gorm:"not null" json:"email" validate:"required"`
	Password string `gorm:"not null" json:"password" validate:"required"`
}
