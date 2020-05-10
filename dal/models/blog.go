package models

type Article struct {
	ID         int
	CategoryID int
	Title      string
	Content    string
	Status     int8
	CreatedBy  int
	CreateTS   int64
	UpdateTS   int64
}

func (*Article) TableName() string {
	return "mememe_article"
}

type ExtendArticle struct {
	ID         int
	CategoryID int
	Title      string
	Content    string
	Status     int8
	CreatedBy  int
	Creator    string
	CreateTS   int64
	UpdateTS   int64
}
