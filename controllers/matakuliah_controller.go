package controllers

import (
	"mahasiswa/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMatkul(ctx *gin.Context) {
	newMatkul := model.MatakuliahCreateRequest{}

	db := ctx.MustGet("db").(*gorm.DB)

	if err := ctx.ShouldBindJSON(&newMatkul); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	matkul := model.Matakuliah{
		KodeMatkul: newMatkul.KodeMatkul,
		NamaMatkul: newMatkul.NamaMatkul,
		NipDosen:   newMatkul.NipDosen,
		JumlahSks:  newMatkul.JumlahSks,
	}

	err := db.Create(&matkul).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	response := model.MatakuliahCreateResponse{
		KodeMatkul: matkul.KodeMatkul,
		NamaMatkul: matkul.NamaMatkul,
		NipDosen:   matkul.NipDosen,
		JumlahSks:  matkul.JumlahSks,
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
}

func GetMatkulId(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)
	kodeMatkul := ctx.Param("kode_matkul")

	matkul := model.Matakuliah{}

	if err := db.Where("kode_matkul = ?", kodeMatkul).First(&matkul).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusBadRequest,
					Message: "Matakuliah dengan Kode " + kodeMatkul + " tidak ditemukan",
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
		Data: matkul,
	})
}

func GetAllMatkul(ctx *gin.Context) {

	matkul := []model.Matakuliah{}

	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Find(&matkul).Error; err != nil {
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
		Data: matkul,
	})
}

func UpdateMatakuliah(ctx *gin.Context) {
	Mk := ctx.Param("kode_matkul")

	db := ctx.MustGet("db").(*gorm.DB)
	matkul := model.Matakuliah{}
	if err := db.Where("kode_matkul= ?", Mk).First(&matkul).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&matkul); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := db.Save(&matkul).Error; err != nil {
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
		Data: model.MatakuliahUpdateResponse{
			KodeMatkul: matkul.KodeMatkul,
			NamaMatkul: matkul.NamaMatkul,
			NipDosen:   matkul.NipDosen,
			JumlahSks:  matkul.JumlahSks,
		},
	})
}

func DeleteMatkul(ctx *gin.Context) {
	kode := ctx.Param("kode_matkul")

	db := ctx.MustGet("db").(*gorm.DB)
	matkul := model.Matakuliah{}

	if err := db.Where("kode_matkul = ?", kode).First(&matkul).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: "Matakuliah dengan Kode " + kode + " tidak ditemukan",
				},
				Error: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	if err := db.Delete(&matkul).Error; err != nil {
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
		Data: "Matakuliah dengan kode " + kode + " telah dihapus",
	})
}
