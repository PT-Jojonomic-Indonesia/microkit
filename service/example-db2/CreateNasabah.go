package main

import (
	"context"

	"bitbucket.org/jojocoders/microkit/database/db2"
	"bitbucket.org/jojocoders/microkit/service/example-db2/entity"
)

var CreateNasabah = func(ctx context.Context, data *entity.Nasabah) error {
	_, err := db2.NamedExec(ctx, "INSERT INTO testdb.nasabah (nama, alamat, no_telpon) VALUES (:NAMA, :ALAMAT, :NO_TELPON)", data)
	return err
}
