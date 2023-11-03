package model

type IPMahasiswa struct {
	Id       int     `gorm:"primaryKey" json:"id_ip"`
	Nim      int     `gorm:"not null" json:"nim_mahasiswa"`
	Semester int     `gorm:"not null" json:"semester"`
	Ip       float32 `json:"ip"`
}
type IpMahasiswaCreateRequest struct {
	Nim      int     `json:"nim_mahasiswa" valid:"required~Student NIM is required"`
	Semester int     `json:"semester" valid:"required~Semester is required"`
	IP       float32 `json:"ip" valid:"required~IP is required"`
}

type IpMahasiswaCreateResponse struct {
	ID       int     `json:"id_ip"`
	Nim      int     `json:"nim_mahasiswa"`
	Semester int     `json:"semester"`
	IP       float32 `json:"ip"`
}

type IpMahasiswaGetResponse struct {
	ID       int     `json:"id_ip"`
	Nim      int     `json:"nim_mahasiswa"`
	Semester int     `json:"semester"`
	IP       float32 `json:"ip"`
}

type IpMahasiswaUpdateRequest struct {
	IP float32 `json:"ip" valid:"required~IP is required"`
}

type IpMahasiswaUpdateResponse struct {
	ID       int     `json:"id_ip"`
	Nim      int     `json:"nim_mahasiswa"`
	Semester int     `json:"semester"`
	IP       float32 `json:"ip"`
}

type IpMahasiswaDeleteResponse struct {
	Message string `json:"message"`
}
