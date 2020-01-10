package test

type ProtoTypes []*ProtoType

func NewPopulatedProtoTypes(r randyThetest) *ProtoTypes {
	return &ProtoTypes{&ProtoType{}}
}
