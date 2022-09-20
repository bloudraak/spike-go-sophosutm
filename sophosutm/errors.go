package sophosutm

import "fmt"

var ErrMissingBaseUrl = fmt.Errorf("missing base url")
var ErrMissingToken = fmt.Errorf("missing token")
var ErrUnauthorized = fmt.Errorf("unauthorized")
var ErrForbidden = fmt.Errorf("forbidden")
