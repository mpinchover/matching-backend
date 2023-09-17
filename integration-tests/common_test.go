package integrationtests

// func getAuthToken() (string, error) {

// 	url := "https://dev-xm3amp8pvdypsq8r.us.auth0.com/oauth/token"

// 	payload := map[string]string{
// 		// TODO, convert to .env
// 		"client_id":     os.Getenv("TEST_AUTH_0_CLIENT_ID"),
// 		"client_secret": os.Getenv("TEST_AUTH_0_CLIENT_SECRET"),
// 		"audience":      "matching-api",
// 		"grant_type":    "client_credentials",
// 	}

// 	postBody, err := json.Marshal(payload)
// 	if err != nil {
// 		return "", err
// 	}

// 	reqBody := bytes.NewBuffer(postBody)
// 	req, err := http.NewRequest("POST", url, reqBody)

// 	req.Header.Add("content-type", "application/json")

// 	if err != nil {
// 		return "", err
// 	}

// 	// req.Header.Add("content-type", "application/json")
// 	// req.Header.Add("authorization", "Bearer MGMT_API_ACCESS_TOKEN")
// 	// req.Header.Add("cache-control", "no-cache")

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}

// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	fmt.Println(res)
// 	fmt.Println(string(body))

// 	// ensure acces token is nonempty
// }
