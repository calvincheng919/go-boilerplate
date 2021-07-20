package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	T "my-service/src/types"
	"net/http"
	"testing"
)

// test name must be of form Test*
func TestVersionRoute(t *testing.T) {
	// make request
	url := `http://localhost:8080/version`
	response, err := http.Get(url)
	if err != nil {
		t.Errorf("request error in test: %v", err)
	}
	defer func() { _ = response.Body.Close() }()

	// read response
	payload, decodeErr := ioutil.ReadAll(response.Body)
	if decodeErr != nil {
		t.Errorf("decode error in test: %v", decodeErr)
	}
	var result T.Map
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Errorf("unmarshal error in test: %v", err)
	}

	// test response
	if result["version"] != "0.0.0" {
		fmt.Println(string(payload))
		t.Errorf("wrong version!\n Expected '0.0.0', but got %s\n Full response: %s", result["version"], string(payload))
	}
}
