package entity

type Nasabah struct {
	Nama     string `json:"nama" db:"NAMA"`
	Alamat   string `json:"alamat" db:"ALAMAT"`
	NoTelpon string `json:"no_telpon" db:"NO_TELPON"`
}
