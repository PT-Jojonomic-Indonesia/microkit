package postgres

import (
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func Init(dsn string, config *gorm.Config) error {

	var initError error

	once.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), config)
		if err != nil {
			initError = err
			return
		}
		DB = db
	})

	return initError
}

var Health = func() error {
	return DB.Raw("select version()").Error
}

var Migrate = func(models ...interface{}) error {
	return DB.AutoMigrate(models...)
}
