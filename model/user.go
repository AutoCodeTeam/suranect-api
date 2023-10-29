package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Username    string `gorm:"unique"`
	Password    string
	Photo       string
	Email       string `gorm:"unique"`
	EmailVerify bool
	EmailCode   int
	Laporans    []Laporan
}
