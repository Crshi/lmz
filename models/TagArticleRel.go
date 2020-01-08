package models

// TagArticle 关联表
type TagArticle struct {
	Id        int
	TagId     int
	ArticleId int
}

// TableName 获取表名
func (a *TagArticle) TableName() string {
	return TagArticleTBName()
}
