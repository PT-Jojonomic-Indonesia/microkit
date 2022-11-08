package entity

// default generate table query is nasabahs
type Nasabah struct {
	ID       int64  `json:"id" gorm:"autoIncrement:false;type:bigint;index:idx_nasabah,primaryKey"` // example custom field in auto migrate
	Nama     string `json:"nama" `
	Alamat   string `json:"alamat"`
	NoTelpon string `json:"no_telpon" gorm:"type:varchar(15);"`
}

// example custom table name
func (Nasabah) TableName() string {
	return "nasabah"
}
