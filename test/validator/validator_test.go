package validatortest

import "testing"

func buildProto3(someString string, someInt uint32, identifier string, someValue int64) *ValidatorMessage3 {
	goodEmbeddedProto3 := &ValidatorMessage3_Embedded{
		Identifier: identifier,
		SomeValue:  someValue,
	}

	goodProto3 := &ValidatorMessage3{
		SomeString:                    someString,
		SomeStringRep:                 []string{someString, "xyz34"},
		SomeInt:                       someInt,
		SomeIntRep:                    []uint32{someInt, 12, 13, 14, 15, 16},
		SomeIntRepNonNull:             []uint32{someInt, 102},
		SomeEmbedded:                  nil,
		SomeEmbeddedNonNullable:       *goodEmbeddedProto3,
		SomeEmbeddedExists:            goodEmbeddedProto3,
		SomeEmbeddedExistsNonNullable: *goodEmbeddedProto3,
		SomeEmbeddedRep:               []*ValidatorMessage3_Embedded{goodEmbeddedProto3},
		SomeEmbeddedRepNonNullable:    []ValidatorMessage3_Embedded{*goodEmbeddedProto3},
	}
	return goodProto3
}

func buildProto2(someString string, someInt uint32, identifier string, someValue int64) *ValidatorMessage {
	goodEmbeddedProto2 := &ValidatorMessage_Embedded{
		Identifier: &identifier,
		SomeValue:  &someValue,
	}

	goodProto2 := &ValidatorMessage{
		StringReq:        &someString,
		StringReqNonNull: someString,

		StringOpt:        nil,
		StringOptNonNull: someString,

		IntReq:        &someInt,
		IntReqNonNull: someInt,
		IntRep:        []uint32{someInt, 12, 13, 14, 15, 16},
		IntRepNonNull: []uint32{someInt, 12, 13, 14, 15, 16},

		EmbeddedReq:            goodEmbeddedProto2,
		EmbeddedNonNull:        *goodEmbeddedProto2,
		EmbeddedRep:            []*ValidatorMessage_Embedded{goodEmbeddedProto2},
		EmbeddedRepNonNullable: []ValidatorMessage_Embedded{*goodEmbeddedProto2},
	}
	return goodProto2
}

func TestGoodProto3(t *testing.T) {
	var err error
	goodProto3 := buildProto3("-%ab", 11, "abba", 99)
	err = goodProto3.Validate()
	if err != nil {
		t.Fatalf("unexpected fail in validator: %v", err)
	}
}

func TestGoodProto2(t *testing.T) {
	var err error
	goodProto2 := buildProto2("-%ab", 11, "abba", 99)
	err = goodProto2.Validate()
	if err != nil {
		t.Fatalf("unexpected fail in validator: %v", err)
	}
}

func TestStringRegex(t *testing.T) {
	tooLong1Proto3 := buildProto3("toolong", 11, "abba", 99)
	if tooLong1Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooLong2Proto3 := buildProto3("-%ab", 11, "bad#", 99)
	if tooLong2Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooLong1Proto2 := buildProto2("toolong", 11, "abba", 99)
	if tooLong1Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	tooLong2Proto2 := buildProto2("-%ab", 11, "bad#", 99)
	if tooLong2Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestIntLowerBounds(t *testing.T) {
	lowerThan10Proto3 := buildProto3("-%ab", 9, "abba", 99)
	if lowerThan10Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan10Proto2 := buildProto2("-%ab", 9, "abba", 99)
	if lowerThan10Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan0Proto3 := buildProto3("-%ab", 11, "abba", -1)
	if lowerThan0Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	lowerThan0Proto2 := buildProto2("-%ab", 11, "abba", -1)
	if lowerThan0Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestIntUpperBounds(t *testing.T) {
	higherThan100Proto3 := buildProto3("-%ab", 11, "abba", 101)
	if higherThan100Proto3.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
	higherThan100Proto2 := buildProto2("-%ab", 11, "abba", 101)
	if higherThan100Proto2.Validate() == nil {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}
}

func TestMsgExist(t *testing.T) {
	someProto3 := buildProto3("-%ab", 11, "abba", 99)
	someProto3.SomeEmbedded = nil
	if err := someProto3.Validate(); err != nil {
		t.Fatalf("valiate shoudlnt fail on missing SomeEmbedded, not annotated")
	}
	someProto3.SomeEmbeddedExists = nil
	if err := someProto3.Validate(); err == nil {
		t.Fatalf("expected fail due to lacking SomeEmbeddedExists")
	}
}
