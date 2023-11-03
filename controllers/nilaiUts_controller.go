package controllers

import (
	"mahasiswa/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNilaiUts(ctx *gin.Context) {
	nilai := model.NilaiUtsCreateRequest{}

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

	nilaiUts := model.NilaiUts{
		Nim:        nilai.Nim,
		KodeMatkul: nilai.KodeMatkul,
		NilaiUts:   nilai.NilaiUts,
	}

	if err := db.Create(&nilaiUts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	response := model.NilaiUtsCreateResponse{
		ID:         nilaiUts.Id,
		Nim:        nilaiUts.Nim,
		KodeMatkul: nilaiUts.KodeMatkul,
		NilaiUts:   nilaiUts.NilaiUts,
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
}

func GetNilaiUtsId(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)
	nilai := ctx.Param("id")
	nilaiUts := model.NilaiUts{}

	if err := db.Where("id = ?", nilai).First(&nilaiUts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusBadRequest,
					Message: "Nilai UTS dengan id " + nilai + " tidak ditemukan",
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
		Data: nilaiUts,
	})

}

func GetAllNilaiUts(ctx *gin.Context) {

	nilai := []model.NilaiUts{}
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

func UpdateNilaiUts(ctx *gin.Context) {
	id := ctx.Param("id")

	db := ctx.MustGet("db").(*gorm.DB)
	var nilai model.NilaiUts

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
		Data: model.NilaiUtsUpdateResponse{
			ID:         nilai.Id,
			Nim:        nilai.Nim,
			KodeMatkul: nilai.KodeMatkul,
			NilaiUts:   nilai.NilaiUts,
		},
	})
}

func DeletenilaUts(ctx *gin.Context) {
	id := ctx.Param("id")

	db := ctx.MustGet("db").(*gorm.DB)
	nilai := model.NilaiUts{}

	if err := db.Where("id = ?", id).First(&nilai).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: "Nilai UTS dengan id " + id + " tidak ditemukan",
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
		Data: "Nilai dengan ID " + id + " telah dihapus",
	})
}
