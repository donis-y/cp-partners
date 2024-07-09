
# 쿠팡 파트너스 API Integration

이 저장소는 도메인 주도 설계(DDD)와 테스트 주도 개발(TDD) 방식을 사용하여 쿠팡 API와 통합하는 Go 애플리케이션의 예제를 제공합니다. 이 애플리케이션은 딥링크를 생성하고 카테고리별 베스트셀러 상품을 조회하는 방법을 설명합니다.

## 디렉토리 구조

```
cp-partners
├── cmd
│   └── main.go
├── domain
│   ├── deeplink.go
│   └── bestcategories.go
├── application
│   ├── deeplink_service.go
│   ├── deeplink_service_test.go
│   ├── bestcategories_service.go
│   └── bestcategories_service_test.go
├── infrastructure
│   ├── http_client.go
│   ├── http_client_test.go
│   └── env_config.go
└── interface
    ├── handler.go
    └── handler_test.go
```

## 시작하기

### 사전 요구사항

- Go 1.20 이상
- `ACCESS_KEY`와 `SECRET_KEY` 환경 변수를 설정하거나 설정 파일을 제공해야 합니다.

### 환경 변수 설정

터미널에서 필요한 환경 변수를 설정할 수 있습니다:

```sh
export ACCESS_KEY=your_access_key
export SECRET_KEY=your_secret_key
```

또는 프로젝트의 루트 디렉토리에 `config.yaml` 파일을 생성할 수 있습니다:

```yaml
access_key: your_access_key
secret_key: your_secret_key
```

### 설치

1. 저장소를 클론합니다:

```sh
git clone https://github.com/donis-y/cp-partners.git
cd cp-partners
```

2. 종속성을 설치합니다:

```sh
go mod tidy
```

### 애플리케이션 실행

`cmd` 디렉토리로 이동하여 메인 애플리케이션을 실행합니다:

```sh
cd cmd
go run main.go
```

### 테스트 실행

다음 명령을 사용하여 테스트를 실행할 수 있습니다:

```sh
go test ./...
```

## 프로젝트 구조

### 도메인 계층

핵심 비즈니스 객체를 정의합니다:

- `domain/deeplink.go`: 딥링크 API와 관련된 구조체를 정의합니다.
- `domain/bestcategories.go`: 베스트 카테고리 API와 관련된 구조체를 정의합니다.

### 애플리케이션 계층

도메인 객체와 상호작용하는 서비스를 포함합니다:

- `application/deeplink_service.go`: 딥링크 생성을 위한 서비스를 포함합니다.
- `application/bestcategories_service.go`: 베스트 카테고리를 가져오는 서비스를 포함합니다.
- `application/deeplink_service_test.go`: 딥링크 서비스에 대한 테스트 케이스를 포함합니다.
- `application/bestcategories_service_test.go`: 베스트 카테고리 서비스에 대한 테스트 케이스를 포함합니다.

### 인프라스트럭처 계층

외부 시스템과의 통신을 처리합니다:

- `infrastructure/http_client.go`: 쿠팡 API와 상호작용하기 위한 HTTP 클라이언트를 포함합니다.
- `infrastructure/http_client_test.go`: HTTP 클라이언트에 대한 테스트 케이스를 포함합니다.
- `infrastructure/env_config.go`: `.env` 파일이나 환경 변수에서 설정을 로드합니다.

### 인터페이스 계층

사용자 입력과 출력을 관리합니다:

- `interface/handler.go`: 다양한 API 상호작용을 위한 핸들러를 포함합니다.
- `interface/handler_test.go`: 핸들러에 대한 테스트 케이스를 포함합니다.

## 예제

`DeeplinkService`를 사용하여 딥링크를 생성하는 간단한 예제입니다:

```go
package main

import (
    "cp-partners/application"
    "cp-partners/infrastructure"
    "fmt"
    "github.com/go-resty/resty/v2"
)

func main() {
    err := infrastructure.LoadEnv()
    if err != nil {
        fmt.Println("Error loading .env file:", err)
        return
    }

    accessKey := infrastructure.GetEnv("ACCESS_KEY")
    secretKey := infrastructure.GetEnv("SECRET_KEY")
    subId := infrastructure.GetEnv("SUB_ID")

    client := infrastructure.HTTPClient{
        Client: resty.New(),
        Domain: "https://api-gateway.coupang.com",
    }

    deeplinkService := application.DeeplinkService{
        Client: client,
    }

    response, err := deeplinkService.GetDeeplink([]string{"https://www.coupang.com/vp/products/7534234442"}, "/v2/providers/affiliate_open_api/apis/openapi/v1/deeplink", secretKey, accessKey, subId)
    if err != nil {
        fmt.Println("Error calling deeplink API:", err)
        return
    }

    fmt.Printf("Deeplink Response: %+v
", response)
}
```

## 기여하기

이 프로젝트에 기여하고 싶다면 포크를 생성하고 풀 리퀘스트를 제출하세요. 모든 기여를 환영합니다!

## 라이선스

이 프로젝트는 MIT 라이선스에 따라 라이선스가 부여됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.
