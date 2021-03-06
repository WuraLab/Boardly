package models

type Base struct {
	ID      uint  `gorm:"primaryKey;column:id;autoIncrement;"`
	Updated int64 `gorm:"autoUpdateTime:milli"`
	Created int64 `gorm:"autoCreateTime:milli"`
}
