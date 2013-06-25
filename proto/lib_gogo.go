package proto

import (
	"encoding/json"
	"strconv"
)

func MarshalJSONEnum(m map[int32]string, value int32) ([]byte, error) {
	s, ok := m[value]
	if !ok {
		s = strconv.Itoa(int(value))
	}
	return json.Marshal(s)
}
