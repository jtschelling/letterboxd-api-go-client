package letterboxd

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	// precomputed request signature
	// Letterboxd team ruby client signature method used as source of truth
	eSig := "9e0e784e7c790b503637e9fa2b8e00e2d681f6d40b9b5949784900a30b8e9185"
	m := "GET"
	e := "/fake"
	b := ""

	c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")
	builder := &Builder{}
	r := builder.
		Method(m).
		Endpoint(e).
		Body(b).
		Build(c)

	if r.method != m {
		t.Errorf(errorMsg("Incorrect method generated.", m, r.method))
	}

	if r.endpoint != e {
		t.Errorf(errorMsg("Incorrect endpoint generated.", e, r.endpoint))
	}

	if r.body != b {
		t.Errorf(errorMsg("Incorrect method generated.", b, r.body))
	}

	if r.signature != eSig {
		t.Errorf(errorMsg("Incorrect signature generated.", eSig, r.signature))
	}
}

func TestBuildSignature(t *testing.T) {
	// precomputed request signature
	// Letterboxd team ruby client signature method used as source of truth
	eSig := "9e0e784e7c790b503637e9fa2b8e00e2d681f6d40b9b5949784900a30b8e9185"
	c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")
	r := c.NewRequest("GET", "/fake", "")

	if r.signature != eSig {
		t.Errorf(errorMsg("Incorrect signature generated.", eSig, r.signature))
	}
}

func errorMsg(m string, w string, g string) string {
	return fmt.Sprintf("%s\n\twanted:\t%s\n\tgot:\t%s", m, w, g)
}
