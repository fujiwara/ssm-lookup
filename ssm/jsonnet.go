package ssm

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func JsonnetNativeFuncs(ctx context.Context, cfg aws.Config) ([]*jsonnet.NativeFunction, error) {
	cache := sync.Map{}
	app := New(cfg, &cache)
	return app.JsonnetNativeFuncs(ctx), nil
}

func (app *App) JsonnetNativeFuncs(ctx context.Context) []*jsonnet.NativeFunction {
	return []*jsonnet.NativeFunction{
		{
			Name: "ssm",
			Func: func(p []interface{}) (interface{}, error) {
				paramName, ok := p[0].(string)
				if !ok {
					return nil, fmt.Errorf("ssm: parameter name must be a string")
				}
				return app.Lookup(ctx, paramName)
			},
			Params: []ast.Identifier{"name"},
		},
		{
			Name: "ssm_list",
			Func: func(p []interface{}) (interface{}, error) {
				paramName, ok := p[0].(string)
				if !ok {
					return nil, fmt.Errorf("ssm: parameter name must be a string")
				}
				index, ok := p[1].(float64)
				if !ok {
					return nil, fmt.Errorf("ssm: index must be a number")
				}
				return app.Lookup(ctx, paramName, int(index))
			},
			Params: []ast.Identifier{"name", "index"},
		},
	}
}
