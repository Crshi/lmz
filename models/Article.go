package models

func (a *Article) TableName() string {
	return ArticleTBName()
}

type Article struct {
	Id			int
	Title       string
	Content 	string
	V
}
