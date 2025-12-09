package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	baseURL string
	token   string
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
	}
}

type LoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (c *Client) Login(user, password string) error {
	data := LoginRequest{
		User:     user,
		Password: password,
	}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.baseURL+"/login", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return nil
}
