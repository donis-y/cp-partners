package domain

type DeeplinkRequestBody struct {
	CoupangUrls []string `json:"coupangUrls"`
	SubId       string   `json:"subId"`
}

type DeeplinkData struct {
	OriginalUrl string `json:"originalUrl"`
	ShortenUrl  string `json:"shortenUrl"`
}

type DeeplinkResponseBody struct {
	RCode    string         `json:"rCode"`
	RMessage string         `json:"rMessage"`
	Data     []DeeplinkData `json:"data"`
}
