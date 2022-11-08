package entity

type Nasabah struct {
	Nama     string `json:"nama" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
	NoTelpon string `json:"no_telpon" validate:"required"`
}
