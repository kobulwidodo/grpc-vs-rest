package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Interface interface {
	GetList() (User, error)
	GetLarge() (Large, error)
}

type userClient struct {
	url string
}

func InitClient() Interface {
	return &userClient{
		url: "http://localhost:8080/api/v1",
	}
}

func (u *userClient) GetList() (User, error) {
	// Make the GET request
	resp, err := http.Get(fmt.Sprintf("%s%s", u.url, "/users"))
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	// Unmarshal the JSON response into a slice of User structs
	var users User
	err = json.Unmarshal(body, &users)
	if err != nil {
		return User{}, err
	}

	return users, nil
}

func (u *userClient) GetLarge() (Large, error) {
	// Make the GET request
	resp, err := http.Get(fmt.Sprintf("%s%s", u.url, "/test-large-data"))
	if err != nil {
		return Large{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Large{}, err
	}

	// Unmarshal the JSON response into a slice of Large structs
	var largeData Large
	err = json.Unmarshal(body, &largeData)
	if err != nil {
		return Large{}, err
	}

	return largeData, nil
}
