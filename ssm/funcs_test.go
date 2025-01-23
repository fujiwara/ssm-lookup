package ssm_test

import (
	"context"
	"strings"
	"testing"
	"text/template"
)

func TestFuncMap(t *testing.T) {
	ctx := context.Background()
	app := newMockApp(mockGetParameter)
	tmpl := template.New("test")
	tmpl.Funcs(app.FuncMap(ctx))
	tmpl.Funcs(app.FuncMapWithName(ctx, "my_ssm"))
	tmpl = template.Must(tmpl.Parse(`{{ssm "/string"}}:{{my_ssm "/string"}}`))

	buf := &strings.Builder{}
	err := tmpl.Execute(buf, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.String() != "string value:string value" {
		t.Errorf("unexpected result: %s", buf.String())
	}
}
