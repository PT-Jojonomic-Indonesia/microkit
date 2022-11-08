package main

import (
	"context"

	"github.com/PT-Jojonomic-Indonesia/microkit/database/db2"
	"github.com/PT-Jojonomic-Indonesia/microkit/service/example-db2/entity"
)

var CreateNasabah = func(ctx context.Context, data *entity.Nasabah) error {
	_, err := db2.NamedExec(ctx, "INSERT INTO testdb.nasabah (nama, alamat, no_telpon) VALUES (:NAMA, :ALAMAT, :NO_TELPON)", data)
	return err
}
