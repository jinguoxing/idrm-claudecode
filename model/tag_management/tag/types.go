package tag

import "time"

// Tag 数据标签实体
type Tag struct {
	Id          int64     `json:"id" gorm:"column:id;primaryKey"`
	Name        string    `json:"name" gorm:"column:name;type:varchar(50);not null"`
	Description string    `json:"description" gorm:"column:description;type:varchar(200)"`
	Color       string    `json:"color" gorm:"column:color;type:varchar(7);default:'#1890ff'"`
	Status      int       `json:"status" gorm:"column:status;type:tinyint;not null;default:1"`
	CreatedBy   int64     `json:"createdBy" gorm:"column:created_by;not null"`
	UpdatedBy   *int64    `json:"updatedBy" gorm:"column:updated_by"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}

// TableName 指定表名
func (Tag) TableName() string {
	return "tags"
}
