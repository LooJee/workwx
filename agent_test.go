package workwx

import (
	"os"
	"testing"
)

func TestApp_GetAgentPermission(t *testing.T) {
	permission, err := New(os.Getenv("corp_id")).WithApp(os.Getenv("secret")).GetAgentPermission()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("permission: %+v", permission)
}
