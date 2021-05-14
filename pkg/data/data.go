// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	// Constants error.
	errDecodeResponseBody = "error in decode response body"
)

// BodyDecode method decode body of the response object.
func BodyDecode(contentBody []byte) ([]interface{}, error) {
	var body []interface{}

	err := json.Unmarshal(contentBody, &body)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", errDecodeResponseBody, err)
	}

	return body, nil
}

// ResponseDecode method transforms struct to []interface{}.
func ResponseDecode(v interface{}) ([]interface{}, error) {
	respBodyByte, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var respBody []interface{}

	err = json.NewDecoder(bytes.NewReader(respBodyByte)).Decode(&respBody)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return respBody, nil
}
