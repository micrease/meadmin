package datetime

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

//15619327487
type DBTime time.Time

func Now() DBTime {
	return DBTime(time.Now())
}

func (t *DBTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = DBTime(t1)
	return err
}

func (t DBTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t DBTime) MarshalBinary() (data []byte, err error) {
	formatted := fmt.Sprintf("%v", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t DBTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *DBTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = DBTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *DBTime) String() string {
	str := fmt.Sprintf("%v", time.Time(*t).Format("2006-01-02 15:04:05"))
	return str
}
