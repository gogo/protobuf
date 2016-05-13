package extrastructcontent

import (
	"testing"
)

func TestExtraStructContent(t *testing.T) {
	first := ExtraStructContent{
		MyMagicValue:      42,
		MyOtherMagicValue: "such magic, wow",
	}
	content, err := first.Marshal()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	other := ExtraStructContent{}
	err = other.Unmarshal(content)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if other.MyMagicValue != 0 || other.MyOtherMagicValue != "" {
		t.Errorf("unexpected content")
	}
}
