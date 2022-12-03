package data

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"
)

type (
	NullTime sql.NullTime
	Time     time.Time
	Bool     bool
)

// DB

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

func (t Time) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format("2006-01-02 15:04:05")), nil
}

func (b *Bool) Scan(value interface{}) error {
	val := value.(int64)
	if val == 1 {
		*b = true
	} else {
		*b = false
	}
	return nil
}

// JSON

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

func (b Bool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}
