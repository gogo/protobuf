package unmarshalmerge

import (
	"code.google.com/p/gogoprotobuf/proto"
	math_rand "math/rand"
	"testing"
	"time"
)

func TestUnmarshalMerge(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBig(popr, true)
	if p.GetSub() == nil {
		p.Sub = &Sub{SubNumber: proto.Int64(12345)}
	}
	data, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	s := &Sub{}
	b := &Big{
		Sub: s,
	}
	err = proto.UnmarshalMerge(data, b)
	if err != nil {
		panic(err)
	}
	if s.GetSubNumber() != p.GetSub().GetSubNumber() {
		t.Fatalf("should have unmarshaled subnumber into sub")
	}
}

func TestUnsafeUnmarshalMerge(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBigUnsafe(popr, true)
	if p.GetSub() == nil {
		p.Sub = &Sub{SubNumber: proto.Int64(12345)}
	}
	data, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	s := &Sub{}
	b := &BigUnsafe{
		Sub: s,
	}
	err = proto.UnmarshalMerge(data, b)
	if err != nil {
		panic(err)
	}

	if s.GetSubNumber() != p.GetSub().GetSubNumber() {
		t.Fatalf("should have unmarshaled subnumber into sub")
	}
}
