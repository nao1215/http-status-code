package http_test

import (
	"fmt"

	"github.com/nao1215/http-status-code/http"
)

func Example() {
	h := http.New()
	h = h.Search("404")
	fmt.Printf("http status code=%s, description=%s, rfc=%s\n", h.Code, h.Description, h.RFC)

	// Output:
	// http status code=404, description=Not Found, rfc=RFC9110, Section 15.5.5
}
