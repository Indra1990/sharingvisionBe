package entity

import "time"

type Post struct {
	ID          uint64    `gorm:"primary_key:auto_increment"`
	Title       string    `gorm:"type:varchar(200);not null"`
	Content     string    `gorm:"type:text;not null"`
	Category    string    `gorm:"type:varchar(100);not null"`
	CreatedDate time.Time `gorm:"type:timestamp;<-:create"`
	UpdatedDate time.Time `gorm:"type:timestamp;<-:update"`
	Status      string    `gorm:"type:varchar(100); not null"`
}
