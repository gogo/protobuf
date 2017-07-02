package issue260

import "time"

type Dropped struct {
	Name string `protobuf:"bytes,1,opt,name=name,json=name"`
	Age  int32  `protobuf:"varint,2,opt,name=age,json=age"`
}

func (d *Dropped) Drop() bool {
	return true
}

type DroppedWithoutGetters struct {
	Width             int64      `protobuf:"varint,1,opt,name=width,json=width"`
	Height            int64      `protobuf:"varint,2,opt,name=height,json=height"`
	Timestamp         time.Time  `protobuf:"bytes,3,opt,name=timestamp,stdtime" json:"timestamp"`
	NullableTimestamp *time.Time `protobuf:"bytes,4,opt,name=nullable_timestamp,json=nullableTimestamp,stdtime" json:"nullable_timestamp,omitempty"`
}
