package handler

import (
	"cp-partners/application"
	"cp-partners/infrastructure"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func setupEnvironment() {
	_ = os.Setenv("ACCESS_KEY", "test-access-key")
	_ = os.Setenv("SECRET_KEY", "test-secret-key")
	_ = os.Setenv("SUB_ID", "test-sub-id")
}

func TestHandleDeeplink(t *testing.T) {
	setupEnvironment()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"rCode":"0","rMessage":"","data":[{"originalUrl":"https://example.com","shortenUrl":"https://short.url"}]}`))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	client := infrastructure.HTTPClient{
		Client: resty.New(),
		Domain: ts.URL,
	}

	deeplinkService := &application.DeeplinkService{
		Client: client,
	}

	h := Handler{
		DeeplinkService: deeplinkService,
	}

	// stdout을 리디렉션하여 출력 캡처
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	h.HandleDeeplink([]string{"https://www.coupang.com/vp/products/7534234442"})

	// stdout을 복원하고 출력 캡처
	err := w.Close()
	if err != nil {
		return
	}
	os.Stdout = old

	var buf strings.Builder
	_, err = io.Copy(&buf, r)
	if err != nil {
		return
	}
	assert.Contains(t, buf.String(), `"rCode": "0"`)
	assert.Contains(t, buf.String(), `"originalUrl": "https://example.com"`)
	assert.Contains(t, buf.String(), `"shortenUrl": "https://short.url"`)
}

func TestHandleBestCategories(t *testing.T) {
	setupEnvironment()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"rCode":"0","rMessage":"","data":[{"categoryName":"Category1","isRocket":true,"isFreeShipping":true,"productId":12345,"productImage":"https://example.com/image.jpg","productName":"Product1","productPrice":1000,"productUrl":"https://example.com/product"}]}`))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	client := infrastructure.HTTPClient{
		Client: resty.New(),
		Domain: ts.URL,
	}

	bestCategoriesService := &application.BestCategoriesService{
		Client: client,
	}

	h := Handler{
		BestCategoriesService: bestCategoriesService,
	}

	// stdout을 리디렉션하여 출력 캡처
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	h.HandleBestCategories()

	// stdout을 복원하고 출력 캡처
	err := w.Close()
	if err != nil {
		return
	}
	os.Stdout = old

	var buf strings.Builder
	_, err = io.Copy(&buf, r)
	if err != nil {
		return
	}
	assert.Contains(t, buf.String(), `"rCode": "0"`)
	assert.Contains(t, buf.String(), `"categoryName": "Category1"`)
	assert.Contains(t, buf.String(), `"productName": "Product1"`)
}
