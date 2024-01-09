package workwx

import (
	"context"
	"os"
	"testing"
)

func TestApp_getSuiteAccessToken(t *testing.T) {
	token, err := New(os.Getenv("suite_id")).WithSuite(os.Getenv("suite_secret"), func(ctx context.Context, key string) (string, error) {
		return os.Getenv("suite_ticket"), nil
	}).getSuiteAccessToken()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", token)
}
