package controllers

import (
	"mahasiswa/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMahasiswa(ctx *gin.Context) {

	newMahasiswa := model.MahasiswaCreateRequest{}

	db := ctx.MustGet("db").(*gorm.DB)

	if err := ctx.ShouldBindJSON(&newMahasiswa); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	mahasiswa := model.Mahasiswa{
		Nim:        newMahasiswa.Nim,
		Nama:       newMahasiswa.Nama,
		KotaAsal:   newMahasiswa.KotaAsal,
		TahunMasuk: newMahasiswa.TahunMasuk,
	}

	if err := db.Create(&mahasiswa).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	response := model.MahasiswacreateResponse{
		Nim:        mahasiswa.Nim,
		Nama:       mahasiswa.Nama,
		KotaAsal:   mahasiswa.KotaAsal,
		TahunMasuk: mahasiswa.TahunMasuk,
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
}

func GetMahasiswaId(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)
	nimMahasiswa := ctx.Param("nim")
	Mahasiswa := model.Mahasiswa{}

	if err := db.Preload("IPMahasiswa").Preload("NilaiUts").Preload("NilaiUas").Where("nim = ?",
		nimMahasiswa).First(&Mahasiswa).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusBadRequest,
					Message: "Mahasiswa dengan Nim " + nimMahasiswa + " tidak ditemukan",
				},
				Error: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: Mahasiswa,
	})
}

func GetAllMahasiswa(ctx *gin.Context) {

	mahasiswa := []model.Mahasiswa{}

	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Preload("IPMahasiswa").Preload("NilaiUts").Preload("NilaiUas").Find(&mahasiswa).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: mahasiswa,
	})
}

func UpdateMahasiswa(ctx *gin.Context) {
	nim := ctx.Param("nim")

	db := ctx.MustGet("db").(*gorm.DB)
	mahasiswa := model.Mahasiswa{}

	if err := db.Where("nim = ?", nim).First(&mahasiswa).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&mahasiswa); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := db.Save(&mahasiswa).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: model.MahasiswaUpdateResponse{
			Nim:        mahasiswa.Nim,
			Nama:       mahasiswa.Nama,
			KotaAsal:   mahasiswa.KotaAsal,
			TahunMasuk: mahasiswa.TahunMasuk,
		},
	})
}

func DeleteMahasiswa(ctx *gin.Context) {

	nim := ctx.Param("nim")

	db := ctx.MustGet("db").(*gorm.DB)

	mahasiswa := model.Mahasiswa{}

	if err := db.Where("nim = ?", nim).First(&mahasiswa).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: "Mahasiswa dengan NIM " + nim + " tidak ditemukan",
				},
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
				},
				Error: err.Error(),
			})
		}
		return
	}

	if err := db.Delete(&mahasiswa).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: "Mahasiswa dengan NIM " + nim + " telah dihapus",
	})
}
