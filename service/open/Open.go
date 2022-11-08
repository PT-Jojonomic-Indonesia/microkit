package main

import (
	"context"

	"github.com/PT-Jojonomic-Indonesia/microkit/service"
)

type OpenInput struct {
	Nama            string  `json:"nama"`
	CIF             string  `json:"cif"`
	NomorRekening   string  `json:"nomor_rekening"`
	JumlahGram      float64 `json:"jumlah_gram"`
	HargaJual       float64 `json:"harga_jual"`
	TotalPembayaran float64 `json:"total_pembayaran"`
}

func Open(ctx context.Context, input OpenInput) (resp any) {

	hargaEmas := service.GetHargaEmas(ctx)

	resp = map[string]any{
		"status": "success",
		"ok":     3,
		"harga":  hargaEmas,
	}
	return
}
