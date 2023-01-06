package lib

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Time struct {
	time.Time
}

// MarshalJSON 序列化为JSON
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// UnmarshalJSON 反序列化为JSON
func (t *Time) UnmarshalJSON(data []byte) error {
	var err error
	t.Time, err = time.Parse("2006-01-02 15:04:05", string(data)[1:20])
	return err
}

// String 重写String方法
func (t *Time) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

// FieldType 数据类型
func (t *Time) FieldType() int {
	return orm.TypeDateTimeField

}

// SetRaw 读取数据库值
func (t *Time) SetRaw(value interface{}) error {
	switch value.(type) {
	case time.Time:
		t.Time = value.(time.Time)
	}
	return nil
}

// RawValue 写入数据库
func (t *Time) RawValue() interface{} {
	str := t.Format("2006-01-02 15:04:05")
	if str == "0001-01-01 00:00:00" {
		return nil
	}
	return str
}

// 这里给写入数据库时间时使用
func NowDbTime() Time {
	libTime := Time{}
	libTime.Time = time.Now()
	return libTime
}
func (t *Time) AddDates(years int, months int, days int) Time {
	libTime := Time{}
	libTime.Time = t.Time.AddDate(years, months, days)
	return libTime
}
func TimeParse(value string) (Time, error) {
	libTime := Time{}
	times, err := time.Parse("2006-01-02 15:04:05", value)
	libTime.Time = times
	return libTime, err
}
func (t *Time) GetStr() string {
	return t.Format("2006-01-02 15:04:05")
}
