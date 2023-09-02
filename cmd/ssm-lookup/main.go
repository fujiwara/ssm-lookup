package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/fujiwara/ssm-lookup/ssm"
)

var cache = &sync.Map{}

var CLI struct {
	ParameterName string `arg:"" required:"" help:"SSM parameter name"`
	Index         *int   `arg:"" optional:"" help:"Index of the value in the StringList parameter"`
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	kctx := kong.Parse(&CLI)
	if kctx.Error != nil {
		return kctx.Error
	}
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		return err
	}
	app := ssm.New(cfg, cache)
	if CLI.Index != nil {
		value, err := app.Lookup(ctx, CLI.ParameterName, *CLI.Index)
		if err != nil {
			return err
		}
		fmt.Println(value)
	} else {
		value, err := app.Lookup(ctx, CLI.ParameterName)
		if err != nil {
			return err
		}
		fmt.Println(value)
	}
	return nil
}
