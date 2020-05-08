package dbactions

import "github.com/antiphy/mememe/dal/models"

func CreateAccount(a *models.Account) error {
	return db.Create(a).Error
}

func QueryAccount(a *models.Account) error {
	return db.Model(a).Where("name = ?", a.Name).Scan(a).Error
}
