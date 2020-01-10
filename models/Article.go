package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

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

type ArticleQueryInput struct {
	PagedInput
	TagId        int
	CatalogId    int
	CreationTime time.Time
}

// TableName 获取表名
func (a *Article) TableName() string {
	return ArticleTBName()
}

// GetAllArticlesByPage 获取分页数据
func GetAllArticlesByPage(input *ArticleQueryInput) ([]*Article, int64) {
	query := orm.NewOrm().QueryTable(ArticleTBName())
	data := make([]*Article, 0)
	//默认排序
	sortorder := "Id"
	query = query.Filter("title__istartswith", input.Filter)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(input.MaxResultCount, input.SkipCount).All(&data)
	return data, total
}

// Get 获取单个文章
func Get(id int) (*Article, error) {
	o := orm.NewOrm()
	a := Article{Id: id}
	err := o.Read(&a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
