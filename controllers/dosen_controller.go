package controllers

import (
	"mahasiswa/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDosen(ctx *gin.Context) {

	newDosen := model.DosenCreateRequest{}

	db := ctx.MustGet("db").(*gorm.DB)

	if err := ctx.ShouldBindJSON(&newDosen); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	dosen := model.Dosen{
		Nip:       newDosen.Nip,
		NamaDosen: newDosen.NamaDosen,
		KotaAsal:  newDosen.KotaAsal,
	}

	err := db.Create(&dosen).Error
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

	response := model.DosenCreateResponse{
		Nip:       dosen.Nip,
		NamaDosen: dosen.NamaDosen,
		KotaAsal:  dosen.KotaAsal,
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
}

func GetDosenId(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)
	nipDosen := ctx.Param("nip")

	Dosen := model.Dosen{}

	if err := db.Preload("Matakuliah").Where("nip = ?", nipDosen).First(&Dosen).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusBadRequest,
					Message: "Dosen dengan NIP " + nipDosen + " tidak ditemukan",
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
		Data: Dosen,
	})
}

func GetAllDosen(ctx *gin.Context) {

	dosen := []model.Dosen{}

	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Preload("Matakuliah").Find(&dosen).Error; err != nil {
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
		Data: dosen,
	})
}

func UpdateDosen(ctx *gin.Context) {
	nim := ctx.Param("nip")

	db := ctx.MustGet("db").(*gorm.DB)
	dosen := model.Dosen{}

	if err := db.Where("nip= ?", nim).First(&dosen).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&dosen); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := db.Save(&dosen).Error; err != nil {
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
		Data: model.DosenUpdateResponse{
			Nip:       dosen.Nip,
			NamaDosen: dosen.NamaDosen,
			KotaAsal:  dosen.KotaAsal,
		},
	})
}

func DeleteDosen(ctx *gin.Context) {
	nip := ctx.Param("nip")

	db := ctx.MustGet("db").(*gorm.DB)
	dosen := model.Dosen{}

	if err := db.Where("nip = ?", nip).First(&dosen).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: "Dosen dengan NIP " + nip + " tidak ditemukan",
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

	if err := db.Delete(&dosen).Error; err != nil {
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
		Data: "Dosen dengan NIP " + nip + " telah dihapus",
	})
}
