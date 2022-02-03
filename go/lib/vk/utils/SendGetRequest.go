package utils

import (
	"encoding/json"
	"net/http"
)

func SendGetRequest(url string, query map[string]string, response interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	//toString := false

	q := req.URL.Query()
	for k, v := range query {
		//if k == "message" {
		//	toString = true
		//}
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return err
	}

	//if toString {
	//	bodyBytes, err := io.ReadAll(resp.Body)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	bodyString := string(bodyBytes)
	//	fmt.Println(bodyString)
	//}

	return json.NewDecoder(resp.Body).Decode(response)
}
