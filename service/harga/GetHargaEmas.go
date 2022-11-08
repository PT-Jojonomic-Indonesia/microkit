package main

import "context"

func GetHargaEmas(ctx context.Context) (resp any) {
	resp = map[string]any{
		"buyback": 850000,
		"jual":    870000,
	}
	return
}
