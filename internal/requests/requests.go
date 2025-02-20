package requests

import (
	"fmt"
	"io"
	"net/http"
)

func MakeGETRequest(URL string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		fmt.Printf("No info about requested entity\n")
		return []byte{}, nil
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return []byte(resBody), nil
}
