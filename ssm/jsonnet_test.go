package ssm_test

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-jsonnet"
)

func TestJsonnetNativeFunc(t *testing.T) {
	ctx := context.Background()
	app := newMockApp(mockGetParameter)

	funcs := app.JsonnetNativeFuncs(ctx)
	vm := jsonnet.MakeVM()
	for _, fn := range funcs {
		vm.NativeFunction(fn)
	}
	out, err := vm.EvaluateAnonymousSnippet("test.jsonnet", `
	local ssm = std.native("ssm");
	local ssm_list = std.native("ssm_list");
	{
		"string": ssm("/string"),
		"stringlist": [ssm_list("/stringlist", 0), ssm_list("/stringlist", 1)],
		"securestring": ssm("/securestring")
	}`+"\n")
	if err != nil {
		t.Fatal(err)
	}
	ob := new(bytes.Buffer)
	if err := json.Indent(ob, []byte(out), "", "  "); err != nil {
		t.Fatal(err)
	}
	eb := new(bytes.Buffer)
	expect := `{
      "securestring": "securestring value",
      "string": "string value",
      "stringlist": [
        "stringlist value 1",
        "stringlist value 2"
      ]
    }` + "\n"
	if err := json.Indent(eb, []byte(expect), "", "  "); err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(ob.String(), eb.String()); diff != "" {
		t.Errorf("unexpected output: %s", diff)
	}
}
