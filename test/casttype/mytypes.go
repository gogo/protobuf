package casttype

import (
	"encoding/json"
)

type MyInt32Type int32

type MyUint64Type uint64

type Bytes []byte

func (this Bytes) MarshalJSON() ([]byte, error) {
	return json.Marshal([]byte(this))
}

func (this *Bytes) UnmarshalJSON(data []byte) error {
	v := new([]byte)
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	*this = *v
	return nil
}

type MyStringType string

type MyMapType map[string]uint64
