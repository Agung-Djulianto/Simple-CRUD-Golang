package controllers

import (
	"mahasiswa/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNilaiUAs(ctx *gin.Context) {

	nilai := model.NilaiUasCreateRequest{}

	db := ctx.MustGet("db").(*gorm.DB)

	if err := ctx.ShouldBindJSON(&nilai); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	nilaiUas := model.NilaiUas{
		Nim:        nilai.Nim,
		KodeMatkul: nilai.KodeMatkul,
		NilaiUas:   nilai.NilaiUas,
	}

	if err := db.Create(&nilaiUas).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	response := model.NilaiUasCreateResponse{
		ID:         nilaiUas.ID,
		Nim:        nilaiUas.Nim,
		KodeMatkul: nilaiUas.KodeMatkul,
		NilaiUas:   nilaiUas.NilaiUas,
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
}

func GetNilaiId(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)
	nilai := ctx.Param("id")
	nilaiUas := model.NilaiUas{}

	if err := db.Where("id = ?", nilai).First(&nilaiUas).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusBadRequest,
					Message: "Nilai UAS dengan id " + nilai + " tidak ditemukan",
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
		Data: nilaiUas,
	})

}

func GetAllNilai(ctx *gin.Context) {

	nilai := []model.NilaiUas{}
	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Find(&nilai).Error; err != nil {
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
		Data: nilai,
	})
}

func UpdateNilai(ctx *gin.Context) {
	id := ctx.Param("id")

	db := ctx.MustGet("db").(*gorm.DB)
	var nilai model.NilaiUas

	if err := db.Where("id = ?", id).First(&nilai).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&nilai); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := db.Save(&nilai).Error; err != nil {
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
		Data: model.NilaiUasUpdateResponse{
			ID:         nilai.ID,
			Nim:        nilai.Nim,
			KodeMatkul: nilai.KodeMatkul,
			NilaiUas:   nilai.NilaiUas,
		},
	})
}

func DeletenilaUas(ctx *gin.Context) {
	id := ctx.Param("id")

	db := ctx.MustGet("db").(*gorm.DB)
	nilai := model.NilaiUas{}

	if err := db.Where("id = ?", id).First(&nilai).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: "Nilai UAS dengan id " + id + " tidak ditemukan",
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

	if err := db.Delete(&nilai).Error; err != nil {
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
		Data: "Nilai UAS dengan ID " + id + " telah dihapus",
	})
}
