package graphqlclient

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/machinebox/graphql"
	"io/ioutil"
	"net/http"
)

type ResponseMetadata struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// APIResponse is a generic wrapper for GraphQL responses, where T is the type of the data field.
type APIResponse[T any] struct {
	Response ResponseMetadata `json:"response"`
	Data     T                `json:"data"`
}

type Client struct {
	httpClient    *http.Client
	graphqlClient *graphql.Client
	url           string
}

func NewClient(url string) *Client {
	return &Client{
		httpClient:    &http.Client{},
		graphqlClient: graphql.NewClient(url),
		url:           url,
	}
}

func (c *Client) DoQuery(query string, response interface{}) error {
	req := graphql.NewRequest(query)
	// Set any additional headers here
	// req.Header.Set("Authorization", "Bearer " + token)
	return c.graphqlClient.Run(context.Background(), req, response)
}

func (c *Client) DoQueryWithHTTP(query string) ([]byte, error) {
	requestBody, err := json.Marshal(map[string]string{"query": query})
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Post(c.url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
