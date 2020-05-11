package dbactions

import "github.com/antiphy/mememe/dal/models"

func CreateTableSetting() error {
	return db.Exec("CREATE TABLE `mememe_setting` (`id` int(11) NOT NULL AUTO_INCREMENT, `setting_key` varchar(100) DEFAULT '', `setting_value` varchar(1024) DEFAULT '', `setting_type` tinyint(4) DEFAULT 0, `setting_status` tinyint(4) DEFAULT 0, `create_ts` bigint(20) DEFAULT 0, `update_ts` bigint(20) DEFAULT 0, PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;").Error
}

func InsertOrUpdateSettings(settings []models.Setting) error {
	for i := range settings {
		err := db.Exec(`insert into mememe_setting (setting_key, setting_value, setting_type, setting_status, 
		create_ts, update_ts) values (?, ?, ?, ?, ?, ?) on duplicate key update setting_value = ?, setting_type = ?, 
		setting_status = ?, create_ts = ?, update_ts = ?`, settings[i].SettingKey, settings[i].SettingValue,
			settings[i].SettingType, settings[i].SettingStatus, settings[i].CreateTS, settings[i].UpdateTS, settings[i].SettingValue,
			settings[i].SettingType, settings[i].SettingStatus, settings[i].CreateTS, settings[i].UpdateTS).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func QuerySettings() ([]models.Setting, error) {
	var settings []models.Setting
	err := db.Raw(`select * from mememe_setting`).Scan(&settings).Error
	if err != nil {
		return nil, err
	}
	return settings, nil
}
