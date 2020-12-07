package letterboxd

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	// // precomputed request signature. haven't figured out how to test this func
	// eSig := ""
	// eURL := "https://api.letterboxd.com/api/v0/fake?apikey=FAKE_KEY&nonce=&timestamp=&signature="
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
		t.Errorf(errorMsg("Incorrect body generated.", b, r.body))
	}

	// if r.signature != eSig {
	// 	t.Errorf(errorMsg("Incorrect signature generated.", eSig, r.signature))
	// }

	// if r.url != eURL {
	// 	t.Errorf(errorMsg("Incorrect URL generated.", eURL, r.url))
	// }
}

func TestBuildSignature(t *testing.T) {
	// // precomputed request signature. haven't figured out how to test this func
	// eSig := ""
	// c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")
	// r := c.NewRequest("GET", "/fake", "")

	// if r.signature != eSig {
	// 	t.Errorf(errorMsg("Incorrect signature generated.", eSig, r.signature))
	// }
}

func errorMsg(m string, w string, g string) string {
	return fmt.Sprintf("%s\n\twanted:\t%s\n\tgot:\t%s", m, w, g)
}
