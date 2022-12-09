package db2

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
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

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	d.Format = DateLayout
	t, _ := time.Parse(d.Format, s)
	d.Time = t
	return nil
}

func (d *Date) IsZero() bool {
	if d == nil {
		return true
	}
	return d.Time.IsZero()
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(d.Format))
}

func (d *Date) Scan(value interface{}) error {
	d.Format = DateLayout
	switch v := value.(type) {
	case []byte:
		date, err := time.Parse(d.Format, string(v))
		if err != nil {
			return err
		}
		d.Time = date
	case string:
		date, err := time.Parse(d.Format, v)
		if err != nil {
			return err
		}
		d.Time = date
	case time.Time:
		d.Time = v
	case nil:
		d.Time = time.Time{}
	default:
		return fmt.Errorf("cannot sql.Scan() time.Time from: %#v", v)
	}
	return nil
}

func (d Date) Value() (driver.Value, error) {
	dateStr := d.Time.Format(d.Format)
	if dateStr == "" {
		return nil, nil
	}

	return dateStr, nil
}

func (d Date) Validate(strfmt.Registry) error {
	return nil
}
