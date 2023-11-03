package model

type Mahasiswa struct {
	Nim         int           `gorm:"primaryKey;size(12)" json:"nim_mahasiswa"`
	Nama        string        `gorm:"not null;unique;type:varchar(30)" json:"nama_mahasiswa"`
	KotaAsal    string        `gorm:"not null;type:varchar(15)" json:"kota_asal"`
	TahunMasuk  int           `gorm:"not null;size(6)" json:"tahun_masuk"`
	NilaiUas    []NilaiUas    `gorm:"foreignKey:Nim" json:"nilai_uas"`
	NilaiUts    []NilaiUts    `gorm:"foreignKey:Nim" json:"nilai_uts"`
	IPMahasiswa []IPMahasiswa `gorm:"foreignKey:Nim" json:"ip_mahasiswa"`
}

type MahasiswaCreateRequest struct {
	Nim        int    `json:"nim" valid:"required~Nim is required"`
	Nama       string `json:"nama_mahasiswa" valid:"required~Name is required"`
	KotaAsal   string `json:"kota_asal" valid:"required~City of origin is required"`
	TahunMasuk int    `json:"tahun_masuk" valid:"required~Year of admission is required"`
}

type MahasiswaUpdateRequest struct {
	Nama       string `json:"nama_mahasiswa" valid:"required~Name is required"`
	KotaAsal   string `json:"kota_asal" valid:"required~Name is required"`
	TahunMasuk int    `json:"tahun_masuk" valid:"required~Year of admission is required"`
}

type MahasiswacreateResponse struct {
	Nim        int    `json:"nim_mahasiswa"`
	Nama       string `json:"nama_mahasiswa"`
	KotaAsal   string `json:"kota_asal"`
	TahunMasuk int    `json:"tahun_masuk"`
}

type MahasiswaGetResponse struct {
	Nim         int           `json:"nim_mahasiswa"`
	Nama        string        `json:"nama_mahasiswa"`
	KotaAsal    string        `json:"kota_asal"`
	TahunMasuk  int           `json:"tahun_masuk"`
	NilaiUts    []NilaiUts    `json:"nilai_uts"`
	NilaiUas    []NilaiUas    `json:"nilai_uas"`
	IpMahasiswa []IPMahasiswa `json:"ip_mahasiswa"`
}

type MahasiswaUpdateResponse struct {
	Nim        int    `json:"nim_mahasiswa"`
	Nama       string `json:"nama_mahasiswa"`
	KotaAsal   string `json:"kota_asal"`
	TahunMasuk int    `json:"tahun_masuk"`
}

type MahasiswaDeleteResponse struct {
	Message string `json:"message"`
}
