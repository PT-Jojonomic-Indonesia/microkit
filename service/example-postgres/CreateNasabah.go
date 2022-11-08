package main

import (
	"context"
	"time"

	"bitbucket.org/jojocoders/microkit/database/postgres"
	"bitbucket.org/jojocoders/microkit/service/example-postgres/entity"
)

var CreateNasabah = func(ctx context.Context, data *entity.Nasabah) error {
	data.ID = time.Now().UnixNano()
	return postgres.DB.Model(data).WithContext(ctx).Save(data).Error
}
