# ssm-lookup

Lookup values in AWS SSM Parameter store. This is a port from [kayac/ecspresso/v2/ssm](https://github.com/kayac/ecspresso/tree/v2/ssm).

## Usage

```go
package main

import (
    "context"
    "fmt"
    "sync"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/fujiwara/ssm-lookup/ssm"
)

func main() {
	ctx := context.Background()
	cfg, _  := config.LoadDefaultConfig(ctx)
	cache := &sync.Map{}
	app := ssm.New(cfg, cache)
	value, _ := app.Lookup(ctx, parameterName)
	fmt.Println(value)
}
```

## LICENSE

MIT

## Authors

Liooo, fujiwara
