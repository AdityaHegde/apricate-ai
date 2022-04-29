package apricate_clients

import (
  "encoding/json"
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

func HttpClient[T any](respObj *ApricateResponse[T], method string, requestUrl string) (*ApricateResponse[T], error) {
  client := &http.Client{}
  req, err := http.NewRequest(method, ApiBase+requestUrl, nil)
  if err != nil {
    return nil, err
  }

  req.Header.Add("Authorization", "Bearer "+os.Getenv("APRICATE_TOKEN"))
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  if err := json.NewDecoder(resp.Body).Decode(respObj); err != nil {
    log.Fatalln(err)
    return nil, err
  }

  return respObj, nil
}
