package awsutils

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/smithy-go/logging"
)

// NoOpRateLimit to prevent limiting of queries to AWS.
type NoOpRateLimit struct{}

// AddTokens to return nil for NoOpRateLimit.
func (NoOpRateLimit) AddTokens(uint) error {
	return nil
}

// GetToken will return nil so that there will be no rate limiting.
func (NoOpRateLimit) GetToken(context.Context, uint) (func() error, error) {
	return noOpToken, nil
}

func noOpToken() error {
	return nil
}

// GetAWSConfig returns a standard AWS config object pre-configured for use with regions, retries, and verbosity.
func GetAWSConfig(ctx context.Context, region string, retries int, verbose bool) (aws.Config, error) {
	emptyConfig := aws.Config{}

	// Pull AWS credentials from the environment.
	conf, err := config.LoadDefaultConfig(
		ctx,
		config.WithDefaultRegion(region),
		config.WithRegion(region),
		config.WithRetryer(func() aws.Retryer {
			retryLogic := retry.NewStandard()
			retry.AddWithMaxAttempts(retryLogic, retries)

			return retryLogic
		}),
	)
	if err != nil {
		return emptyConfig, fmt.Errorf("AWS configuration error: %w", err)
	}

	if verbose {
		conf.Logger = logging.NewStandardLogger(os.Stderr)
	}

	return conf, nil
}
