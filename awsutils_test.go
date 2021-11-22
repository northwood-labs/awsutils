package awsutils

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAWSConfigRegionExplicit(t *testing.T) {
	ctx := context.Background()
	region := "us-east-1"
	profile := "default"
	retries := 3
	verbose := false

	expected := "us-east-1"
	tmp, _ := GetAWSConfig(ctx, region, profile, retries, verbose)
	actual := tmp.Region

	if actual != expected {
		t.Errorf("Result was `%s` instead of `%s`.", actual, expected)
	}
}

func TestGetAWSConfigRegionImplicit1(t *testing.T) {
	os.Setenv("AWS_REGION", "us-east-1")

	ctx := context.Background()
	region := ""
	profile := "default"
	retries := 3
	verbose := false

	expected := "us-east-1"
	tmp, _ := GetAWSConfig(ctx, region, profile, retries, verbose)
	actual := tmp.Region

	if actual != expected {
		t.Errorf("Result was `%s` instead of `%s`.", actual, expected)
	}
}

func TestGetAWSConfigRegionImplicit2(t *testing.T) {
	os.Setenv("AWS_REGION", "us-east-2")

	ctx := context.Background()
	region := ""
	profile := "default"
	retries := 3
	verbose := false

	expected := "us-east-2"
	tmp, _ := GetAWSConfig(ctx, region, profile, retries, verbose)
	actual := tmp.Region

	if actual != expected {
		t.Errorf("Result was `%s` instead of `%s`.", actual, expected)
	}
}

func TestGetAWSConfigProviderDontPanic(t *testing.T) {
	assert.NotPanics(t, panicMode, "The code did panicked!")
}

func panicMode() {
	ctx := context.Background()
	region := "us-east-1"
	profile := ""
	retries := 3
	verbose := false

	_, _ = GetAWSConfig(ctx, region, profile, retries, verbose)
}
