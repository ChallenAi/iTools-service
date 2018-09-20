package models

type Tag struct {
	TagId    int    `gorm:"type:int(8);PRIMARY_KEY" json:"tagId"`
	ParentId int    `gorm:"type:int(8)" json:"parentId"`
	Content  string `json:"content"`
}

func (Tag) TableName() string {
	return "tag"
}

func GetTags() ([]Tag, error) {
	var tags []Tag
	err := DB.Find(&tags).Error
	return tags, err
}
