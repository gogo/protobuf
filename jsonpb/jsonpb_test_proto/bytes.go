package jsonpb

// Byte is used to test that []byte type aliases are serialized to base64.
type Byte byte

// Bytes is used to test that []byte type aliases are serialized to base64.
type Bytes []Byte

// FourByteA is used to test that array cast types behave like slice types.
type FourByteA [4]byte
