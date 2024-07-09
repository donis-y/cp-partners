package main

import (
	"cp-partners/application"
	"cp-partners/infrastructure"
	"cp-partners/interface"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	if err := infrastructure.LoadEnv(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	client := resty.New()
	httpClient := &infrastructure.HTTPClient{
		Client: client,
		Domain: "https://api-gateway.coupang.com",
	}

	deeplinkService := &application.DeeplinkService{
		Client: *httpClient,
	}

	bestCategoriesService := &application.BestCategoriesService{
		Client: *httpClient,
	}

	h := handler.Handler{
		DeeplinkService:       deeplinkService,
		BestCategoriesService: bestCategoriesService,
	}

	// Deeplink API 호출
	h.HandleDeeplink([]string{"https://www.coupang.com/vp/products/7534234442"})

	// Get Best Categories API 호출
	h.HandleBestCategories()

}
