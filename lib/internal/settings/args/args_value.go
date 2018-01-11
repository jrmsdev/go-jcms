package args

type Value struct {
	val string
}

func newValue(val string) *Value {
	return &Value{val}
}

func (v *Value) String() string {
	return v.val
}
