package casttypewith

import (
	"math"
	"math/big"
	"testing"
)

func TestVals(t *testing.T) {
	orig := &Castaway{
		MyBigInt: big.NewInt(math.MinInt64),
	}

	t.Logf("%#v", orig)

	buf, err := orig.Marshal()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	final := &Castaway{}

	err = final.Unmarshal(buf)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = orig.VerboseEqual(final)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
