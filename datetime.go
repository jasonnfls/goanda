package goanda

import (
	"time"
)

type DateTime time.Time

func (self *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*self).In(time.UTC).Format("\"2006-01-02T15:04:05.999999999Z\"")), nil
}

func (self *DateTime) UnmarshalJSON(data []byte) error {
	dt, err := time.ParseInLocation("\"2006-01-02T15:04:05.999999999Z\"", string(data), time.UTC)
	if err == nil {
		*self = DateTime(dt)
	}
	return err
}
