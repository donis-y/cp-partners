package application

import (
	"cp-partners/infrastructure"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetDeeplink(t *testing.T) {
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

	service := DeeplinkService{
		Client: client,
	}

	response, err := service.GetDeeplink([]string{"https://www.coupang.com/vp/products/7534234442"}, "/deeplink", "secret", "access", "subId")
	assert.NoError(t, err)
	assert.Equal(t, "0", response.RCode)
	assert.Equal(t, "", response.RMessage)
	assert.Len(t, response.Data, 1)
	assert.Equal(t, "https://example.com", response.Data[0].OriginalUrl)
	assert.Equal(t, "https://short.url", response.Data[0].ShortenUrl)
}
