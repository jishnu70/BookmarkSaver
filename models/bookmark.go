package models

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model
	Title  string
	URL    string
	Tags   []Tag `gorm:"many2many:bookmark_tags"`
	UserID uint
}
