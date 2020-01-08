package models

// CatalogArticle 表
type CatalogArticle struct {
	Id        int
	CatalogId int
	ArticleId int
}

// TableName 表名
func (a *CatalogArticle) TableName() string {
	return CatalogArticleTBName()
}
