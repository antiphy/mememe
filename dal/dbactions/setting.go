package dbactions

func CreateTableSetting() error {
	return db.Exec("CREATE TABLE `mememe_setting` (`id` int(11) NOT NULL AUTO_INCREMENT, `setting_key` varchar(100) DEFAULT '', `setting_value` varchar(1024) DEFAULT '', `setting_type` tinyint(4) DEFAULT 0, `setting_status` tinyint(4) DEFAULT 0, `create_ts` bigint(20) DEFAULT 0, `update_ts` bigint(20) DEFAULT 0, PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;").Error
}
