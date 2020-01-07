package models

// Tag 表
type Tag struct {
	ID   int
	Name string `orm:"size(32)"`
}

// TableName 获取表名
func (a *Tag) TableName() string {
	return TagTBName()
}
