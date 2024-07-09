package handler

import (
	"cp-partners/application"
	"cp-partners/infrastructure"
	"encoding/json"
	"fmt"
)

type Handler struct {
	DeeplinkService       *application.DeeplinkService
	BestCategoriesService *application.BestCategoriesService
}

func (h *Handler) HandleDeeplink(goodsUrls []string) {
	accessKey := infrastructure.GetEnv("ACCESS_KEY")
	secretKey := infrastructure.GetEnv("SECRET_KEY")
	subId := infrastructure.GetEnv("SUB_ID")

	if accessKey == "" || secretKey == "" || subId == "" {
		fmt.Println("Error: ACCESS_KEY, SECRET_KEY or SUB_ID not set in environment")
		return
	}

	response, err := h.DeeplinkService.GetDeeplink(goodsUrls, "/v2/providers/affiliate_open_api/apis/openapi/v1/deeplink", secretKey, accessKey, subId)
	if err != nil {
		fmt.Println("Error calling deeplink API:", err)
		return
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling deeplink response:", err)
		return
	}
	fmt.Printf("Deeplink Response:\n%s\n", string(responseJSON))
}

func (h *Handler) HandleBestCategories() {
	accessKey := infrastructure.GetEnv("ACCESS_KEY")
	secretKey := infrastructure.GetEnv("SECRET_KEY")
	subId := infrastructure.GetEnv("SUB_ID")

	if accessKey == "" || secretKey == "" || subId == "" {
		fmt.Println("Error: ACCESS_KEY, SECRET_KEY or SUB_ID not set in environment")
		return
	}

	categoryId := "1001"
	limit := "50"
	imageSize := "512x512"

	response, err := h.BestCategoriesService.GetBestCategories(categoryId, limit, subId, imageSize, secretKey, accessKey)
	if err != nil {
		fmt.Println("Error calling getBestCategories API:", err)
		return
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling best categories response:", err)
		return
	}
	fmt.Printf("Best Categories Response:\n%s\n", string(responseJSON))
}
