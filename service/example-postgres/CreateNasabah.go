package main

import (
	"context"
	"time"

	"github.com/PT-Jojonomic-Indonesia/microkit/database/postgres"
	"github.com/PT-Jojonomic-Indonesia/microkit/service/example-postgres/entity"
)

var CreateNasabah = func(ctx context.Context, data *entity.Nasabah) error {
	data.ID = time.Now().UnixNano()
	return postgres.DB.Model(data).WithContext(ctx).Save(data).Error
}
