package sophosutm

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

type StatusClient struct {
    client              *http.Client
    baseUrl             *url.URL
    authorizationHeader string
}

func (c *StatusClient) Version() (*Version, error) {
    client := c.client

    var version Version
    var request *http.Request
    var err error
    var response *http.Response

    b := strings.Builder{}
    b.WriteString(strings.TrimSuffix(c.baseUrl.String(), "/"))
    b.WriteString("/api/status/version")
    u := b.String()

    request, err = http.NewRequest(http.MethodGet, u, nil)
    request.Header.Set("Accept", "application/json")
    request.Header.Set("User-Agent", "go-sophosutm")
    request.Header.Set("Authorization", "Basic "+c.authorizationHeader)
    request.Header.Set("Content-Type", "application/json")

    if err != nil {
        return nil, err
    }
    response, err = client.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()
    switch response.StatusCode {
    case 200:
        err = json.NewDecoder(response.Body).Decode(&version)
        if err != nil {
            return nil, err
        }
        return &version, nil
    case 401:
        return nil, ErrUnauthorized
    case 403:
        return nil, ErrForbidden
    default:
        return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
    }
    return nil, nil
}
