package awsutils

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAWSConfigRegionExplicit(t *testing.T) {
	ctx := context.Background()

	tmp, err := GetAWSConfig(ctx, AWSConfigOptions{
		Region:  "us-east-1",
		Retries: 3,
		Verbose: false,
	})
	if err != nil {
		t.Errorf("Error was `%s`.", err)
	}

	actual := tmp.Region
	expected := "us-east-1"

	if actual != expected {
		t.Errorf("Result was `%s` instead of `%s`.", actual, expected)
	}
}

func TestGetAWSConfigRegionImplicit1(t *testing.T) {
	os.Setenv("AWS_REGION", "us-east-1")

	ctx := context.Background()

	tmp, err := GetAWSConfig(ctx, AWSConfigOptions{
		Retries: 3,
		Verbose: false,
	})
	if err != nil {
		t.Errorf("Error was `%s`.", err)
	}

	actual := tmp.Region
	expected := "us-east-1"

	if actual != expected {
		t.Errorf("Result was `%s` instead of `%s`.", actual, expected)
	}
}

func TestGetAWSConfigRegionImplicit2(t *testing.T) {
	os.Setenv("AWS_REGION", "us-east-2")

	ctx := context.Background()

	tmp, err := GetAWSConfig(ctx, AWSConfigOptions{
		Retries: 3,
		Verbose: false,
	})
	if err != nil {
		t.Errorf("Error was `%s`.", err)
	}

	actual := tmp.Region
	expected := "us-east-2"

	if actual != expected {
		t.Errorf("Result was `%s` instead of `%s`.", actual, expected)
	}
}

func TestGetAWSConfigProviderDontPanic(t *testing.T) {
	assert.NotPanics(t, panicMode, "The code panicked!")
}

func panicMode() {
	ctx := context.Background()

	_, _ = GetAWSConfig(ctx, AWSConfigOptions{
		Region:  "us-east-1",
		Retries: 3,
		Verbose: false,
	})
}
