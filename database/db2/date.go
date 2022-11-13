package db2

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

var ISODateLayout = "2006-01-02"

type ISODate struct {
	Format string
	time.Time
}

func NewISODate() ISODate {
	return ISODate{
		Format: ISODateLayout,
		Time:   time.Now(),
	}
}

func (Date *ISODate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	Date.Format = ISODateLayout
	t, _ := time.Parse(Date.Format, s)
	Date.Time = t
	return nil
}

func (Date *ISODate) MarshalJSON() ([]byte, error) {
	return json.Marshal(Date.Time.Format(Date.Format))
}

func (Date *ISODate) Scan(value interface{}) error {
	Date.Format = ISODateLayout
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

func (Date ISODate) Value() (driver.Value, error) {
	dateStr := Date.Time.Format(Date.Format)
	if dateStr == "" {
		return nil, nil
	}

	return dateStr, nil
}
