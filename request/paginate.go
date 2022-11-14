package request

import (
	"net/http"
	"strconv"
)

func GetPaginateQuery(r *http.Request) (page, offset, limit int) {
	q := r.URL.Query()
	page, offset, limit, rPage, rLimit := 1, 0, 10, q.Get("page"), q.Get("limit")

	if rPage != "" && rLimit != "" {
		pageInt, errParsePage := strconv.Atoi(rPage)
		limitInt, errParseLimit := strconv.Atoi(rLimit)

		if errParsePage == nil && errParseLimit == nil {
			offset = (pageInt - 1) * limitInt
			limit = limitInt
			page = pageInt
		}
	}

	return page, offset, limit
}
