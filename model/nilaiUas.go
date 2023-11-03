package model

type NilaiUas struct {
	ID         uint   `gorm:"primaryKey" json:"id_nilai_uas"`
	Nim        int    `gorm:"not null" json:"nim_mahasiswa"`
	KodeMatkul string `gorm:"not null" json:"kode_matakuliah"`
	NilaiUas   int    `gorm:"not null" json:"nilai_uas"`
}

type NilaiUasCreateRequest struct {
	Nim        int    `json:"nim_mahasiswa" valid:"required~Student NIM is required"`
	KodeMatkul string `json:"kode_matakuliah" valid:"required~Course code is required"`
	NilaiUas   int    `json:"nilai_uas" valid:"required~Final exam score is required"`
}

type NilaiUasCreateResponse struct {
	ID         uint   `json:"id_nilai_uas"`
	Nim        int    `json:"nim_mahasiswa"`
	KodeMatkul string `json:"kode_matakuliah"`
	NilaiUas   int    `json:"nilai_uas"`
}

type NilaiUasGetResponse struct {
	ID         uint   `json:"id_nilai_uas"`
	Nim        int    `json:"nim_mahasiswa"`
	KodeMatkul string `json:"kode_matakuliah"`
	NilaiUas   int    `json:"nilai_uas"`
}

type NilaiUasUpdateRequest struct {
	NilaiUas int `json:"nilai_uas" valid:"required~Final exam score is required"`
}

type NilaiUasUpdateResponse struct {
	ID         uint   `json:"id_nilai_uas"`
	Nim        int    `json:"nim_mahasiswa"`
	KodeMatkul string `json:"kode_matakuliah"`
	NilaiUas   int    `json:"nilai_uas"`
}

type NilaiUasDeleteResponse struct {
	Message string `json:"message"`
}
