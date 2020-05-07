package models

type Account struct {
	ID       int
	Name     string
	Email    string
	Password string
	Group    int8
	Status   int8
	CreateTS int64
	UpdateTS int64
}

func (*Account) TableName() string {
	return "blog_account"
}

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
	return "blog_article"
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
