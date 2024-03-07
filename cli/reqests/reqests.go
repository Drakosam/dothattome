package reqests

import (
	"fmt"
	"io"
	"net/http"
)

// CallTestGet is a function.
func CallTestGet() {
	response, err := http.Get("http://127.0.0.1:8081/test")
	if err != nil {
		fmt.Println(err)
	}

	rdata, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(rdata))
}
