package model

type Matakuliah struct {
	KodeMatkul string `gorm:"primaryKey;type:varchar(10)" json:"kode_matakuliah"`
	NamaMatkul string `gorm:"not null" json:"nama_matakuliah"`
	NipDosen   string `gorm:"not null" json:"nip_dosen"`
	JumlahSks  int    `gorm:"not null" json:"jumlah_sks"`
}

type MatakuliahCreateRequest struct {
	KodeMatkul string `json:"kode_matakuliah" valid:"required~Course code is required"`
	NamaMatkul string `json:"nama_matakuliah" valid:"required~Course name is required"`
	NipDosen   string `json:"nip_dosen" valid:"required~Lecturer NIP is required"`
	JumlahSks  int    `json:"jumlah_sks" valid:"required~Credit hours is required"`
}

type MatakuliahCreateResponse struct {
	KodeMatkul string `json:"kode_matakuliah"`
	NamaMatkul string `json:"nama_matakuliah"`
	NipDosen   string `json:"nip_dosen"`
	JumlahSks  int    `json:"jumlah_sks"`
}

type MatakuliahGetResponse struct {
	KodeMatkul string `json:"kode_matakuliah"`
	NamaMatkul string `json:"nama_matakuliah"`
	NipDosen   string `json:"nip_dosen"`
	JumlahSks  int    `json:"jumlah_sks"`
}

type MatakuliahUpdateRequest struct {
	NamaMatkul string `json:"nama_matakuliah" valid:"required~Course name is required"`
	NipDosen   string `json:"nip_dosen" valid:"required~Lecturer NIP is required"`
	JumlahSks  int    `json:"jumlah_sks" valid:"required~Credit hours is required"`
}

type MatakuliahUpdateResponse struct {
	KodeMatkul string `json:"kode_matakuliah"`
	NamaMatkul string `json:"nama_matakuliah"`
	NipDosen   string `json:"nip_dosen"`
	JumlahSks  int    `json:"jumlah_sks"`
}

type MatakuliahDeleteResponse struct {
	Message string `json:"message"`
}
