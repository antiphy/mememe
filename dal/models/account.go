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
	return "mememe_account"
}

func (a *Account) IsAdmin() bool {
	return true
}
