package users

import (
  apricate_clients "AdityaHegde/apricate-ai/src/main/go/apricate-clients"
)

func GetUser() (*UserEntity, error) {
  var userResp apricate_clients.ApricateResponse[UserEntity]
  _, err := apricate_clients.HttpClient(&userResp, "GET", "/my/user")
  if err != nil {
    return nil, err
  }
  return &userResp.Data, nil
}
