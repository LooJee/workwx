package workwx

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestApp_GetAppQrCode(t *testing.T) {
	qrcode, err := New(os.Getenv("corp_id")).WithProvider(os.Getenv("secret")).GetCustomizedAuthUrl(context.Background(), "hello", strings.Split(os.Getenv("temp_ids"), ","))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", qrcode)
}
