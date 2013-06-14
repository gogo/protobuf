package unrecognized

import (
	code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"
	math_rand "math/rand"
	"testing"
	time "time"
)

func TestNewOld(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	newer := NewPopulatedA(popr)
	data1, err := code_google_com_p_gogoprotobuf_proto.Marshal(newer)
	if err != nil {
		panic(err)
	}
	older := &OldA{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data1, older); err != nil {
		panic(err)
	}
	data2, err := code_google_com_p_gogoprotobuf_proto.Marshal(older)
	if err != nil {
		panic(err)
	}
	bluer := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data2, bluer); err != nil {
		panic(err)
	}
	if err := newer.VerboseEqual(bluer); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", newer, bluer, err)
	}
}

func TestOldNew(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	older := NewPopulatedOldA(popr)
	data1, err := code_google_com_p_gogoprotobuf_proto.Marshal(older)
	if err != nil {
		panic(err)
	}
	newer := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data1, newer); err != nil {
		panic(err)
	}
	data2, err := code_google_com_p_gogoprotobuf_proto.Marshal(newer)
	if err != nil {
		panic(err)
	}
	bluer := &OldA{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data2, bluer); err != nil {
		panic(err)
	}
	if err := older.VerboseEqual(bluer); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", older, bluer, err)
	}
}

func TestOldNewOldNew(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	older := NewPopulatedOldA(popr)
	data1, err := code_google_com_p_gogoprotobuf_proto.Marshal(older)
	if err != nil {
		panic(err)
	}
	newer := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data1, newer); err != nil {
		panic(err)
	}
	data2, err := code_google_com_p_gogoprotobuf_proto.Marshal(newer)
	if err != nil {
		panic(err)
	}
	bluer := &OldA{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data2, bluer); err != nil {
		panic(err)
	}
	if err := older.VerboseEqual(bluer); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", older, bluer, err)
	}

	data3, err := code_google_com_p_gogoprotobuf_proto.Marshal(bluer)
	if err != nil {
		panic(err)
	}
	purple := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data3, purple); err != nil {
		panic(err)
	}
	data4, err := code_google_com_p_gogoprotobuf_proto.Marshal(purple)
	if err != nil {
		panic(err)
	}
	magenta := &OldA{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data4, magenta); err != nil {
		panic(err)
	}
	if err := older.VerboseEqual(magenta); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", older, magenta, err)
	}
}
