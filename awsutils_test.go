package awsutils

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sts"
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
	t.Setenv("AWS_REGION", "us-east-1")

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
	t.Setenv("AWS_REGION", "us-east-2")

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

// Requires a valid AWS profile to be set. It's a bit fragile for anybody but
// me. Need to move this into GHA with proper automation.
func TestGetAWSConfigProfile(t *testing.T) {
	t.Setenv("AWS_PROFILE", "sandbox")
	t.Setenv("AWS_REGION", "us-east-2")

	ctx := context.Background()

	tmp, err := GetAWSConfig(ctx, AWSConfigOptions{
		Retries: 3,
		Verbose: false,
	})
	if err != nil {
		t.Errorf("Error was `%s`.", err)
	}

	client := sts.NewFromConfig(tmp)

	response, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		t.Fatalf(
			"Error with AWS SDK request: If you get this error, try running `aws sso login --sso-session nwl`.\n%s",
			err,
		)
	}

	actual := *response.Account
	expected := "590184084631"

	if actual != expected {
		t.Errorf("Result was `%s` instead of `%s`.", actual, expected)
	}
}

// func TestGetAWSConfigProviderDontPanic(t *testing.T) {
// 	assert.NotPanics(t, panicMode, "The code panicked!")
// }

// func panicMode() {
// 	ctx := context.Background()

// 	_, _ = GetAWSConfig(ctx, AWSConfigOptions{
// 		Region:  "us-east-1",
// 		Retries: 3,
// 		Verbose: false,
// 	})
// }
