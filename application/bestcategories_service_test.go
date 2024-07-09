package application

import (
	"cp-partners/infrastructure"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetBestCategories(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		//response domain.BestCategoriesResponseBody
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

	service := BestCategoriesService{
		Client: client,
	}

	response, err := service.GetBestCategories("1001", "50", "subId", "512x512", "secret", "access")
	assert.NoError(t, err)
	assert.Equal(t, "0", response.RCode)
	assert.Equal(t, "", response.RMessage)
	assert.Len(t, response.Data, 1)
	assert.Equal(t, "Category1", response.Data[0].CategoryName)
	assert.Equal(t, true, response.Data[0].IsRocket)
	assert.Equal(t, true, response.Data[0].IsFreeShipping)
	assert.Equal(t, 12345, response.Data[0].ProductId)
	assert.Equal(t, "https://example.com/image.jpg", response.Data[0].ProductImage)
	assert.Equal(t, "Product1", response.Data[0].ProductName)
	assert.Equal(t, 1000, response.Data[0].ProductPrice)
	assert.Equal(t, "https://example.com/product", response.Data[0].ProductUrl)
}
