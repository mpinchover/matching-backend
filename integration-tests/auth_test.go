package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	url := "https://dev-xm3amp8pvdypsq8r.us.auth0.com/oauth/token"

	// payload := strings.NewReader("{ \"scopes\": [ { \"value\": \"PERMISSION_NAME\", \"description\": \"PERMISSION_DESC\" }, { \"value\": \"PERMISSION_NAME\", \"description\": \"PERMISSION_DESC\" } ] }")

	payload := map[string]string{

		"audience":   "matching-api",
		"grant_type": "client_credentials",
	}

	postBody, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	reqBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, reqBody)

	req.Header.Add("content-type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	// req.Header.Add("content-type", "application/json")
	// req.Header.Add("authorization", "Bearer MGMT_API_ACCESS_TOKEN")
	// req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(string(body))

	// ensure acces token is nonempty
}
