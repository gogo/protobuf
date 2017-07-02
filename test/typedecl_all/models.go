package typedeclall

type Dropped struct {
	Name string `protobuf:"bytes,1,opt,name=name,json=name"`
	Age  int32  `protobuf:"varint,2,opt,name=age,json=age"`
}

func (d *Dropped) Drop() bool {
	return true
}

type DroppedWithoutGetters struct {
	Width  int64 `protobuf:"varint,1,opt,name=width,json=width"`
	Height int64 `protobuf:"varint,2,opt,name=height,json=height"`
}

func (d *DroppedWithoutGetters) GetHeight() int64 {
	return d.Height
}
