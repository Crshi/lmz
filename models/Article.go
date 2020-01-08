package models

import "time"

// Article 表
type Article struct {
	Id           int
	Title        string `orm:"size(128)"`
	Content      string `orm:"size(12800)"`
	Viewcount    int
	CreationTime time.Time
	DeletionTime time.Time
	//CreatorUserId
	//LastModificationTime
	//LastModifierUserId
	//IsDeleted
	//DeleterUserId
}

// TableName 获取表名
func (a *Article) TableName() string {
	return ArticleTBName()
}
