package db2

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

var DateLayout = "2006-01-02"

type Date struct {
	Format string
	time.Time
}

func NewDate() Date {
	return Date{
		Format: DateLayout,
		Time:   time.Now(),
	}
}

func (Date *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	Date.Format = DateLayout
	t, _ := time.Parse(Date.Format, s)
	Date.Time = t
	return nil
}

func (Date *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(Date.Time.Format(Date.Format))
}

func (Date *Date) Scan(value interface{}) error {
	Date.Format = DateLayout
	switch v := value.(type) {
	case []byte:
		d, err := time.Parse(Date.Format, string(v))
		if err != nil {
			return err
		}
		Date.Time = d
	case string:
		d, err := time.Parse(Date.Format, v)
		if err != nil {
			return err
		}
		Date.Time = d
	case time.Time:
		Date.Time = v
	case nil:
		Date.Time = time.Time{}
	default:
		return fmt.Errorf("cannot sql.Scan() time.Time from: %#v", v)
	}
	return nil
}

func (Date Date) Value() (driver.Value, error) {
	dateStr := Date.Time.Format(Date.Format)
	if dateStr == "" {
		return nil, nil
	}

	return dateStr, nil
}
