package model

type Dosen struct {
	Nip        string       `gorm:"primaryKey;size:(12)" json:"nip_dosen"`
	NamaDosen  string       `gorm:"not null;unique;type:varchar(30)" json:"nama_dosen"`
	KotaAsal   string       `gorm:"not null;type:varchar(30)" json:"kota_asal"`
	Matakuliah []Matakuliah `gorm:"foreignKey:NipDosen" json:"matakuliah"`
}

type DosenCreateRequest struct {
	Nip       string `json:"nip_dosen" valid:"required~Nip is required"`
	NamaDosen string `json:"nama_dosen" valid:"required~Name is required"`
	KotaAsal  string `json:"kota_asal" valid:"required~City of origin is required"`
}

type DosenCreateResponse struct {
	Nip       string `json:"nip_dosen"`
	NamaDosen string `json:"nama_dosen"`
	KotaAsal  string `json:"kota_asal"`
}

type DosenGetResponse struct {
	Nip        string       `json:"nip_dosen"`
	NamaDosen  string       `json:"nama_dosen"`
	KotaAsal   string       `json:"kota_asal"`
	Matakuliah []Matakuliah `json:"matakuliah"`
}

type DosenUpdateRequest struct {
	NamaDosen string `json:"nama_dosen" valid:"required~Name is required"`
	KotaAsal  string `json:"kota_asal" valid:"required~City of origin is required"`
}

type DosenUpdateResponse struct {
	Nip       string `json:"nip_dosen"`
	NamaDosen string `json:"nama_dosen"`
	KotaAsal  string `json:"kota_asal"`
}

type DosenDeleteResponse struct {
	Message string `json:"message"`
}
