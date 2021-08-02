package pagecall

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type user struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	Name           string `json:"name"`
	ApplicationID  string `json:"application_id"`
	OrganizationID string `json:"organization_id"`
	AccessToken    string `json:"access_token"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func (p pageCallClient) CreateUser(id string, name string) (*user, error) {
	reqBody := make(map[string]interface{})
	reqBody["user_id"] = id
	reqBody["name"] = name

	ubytes, _ := json.Marshal(reqBody)
	payload := bytes.NewBuffer(ubytes)

	body, err := p.request("POST", "/users", payload)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		User user `json:"user"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return &respBody.User, nil
}

func (p pageCallClient) GetUser(userID string) (*user, error) {
	path := fmt.Sprintf("/users/%s", userID)
	body, err := p.request("GET", path, nil)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		User user `json:"user"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return &respBody.User, nil
}

func (p pageCallClient) GetUsers(offset int, limit int) ([]user, error) {
	path := fmt.Sprintf("/users?offset=%d&limit=%d", offset, limit)
	body, err := p.request("GET", path, nil)

	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		Users []user `json:"users"`
	}

	respBody := &ResponseBody{}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		return nil, err
	}

	return respBody.Users, nil
}
