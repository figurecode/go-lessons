package main

import (
	"encoding/json"
	"fmt"
)

type (
	Response struct {
		Header ResponseHeader `json:"header"`
		Data   ResponseData   `json:"data,omitempty"`
	}

	ResponseHeader struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	}

	ResponseData []ResponseDataItem

	ResponseDataItem struct {
		Type       string                `json:"type"`
		Id         int                   `json:"id"`
		Attributes ResponseDataItemAttrs `json:"attributes"`
	}

	ResponseDataItemAttrs struct {
		Email      string `json:"email"`
		ArticleIds []int  `json:"article_ids"`
	}
)

func ReadResponse(rawResp string) (Response, error) {
	resp := Response{}

	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
	}

	return resp, nil
}

const resp = `{
	"header": {
		"code": 0,
		"message": ""
	},
	"data": [{
		"type": "user",
		"id": 100,
		"attributes": {
			"email": "bob@yandex.ru",
			"article_ids": [10, 11, 12]
		}
	}]
}`

func main() {

	fmt.Println(ReadResponse(resp))
}
