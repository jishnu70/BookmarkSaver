package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name      string
	Bookmarks []Bookmark `gorm:"many2many:bookmark_tags"`
}
