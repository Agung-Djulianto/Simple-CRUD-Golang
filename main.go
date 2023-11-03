package main

import (
	database "mahasiswa/Database"
	"mahasiswa/model"
	"mahasiswa/routers"
)

func main() {
	db := database.ReadDB()
	db.AutoMigrate(model.Mahasiswa{}, model.Dosen{}, model.Matakuliah{}, model.NilaiUas{}, model.NilaiUts{}, model.IPMahasiswa{})

	x := routers.ReadRouters(db)
	x.Run(":8088")
}
