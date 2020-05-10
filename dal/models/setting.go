package models

type Setting struct {
	ID            int
	SettingKey    string
	SettingValue  string
	SettingType   int8
	SettingStatus int8
	CreateTS      int64
	UpdateTS      int64
}

func (*Setting) TableName() string {
	return "mememe_setting"
}
