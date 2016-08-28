package types

func NewPopulatedTimestamp(r interface {
	Int63() int64
}, easy bool) *Timestamp {
	this := &Timestamp{}
	ns := int64(r.Int63())
	this.Seconds = ns / 1e9
	this.Nanos = int32(ns % 1e9)
	return this
}
