package routers

import (
	"mahasiswa/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadRouters(db *gorm.DB) *gin.Engine {
	x := gin.Default()
	x.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	mahasiswaGroup := x.Group("/mahasiswa")
	{
		mahasiswaGroup.POST("/create", controllers.CreateMahasiswa)
		mahasiswaGroup.GET("/:nim", controllers.GetMahasiswaId)
		mahasiswaGroup.GET("/", controllers.GetAllMahasiswa)
		mahasiswaGroup.PUT("/:nim", controllers.UpdateMahasiswa)
		mahasiswaGroup.DELETE("/:nim", controllers.DeleteMahasiswa)
	}

	dosenGroup := x.Group("/dosen")
	{
		dosenGroup.POST("/create", controllers.CreateDosen)
		dosenGroup.GET("/:nip", controllers.GetDosenId)
		dosenGroup.GET("/", controllers.GetAllDosen)
		dosenGroup.PUT("/:nip", controllers.UpdateDosen)
		dosenGroup.DELETE("/:nip", controllers.DeleteDosen)
	}

	matkulGroup := x.Group("/matakuliah")
	{
		matkulGroup.POST("/create", controllers.CreateMatkul)
		matkulGroup.GET("/:kode_matkul", controllers.GetMatkulId)
		matkulGroup.GET("/", controllers.GetAllMatkul)
		matkulGroup.PUT("/:kode_matkul", controllers.UpdateMatakuliah)
		matkulGroup.DELETE("/:kode_matkul", controllers.DeleteMatkul)
	}

	nilaiUasGroup := x.Group("/nilai_uas")
	{
		nilaiUasGroup.POST("/create", controllers.CreateNilaiUAs)
		nilaiUasGroup.GET("/:id", controllers.GetNilaiId)
		nilaiUasGroup.GET("/", controllers.GetAllNilai)
		nilaiUasGroup.PUT("/:id", controllers.UpdateNilai)
		nilaiUasGroup.DELETE("/:id", controllers.DeletenilaUas)
	}

	nilaiUtsGroup := x.Group("/nilai_uts")
	{
		nilaiUtsGroup.POST("/create", controllers.CreateNilaiUts)
		nilaiUtsGroup.GET("/:id", controllers.GetNilaiUtsId)
		nilaiUtsGroup.GET("/", controllers.GetAllNilaiUts)
		nilaiUtsGroup.PUT("/:id", controllers.UpdateNilaiUts)
		nilaiUtsGroup.DELETE("/:id", controllers.DeletenilaUts)
	}

	IpGroup := x.Group("/ip_Mahasiswa")
	{
		IpGroup.POST("/create", controllers.CreateIpMahasiswa)
		IpGroup.GET("/:id", controllers.GetIpMahasiswaId)
		IpGroup.GET("/", controllers.GetAllIpMahasiswa)
		IpGroup.PUT("/:id", controllers.UpdateIpMahasiswa)
		IpGroup.DELETE("/:id", controllers.DeleteIpMahasiswa)

	}

	return x
}
