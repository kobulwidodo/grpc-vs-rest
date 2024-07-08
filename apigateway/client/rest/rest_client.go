package rest_client

import "apigateway/client/rest/user"

type RestClient struct {
	User user.Interface
}

func Init() *RestClient {
	rc := &RestClient{
		User: user.InitClient(),
	}

	return rc
}
