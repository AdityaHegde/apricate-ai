package commons

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type ApricateResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

const ApiBase = "https://apricate.io/api"

var client = &http.Client{}

func HttpClient[T any](
	respObj *ApricateResponse[T], method string,
	requestUrl string, body io.Reader,
) (*ApricateResponse[T], error) {
	req, err := http.NewRequest(method, ApiBase+requestUrl, body)
	if err != nil {
		return nil, err
	}

	log.Printf("[%s] %s%s", method, ApiBase, requestUrl)

	req.Header.Add("Authorization", "Bearer "+os.Getenv("APRICATE_TOKEN"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(respObj); err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if respObj.Code > 1 {
		return nil, errors.New(respObj.Message)
	}

	return respObj, nil
}

func HttpClientGet[T any](
	respObj *ApricateResponse[T], requestUrl string,
) (*ApricateResponse[T], error) {
	return HttpClient(respObj, "GET", requestUrl, nil)
}
func HttpClientPost[T any](
	respObj *ApricateResponse[T], requestUrl string,
	requestBody interface{},
) (*ApricateResponse[T], error) {
	jsonValue, _ := json.Marshal(requestBody)
	return HttpClient(respObj, "POST", requestUrl, bytes.NewBuffer(jsonValue))
}
func HttpClientPatch[T any](
	respObj *ApricateResponse[T], requestUrl string,
	requestBody interface{},
) (*ApricateResponse[T], error) {
	jsonValue, _ := json.Marshal(requestBody)
	return HttpClient(respObj, "PATCH", requestUrl, bytes.NewBuffer(jsonValue))
}
func HttpClientDelete[T any](
	respObj *ApricateResponse[T], requestUrl string,
) (*ApricateResponse[T], error) {
	return HttpClient(respObj, "Delete", requestUrl, nil)
}
