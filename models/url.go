package models

type URL struct {
	ID            uint        `gorm:"primaryKey"`
	EnvironmentID uint        `gorm:"uniqueIndex:idx_env_target"`
	Environment   Environment `gorm:"foreignKey:EnvironmentID"`
	Target        string      `gorm:"uniqueIndex:idx_env_target"`

	Tags        []Tag `gorm:"many2many:url_tags"`
	LocationID  uint
	Location    Location `gorm:"foreignKey:LocationID"`
	CheckTypeID uint
	CheckType   CheckType `gorm:"foreignKey:CheckTypeID"`
}

type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}

// URLTag is the join table for the many-to-many relation between URLs and tags.
type URLTag struct {
	URLID uint `gorm:"primaryKey"`
	TagID uint `gorm:"primaryKey"`
}

type Location struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}

type CheckType struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}
