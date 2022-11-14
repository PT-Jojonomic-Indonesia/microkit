package response

import "math"

type BaseResponse struct {
	Error    bool        `json:"error"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

type BasePaginate struct {
	Page      int `json:"page"`
	TotalData int `json:"totalData"`
	TotalPage int `json:"totalPage"`
}

func NewPaginate(page, limit, totalData int) BasePaginate {
	totalPage := math.Ceil(float64(totalData) / float64(limit))
	return BasePaginate{
		Page:      page,
		TotalData: totalData,
		TotalPage: int(totalPage),
	}
}
