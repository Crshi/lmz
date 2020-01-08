package models

// Catalog 表,规定ParentID为0的是根节点
type Catalog struct {
	Id          int
	ParentId    int
	Name        string `orm:"size(32)"`
	Description string `orm:"size(128)"`
}

// TableName 获取表名
func (a *Catalog) TableName() string {
	return CatalogTBName()
}
