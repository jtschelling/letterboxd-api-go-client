package letterboxd

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	m := "GET"
	e := "/fake"
	d := ""
	b := ""

	c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")
	builder := &Builder{}
	r := builder.
		Method(m).
		Endpoint(e).
		Data(d).
		Body(b).
		Build(c)

	if r.method != m {
		t.Errorf(errorMsg("Incorrect method generated.", m, r.method))
	}

	if r.endpoint != e {
		t.Errorf(errorMsg("Incorrect endpoint generated.", e, r.endpoint))
	}

	if r.body != b {
		t.Errorf(errorMsg("Incorrect body generated.", b, r.body))
	}
}

// TODO: This definitely needs to be tested specifically. Not sure how to get a baseline "correct" signature because of the timestamp and UUID requirements
func TestBuildSignature(t *testing.T) {
}

func errorMsg(m string, w string, g string) string {
	return fmt.Sprintf("%s\n\twanted:\t%s\n\tgot:\t%s", m, w, g)
}
