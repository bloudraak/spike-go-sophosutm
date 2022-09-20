package sophosutm

import (
    "net/http"
    "net/url"
    "os"
)

type Client struct {
    baseUrl *url.URL
    status  *StatusClient
}

func NewClient(baseUrl, token string) (*Client, error) {
    client := http.DefaultClient

    if token == "" {
        token = os.Getenv("SOPHOS_UTM_TOKEN")
    }
    if baseUrl == "" {
        baseUrl = os.Getenv("SOPHOS_UTM_BASEURL")
    }
    if baseUrl == "" {
        return nil, ErrMissingBaseUrl
    }
    if token == "" {
        return nil, ErrMissingToken
    }

    u, _ := url.Parse(baseUrl)
    authorizationHeader := encodeBase64("token:" + token)
    return &Client{
        status: &StatusClient{
            baseUrl:             u,
            client:              client,
            authorizationHeader: authorizationHeader,
        },
    }, nil
}

func (c *Client) Status() *StatusClient {
    return c.status
}
