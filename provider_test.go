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

func TestApp_GetAppLicenseInfo(t *testing.T) {
	licenseInfo, err := New(os.Getenv("corp_id")).WithProvider(os.Getenv("secret")).GetAppLicenseInfo(context.Background(), os.Getenv("auth_corp_id"), os.Getenv("suite_id"))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("license info: %+v", licenseInfo)
}
