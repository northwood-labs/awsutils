# Shared AWS Utilities for Northwood Labs

This is only a library. It does not work on its own. It is consumed by the various apps.

## Usage

```go
import (
    "github.com/northwood-labs/awsutils"
    "github.com/northwood-labs/golang-utils/exiterrorf"
)

func main() {
    ctx := context.Background()
    region := "us-west-2"
    retries := 5
    verbose := false

    config, err := awsutils.GetAWSConfig(ctx, region, retries, verbose)
    if err != nil {
        exiterrorf.ExitErrorf(err)
    }
}
```
