package resource_tag

import "time"

// ResourceTag 资源标签关联实体
type ResourceTag struct {
	Id           int64     `json:"id" gorm:"column:id;primaryKey"`
	ResourceId   int64     `json:"resourceId" gorm:"column:resource_id;not null"`
	ResourceType string    `json:"resourceType" gorm:"column:resource_type;type:varchar(50);not null"`
	TagId        int64     `json:"tagId" gorm:"column:tag_id;not null"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
}

// TableName 指定表名
func (ResourceTag) TableName() string {
	return "resource_tags"
}
