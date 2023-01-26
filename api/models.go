package main

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name  string
	Files []File
}

type File struct {
	gorm.Model
	Name string
	Size int64
}
