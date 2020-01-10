package test

type ProtoTypesPointer []*ProtoType
type ProtoTypesNotPointer []ProtoType

func NewPopulatedProtoTypesPointer(r randyThetest) *ProtoTypesPointer {
	return &ProtoTypesPointer{&ProtoType{}}
}

func NewPopulatedProtoTypesNotPointer(r randyThetest) *ProtoTypesNotPointer {
	return &ProtoTypesNotPointer{ProtoType{}}
}
