package model

import "gorm.io/gorm"

type Laporan struct {
	gorm.Model
	UserID         int
	User           User `gorm:"foreignKey:UserID"`
	Lokasi         string
	LokasiSpesifik string
	Permasalahan   string
	Kategori       string
	Photo          string
	Status         string
}
