package main

import (
	"strings"
	"time"
	"fmt"

	driver "database/sql/driver"
)

const timestampFormat = "2006-01-02 15:04:05"

type Timestamp struct {
	time.Time
}

func (c *Timestamp) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`) // remove quotes
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(timestampFormat, s)
	return
}

func (c Timestamp) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(timestampFormat))), nil
}

func (c Timestamp) Value() (driver.Value, error) {
	return c.Time, nil
}