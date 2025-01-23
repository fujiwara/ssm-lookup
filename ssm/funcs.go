package ssm

import (
	"context"
	"fmt"
	"sync"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func FuncMap(ctx context.Context, cfg aws.Config) (template.FuncMap, error) {
	cache := sync.Map{}
	app := New(cfg, &cache)
	return app.FuncMap(ctx), nil
}

func (app *App) FuncMap(ctx context.Context) template.FuncMap {
	return app.FuncMapWithName(ctx, "ssm")
}

func (app *App) FuncMapWithName(ctx context.Context, name string) template.FuncMap {
	return template.FuncMap{
		name: func(paramName string, index ...int) (string, error) {
			value, err := app.Lookup(ctx, paramName, index...)
			if err != nil {
				return "", fmt.Errorf("failed to lookup ssm parameter: %w", err)
			}
			return value, nil
		},
	}
}
