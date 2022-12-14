package postgres

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
)

type Date struct {
	Format string
	time.Time
}

func (Date *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	Date.Format = "2006-01-02"
	t, _ := time.Parse(Date.Format, s)
	Date.Time = t
	return nil
}

func (Date *Date) String() string {
	return Date.Time.Format(Date.Format)
}

func (Date *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(Date.Time.Format(Date.Format))
}

func (Date *Date) Scan(value interface{}) error {
	Date.Format = "2006-01-02"
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
	return Date.Time.Format(Date.Format), nil
}

func (Date Date) Validate(strfmt.Registry) error {
	return nil
}
