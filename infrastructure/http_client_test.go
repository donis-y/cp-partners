package infrastructure

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGenerateHMAC(t *testing.T) {
	client := HTTPClient{}
	method := "GET"
	uri := "/test"
	secretKey := "SECRET_KEY"
	accessKey := "ACCESS_KEY"

	expectedDate := time.Now().UTC().Format("060102T150405Z")
	message := expectedDate + method + uri
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	expectedSignature := hex.EncodeToString(mac.Sum(nil))
	expectedAuthorization := "CEA algorithm=HmacSHA256, access-key=" + accessKey + ", signed-date=" + expectedDate + ", signature=" + expectedSignature

	authorization, err := client.GenerateHMAC(method, uri, secretKey, accessKey)
	assert.NoError(t, err)
	assert.Equal(t, expectedAuthorization, authorization)
}

func TestExecuteRequest(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	secretKey := os.Getenv("SECRET_KEY")
	accessKey := os.Getenv("ACCESS_KEY")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// HMAC 서명 생성
		expectedDate := time.Now().UTC().Format("060102T150405Z")
		message := expectedDate + "GET" + "/"
		mac := hmac.New(sha256.New, []byte(secretKey))
		mac.Write([]byte(message))
		expectedSignature := hex.EncodeToString(mac.Sum(nil))
		expectedAuthorization := "CEA algorithm=HmacSHA256, access-key=" + accessKey + ", signed-date=" + expectedDate + ", signature=" + expectedSignature

		assert.Equal(t, expectedAuthorization, r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"rCode":"0","rMessage":"","data":[]}`))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	client := HTTPClient{
		Client: resty.New(),
		Domain: ts.URL,
	}

	var result struct {
		RCode    string     `json:"rCode"`
		RMessage string     `json:"rMessage"`
		Data     []struct{} `json:"data"`
	}

	err := client.ExecuteRequest("GET", "/", secretKey, accessKey, nil, &result)
	assert.NoError(t, err)
	assert.Equal(t, "0", result.RCode)
	assert.Equal(t, "", result.RMessage)
	assert.NotNil(t, result.Data)
}
