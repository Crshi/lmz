package models

// TagArticle 关联表
type TagArticle struct {
	ID        int
	TagID     int
	ArticleID int
}

// TableName 获取表名
func (a *TagArticle) TableName() string {
	return TagArticleTBName()
}
