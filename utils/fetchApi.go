package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchApi(url string) any {
	apiUrl := url
	request, error := http.NewRequest("GET", apiUrl, nil)

	if error != nil {
		fmt.Println(error)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	var responseJson ResponseAirPollution
	json.Unmarshal(responseBody, &responseJson)

	// clean up memory after execution
	defer response.Body.Close()

	return responseJson
}
