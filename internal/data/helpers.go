package data

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type (
	NullTime sql.NullTime
	Time     time.Time
)

func (nt *NullTime) Scan(value interface{}) error {
	var t sql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{t.Time, false}
	} else {
		*nt = NullTime{t.Time, true}
	}

	return nil
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format("2006-01-02 15:04:05"))
	return []byte(val), nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	val := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(val), nil
}
