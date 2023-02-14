package db2

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
)

var DefaultTimestampLayout = "2006-01-02-15.04.05.000000"

type Timestamp struct {
	Format string
	time.Time
}

func NewTimestamp() *Timestamp {
	return &Timestamp{
		Format: DefaultTimestampLayout,
		Time:   time.Now(),
	}
}

func (dt *Timestamp) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	dt.Format = time.RFC3339
	t, _ := time.Parse(dt.Format, s)
	dt.Time = t
	return nil
}

func (dt *Timestamp) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return []byte{}, nil
	}
	return json.Marshal(dt.Time.Format(dt.Format))
}

func (d *Timestamp) IsZero() bool {
	if d == nil {
		return true
	}
	return d.Time.IsZero()
}

func (dt *Timestamp) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		d, err := time.Parse(DefaultTimestampLayout, string(v))
		if err != nil {
			return err
		}
		dt.Time = d
	case string:
		d, err := time.Parse(DefaultTimestampLayout, v)
		if err != nil {
			return err
		}
		dt.Time = d
	case time.Time:
		dt.Time = v
	case nil:
		dt.Time = time.Time{}
	default:
		return fmt.Errorf("cannot sql.Scan() time.Time from: %#v", v)
	}
	return nil
}

func (dt Timestamp) Value() (driver.Value, error) {
	dateStr := dt.Time.Format(DefaultTimestampLayout)
	return dateStr, nil
}

func (Date Timestamp) Validate(strfmt.Registry) error {
	return nil
}
