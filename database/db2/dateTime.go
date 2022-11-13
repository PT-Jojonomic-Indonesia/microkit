package db2

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ISODateTime struct {
	Format string
	time.Time
}

func NewISODateTime() *ISODateTime {
	return &ISODateTime{
		Format: time.RFC3339,
		Time:   time.Now(),
	}
}

func (Date *ISODateTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	Date.Format = time.RFC3339
	t, _ := time.Parse(Date.Format, s)
	Date.Time = t
	return nil
}

func (Date *ISODateTime) MarshalJSON() ([]byte, error) {
	if Date.Time.IsZero() {
		return []byte{}, nil
	}
	return json.Marshal(Date.Time.Format(Date.Format))
}

func (Date *ISODateTime) Scan(value interface{}) error {
	Date.Format = time.RFC3339

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

func (Date ISODateTime) Value() (driver.Value, error) {
	dateStr := Date.Time.Format(Date.Format)
	if dateStr == "" {
		return nil, nil
	}

	dateStr = strings.Split(dateStr, "+")[0]
	dateStr = strings.TrimRight(dateStr, "Z")

	dateStr = strings.Replace(dateStr, "T", "-", 1)
	dateStr = strings.Replace(dateStr, ":", ".", 2)

	return dateStr, nil
}
