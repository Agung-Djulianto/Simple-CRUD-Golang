package model

type NilaiUts struct {
	Id         uint   `gorm:"primaryKey" json:"id_nilai_uts"`
	Nim        int    `gorm:"not null" json:"nim_mahasiswa"`
	KodeMatkul string `gorm:"not null" json:"kode_matakuliah"`
	NilaiUts   int    `gorm:"not null" json:"nilai_uts"`
}

type NilaiUtsCreateRequest struct {
	Nim        int    `json:"nim_mahasiswa" valid:"required~Student NIM is required"`
	KodeMatkul string `json:"kode_matakuliah" valid:"required~Course code is required"`
	NilaiUts   int    `json:"nilai_uts" valid:"required~Final exam score is required"`
}

type NilaiUtsCreateResponse struct {
	ID         uint   `json:"id_nilai_uts"`
	Nim        int    `json:"nim_mahasiswa"`
	KodeMatkul string `json:"kode_matakuliah"`
	NilaiUts   int    `json:"nilai_uts"`
}

type NilaiUtsGetResponse struct {
	ID         uint   `json:"id_nilai_uts"`
	Nim        int    `json:"nim_mahasiswa"`
	KodeMatkul string `json:"kode_matakuliah"`
	NilaiUts   int    `json:"nilai_uts"`
}

type NilaiUtsUpdateRequest struct {
	NilaiUts int `json:"nilai_uts" valid:"required~Final exam score is required"`
}

type NilaiUtsUpdateResponse struct {
	ID         uint   `json:"id_nilai_uts"`
	Nim        int    `json:"nim_mahasiswa"`
	KodeMatkul string `json:"kode_matakuliah"`
	NilaiUts   int    `json:"nilai_uts"`
}

type NilaiUatDeleteResponse struct {
	Message string `json:"message"`
}
