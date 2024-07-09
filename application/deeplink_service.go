package application

import (
	"cp-partners/domain"
	"cp-partners/infrastructure"
)

type DeeplinkService struct {
	Client infrastructure.HTTPClient
}

func (s *DeeplinkService) GetDeeplink(goodsUrls []string, url, secretKey, accessKey, subId string) (*domain.DeeplinkResponseBody, error) {
	requestBody := domain.DeeplinkRequestBody{
		CoupangUrls: goodsUrls,
		SubId:       subId,
	}
	var response domain.DeeplinkResponseBody
	err := s.Client.ExecuteRequest("POST", url, secretKey, accessKey, requestBody, &response)
	return &response, err
}
