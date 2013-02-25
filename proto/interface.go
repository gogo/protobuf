package proto

type Union interface {
	GetField() interface{}
	SetField(v interface{})
}
