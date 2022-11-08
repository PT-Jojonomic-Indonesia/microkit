package service

import (
	"context"
	"os"

	"bitbucket.org/jojocoders/microkit/request"
)

func init() {

}

type HargaEmas struct {
	Error bool `json:"error"`
	Data  struct {
		Buyback int `json:"buyback"`
		Jual    int `json:"jual"`
	} `json:"data"`
}

func GetHargaEmas(ctx context.Context) (resp HargaEmas) {
	url := os.Getenv("HARGA_EMAS_ENDPOINT")
	request.Get(ctx, url+"/get-harga-emas", &resp)
	return
}
