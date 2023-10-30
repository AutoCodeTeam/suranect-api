package controller

import (
	"mime/multipart"
	"strconv"
	"suranect_api/database"
	"suranect_api/model"
	"suranect_api/utils"

	"github.com/gin-gonic/gin"
)

type LaporanController struct{}

func (ctrl LaporanController) Index(c *gin.Context) {
	db, _ := database.ConnectMysql()

	var laporan []model.Laporan
	db.Model(&model.Laporan{}).Preload("User").Find(&laporan)

	c.JSON(200, gin.H{
		"Data":    laporan,
		"Status":  "success",
		"Message": "Success Get Laporan",
	})
}

type storeForm struct {
	Lokasi         string               `form:"lokasi" binding:"required"`
	LokasiSpesifik string               `form:"lokasiSpesifik"`
	Permasalahan   string               `form:"permasalahan" binding:"required"`
	Kategori       string               `form:"kategori" binding:"required"`
	Status         string               `form:"status"`
	Photo          multipart.FileHeader `form:"photo" binding:"required"`
}

func (ctrl LaporanController) Store(c *gin.Context) {
	is_login := c.MustGet("is_login").(bool)
	if !is_login {
		return
	}

	userinfo := c.MustGet("userinfo").(map[string]interface{})
	userID := uint(userinfo["User_id"].(float64))

	db, _ := database.ConnectMysql()
	var request storeForm
	request.Status = "pending"
	err := c.Bind(&request)

	if err != nil {
		c.JSON(400, gin.H{
			"Status":         "error",
			"Message":        "Please do not leave the form blank",
			"Message_server": err.Error(),
		})
		return
	}

	file, _ := c.FormFile("photo")
	idObject, err := utils.Upload(file, "laporan-image")

	db.Create(&model.Laporan{
		UserID:         int(userID),
		Permasalahan:   request.Permasalahan,
		Lokasi:         request.Lokasi,
		LokasiSpesifik: request.LokasiSpesifik,
		Kategori:       request.Kategori,
		Photo:          idObject,
		Status:         request.Status,
	})

	c.JSON(200, gin.H{
		"Message": "Berhasil Membuat Laporan",
		"Status":  "success",
	})
}

type updateForm struct {
	Lokasi         string               `form:"lokasi" binding:"required"`
	LokasiSpesifik string               `form:"lokasiSpesifik"`
	Permasalahan   string               `form:"permasalahan" binding:"required"`
	Kategori       string               `form:"kategori" binding:"required"`
	Photo          multipart.FileHeader `form:"photo"`
}

func (ctrl LaporanController) Update(c *gin.Context) {
	is_login := c.MustGet("is_login").(bool)
	if !is_login {
		return
	}

	db, _ := database.ConnectMysql()
	var request updateForm
	err := c.Bind(&request)

	if err != nil {
		c.JSON(400, gin.H{
			"Status":         "error",
			"Message":        "Please do not leave the form blank",
			"Message_server": err.Error(),
		})
		return
	}

	laporan := model.Laporan{}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	laporan.ID = uint(id)

	db.Model(&model.Laporan{}).Preload("User").First(&laporan)

	file, _ := c.FormFile("photo")

	if file != nil {
		idObject, _ := utils.Upload(&request.Photo, "laporan-image")
		laporan.Photo = idObject
	}

	laporan.Lokasi = request.Lokasi
	laporan.LokasiSpesifik = request.LokasiSpesifik
	laporan.Permasalahan = request.Permasalahan
	laporan.Kategori = request.Kategori

	db.Save(&laporan)

	c.JSON(200, gin.H{
		"Data":    laporan,
		"Message": "Berhasil Edit Laporan",
		"Status":  "success",
	})

}

func (ctrl LaporanController) Delete(c *gin.Context) {
	is_login := c.MustGet("is_login").(bool)
	if !is_login {
		return
	}

	db, _ := database.ConnectMysql()

	laporan := model.Laporan{}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	laporan.ID = uint(id)

	db.Delete(&laporan)

	c.JSON(200, gin.H{
		"Message": "Berhasil Menghapus Laporan",
		"Status":  "success",
	})
}

func (ctrl LaporanController) Show(c *gin.Context) {
	is_login := c.MustGet("is_login").(bool)
	if !is_login {
		return
	}

	db, _ := database.ConnectMysql()

	laporan := model.Laporan{}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	laporan.ID = uint(id)

	db.Model(&model.Laporan{}).Preload("User").First(&laporan)

	c.JSON(200, gin.H{
		"Data":    laporan,
		"Message": "Berhasil Get Detail Laporan",
		"Status":  "success",
	})
}
