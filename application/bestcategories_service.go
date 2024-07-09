package application

import (
	"cp-partners/domain"
	"cp-partners/infrastructure"
	"fmt"
	"strings"
)

type BestCategoriesService struct {
	Client infrastructure.HTTPClient
}

func (s *BestCategoriesService) GetBestCategories(categoryId, limit, subId, imageSize, secretKey, accessKey string) (*domain.BestCategoriesResponseBody, error) {
	url := fmt.Sprintf("/v2/providers/affiliate_open_api/apis/openapi/v1/products/bestcategories/%s", categoryId)
	queryParams := make([]string, 0)
	if limit != "" {
		queryParams = append(queryParams, "limit="+limit)
	}
	if subId != "" {
		queryParams = append(queryParams, "subId="+subId)
	}
	if imageSize != "" {
		queryParams = append(queryParams, "imageSize="+imageSize)
	}
	if len(queryParams) > 0 {
		url = url + "?" + strings.Join(queryParams, "&")
	}

	var response domain.BestCategoriesResponseBody
	err := s.Client.ExecuteRequest("GET", url, secretKey, accessKey, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
