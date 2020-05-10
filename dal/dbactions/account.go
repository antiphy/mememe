package dbactions

import "github.com/antiphy/mememe/dal/models"

func CreateAccount(a *models.Account) error {
	return db.Create(a).Error
}

func QueryAccount(a *models.Account) error {
	return db.Model(a).Where("name = ?", a.Name).Scan(a).Error
}

func CreateTableAccount() error {
	return db.Exec("CREATE TABLE `mememe_account` (  `id` int(11) NOT NULL AUTO_INCREMENT,  `name` varchar(50) DEFAULT '',  `email` varchar(50) DEFAULT '',  `password` varchar(40) DEFAULT '',  `group` tinyint(4) DEFAULT 0,  `status` tinyint(4) DEFAULT 0,  `create_ts` bigint(20) DEFAULT 0,  `update_ts` bigint(20) DEFAULT 0,  PRIMARY KEY (`id`),  KEY `idx_name` (`name`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;").Error
}
