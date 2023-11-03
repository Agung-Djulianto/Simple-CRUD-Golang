package controllers

import (
	"mahasiswa/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateIpMahasiswa(ctx *gin.Context) {
	newIp := model.IpMahasiswaCreateRequest{}

	db := ctx.MustGet("db").(*gorm.DB)

	if err := ctx.ShouldBindJSON(&newIp); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	Ip := model.IPMahasiswa{
		Nim:      newIp.Nim,
		Semester: newIp.Semester,
		Ip:       newIp.IP,
	}

	err := db.Create(&Ip).Error
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

	response := model.IpMahasiswaCreateResponse{
		ID:       Ip.Id,
		Nim:      Ip.Nim,
		Semester: Ip.Semester,
		IP:       Ip.Ip,
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
}

func GetIpMahasiswaId(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)
	IpMahasiswa := ctx.Param("id")
	Ip := model.IPMahasiswa{}

	if err := db.Where("id = ?", IpMahasiswa).First(&Ip).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusBadRequest,
					Message: "IP dengan id " + IpMahasiswa + " tidak ditemukan",
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
		Data: Ip,
	})

}

func GetAllIpMahasiswa(ctx *gin.Context) {
	Ip := []model.IPMahasiswa{}
	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Find(&Ip).Error; err != nil {
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
		Data: Ip,
	})
}

func UpdateIpMahasiswa(ctx *gin.Context) {
	Ip := ctx.Param("id")

	db := ctx.MustGet("db").(*gorm.DB)
	var Ipmahasiswa model.IPMahasiswa

	if err := db.Where("id = ?", Ip).First(&Ipmahasiswa).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&Ipmahasiswa); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if err := db.Save(&Ipmahasiswa).Error; err != nil {
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
		Data: model.IpMahasiswaUpdateResponse{
			ID:       Ipmahasiswa.Id,
			Nim:      Ipmahasiswa.Nim,
			Semester: Ipmahasiswa.Semester,
			IP:       Ipmahasiswa.Ip,
		},
	})
}

func DeleteIpMahasiswa(ctx *gin.Context) {

	Ip := ctx.Param("id")

	db := ctx.MustGet("db").(*gorm.DB)
	Ipmahasiswa := model.IPMahasiswa{}
	if err := db.Where("id = ?", Ip).First(&Ipmahasiswa).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: "IP dengan id " + Ip + " tidak ditemukan",
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

	if err := db.Delete(&Ipmahasiswa).Error; err != nil {
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
		Data: "IP dengan id " + Ip + " telah dihapus",
	})
}
