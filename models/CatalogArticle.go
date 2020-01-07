package models

// CatalogArticle 表
type CatalogArticle struct {
	ID        int
	CatalogID int
	ArticleID int
}

// TableName 表名
func (a *CatalogArticle) TableName() string {
	return CatalogArticleTBName()
}
