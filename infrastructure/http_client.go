package infrastructure

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

type HTTPClient struct {
	Client *resty.Client
	Domain string
}

func (c *HTTPClient) GenerateHMAC(method, uri, secretKey, accessKey string) (string, error) {
	parts := strings.SplitN(uri, "?", 2)
	if len(parts) > 2 {
		return "", fmt.Errorf("incorrect uri format")
	}
	path := parts[0]
	query := ""
	if len(parts) == 2 {
		query = parts[1]
	}

	datetime := time.Now().UTC().Format("060102T150405Z")
	message := datetime + method + path + query

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	rawHmac := mac.Sum(nil)
	signature := hex.EncodeToString(rawHmac)

	return fmt.Sprintf("CEA algorithm=%s, access-key=%s, signed-date=%s, signature=%s", "HmacSHA256", accessKey, datetime, signature), nil
}

func (c *HTTPClient) ExecuteRequest(method, url, secretKey, accessKey string, body interface{}, result interface{}) error {
	authorization, err := c.GenerateHMAC(method, url, secretKey, accessKey)
	if err != nil {
		return fmt.Errorf("error generating HMAC: %v", err)
	}

	fullUrl := c.Domain + url
	request := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", authorization)

	if body != nil {
		request = request.SetBody(body)
	}

	var resp *resty.Response
	if method == "POST" {
		resp, err = request.Post(fullUrl)
	} else {
		resp, err = request.Get(fullUrl)
	}

	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

	// 응답 본문을 디코딩하여 result에 매핑
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}

	return nil
}
