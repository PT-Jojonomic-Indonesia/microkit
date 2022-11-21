package db2

type EntityTrackTime struct {
	CreateBy   *string   `json:"create_by" db:"CREATE_BY"`
	CreateDate *DateTime `json:"create_date" db:"CREATE_DATE"`
	UpdateBy   *string   `json:"update_by" db:"UPDATE_BY"`
	UpdateDate *DateTime `json:"update_date" db:"UPDATE_DATE"`
}
