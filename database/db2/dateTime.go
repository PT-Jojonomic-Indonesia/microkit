package db2

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
)

type DateTime struct {
	Format string
	time.Time
}

func NewDateTime() *DateTime {
	return &DateTime{
		Format: time.RFC3339,
		Time:   time.Now(),
	}
}

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	dt.Format = time.RFC3339
	t, _ := time.Parse(dt.Format, s)
	dt.Time = t
	return nil
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return []byte{}, nil
	}
	return json.Marshal(dt.Time.Format(dt.Format))
}

func (d *DateTime) IsZero() bool {
	if d == nil {
		return true
	}
	return d.Time.IsZero()
}

func (dt *DateTime) Scan(value interface{}) error {
	dt.Format = time.RFC3339

	switch v := value.(type) {
	case []byte:
		d, err := time.Parse(dt.Format, string(v))
		if err != nil {
			return err
		}
		dt.Time = d
	case string:
		d, err := time.Parse(dt.Format, v)
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

func (dt DateTime) Value() (driver.Value, error) {
	dateStr := dt.Time.Format(dt.Format)
	if dateStr == "" {
		return nil, nil
	}

	dateStr = strings.Split(dateStr, "+")[0]
	dateStr = strings.TrimRight(dateStr, "Z")

	dateStr = strings.Replace(dateStr, "T", "-", 1)
	dateStr = strings.Replace(dateStr, ":", ".", 2)

	return dateStr, nil
}

func (Date DateTime) Validate(strfmt.Registry) error {
	return nil
}
