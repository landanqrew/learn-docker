package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyWithValidHeader(t *testing.T) {
	var header http.Header = http.Header{}
	header.Set("Authorization", "ApiKey some_auth_value")
	key, err := GetAPIKey(header)
	if err != nil {
		t.Errorf("GetAPIKey Failed. Expected Pass for Valid header, %s", err.Error())
	}
	t.Logf("Test Pass | apiKey: %s", key)
}

func TestGetAPIKeyWithInvalidHeader(t *testing.T) {
	var header http.Header = http.Header{}
	key, err := GetAPIKey(header)
	if err != nil {
		t.Logf("GetAPIKey Failed. Expected Failure for Invalid header, %s", err.Error())
	} else {
		t.Errorf("Test Pass, Expected Failure | apiKey: %s", key)
	}
}
