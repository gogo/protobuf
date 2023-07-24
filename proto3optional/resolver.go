package proto3optional

import (
	"fmt"
	"strconv"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

type Resolver struct {
	proto3     bool
	fakeOneOfs map[string]struct{}
	allOneOfs  map[string]*int32
}

func NewResolver(proto3 bool, fields []*descriptor.FieldDescriptorProto) *Resolver {
	if !proto3 {
		return &Resolver{
			proto3: false,
		}
	}

	r := &Resolver{
		proto3:     true,
		fakeOneOfs: make(map[string]struct{}),
		allOneOfs:  make(map[string]*int32),
	}

	tmp := make(map[int32][]*descriptor.FieldDescriptorProto)
	for _, f := range fields {
		r.allOneOfs[getMapKey(f)] = f.OneofIndex

		if f.OneofIndex != nil {
			tmp[*f.OneofIndex] = append(tmp[*f.OneofIndex], f)
		}
	}

	for _, v := range tmp {
		if len(v) == 1 && v[0].GetProto3Optional() {
			r.fakeOneOfs[getMapKey(v[0])] = struct{}{}
		}
	}

	return r
}

func (r *Resolver) IsFakeOneOf(f *descriptor.FieldDescriptorProto) bool {
	if !r.proto3 {
		return false
	}

	_, ok := r.fakeOneOfs[getMapKey(f)]
	return ok
}

func (r *Resolver) IsRealOneOf(f *descriptor.FieldDescriptorProto) bool {
	return !r.IsFakeOneOf(f) && r.allOneOfs[getMapKey(f)] != nil
}

func (r *Resolver) IsProto3WithoutOptional(f *descriptor.FieldDescriptorProto) bool {
	return r.proto3 && !r.IsFakeOneOf(f)
}

func getMapKey(f *descriptor.FieldDescriptorProto) string {
	if f.GetNumber() == 0 {
		panic(fmt.Sprintf("field %s has no number", f.GetName()))
	}

	return strconv.Itoa(int(f.GetNumber())) + f.GetType().String() + f.GetName()
}
