package helpers

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HttpRequest(endPoint string) ([]byte, error) {
	apiUrl := os.Getenv("INFO_API_URL")
	apiToken := os.Getenv("INFO_API_TOKEN")
	apiMethod := os.Getenv("INFO_API_METHOD")
	client := &http.Client{}
	request, err := http.NewRequest(apiMethod, apiUrl+endPoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Auth-Token", apiToken)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
