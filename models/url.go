package models

type URL struct {
	ID            uint        `gorm:"primaryKey"`
	EnvironmentID uint        `gorm:"uniqueIndex:idx_env_target"`
	Environment   Environment `gorm:"foreignKey:EnvironmentID"`
	Target        string      `gorm:"uniqueIndex:idx_env_target"`
}
