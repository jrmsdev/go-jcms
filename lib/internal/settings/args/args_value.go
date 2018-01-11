package args

import (
	"strconv"
)

type Value struct {
	val string
}

func newValue(val string) *Value {
	return &Value{val}
}

func (v *Value) String() string {
	return v.val
}

func (v *Value) Int() (int, error) {
	return strconv.Atoi(v.val)
}

func (v *Value) Int64() (int64, error) {
	return strconv.ParseInt(v.val, 10, 64)
}

func (v *Value) Float() (float64, error) {
	return strconv.ParseFloat(v.val, 64)
}

func (v *Value) Bool() (bool, error) {
	return strconv.ParseBool(v.val)
}
